/*
Copyright 2016 The Kubernetes Authors.

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

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/sample-apiserver/pkg/apis/kafee"
)

// ValidateEspresso validates a Espresso.
func ValidateEspresso(f *kafee.Espresso) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidateEspressoSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateEspressoSpec validates a EspressoSpec.
func ValidateEspressoSpec(s *kafee.EspressoSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(s.EspressoReference) != 0 && len(s.CapuccinoReference) != 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("capuccinoReference"), s.CapuccinoReference, "cannot be set with espressoReference at the same time"))
	} else if len(s.EspressoReference) != 0 && s.ReferenceType != kafee.EspressoReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("espressoReference"), s.EspressoReference, "cannot be set if referenceType is not Espresso"))
	} else if len(s.CapuccinoReference) != 0 && s.ReferenceType != kafee.CapuccinoReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("capuccinoReference"), s.CapuccinoReference, "cannot be set if referenceType is not Capuccino"))
	} else if len(s.CapuccinoReference) == 0 && s.ReferenceType == kafee.CapuccinoReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("capuccinoReference"), s.CapuccinoReference, "cannot be empty if referenceType is Capuccino"))
	} else if len(s.EspressoReference) == 0 && s.ReferenceType == kafee.EspressoReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("espressoReference"), s.EspressoReference, "cannot be empty if referenceType is Espresso"))
	}

	if len(s.ReferenceType) != 0 && s.ReferenceType != kafee.CapuccinoReferenceType && s.ReferenceType != kafee.EspressoReferenceType {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("referenceType"), s.ReferenceType, "must be Espresso or Capuccino"))
	}

	return allErrs
}
