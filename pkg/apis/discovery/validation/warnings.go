/*
Copyright 2022 The Kubernetes Authors.

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
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/apis/discovery"
)

func GetWarningsForEndpointSliceCreation(eps *discovery.EndpointSlice) []string {
	var warnings []string
	warnings = append(warnings, warnOnDeprecatedAddressType(eps.AddressType)...)
	return warnings
}

func warnOnDeprecatedAddressType(addressType discovery.AddressType) []string {
	if addressType == discovery.AddressTypeFQDN {
		return []string{fmt.Sprintf("%s: FQDN type is deprecated.", field.NewPath("spec").Child("addressType"))}
	}
	return nil
}
