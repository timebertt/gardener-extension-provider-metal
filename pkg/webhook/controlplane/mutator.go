package controlplane

import (
	"context"
	"fmt"

	"github.com/gardener/gardener/extensions/pkg/controller/operatingsystemconfig/oscommon/cloudinit"
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	gcontext "github.com/gardener/gardener/extensions/pkg/webhook/context"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/original/components/kubelet"
	"github.com/gardener/gardener/pkg/operation/botanist/component/extensions/operatingsystemconfig/utils"

	"errors"

	"github.com/coreos/go-systemd/v22/unit"
	druidv1alpha1 "github.com/gardener/etcd-druid/api/v1alpha1"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

const (
	KonnectivityDeploymentName = "konnectivity-server"
	KonnectivityContainerName  = "konnectivity-server"
)

/*
FIXME:
This is a copy of gardener's controlplane generic mutator https://github.com/gardener/gardener/blob/master/extensions/pkg/webhook/controlplane/genericmutator/mutator.go
It is necessary so that we can fix the hostPort property of the konnectivity-server deployment, which is not possible with gardener's version

This should be removed once we drop the konnectivity feature.
*/

// Ensurer ensures that various standard Kubernets controlplane objects conform to the provider requirements.
// If they don't initially, they are mutated accordingly.
type Ensurer interface {
	// EnsureKubeAPIServerService ensures that the kube-apiserver service conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubeAPIServerService(ctx context.Context, gctx gcontext.GardenContext, new, old *corev1.Service) error
	// EnsureKubeAPIServerDeployment ensures that the kube-apiserver deployment conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubeAPIServerDeployment(ctx context.Context, gctx gcontext.GardenContext, new, old *appsv1.Deployment) error
	// EnsureKubeControllerManagerDeployment ensures that the kube-controller-manager deployment conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubeControllerManagerDeployment(ctx context.Context, gctx gcontext.GardenContext, new, old *appsv1.Deployment) error
	// EnsureKubeSchedulerDeployment ensures that the kube-scheduler deployment conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubeSchedulerDeployment(ctx context.Context, gctx gcontext.GardenContext, new, old *appsv1.Deployment) error
	// EnsureETCD ensures that the etcds conform to the respective provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureETCD(ctx context.Context, gctx gcontext.GardenContext, new, old *druidv1alpha1.Etcd) error
	// EnsureKubeletServiceUnitOptions ensures that the kubelet.service unit options conform to the provider requirements.
	EnsureKubeletServiceUnitOptions(ctx context.Context, gctx gcontext.GardenContext, new, old []*unit.UnitOption) ([]*unit.UnitOption, error)
	// EnsureKubeletConfiguration ensures that the kubelet configuration conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubeletConfiguration(ctx context.Context, gctx gcontext.GardenContext, new, old *kubeletconfigv1beta1.KubeletConfiguration) error
	// EnsureKubernetesGeneralConfiguration ensures that the kubernetes general configuration conforms to the provider requirements.
	// "old" might be "nil" and must always be checked.
	EnsureKubernetesGeneralConfiguration(ctx context.Context, gctx gcontext.GardenContext, new, old *string) error
	// ShouldProvisionKubeletCloudProviderConfig returns true if the cloud provider config file should be added to the kubelet configuration.
	ShouldProvisionKubeletCloudProviderConfig(ctx context.Context, gctx gcontext.GardenContext) bool
	// EnsureKubeletCloudProviderConfig ensures that the cloud provider config file content conforms to the provider requirements.
	EnsureKubeletCloudProviderConfig(context.Context, gcontext.GardenContext, *string, string) error
	// EnsureAdditionalUnits ensures additional systemd units
	// "old" might be "nil" and must always be checked.
	EnsureAdditionalUnits(ctx context.Context, gctx gcontext.GardenContext, new, old *[]extensionsv1alpha1.Unit) error
	// EnsureAdditionalFile ensures additional systemd files
	// "old" might be "nil" and must always be checked.
	EnsureAdditionalFiles(ctx context.Context, gctx gcontext.GardenContext, new, old *[]extensionsv1alpha1.File) error
}

// NewMutator creates a new controlplane mutator.
func NewMutator(
	ensurer Ensurer,
	unitSerializer utils.UnitSerializer,
	kubeletConfigCodec kubelet.ConfigCodec,
	fciCodec utils.FileContentInlineCodec,
	logger logr.Logger,
) extensionswebhook.Mutator {
	return &mutator{
		ensurer:            ensurer,
		unitSerializer:     unitSerializer,
		kubeletConfigCodec: kubeletConfigCodec,
		fciCodec:           fciCodec,
		logger:             logger.WithName("mutator"),
	}
}

type mutator struct {
	client             client.Client
	ensurer            Ensurer
	unitSerializer     utils.UnitSerializer
	kubeletConfigCodec kubelet.ConfigCodec
	fciCodec           utils.FileContentInlineCodec
	logger             logr.Logger
}

// InjectClient injects the given client into the ensurer.
func (m *mutator) InjectClient(client client.Client) error {
	m.client = client
	return nil
}

// InjectFunc injects stuff into the ensurer.
func (m *mutator) InjectFunc(f inject.Func) error {
	return f(m.ensurer)
}

// Mutate validates and if needed mutates the given object.
func (m *mutator) Mutate(ctx context.Context, new, old client.Object) error {
	// If the object does have a deletion timestamp then we don't want to mutate anything.
	if new.GetDeletionTimestamp() != nil {
		return nil
	}
	gctx := gcontext.NewGardenContext(m.client, new)

	switch x := new.(type) {
	case *corev1.Service:
		switch x.Name {
		case v1beta1constants.DeploymentNameKubeAPIServer:
			var oldSvc *corev1.Service
			if old != nil {
				var ok bool
				oldSvc, ok = old.(*corev1.Service)
				if !ok {
					return errors.New("could not cast old object to corev1.Service")
				}
			}

			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.ensurer.EnsureKubeAPIServerService(ctx, gctx, x, oldSvc)
		}
	case *appsv1.Deployment:
		var oldDep *appsv1.Deployment
		if old != nil {
			var ok bool
			oldDep, ok = old.(*appsv1.Deployment)
			if !ok {
				return errors.New("could not cast old object to appsv1.Deployment")
			}
		}

		switch x.Name {
		case KonnectivityDeploymentName:
			// konnektivity support will be dropped in the near future, so this hack / mutator hook should be removed
			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return fixKonnektivityHostPort(&x.Spec.Template.Spec)
		case v1beta1constants.DeploymentNameKubeAPIServer:
			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.ensurer.EnsureKubeAPIServerDeployment(ctx, gctx, x, oldDep)
		case v1beta1constants.DeploymentNameKubeControllerManager:
			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.ensurer.EnsureKubeControllerManagerDeployment(ctx, gctx, x, oldDep)
		case v1beta1constants.DeploymentNameKubeScheduler:
			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.ensurer.EnsureKubeSchedulerDeployment(ctx, gctx, x, oldDep)
		}
	case *druidv1alpha1.Etcd:
		switch x.Name {
		case v1beta1constants.ETCDMain, v1beta1constants.ETCDEvents:
			var oldEtcd *druidv1alpha1.Etcd
			if old != nil {
				var ok bool
				oldEtcd, ok = old.(*druidv1alpha1.Etcd)
				if !ok {
					return errors.New("could not cast old object to druidv1alpha1.Etcd")
				}
			}

			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.ensurer.EnsureETCD(ctx, gctx, x, oldEtcd)
		}
	case *extensionsv1alpha1.OperatingSystemConfig:
		if x.Spec.Purpose == extensionsv1alpha1.OperatingSystemConfigPurposeReconcile {
			var oldOSC *extensionsv1alpha1.OperatingSystemConfig
			if old != nil {
				var ok bool
				oldOSC, ok = old.(*extensionsv1alpha1.OperatingSystemConfig)
				if !ok {
					return errors.New("could not cast old object to extensionsv1alpha1.OperatingSystemConfig")
				}
			}

			extensionswebhook.LogMutation(m.logger, x.Kind, x.Namespace, x.Name)
			return m.mutateOperatingSystemConfig(ctx, gctx, x, oldOSC)
		}
		return nil
	}
	return nil
}

func getKubeletService(osc *extensionsv1alpha1.OperatingSystemConfig) *string {
	if osc != nil {
		if u := extensionswebhook.UnitWithName(osc.Spec.Units, v1beta1constants.OperatingSystemConfigUnitNameKubeletService); u != nil {
			return u.Content
		}
	}

	return nil
}

func getKubeletConfigFile(osc *extensionsv1alpha1.OperatingSystemConfig) *extensionsv1alpha1.FileContentInline {
	return findFileWithPath(osc, v1beta1constants.OperatingSystemConfigFilePathKubeletConfig)
}

func getKubernetesGeneralConfiguration(osc *extensionsv1alpha1.OperatingSystemConfig) *extensionsv1alpha1.FileContentInline {
	return findFileWithPath(osc, v1beta1constants.OperatingSystemConfigFilePathKernelSettings)
}

func findFileWithPath(osc *extensionsv1alpha1.OperatingSystemConfig, path string) *extensionsv1alpha1.FileContentInline {
	if osc != nil {
		if f := extensionswebhook.FileWithPath(osc.Spec.Files, path); f != nil {
			return f.Content.Inline
		}
	}

	return nil
}

func (m *mutator) mutateOperatingSystemConfig(ctx context.Context, gctx gcontext.GardenContext, osc, oldOSC *extensionsv1alpha1.OperatingSystemConfig) error {
	// Mutate kubelet.service unit, if present
	if content := getKubeletService(osc); content != nil {
		if err := m.ensureKubeletServiceUnitContent(ctx, gctx, content, getKubeletService(oldOSC)); err != nil {
			return err
		}
	}

	// Mutate kubelet configuration file, if present
	if content := getKubeletConfigFile(osc); content != nil {
		if err := m.ensureKubeletConfigFileContent(ctx, gctx, content, getKubeletConfigFile(oldOSC)); err != nil {
			return err
		}
	}

	// Mutate 99 kubernetes general configuration file, if present
	if content := getKubernetesGeneralConfiguration(osc); content != nil {
		if err := m.ensureKubernetesGeneralConfiguration(ctx, gctx, content, getKubernetesGeneralConfiguration(oldOSC)); err != nil {
			return err
		}
	}

	// Check if cloud provider config needs to be ensured
	if m.ensurer.ShouldProvisionKubeletCloudProviderConfig(ctx, gctx) {
		if err := m.ensureKubeletCloudProviderConfig(ctx, gctx, osc); err != nil {
			return err
		}
	}

	var (
		oldFiles *[]extensionsv1alpha1.File
		oldUnits *[]extensionsv1alpha1.Unit
	)

	if oldOSC != nil {
		oldFiles = &oldOSC.Spec.Files
		oldUnits = &oldOSC.Spec.Units
	}

	if err := m.ensurer.EnsureAdditionalFiles(ctx, gctx, &osc.Spec.Files, oldFiles); err != nil {
		return err
	}

	if err := m.ensurer.EnsureAdditionalUnits(ctx, gctx, &osc.Spec.Units, oldUnits); err != nil {
		return err
	}

	return nil
}

func (m *mutator) ensureKubeletServiceUnitContent(ctx context.Context, gctx gcontext.GardenContext, content, oldContent *string) error {
	var (
		opts, oldOpts []*unit.UnitOption
		err           error
	)

	// Deserialize unit options
	if opts, err = m.unitSerializer.Deserialize(*content); err != nil {
		return fmt.Errorf("could not deserialize kubelet.service unit content %w", err)
	}

	if oldContent != nil {
		// Deserialize old unit options
		if oldOpts, err = m.unitSerializer.Deserialize(*oldContent); err != nil {
			return fmt.Errorf("could not deserialize old kubelet.service unit content %w", err)
		}
	}

	if opts, err = m.ensurer.EnsureKubeletServiceUnitOptions(ctx, gctx, opts, oldOpts); err != nil {
		return err
	}

	// Serialize unit options
	if *content, err = m.unitSerializer.Serialize(opts); err != nil {
		return fmt.Errorf("could not serialize kubelet.service unit options %w", err)
	}

	return nil
}

func (m *mutator) ensureKubeletConfigFileContent(ctx context.Context, gctx gcontext.GardenContext, fci, oldFCI *extensionsv1alpha1.FileContentInline) error {
	var (
		kubeletConfig, oldKubeletConfig *kubeletconfigv1beta1.KubeletConfiguration
		err                             error
	)

	// Decode kubelet configuration from inline content
	if kubeletConfig, err = m.kubeletConfigCodec.Decode(fci); err != nil {
		return fmt.Errorf("could not decode kubelet configuration %w", err)
	}

	if oldFCI != nil {
		// Decode old kubelet configuration from inline content
		if oldKubeletConfig, err = m.kubeletConfigCodec.Decode(oldFCI); err != nil {
			return fmt.Errorf("could not decode old kubelet configuration %w", err)
		}
	}

	if err = m.ensurer.EnsureKubeletConfiguration(ctx, gctx, kubeletConfig, oldKubeletConfig); err != nil {
		return err
	}

	// Encode kubelet configuration into inline content
	var newFCI *extensionsv1alpha1.FileContentInline
	if newFCI, err = m.kubeletConfigCodec.Encode(kubeletConfig, fci.Encoding); err != nil {
		return fmt.Errorf("could not encode kubelet configuration %w", err)
	}
	*fci = *newFCI

	return nil
}

func (m *mutator) ensureKubernetesGeneralConfiguration(ctx context.Context, gctx gcontext.GardenContext, fci, oldFCI *extensionsv1alpha1.FileContentInline) error {
	var (
		data, oldData []byte
		err           error
	)

	// Decode kubernetes general configuration from inline content
	if data, err = m.fciCodec.Decode(fci); err != nil {
		return fmt.Errorf("could not decode kubernetes general configuration %w", err)
	}

	if oldFCI != nil {
		// Decode kubernetes general configuration from inline content
		if oldData, err = m.fciCodec.Decode(oldFCI); err != nil {
			return fmt.Errorf("could not decode old kubernetes general configuration %w", err)
		}
	}

	s := string(data)
	oldS := string(oldData)
	if err = m.ensurer.EnsureKubernetesGeneralConfiguration(ctx, gctx, &s, &oldS); err != nil {
		return err
	}

	// Encode kubernetes general configuration into inline content
	var newFCI *extensionsv1alpha1.FileContentInline
	if newFCI, err = m.fciCodec.Encode([]byte(s), fci.Encoding); err != nil {
		return fmt.Errorf("could not encode kubernetes general configuration %w", err)
	}
	*fci = *newFCI

	return nil
}

const CloudProviderConfigPath = "/var/lib/kubelet/cloudprovider.conf"

func (m *mutator) ensureKubeletCloudProviderConfig(ctx context.Context, gctx gcontext.GardenContext, osc *extensionsv1alpha1.OperatingSystemConfig) error {
	var err error

	// Ensure kubelet cloud provider config
	var s string
	if err = m.ensurer.EnsureKubeletCloudProviderConfig(ctx, gctx, &s, osc.Namespace); err != nil {
		return err
	}

	// Encode cloud provider config into inline content
	var fci *extensionsv1alpha1.FileContentInline
	if fci, err = m.fciCodec.Encode([]byte(s), string(cloudinit.B64FileCodecID)); err != nil {
		return fmt.Errorf("could not encode kubelet cloud provider config %w", err)
	}

	// Ensure the cloud provider config file is part of the OperatingSystemConfig
	osc.Spec.Files = extensionswebhook.EnsureFileWithPath(osc.Spec.Files, extensionsv1alpha1.File{
		Path:        CloudProviderConfigPath,
		Permissions: pointer.Int32Ptr(0644),
		Content: extensionsv1alpha1.FileContent{
			Inline: fci,
		},
	})
	return nil
}

// fixKonnektivityHostPort fixes a Gardener bug introduced in v1.16 where host port is preventing multiple
// API servers in a seed to be scheduled because host ports can only be taken once
// TODO: Remove when a fix is available from Gardener upstream
func fixKonnektivityHostPort(ps *corev1.PodSpec) error {
	var containers []corev1.Container
	for _, c := range ps.Containers {
		if c.Name != KonnectivityContainerName {
			containers = append(containers, c)
			continue
		}

		var ports []corev1.ContainerPort
		for _, p := range c.Ports {
			p := p

			if p.Name == "server" || p.Name == "agent" || p.Name == "admin" || p.Name == "health" {
				p = corev1.ContainerPort{
					Name:          p.Name,
					Protocol:      p.Protocol,
					ContainerPort: p.ContainerPort,
				}
			}

			ports = append(ports, p)
		}

		c.Ports = ports
		c.LivenessProbe.HTTPGet.Host = ""

		containers = append(containers, c)
	}

	ps.Containers = containers

	return nil
}
