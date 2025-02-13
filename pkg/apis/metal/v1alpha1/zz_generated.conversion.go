//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/config"
	metal "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/metal"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfig)(nil), (*metal.CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfig_To_metal_CloudControllerManagerConfig(a.(*CloudControllerManagerConfig), b.(*metal.CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.CloudControllerManagerConfig)(nil), (*CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(a.(*metal.CloudControllerManagerConfig), b.(*CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudProfileConfig)(nil), (*metal.CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudProfileConfig_To_metal_CloudProfileConfig(a.(*CloudProfileConfig), b.(*metal.CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.CloudProfileConfig)(nil), (*CloudProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(a.(*metal.CloudProfileConfig), b.(*CloudProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConnectorConfig)(nil), (*metal.ConnectorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConnectorConfig_To_metal_ConnectorConfig(a.(*ConnectorConfig), b.(*metal.ConnectorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.ConnectorConfig)(nil), (*ConnectorConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_ConnectorConfig_To_v1alpha1_ConnectorConfig(a.(*metal.ConnectorConfig), b.(*ConnectorConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneConfig)(nil), (*metal.ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneConfig_To_metal_ControlPlaneConfig(a.(*ControlPlaneConfig), b.(*metal.ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.ControlPlaneConfig)(nil), (*ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(a.(*metal.ControlPlaneConfig), b.(*ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneFeatures)(nil), (*metal.ControlPlaneFeatures)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures(a.(*ControlPlaneFeatures), b.(*metal.ControlPlaneFeatures), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.ControlPlaneFeatures)(nil), (*ControlPlaneFeatures)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures(a.(*metal.ControlPlaneFeatures), b.(*ControlPlaneFeatures), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressRule)(nil), (*metal.EgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EgressRule_To_metal_EgressRule(a.(*EgressRule), b.(*metal.EgressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.EgressRule)(nil), (*EgressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_EgressRule_To_v1alpha1_EgressRule(a.(*metal.EgressRule), b.(*EgressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Firewall)(nil), (*metal.Firewall)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Firewall_To_metal_Firewall(a.(*Firewall), b.(*metal.Firewall), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.Firewall)(nil), (*Firewall)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_Firewall_To_v1alpha1_Firewall(a.(*metal.Firewall), b.(*Firewall), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FirewallControllerVersion)(nil), (*metal.FirewallControllerVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FirewallControllerVersion_To_metal_FirewallControllerVersion(a.(*FirewallControllerVersion), b.(*metal.FirewallControllerVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.FirewallControllerVersion)(nil), (*FirewallControllerVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_FirewallControllerVersion_To_v1alpha1_FirewallControllerVersion(a.(*metal.FirewallControllerVersion), b.(*FirewallControllerVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FirewallStatus)(nil), (*metal.FirewallStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FirewallStatus_To_metal_FirewallStatus(a.(*FirewallStatus), b.(*metal.FirewallStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.FirewallStatus)(nil), (*FirewallStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_FirewallStatus_To_v1alpha1_FirewallStatus(a.(*metal.FirewallStatus), b.(*FirewallStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IAMConfig)(nil), (*metal.IAMConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IAMConfig_To_metal_IAMConfig(a.(*IAMConfig), b.(*metal.IAMConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.IAMConfig)(nil), (*IAMConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_IAMConfig_To_v1alpha1_IAMConfig(a.(*metal.IAMConfig), b.(*IAMConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IDMConfig)(nil), (*metal.IDMConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IDMConfig_To_metal_IDMConfig(a.(*IDMConfig), b.(*metal.IDMConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.IDMConfig)(nil), (*IDMConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_IDMConfig_To_v1alpha1_IDMConfig(a.(*metal.IDMConfig), b.(*IDMConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*metal.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_metal_InfrastructureConfig(a.(*InfrastructureConfig), b.(*metal.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*metal.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*metal.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_metal_InfrastructureStatus(a.(*InfrastructureStatus), b.(*metal.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*metal.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IssuerConfig)(nil), (*metal.IssuerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IssuerConfig_To_metal_IssuerConfig(a.(*IssuerConfig), b.(*metal.IssuerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.IssuerConfig)(nil), (*IssuerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_IssuerConfig_To_v1alpha1_IssuerConfig(a.(*metal.IssuerConfig), b.(*IssuerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalControlPlane)(nil), (*metal.MetalControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MetalControlPlane_To_metal_MetalControlPlane(a.(*MetalControlPlane), b.(*metal.MetalControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.MetalControlPlane)(nil), (*MetalControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_MetalControlPlane_To_v1alpha1_MetalControlPlane(a.(*metal.MetalControlPlane), b.(*MetalControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NamespaceGroupConfig)(nil), (*metal.NamespaceGroupConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NamespaceGroupConfig_To_metal_NamespaceGroupConfig(a.(*NamespaceGroupConfig), b.(*metal.NamespaceGroupConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.NamespaceGroupConfig)(nil), (*NamespaceGroupConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_NamespaceGroupConfig_To_v1alpha1_NamespaceGroupConfig(a.(*metal.NamespaceGroupConfig), b.(*NamespaceGroupConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Partition)(nil), (*metal.Partition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Partition_To_metal_Partition(a.(*Partition), b.(*metal.Partition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.Partition)(nil), (*Partition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_Partition_To_v1alpha1_Partition(a.(*metal.Partition), b.(*Partition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RateLimit)(nil), (*metal.RateLimit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RateLimit_To_metal_RateLimit(a.(*RateLimit), b.(*metal.RateLimit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.RateLimit)(nil), (*RateLimit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_RateLimit_To_v1alpha1_RateLimit(a.(*metal.RateLimit), b.(*RateLimit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerStatus)(nil), (*metal.WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerStatus_To_metal_WorkerStatus(a.(*WorkerStatus), b.(*metal.WorkerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metal.WorkerStatus)(nil), (*WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metal_WorkerStatus_To_v1alpha1_WorkerStatus(a.(*metal.WorkerStatus), b.(*WorkerStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudControllerManagerConfig_To_metal_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *metal.CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.DefaultExternalNetwork = (*string)(unsafe.Pointer(in.DefaultExternalNetwork))
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfig_To_metal_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfig_To_metal_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *metal.CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfig_To_metal_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_metal_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *metal.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.DefaultExternalNetwork = (*string)(unsafe.Pointer(in.DefaultExternalNetwork))
	return nil
}

// Convert_metal_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_metal_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *metal.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_metal_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_v1alpha1_CloudProfileConfig_To_metal_CloudProfileConfig(in *CloudProfileConfig, out *metal.CloudProfileConfig, s conversion.Scope) error {
	out.MetalControlPlanes = *(*map[string]metal.MetalControlPlane)(unsafe.Pointer(&in.MetalControlPlanes))
	return nil
}

// Convert_v1alpha1_CloudProfileConfig_To_metal_CloudProfileConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudProfileConfig_To_metal_CloudProfileConfig(in *CloudProfileConfig, out *metal.CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudProfileConfig_To_metal_CloudProfileConfig(in, out, s)
}

func autoConvert_metal_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *metal.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	out.MetalControlPlanes = *(*map[string]MetalControlPlane)(unsafe.Pointer(&in.MetalControlPlanes))
	return nil
}

// Convert_metal_CloudProfileConfig_To_v1alpha1_CloudProfileConfig is an autogenerated conversion function.
func Convert_metal_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in *metal.CloudProfileConfig, out *CloudProfileConfig, s conversion.Scope) error {
	return autoConvert_metal_CloudProfileConfig_To_v1alpha1_CloudProfileConfig(in, out, s)
}

func autoConvert_v1alpha1_ConnectorConfig_To_metal_ConnectorConfig(in *ConnectorConfig, out *metal.ConnectorConfig, s conversion.Scope) error {
	out.IdmApiUrl = in.IdmApiUrl
	out.IdmApiUser = in.IdmApiUser
	out.IdmApiPassword = in.IdmApiPassword
	out.IdmSystemId = in.IdmSystemId
	out.IdmAccessCode = in.IdmAccessCode
	out.IdmCustomerId = in.IdmCustomerId
	out.IdmGroupOU = in.IdmGroupOU
	out.IdmGroupnameTemplate = in.IdmGroupnameTemplate
	out.IdmDomainName = in.IdmDomainName
	out.IdmTenantPrefix = in.IdmTenantPrefix
	out.IdmSubmitter = in.IdmSubmitter
	out.IdmJobInfo = in.IdmJobInfo
	out.IdmReqSystem = in.IdmReqSystem
	out.IdmReqUser = in.IdmReqUser
	out.IdmReqEMail = in.IdmReqEMail
	return nil
}

// Convert_v1alpha1_ConnectorConfig_To_metal_ConnectorConfig is an autogenerated conversion function.
func Convert_v1alpha1_ConnectorConfig_To_metal_ConnectorConfig(in *ConnectorConfig, out *metal.ConnectorConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConnectorConfig_To_metal_ConnectorConfig(in, out, s)
}

func autoConvert_metal_ConnectorConfig_To_v1alpha1_ConnectorConfig(in *metal.ConnectorConfig, out *ConnectorConfig, s conversion.Scope) error {
	out.IdmApiUrl = in.IdmApiUrl
	out.IdmApiUser = in.IdmApiUser
	out.IdmApiPassword = in.IdmApiPassword
	out.IdmSystemId = in.IdmSystemId
	out.IdmAccessCode = in.IdmAccessCode
	out.IdmCustomerId = in.IdmCustomerId
	out.IdmGroupOU = in.IdmGroupOU
	out.IdmGroupnameTemplate = in.IdmGroupnameTemplate
	out.IdmDomainName = in.IdmDomainName
	out.IdmTenantPrefix = in.IdmTenantPrefix
	out.IdmSubmitter = in.IdmSubmitter
	out.IdmJobInfo = in.IdmJobInfo
	out.IdmReqSystem = in.IdmReqSystem
	out.IdmReqUser = in.IdmReqUser
	out.IdmReqEMail = in.IdmReqEMail
	return nil
}

// Convert_metal_ConnectorConfig_To_v1alpha1_ConnectorConfig is an autogenerated conversion function.
func Convert_metal_ConnectorConfig_To_v1alpha1_ConnectorConfig(in *metal.ConnectorConfig, out *ConnectorConfig, s conversion.Scope) error {
	return autoConvert_metal_ConnectorConfig_To_v1alpha1_ConnectorConfig(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneConfig_To_metal_ControlPlaneConfig(in *ControlPlaneConfig, out *metal.ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*metal.CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	out.IAMConfig = (*metal.IAMConfig)(unsafe.Pointer(in.IAMConfig))
	if err := Convert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures(&in.FeatureGates, &out.FeatureGates, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControlPlaneConfig_To_metal_ControlPlaneConfig is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneConfig_To_metal_ControlPlaneConfig(in *ControlPlaneConfig, out *metal.ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneConfig_To_metal_ControlPlaneConfig(in, out, s)
}

func autoConvert_metal_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *metal.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	out.IAMConfig = (*IAMConfig)(unsafe.Pointer(in.IAMConfig))
	if err := Convert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures(&in.FeatureGates, &out.FeatureGates, s); err != nil {
		return err
	}
	return nil
}

// Convert_metal_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig is an autogenerated conversion function.
func Convert_metal_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *metal.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_metal_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures(in *ControlPlaneFeatures, out *metal.ControlPlaneFeatures, s conversion.Scope) error {
	out.MachineControllerManagerOOT = (*bool)(unsafe.Pointer(in.MachineControllerManagerOOT))
	out.ClusterAudit = (*bool)(unsafe.Pointer(in.ClusterAudit))
	out.AuditToSplunk = (*bool)(unsafe.Pointer(in.AuditToSplunk))
	return nil
}

// Convert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures(in *ControlPlaneFeatures, out *metal.ControlPlaneFeatures, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneFeatures_To_metal_ControlPlaneFeatures(in, out, s)
}

func autoConvert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures(in *metal.ControlPlaneFeatures, out *ControlPlaneFeatures, s conversion.Scope) error {
	out.MachineControllerManagerOOT = (*bool)(unsafe.Pointer(in.MachineControllerManagerOOT))
	out.ClusterAudit = (*bool)(unsafe.Pointer(in.ClusterAudit))
	out.AuditToSplunk = (*bool)(unsafe.Pointer(in.AuditToSplunk))
	return nil
}

// Convert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures is an autogenerated conversion function.
func Convert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures(in *metal.ControlPlaneFeatures, out *ControlPlaneFeatures, s conversion.Scope) error {
	return autoConvert_metal_ControlPlaneFeatures_To_v1alpha1_ControlPlaneFeatures(in, out, s)
}

func autoConvert_v1alpha1_EgressRule_To_metal_EgressRule(in *EgressRule, out *metal.EgressRule, s conversion.Scope) error {
	out.NetworkID = in.NetworkID
	out.IPs = *(*[]string)(unsafe.Pointer(&in.IPs))
	return nil
}

// Convert_v1alpha1_EgressRule_To_metal_EgressRule is an autogenerated conversion function.
func Convert_v1alpha1_EgressRule_To_metal_EgressRule(in *EgressRule, out *metal.EgressRule, s conversion.Scope) error {
	return autoConvert_v1alpha1_EgressRule_To_metal_EgressRule(in, out, s)
}

func autoConvert_metal_EgressRule_To_v1alpha1_EgressRule(in *metal.EgressRule, out *EgressRule, s conversion.Scope) error {
	out.NetworkID = in.NetworkID
	out.IPs = *(*[]string)(unsafe.Pointer(&in.IPs))
	return nil
}

// Convert_metal_EgressRule_To_v1alpha1_EgressRule is an autogenerated conversion function.
func Convert_metal_EgressRule_To_v1alpha1_EgressRule(in *metal.EgressRule, out *EgressRule, s conversion.Scope) error {
	return autoConvert_metal_EgressRule_To_v1alpha1_EgressRule(in, out, s)
}

func autoConvert_v1alpha1_Firewall_To_metal_Firewall(in *Firewall, out *metal.Firewall, s conversion.Scope) error {
	out.Size = in.Size
	out.Image = in.Image
	out.Networks = *(*[]string)(unsafe.Pointer(&in.Networks))
	out.RateLimits = *(*[]metal.RateLimit)(unsafe.Pointer(&in.RateLimits))
	out.EgressRules = *(*[]metal.EgressRule)(unsafe.Pointer(&in.EgressRules))
	out.ControllerVersion = in.ControllerVersion
	return nil
}

// Convert_v1alpha1_Firewall_To_metal_Firewall is an autogenerated conversion function.
func Convert_v1alpha1_Firewall_To_metal_Firewall(in *Firewall, out *metal.Firewall, s conversion.Scope) error {
	return autoConvert_v1alpha1_Firewall_To_metal_Firewall(in, out, s)
}

func autoConvert_metal_Firewall_To_v1alpha1_Firewall(in *metal.Firewall, out *Firewall, s conversion.Scope) error {
	out.Size = in.Size
	out.Image = in.Image
	out.Networks = *(*[]string)(unsafe.Pointer(&in.Networks))
	out.RateLimits = *(*[]RateLimit)(unsafe.Pointer(&in.RateLimits))
	out.EgressRules = *(*[]EgressRule)(unsafe.Pointer(&in.EgressRules))
	out.ControllerVersion = in.ControllerVersion
	return nil
}

// Convert_metal_Firewall_To_v1alpha1_Firewall is an autogenerated conversion function.
func Convert_metal_Firewall_To_v1alpha1_Firewall(in *metal.Firewall, out *Firewall, s conversion.Scope) error {
	return autoConvert_metal_Firewall_To_v1alpha1_Firewall(in, out, s)
}

func autoConvert_v1alpha1_FirewallControllerVersion_To_metal_FirewallControllerVersion(in *FirewallControllerVersion, out *metal.FirewallControllerVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.URL = in.URL
	out.Classification = (*metal.VersionClassification)(unsafe.Pointer(in.Classification))
	return nil
}

// Convert_v1alpha1_FirewallControllerVersion_To_metal_FirewallControllerVersion is an autogenerated conversion function.
func Convert_v1alpha1_FirewallControllerVersion_To_metal_FirewallControllerVersion(in *FirewallControllerVersion, out *metal.FirewallControllerVersion, s conversion.Scope) error {
	return autoConvert_v1alpha1_FirewallControllerVersion_To_metal_FirewallControllerVersion(in, out, s)
}

func autoConvert_metal_FirewallControllerVersion_To_v1alpha1_FirewallControllerVersion(in *metal.FirewallControllerVersion, out *FirewallControllerVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.URL = in.URL
	out.Classification = (*VersionClassification)(unsafe.Pointer(in.Classification))
	return nil
}

// Convert_metal_FirewallControllerVersion_To_v1alpha1_FirewallControllerVersion is an autogenerated conversion function.
func Convert_metal_FirewallControllerVersion_To_v1alpha1_FirewallControllerVersion(in *metal.FirewallControllerVersion, out *FirewallControllerVersion, s conversion.Scope) error {
	return autoConvert_metal_FirewallControllerVersion_To_v1alpha1_FirewallControllerVersion(in, out, s)
}

func autoConvert_v1alpha1_FirewallStatus_To_metal_FirewallStatus(in *FirewallStatus, out *metal.FirewallStatus, s conversion.Scope) error {
	out.Succeeded = in.Succeeded
	out.MachineID = in.MachineID
	return nil
}

// Convert_v1alpha1_FirewallStatus_To_metal_FirewallStatus is an autogenerated conversion function.
func Convert_v1alpha1_FirewallStatus_To_metal_FirewallStatus(in *FirewallStatus, out *metal.FirewallStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FirewallStatus_To_metal_FirewallStatus(in, out, s)
}

func autoConvert_metal_FirewallStatus_To_v1alpha1_FirewallStatus(in *metal.FirewallStatus, out *FirewallStatus, s conversion.Scope) error {
	out.Succeeded = in.Succeeded
	out.MachineID = in.MachineID
	return nil
}

// Convert_metal_FirewallStatus_To_v1alpha1_FirewallStatus is an autogenerated conversion function.
func Convert_metal_FirewallStatus_To_v1alpha1_FirewallStatus(in *metal.FirewallStatus, out *FirewallStatus, s conversion.Scope) error {
	return autoConvert_metal_FirewallStatus_To_v1alpha1_FirewallStatus(in, out, s)
}

func autoConvert_v1alpha1_IAMConfig_To_metal_IAMConfig(in *IAMConfig, out *metal.IAMConfig, s conversion.Scope) error {
	out.IssuerConfig = (*metal.IssuerConfig)(unsafe.Pointer(in.IssuerConfig))
	out.IdmConfig = (*metal.IDMConfig)(unsafe.Pointer(in.IdmConfig))
	out.GroupConfig = (*metal.NamespaceGroupConfig)(unsafe.Pointer(in.GroupConfig))
	return nil
}

// Convert_v1alpha1_IAMConfig_To_metal_IAMConfig is an autogenerated conversion function.
func Convert_v1alpha1_IAMConfig_To_metal_IAMConfig(in *IAMConfig, out *metal.IAMConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_IAMConfig_To_metal_IAMConfig(in, out, s)
}

func autoConvert_metal_IAMConfig_To_v1alpha1_IAMConfig(in *metal.IAMConfig, out *IAMConfig, s conversion.Scope) error {
	out.IssuerConfig = (*IssuerConfig)(unsafe.Pointer(in.IssuerConfig))
	out.IdmConfig = (*IDMConfig)(unsafe.Pointer(in.IdmConfig))
	out.GroupConfig = (*NamespaceGroupConfig)(unsafe.Pointer(in.GroupConfig))
	return nil
}

// Convert_metal_IAMConfig_To_v1alpha1_IAMConfig is an autogenerated conversion function.
func Convert_metal_IAMConfig_To_v1alpha1_IAMConfig(in *metal.IAMConfig, out *IAMConfig, s conversion.Scope) error {
	return autoConvert_metal_IAMConfig_To_v1alpha1_IAMConfig(in, out, s)
}

func autoConvert_v1alpha1_IDMConfig_To_metal_IDMConfig(in *IDMConfig, out *metal.IDMConfig, s conversion.Scope) error {
	out.Idmtype = in.Idmtype
	out.ConnectorConfig = (*metal.ConnectorConfig)(unsafe.Pointer(in.ConnectorConfig))
	return nil
}

// Convert_v1alpha1_IDMConfig_To_metal_IDMConfig is an autogenerated conversion function.
func Convert_v1alpha1_IDMConfig_To_metal_IDMConfig(in *IDMConfig, out *metal.IDMConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_IDMConfig_To_metal_IDMConfig(in, out, s)
}

func autoConvert_metal_IDMConfig_To_v1alpha1_IDMConfig(in *metal.IDMConfig, out *IDMConfig, s conversion.Scope) error {
	out.Idmtype = in.Idmtype
	out.ConnectorConfig = (*ConnectorConfig)(unsafe.Pointer(in.ConnectorConfig))
	return nil
}

// Convert_metal_IDMConfig_To_v1alpha1_IDMConfig is an autogenerated conversion function.
func Convert_metal_IDMConfig_To_v1alpha1_IDMConfig(in *metal.IDMConfig, out *IDMConfig, s conversion.Scope) error {
	return autoConvert_metal_IDMConfig_To_v1alpha1_IDMConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_metal_InfrastructureConfig(in *InfrastructureConfig, out *metal.InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_Firewall_To_metal_Firewall(&in.Firewall, &out.Firewall, s); err != nil {
		return err
	}
	out.PartitionID = in.PartitionID
	out.ProjectID = in.ProjectID
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_metal_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_metal_InfrastructureConfig(in *InfrastructureConfig, out *metal.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_metal_InfrastructureConfig(in, out, s)
}

func autoConvert_metal_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *metal.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_metal_Firewall_To_v1alpha1_Firewall(&in.Firewall, &out.Firewall, s); err != nil {
		return err
	}
	out.PartitionID = in.PartitionID
	out.ProjectID = in.ProjectID
	return nil
}

// Convert_metal_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_metal_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *metal.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_metal_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_metal_InfrastructureStatus(in *InfrastructureStatus, out *metal.InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_FirewallStatus_To_metal_FirewallStatus(&in.Firewall, &out.Firewall, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_metal_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_metal_InfrastructureStatus(in *InfrastructureStatus, out *metal.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_metal_InfrastructureStatus(in, out, s)
}

func autoConvert_metal_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *metal.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_metal_FirewallStatus_To_v1alpha1_FirewallStatus(&in.Firewall, &out.Firewall, s); err != nil {
		return err
	}
	return nil
}

// Convert_metal_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_metal_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *metal.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_metal_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_IssuerConfig_To_metal_IssuerConfig(in *IssuerConfig, out *metal.IssuerConfig, s conversion.Scope) error {
	out.Url = in.Url
	out.ClientId = in.ClientId
	return nil
}

// Convert_v1alpha1_IssuerConfig_To_metal_IssuerConfig is an autogenerated conversion function.
func Convert_v1alpha1_IssuerConfig_To_metal_IssuerConfig(in *IssuerConfig, out *metal.IssuerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_IssuerConfig_To_metal_IssuerConfig(in, out, s)
}

func autoConvert_metal_IssuerConfig_To_v1alpha1_IssuerConfig(in *metal.IssuerConfig, out *IssuerConfig, s conversion.Scope) error {
	out.Url = in.Url
	out.ClientId = in.ClientId
	return nil
}

// Convert_metal_IssuerConfig_To_v1alpha1_IssuerConfig is an autogenerated conversion function.
func Convert_metal_IssuerConfig_To_v1alpha1_IssuerConfig(in *metal.IssuerConfig, out *IssuerConfig, s conversion.Scope) error {
	return autoConvert_metal_IssuerConfig_To_v1alpha1_IssuerConfig(in, out, s)
}

func autoConvert_v1alpha1_MetalControlPlane_To_metal_MetalControlPlane(in *MetalControlPlane, out *metal.MetalControlPlane, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.IAMConfig = (*metal.IAMConfig)(unsafe.Pointer(in.IAMConfig))
	out.Partitions = *(*map[string]metal.Partition)(unsafe.Pointer(&in.Partitions))
	out.FirewallImages = *(*[]string)(unsafe.Pointer(&in.FirewallImages))
	out.FirewallControllerVersions = *(*[]metal.FirewallControllerVersion)(unsafe.Pointer(&in.FirewallControllerVersions))
	return nil
}

// Convert_v1alpha1_MetalControlPlane_To_metal_MetalControlPlane is an autogenerated conversion function.
func Convert_v1alpha1_MetalControlPlane_To_metal_MetalControlPlane(in *MetalControlPlane, out *metal.MetalControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha1_MetalControlPlane_To_metal_MetalControlPlane(in, out, s)
}

func autoConvert_metal_MetalControlPlane_To_v1alpha1_MetalControlPlane(in *metal.MetalControlPlane, out *MetalControlPlane, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.IAMConfig = (*IAMConfig)(unsafe.Pointer(in.IAMConfig))
	out.Partitions = *(*map[string]Partition)(unsafe.Pointer(&in.Partitions))
	out.FirewallImages = *(*[]string)(unsafe.Pointer(&in.FirewallImages))
	out.FirewallControllerVersions = *(*[]FirewallControllerVersion)(unsafe.Pointer(&in.FirewallControllerVersions))
	return nil
}

// Convert_metal_MetalControlPlane_To_v1alpha1_MetalControlPlane is an autogenerated conversion function.
func Convert_metal_MetalControlPlane_To_v1alpha1_MetalControlPlane(in *metal.MetalControlPlane, out *MetalControlPlane, s conversion.Scope) error {
	return autoConvert_metal_MetalControlPlane_To_v1alpha1_MetalControlPlane(in, out, s)
}

func autoConvert_v1alpha1_NamespaceGroupConfig_To_metal_NamespaceGroupConfig(in *NamespaceGroupConfig, out *metal.NamespaceGroupConfig, s conversion.Scope) error {
	out.ExcludedNamespaces = in.ExcludedNamespaces
	out.ExpectedGroupsList = in.ExpectedGroupsList
	out.NamespaceMaxLength = in.NamespaceMaxLength
	out.ClusterGroupnameTemplate = in.ClusterGroupnameTemplate
	out.RoleBindingNameTemplate = in.RoleBindingNameTemplate
	return nil
}

// Convert_v1alpha1_NamespaceGroupConfig_To_metal_NamespaceGroupConfig is an autogenerated conversion function.
func Convert_v1alpha1_NamespaceGroupConfig_To_metal_NamespaceGroupConfig(in *NamespaceGroupConfig, out *metal.NamespaceGroupConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NamespaceGroupConfig_To_metal_NamespaceGroupConfig(in, out, s)
}

func autoConvert_metal_NamespaceGroupConfig_To_v1alpha1_NamespaceGroupConfig(in *metal.NamespaceGroupConfig, out *NamespaceGroupConfig, s conversion.Scope) error {
	out.ExcludedNamespaces = in.ExcludedNamespaces
	out.ExpectedGroupsList = in.ExpectedGroupsList
	out.NamespaceMaxLength = in.NamespaceMaxLength
	out.ClusterGroupnameTemplate = in.ClusterGroupnameTemplate
	out.RoleBindingNameTemplate = in.RoleBindingNameTemplate
	return nil
}

// Convert_metal_NamespaceGroupConfig_To_v1alpha1_NamespaceGroupConfig is an autogenerated conversion function.
func Convert_metal_NamespaceGroupConfig_To_v1alpha1_NamespaceGroupConfig(in *metal.NamespaceGroupConfig, out *NamespaceGroupConfig, s conversion.Scope) error {
	return autoConvert_metal_NamespaceGroupConfig_To_v1alpha1_NamespaceGroupConfig(in, out, s)
}

func autoConvert_v1alpha1_Partition_To_metal_Partition(in *Partition, out *metal.Partition, s conversion.Scope) error {
	out.FirewallTypes = *(*[]string)(unsafe.Pointer(&in.FirewallTypes))
	return nil
}

// Convert_v1alpha1_Partition_To_metal_Partition is an autogenerated conversion function.
func Convert_v1alpha1_Partition_To_metal_Partition(in *Partition, out *metal.Partition, s conversion.Scope) error {
	return autoConvert_v1alpha1_Partition_To_metal_Partition(in, out, s)
}

func autoConvert_metal_Partition_To_v1alpha1_Partition(in *metal.Partition, out *Partition, s conversion.Scope) error {
	out.FirewallTypes = *(*[]string)(unsafe.Pointer(&in.FirewallTypes))
	return nil
}

// Convert_metal_Partition_To_v1alpha1_Partition is an autogenerated conversion function.
func Convert_metal_Partition_To_v1alpha1_Partition(in *metal.Partition, out *Partition, s conversion.Scope) error {
	return autoConvert_metal_Partition_To_v1alpha1_Partition(in, out, s)
}

func autoConvert_v1alpha1_RateLimit_To_metal_RateLimit(in *RateLimit, out *metal.RateLimit, s conversion.Scope) error {
	out.NetworkID = in.NetworkID
	out.RateLimit = in.RateLimit
	return nil
}

// Convert_v1alpha1_RateLimit_To_metal_RateLimit is an autogenerated conversion function.
func Convert_v1alpha1_RateLimit_To_metal_RateLimit(in *RateLimit, out *metal.RateLimit, s conversion.Scope) error {
	return autoConvert_v1alpha1_RateLimit_To_metal_RateLimit(in, out, s)
}

func autoConvert_metal_RateLimit_To_v1alpha1_RateLimit(in *metal.RateLimit, out *RateLimit, s conversion.Scope) error {
	out.NetworkID = in.NetworkID
	out.RateLimit = in.RateLimit
	return nil
}

// Convert_metal_RateLimit_To_v1alpha1_RateLimit is an autogenerated conversion function.
func Convert_metal_RateLimit_To_v1alpha1_RateLimit(in *metal.RateLimit, out *RateLimit, s conversion.Scope) error {
	return autoConvert_metal_RateLimit_To_v1alpha1_RateLimit(in, out, s)
}

func autoConvert_v1alpha1_WorkerStatus_To_metal_WorkerStatus(in *WorkerStatus, out *metal.WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]metal.MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_WorkerStatus_To_metal_WorkerStatus is an autogenerated conversion function.
func Convert_v1alpha1_WorkerStatus_To_metal_WorkerStatus(in *WorkerStatus, out *metal.WorkerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerStatus_To_metal_WorkerStatus(in, out, s)
}

func autoConvert_metal_WorkerStatus_To_v1alpha1_WorkerStatus(in *metal.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]config.MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_metal_WorkerStatus_To_v1alpha1_WorkerStatus is an autogenerated conversion function.
func Convert_metal_WorkerStatus_To_v1alpha1_WorkerStatus(in *metal.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return autoConvert_metal_WorkerStatus_To_v1alpha1_WorkerStatus(in, out, s)
}
