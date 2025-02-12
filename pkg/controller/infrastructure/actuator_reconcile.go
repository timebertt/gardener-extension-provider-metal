package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/metal-stack/metal-lib/pkg/tag"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/google/uuid"
	metalclient "github.com/metal-stack/gardener-extension-provider-metal/pkg/metal/client"

	metalapi "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/metal"
	"github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/metal/helper"
	metalgo "github.com/metal-stack/metal-go"
	metalfirewall "github.com/metal-stack/metal-go/api/client/firewall"
	mn "github.com/metal-stack/metal-lib/pkg/net"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	controllererrors "github.com/gardener/gardener/extensions/pkg/controller/error"

	v1alpha1constants "github.com/gardener/gardener/pkg/apis/core/v1alpha1/constants"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/utils/secrets"

	"github.com/coreos/container-linux-config-transpiler/config/types"
)

type firewallReconciler struct {
	logger               logr.Logger
	restConfig           *rest.Config
	c                    client.Client
	clientset            kubernetes.Interface
	gc                   gardenerkubernetes.Interface
	infrastructure       *extensionsv1alpha1.Infrastructure
	infrastructureConfig *metalapi.InfrastructureConfig
	providerStatus       *metalapi.InfrastructureStatus
	cluster              *extensionscontroller.Cluster
	mclient              *metalgo.Driver
	clusterID            string
	clusterTag           string
	machineID            string
}

type egressIPReconciler struct {
	logger               logr.Logger
	infrastructureConfig *metalapi.InfrastructureConfig
	mclient              *metalgo.Driver
	clusterID            string
	egressTag            string
}

type firewallReconcileAction string

const (
	firewallControllerName = "firewall-controller"
	droptailerClientName   = "droptailer"
)

var (
	firewallActionRecreate               firewallReconcileAction = "recreate"
	firewallActionDeleteAndRecreate      firewallReconcileAction = "delete"
	firewallActionDoNothing              firewallReconcileAction = "nothing"
	firewallActionUpdateCreationProgress firewallReconcileAction = "update-creation-progress"
	firewallActionCreate                 firewallReconcileAction = "create"
)

func (a *actuator) Reconcile(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	internalInfrastructureConfig, internalInfrastructureStatus, err := decodeInfrastructure(infrastructure, a.decoder)
	if err != nil {
		return err
	}

	cloudProfileConfig, err := helper.CloudProfileConfigFromCluster(cluster)
	if err != nil {
		return err
	}

	metalControlPlane, _, err := helper.FindMetalControlPlane(cloudProfileConfig, internalInfrastructureConfig.PartitionID)
	if err != nil {
		return err
	}

	mclient, err := metalclient.NewClient(ctx, a.client, metalControlPlane.Endpoint, &infrastructure.Spec.SecretRef)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(a.restConfig)
	if err != nil {
		return errors.Wrap(err, "could not create kubernetes clientset")
	}

	gardenerClientset, err := gardenerkubernetes.NewWithConfig(gardenerkubernetes.WithRESTConfig(a.restConfig))
	if err != nil {
		return errors.Wrap(err, "could not create gardener clientset")
	}

	egressIPReconciler := &egressIPReconciler{
		logger:               a.logger,
		infrastructureConfig: internalInfrastructureConfig,
		mclient:              mclient,
		clusterID:            string(cluster.Shoot.GetUID()),
		egressTag:            egressTag(string(cluster.Shoot.GetUID())),
	}
	err = reconcileEgressIPs(ctx, egressIPReconciler)
	if err != nil {
		return err
	}

	reconciler := &firewallReconciler{
		logger:               a.logger,
		restConfig:           a.restConfig,
		c:                    a.client,
		clientset:            clientset,
		gc:                   gardenerClientset,
		infrastructure:       infrastructure,
		infrastructureConfig: internalInfrastructureConfig,
		providerStatus:       internalInfrastructureStatus,
		cluster:              cluster,
		mclient:              mclient,
		clusterID:            string(cluster.Shoot.GetUID()),
		clusterTag:           ClusterTag(string(cluster.Shoot.GetUID())),
		machineID:            decodeMachineID(internalInfrastructureStatus.Firewall.MachineID),
	}

	err = reconcileFirewall(ctx, reconciler)
	if err != nil {
		return err
	}

	return nil
}

