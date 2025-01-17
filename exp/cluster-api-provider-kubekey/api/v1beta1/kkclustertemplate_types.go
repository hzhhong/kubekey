/*
Copyright 2022.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// KKClusterTemplateSpec defines the desired state of KKClusterTemplate
type KKClusterTemplateSpec struct {
	Template KKClusterTemplateResource `json:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kkclustertemplates,scope=Namespaced,categories=cluster-api,shortName=kkct
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Time duration since creation of KKClusterTemplate"
// +k8s:defaulter-gen=true

// KKClusterTemplate is the Schema for the kkclustertemplates API
type KKClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KKClusterTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// KKClusterTemplateList contains a list of KKClusterTemplate
type KKClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KKClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KKClusterTemplate{}, &KKClusterTemplateList{})
}

// KKClusterTemplateResource Standard object's metadata
type KKClusterTemplateResource struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata,omitempty"`
	Spec       KKClusterSpec        `json:"spec"`
}
