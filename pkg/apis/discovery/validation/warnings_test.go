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
	"testing"

	"k8s.io/kubernetes/pkg/apis/discovery"
)

func TestWarningsOnEndpointSliceAddressType(t *testing.T) {
	tests := []struct {
		name        string
		addressType discovery.AddressType
		wantWarning bool
	}{
		{
			name:        "AddressType = FQDN",
			addressType: discovery.AddressTypeFQDN,
			wantWarning: true,
		},
		{
			name:        "AddressType = IPV4",
			addressType: discovery.AddressTypeIPv4,
			wantWarning: false,
		},
		{
			name:        "AddressType = IPV6",
			addressType: discovery.AddressTypeIPv6,
			wantWarning: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			edp := discovery.EndpointSlice{AddressType: tc.addressType}
			got := GetWarningsForEndpointSliceCreation(&edp)
			if tc.wantWarning && len(got) == 0 {
				t.Fatal("Failed warning was not returned")
			} else if !tc.wantWarning && len(got) != 0 {
				t.Fatalf("Failed warning  was returned (%v)", got)
			}
		})
	}
}