func reconcileFirewall(ctx context.Context, r *firewallReconciler) error {
	action, err := firewallNextAction(ctx, r)
	if err != nil {
		return err
	}

	switch action {
	case firewallActionDoNothing:
		r.logger.Info("firewall reconciled, nothing to be done", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.providerStatus.Firewall.MachineID)
		return nil
	case firewallActionCreate:
		err := createFirewall(ctx, r)
		if err != nil {
			return err
		}
		r.logger.Info("firewall created", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.providerStatus.Firewall.MachineID)
		return nil
	case firewallActionUpdateCreationProgress:
		succeeded, err := hasFirewallSucceeded(r.machineID, r.mclient)
		if err != nil {
			return err
		}

		r.logger.Info("firewall creation in progress", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "succeeded", succeeded)
		r.providerStatus.Firewall.Succeeded = succeeded
		return updateProviderStatus(ctx, r.c, r.infrastructure, r.providerStatus, r.infrastructure.Status.NodesCIDR)
	case firewallActionRecreate:
		err := deleteFirewallFromStatus(ctx, r)
		if err != nil {
			return err
		}
		r.logger.Info("firewall removed from status", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.machineID)
		err = createFirewall(ctx, r)
		if err != nil {
			return err
		}
		r.logger.Info("firewall created", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.providerStatus.Firewall.MachineID)
		return nil
	case firewallActionDeleteAndRecreate:
		err := deleteFirewall(r.logger, r.machineID, r.infrastructureConfig.ProjectID, r.clusterTag, r.mclient)
		if err != nil {
			return err
		}
		r.logger.Info("firewall deleted", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.machineID)
		err = deleteFirewallFromStatus(ctx, r)
		if err != nil {
			return err
		}
		r.logger.Info("firewall removed from status", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.machineID)
		err = createFirewall(ctx, r)
		if err != nil {
			return err
		}
		r.logger.Info("firewall created", "cluster-id", r.clusterID, "cluster", r.cluster.Shoot.Name, "machine-id", r.providerStatus.Firewall.MachineID)
		return nil
	default:
		return fmt.Errorf("unsupported firewall reconcile action: %s", action)
	}
}

