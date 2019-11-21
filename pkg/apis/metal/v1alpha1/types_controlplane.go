// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	gardenv1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlPlaneConfig contains configuration settings for the control plane.
type ControlPlaneConfig struct {
	metav1.TypeMeta `json:",inline"`

	// CloudControllerManager contains configuration settings for the cloud-controller-manager.
	// +optional
	CloudControllerManager *CloudControllerManagerConfig `json:"cloudControllerManager,omitempty"`

	// TokenIssuer contains configuration settings for the token issuer.
	TokenIssuer TokenIssuer `json:"tokenissuer" optional:"false"`
}

// CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.
type CloudControllerManagerConfig struct {
	gardenv1beta1.KubernetesConfig `json:",inline"`
}

// TokenIssuer for configuration of AuthNWebhook
type TokenIssuer struct {
	// IssuerURL the url of the issuer
	IssuerURL string `json:"issuerUrl" optional:"false"`
	// ClientId the id of the client
	ClientID string `json:"clientId" optional:"false"`
}
