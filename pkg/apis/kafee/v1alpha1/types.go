/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlunderList is a list of Flunder objects.
type EspressoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Espresso `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ReferenceType string

const (
	EspressoReferenceType = ReferenceType("Espresso")
	CapuccinoReferenceType = ReferenceType("Capuccino")
)

type EspressoSpec struct {
	// A name of another flunder or fischer, depending on the reference type.
	Reference string `json:"reference,omitempty" protobuf:"bytes,1,opt,name=reference"`
	// The reference type, defaults to "Flunder" if reference is set.
	ReferenceType *ReferenceType `json:"referenceType,omitempty" protobuf:"bytes,2,opt,name=referenceType"`
}

type EspressoStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Espresso struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   EspressoSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status EspressoStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Capuccino struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// DisallowedFlunders holds a list of Flunder.Names that are disallowed.
	DisallowedEspressos []string `json:"disallowedEspressos,omitempty" protobuf:"bytes,2,rep,name=disallowedEspressos"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FischerList is a list of Fischer objects.
type CapuccinoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Capuccino `json:"items" protobuf:"bytes,2,rep,name=items"`
}