func firewallNextAction(ctx context.Context, r *firewallReconciler) (firewallReconcileAction, error) {
	if r.providerStatus.Firewall.Succeeded {
		firewalls, err := metalclient.FindClusterFirewalls(r.mclient, r.clusterTag, r.infrastructureConfig.ProjectID)
		if err != nil {
			r.logger.Error(err, "firewalls not found", "clustertag", r.clusterTag, "projectid", r.infrastructureConfig.ProjectID)
			return firewallActionDoNothing, &controllererrors.RequeueAfterError{
				Cause:        err,
				RequeueAfter: 30 * time.Second,
			}
		}

		clusterFirewallAmount := len(firewalls)
		switch clusterFirewallAmount {
		case 0:
			r.logger.Info("firewall does not exist anymore, recreation required", "clusterid", r.clusterID, "machineid", r.machineID)
			return firewallActionRecreate, nil
		case 1:
			fw := firewalls[0]
			if *fw.ID != r.machineID {
				r.logger.Error(
					fmt.Errorf("machine id of this cluster's firewall differs from infrastructure status"),
					"leaving as it is, but something unexpected must have happened in the past. if you want to get to a clean state, remove the firewall by hand (causes downtime!) and reconcile infrastructure again",
					"clusterID", r.clusterID,
					"expectedMachineID", r.machineID,
					"actualMachineID", *fw.ID,
				)
				return firewallActionDoNothing, nil
			}

			if fw.Size.ID != nil && *fw.Size.ID != r.infrastructureConfig.Firewall.Size {
				r.logger.Info("firewall size has changed", "clusterid", r.clusterID, "machineid", r.machineID, "current", *fw.Size.ID, "new", r.infrastructureConfig.Firewall.Size)
				return firewallActionDeleteAndRecreate, nil
			}

			if fw.Allocation != nil && fw.Allocation.Image != nil && fw.Allocation.Image.ID != nil && *fw.Allocation.Image.ID != r.infrastructureConfig.Firewall.Image {
				want := r.infrastructureConfig.Firewall.Image
				image, err := r.mclient.ImageGetLatest(want)
				if err != nil {
					r.logger.Error(err, "firewall latest image not found", "clustertag", r.clusterTag, "projectid", r.infrastructureConfig.ProjectID, "image", want)
					return firewallActionDoNothing, &controllererrors.RequeueAfterError{
						Cause:        err,
						RequeueAfter: 30 * time.Second,
					}
				}

				if image.Image != nil && image.Image.ID != nil && *image.Image.ID != *fw.Allocation.Image.ID {
					r.logger.Info("firewall image has changed", "clusterid", r.clusterID, "machineid", r.machineID, "current", *fw.Allocation.Image.ID, "new", *image.Image.ID)
					return firewallActionDeleteAndRecreate, nil
				}
			}

			currentNetworks := sets.NewString()
			for _, n := range fw.Allocation.Networks {
				if *n.Networktype == mn.PrivatePrimaryUnshared || *n.Networktype == mn.PrivatePrimaryShared {
					continue
				}
				if *n.Underlay {
					continue
				}
				currentNetworks.Insert(*n.Networkid)
			}
			wantNetworks := sets.NewString()
			for _, n := range r.infrastructureConfig.Firewall.Networks {
				wantNetworks.Insert(n)
			}
			if !currentNetworks.Equal(wantNetworks) {
				r.logger.Info("firewall networks have changed", "clusterid", r.clusterID, "machineid", r.machineID, "current", currentNetworks.List(), "new", wantNetworks.List())
				return firewallActionDeleteAndRecreate, nil
			}

			return firewallActionDoNothing, nil
		default:
			err := fmt.Errorf("multiple firewalls exist for this cluster, which is currently unsupported. delete these firewalls and try to keep the one with machine id %q", r.machineID)
			r.logger.Error(
				err,
				"clusterID", r.clusterID,
				"expectedMachineID", r.machineID,
			)
			return firewallActionDoNothing, err
		}
	}

	firewallInProgress := r.machineID != "" && !r.providerStatus.Firewall.Succeeded
	if firewallInProgress {
		return firewallActionUpdateCreationProgress, nil
	}

	return firewallActionCreate, nil
}

func hasFirewallSucceeded(machineID string, mclient *metalgo.Driver) (bool, error) {
	resp, err := mclient.FirewallGet(machineID)
	if err != nil {
		switch e := err.(type) {
		case *metalfirewall.FindFirewallDefault:
			if e.Code() >= 500 {
				return false, &controllererrors.RequeueAfterError{
					Cause:        e,
					RequeueAfter: 5 * time.Second,
				}
			}
		default:
			return false, e
		}
	}

	if resp.Firewall == nil || resp.Firewall.Allocation == nil || resp.Firewall.Allocation.Succeeded == nil {
		return false, fmt.Errorf("firewall %q was created but has no allocation", machineID)
	}

	return *resp.Firewall.Allocation.Succeeded, nil
}

