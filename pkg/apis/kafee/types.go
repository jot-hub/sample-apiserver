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

package kafee

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EspressoList is a list of Espresso objects.
type EspressoList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Espresso
}

// ReferenceType defines the type of an object reference.
type ReferenceType string

const (
	// EspressoReferenceType is used for Espresso references.
	EspressoReferenceType = ReferenceType("Espresso")
	// CapuccinoReferenceType is used for Capuccino references.
	CapuccinoReferenceType = ReferenceType("Capuccino")
)

// EspressoSpec is the specification of a Espresso.
type EspressoSpec struct {
	// A name of another Espresso, mutually exclusive to the CapuccinoReference.
	EspressoReference string
	// A name of a Capuccino, mutually exclusive to the EspressoReference.
	CapuccinoReference string
	// The reference type.
	ReferenceType ReferenceType
}

// EspressoStatus is the status of a Espresso.
type EspressoStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Espresso is an example type with a spec and a status.
type Espresso struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   EspressoSpec
	Status EspressoStatus
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Capuccino is an example type with a list of disallowed Espresso.Names
type Capuccino struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	// DisallowedEspressos holds a list of Espresso.Names that are disallowed.
	DisallowedEspressos []string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CapuccinoList is a list of Capuccino objects.
type CapuccinoList struct {
	metav1.TypeMeta
	metav1.ListMeta

	// Items is a list of Capuccinos
	Items []Capuccino
}