func createFirewall(ctx context.Context, r *firewallReconciler) error {
	nodeCIDR, err := ensureNodeNetwork(ctx, r)
	if err != nil {
		r.logger.Error(err, "firewalls node network", "nodecidr", nodeCIDR)
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	err = updateProviderStatus(ctx, r.c, r.infrastructure, r.providerStatus, &nodeCIDR)
	if err != nil {
		r.logger.Error(err, "firewalls update provider status", "nodecidr", nodeCIDR)
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	clusterName := r.cluster.ObjectMeta.Name
	name := clusterName + "-firewall-" + uuid.String()[:5]

	// find private network
	privateNetwork, err := metalclient.GetPrivateNetworkFromNodeNetwork(r.mclient, r.infrastructureConfig.ProjectID, nodeCIDR)
	if err != nil {
		return err
	}

	kubeconfig, err := createFirewallControllerKubeconfig(ctx, r)
	if err != nil {
		return err
	}

	firewallUserData, err := renderFirewallUserData(kubeconfig)
	if err != nil {
		return err
	}

	// assemble firewall allocation request
	var networks []metalgo.MachineAllocationNetwork
	network := metalgo.MachineAllocationNetwork{
		NetworkID:   *privateNetwork.ID,
		Autoacquire: true,
	}
	networks = append(networks, network)
	for _, n := range r.infrastructureConfig.Firewall.Networks {
		network := metalgo.MachineAllocationNetwork{
			NetworkID:   n,
			Autoacquire: true,
		}
		networks = append(networks, network)
	}

	createRequest := &metalgo.FirewallCreateRequest{
		MachineCreateRequest: metalgo.MachineCreateRequest{
			Description:   name + " created by Gardener",
			Name:          name,
			Hostname:      name,
			Size:          r.infrastructureConfig.Firewall.Size,
			Project:       r.infrastructureConfig.ProjectID,
			Partition:     r.infrastructureConfig.PartitionID,
			Image:         r.infrastructureConfig.Firewall.Image,
			SSHPublicKeys: []string{string(r.infrastructure.Spec.SSHPublicKey)},
			Networks:      networks,
			UserData:      firewallUserData,
			Tags:          []string{r.clusterTag},
		},
	}

	fcr, err := r.mclient.FirewallCreate(createRequest)
	if err != nil {
		r.logger.Error(err, "firewall create error")
		return &controllererrors.RequeueAfterError{
			Cause:        errors.Wrap(err, "failed to create firewall"),
			RequeueAfter: 30 * time.Second,
		}
	}

	machineID := encodeMachineID(*fcr.Firewall.Partition.ID, *fcr.Firewall.ID)

	allocation := fcr.Firewall.Allocation
	if allocation == nil {
		return fmt.Errorf("firewall %q was created but has no allocation", machineID)
	}

	r.providerStatus.Firewall.MachineID = machineID
	r.providerStatus.Firewall.Succeeded = true

	return updateProviderStatus(ctx, r.c, r.infrastructure, r.providerStatus, &nodeCIDR)
}

func reconcileEgressIPs(ctx context.Context, r *egressIPReconciler) error {
	static := metalgo.IPTypeStatic
	currentEgressIPs := sets.NewString()
	resp, err := r.mclient.IPFind(&metalgo.IPFindRequest{
		ProjectID: &r.infrastructureConfig.ProjectID,
		Tags:      []string{r.egressTag},
		Type:      &static,
	})
	if err != nil {
		return &controllererrors.RequeueAfterError{
			Cause:        errors.Wrap(err, "failed to list egress ips of cluster"),
			RequeueAfter: 30 * time.Second,
		}
	}
	for _, ip := range resp.IPs {
		currentEgressIPs.Insert(*ip.Ipaddress)
	}

	wantEgressIPs := sets.NewString()
	for _, egressRule := range r.infrastructureConfig.Firewall.EgressRules {
		wantEgressIPs.Insert(egressRule.IPs...)

		for _, ip := range egressRule.IPs {
			if currentEgressIPs.Has(ip) {
				continue
			}

			resp, err := r.mclient.IPFind(&metalgo.IPFindRequest{
				IPAddress: &ip,
				ProjectID: &r.infrastructureConfig.ProjectID,
				NetworkID: &egressRule.NetworkID,
			})

			if err != nil {
				return &controllererrors.RequeueAfterError{
					Cause:        errors.Wrap(err, fmt.Sprintf("error when retrieving ip %s for egress rule", ip)),
					RequeueAfter: 30 * time.Second,
				}
			}

			switch len(resp.IPs) {
			case 0:
				return &controllererrors.RequeueAfterError{
					Cause:        fmt.Errorf("ip %s for egress rule does not exist", ip),
					RequeueAfter: 30 * time.Second,
				}
			case 1:
			default:
				return fmt.Errorf("ip %s found multiple times", ip)
			}

			dbIP := resp.IPs[0]
			if dbIP.Type != nil && *dbIP.Type != metalgo.IPTypeStatic {
				return &controllererrors.RequeueAfterError{
					Cause:        fmt.Errorf("ips for egress rule must be static, but %s is not static", ip),
					RequeueAfter: 30 * time.Second,
				}
			}

			if len(dbIP.Tags) > 0 {
				return &controllererrors.RequeueAfterError{
					Cause:        fmt.Errorf("won't use ip %s for egress rules because it does not have an egress tag but it has other tags", *dbIP.Ipaddress),
					RequeueAfter: 30 * time.Second,
				}
			}

			iur := metalgo.IPUpdateRequest{
				IPAddress: *dbIP.Ipaddress,
				Tags:      []string{r.egressTag},
			}

			_, err = r.mclient.IPUpdate(&iur)
			if err != nil {
				return &controllererrors.RequeueAfterError{
					Cause:        errors.Wrap(err, fmt.Sprintf("could not tag ip %s for egress usage", ip)),
					RequeueAfter: 30 * time.Second,
				}
			}
		}
	}

	if !currentEgressIPs.Equal(wantEgressIPs) {
		toUnTag := currentEgressIPs.Difference(wantEgressIPs)
		for _, ip := range toUnTag.List() {
			err := clearIPTags(r.mclient, ip)
			if err != nil {
				return &controllererrors.RequeueAfterError{
					Cause:        errors.Wrap(err, fmt.Sprintf("could not remove egress tag from ip %s", ip)),
					RequeueAfter: 30 * time.Second,
				}
			}
		}
	}

	return nil
}

func egressTag(clusterID string) string {
	return fmt.Sprintf("%s=%s", tag.ClusterEgress, clusterID)
}

func clearIPTags(mclient *metalgo.Driver, ip string) error {
	iur := metalgo.IPUpdateRequest{
		IPAddress: ip,
		Tags:      []string{},
	}
	_, err := mclient.IPUpdate(&iur)
	return err
}

func ensureNodeNetwork(ctx context.Context, r *firewallReconciler) (string, error) {
	if r.cluster.Shoot.Spec.Networking.Nodes != nil {
		return *r.cluster.Shoot.Spec.Networking.Nodes, nil
	}
	if r.infrastructure.Status.NodesCIDR != nil {
		resp, err := r.mclient.NetworkFind(&metalgo.NetworkFindRequest{
			ProjectID:   &r.infrastructureConfig.ProjectID,
			PartitionID: &r.infrastructureConfig.PartitionID,
			Labels:      map[string]string{tag.ClusterID: r.clusterID},
		})
		if err != nil {
			return "", err
		}

		if len(resp.Networks) != 0 {
			return *r.infrastructure.Status.NodesCIDR, nil
		}

		return "", fmt.Errorf("node network disappeared from cloud provider: %s", *r.infrastructure.Status.NodesCIDR)
	}

	resp, err := r.mclient.NetworkAllocate(&metalgo.NetworkAllocateRequest{
		ProjectID:   r.infrastructureConfig.ProjectID,
		PartitionID: r.infrastructureConfig.PartitionID,
		Name:        r.cluster.Shoot.GetName(),
		Description: r.clusterID,
		Labels:      map[string]string{tag.ClusterID: r.clusterID},
	})
	if err != nil {
		return "", err
	}

	nodeCIDR := resp.Network.Prefixes[0]

	return nodeCIDR, nil
}

func createFirewallControllerKubeconfig(ctx context.Context, r *firewallReconciler) (string, error) {
	apiServerURL := fmt.Sprintf("api.%s", *r.cluster.Shoot.Spec.DNS.Domain)
	infrastructureSecrets := &secrets.Secrets{
		CertificateSecretConfigs: map[string]*secrets.CertificateSecretConfig{
			v1alpha1constants.SecretNameCACluster: {
				Name:       v1alpha1constants.SecretNameCACluster,
				CommonName: "kubernetes",
				CertType:   secrets.CACert,
			},
		},
		SecretConfigsFunc: func(cas map[string]*secrets.Certificate, clusterName string) []secrets.ConfigInterface {
			return []secrets.ConfigInterface{
				&secrets.ControlPlaneSecretConfig{
					CertificateSecretConfig: &secrets.CertificateSecretConfig{
						Name:         firewallControllerName,
						CommonName:   fmt.Sprintf("system:%s", firewallControllerName),
						Organization: []string{firewallControllerName},
						CertType:     secrets.ClientCert,
						SigningCA:    cas[v1alpha1constants.SecretNameCACluster],
					},
					KubeConfigRequests: []secrets.KubeConfigRequest{
						{
							ClusterName:   clusterName,
							APIServerHost: apiServerURL,
						},
					},
				},
			}
		},
	}

	secret, err := infrastructureSecrets.Deploy(ctx, r.clientset, r.gc, r.infrastructure.Namespace)
	if err != nil {
		return "", err
	}

	kubeconfig, ok := secret[firewallControllerName].Data[secrets.DataKeyKubeconfig]
	if !ok {
		return "", fmt.Errorf("kubeconfig not part of generated firewall controller secret")
	}

	return string(kubeconfig), nil
}

func renderFirewallUserData(kubeconfig string) (string, error) {
	cfg := types.Config{}
	cfg.Systemd = types.Systemd{}

	enabled := true
	fcUnit := types.SystemdUnit{
		Name:    fmt.Sprintf("%s.service", firewallControllerName),
		Enable:  enabled,
		Enabled: &enabled,
	}
	dcUnit := types.SystemdUnit{
		Name:    fmt.Sprintf("%s.service", droptailerClientName),
		Enable:  enabled,
		Enabled: &enabled,
	}

	cfg.Systemd.Units = append(cfg.Systemd.Units, fcUnit, dcUnit)

	cfg.Storage = types.Storage{}

	mode := 0600
	id := 0
	ignitionFile := types.File{
		Path:       "/etc/firewall-controller/.kubeconfig",
		Filesystem: "root",
		Mode:       &mode,
		User: &types.FileUser{
			Id: &id,
		},
		Group: &types.FileGroup{
			Id: &id,
		},
		Contents: types.FileContents{
			Inline: string(kubeconfig),
		},
	}
	cfg.Storage.Files = append(cfg.Storage.Files, ignitionFile)

	outCfg, report := types.Convert(cfg, "", nil)
	if report.IsFatal() {
		return "", fmt.Errorf("could not transpile ignition config: %s", report.String())
	}

	userData, err := json.Marshal(outCfg)
	if err != nil {
		return "", err
	}

	return string(userData), nil
}

func deleteFirewallFromStatus(ctx context.Context, r *firewallReconciler) error {
	r.providerStatus.Firewall.MachineID = ""
	r.providerStatus.Firewall.Succeeded = false
	err := updateProviderStatus(ctx, r.c, r.infrastructure, r.providerStatus, r.infrastructure.Status.NodesCIDR)
	if err != nil {
		return err
	}
	return nil
}

func ClusterTag(clusterID string) string {
	return fmt.Sprintf("%s=%s", tag.ClusterID, clusterID)
}
