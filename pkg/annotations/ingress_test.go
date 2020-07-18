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

package annotations

import (
	"testing"

	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIngress(t *testing.T) {
	for _, tc := range []struct {
		ing                       *v1beta1.Ingress
		allowHTTP                 bool
		useNamedTLS               string
		staticIPName              string
		ingressClass              string
		reserveGlobalStaticIPName string
	}{
		{
			ing:       &v1beta1.Ingress{},
			allowHTTP: true, // defaults to true.
		},
		{
			ing: &v1beta1.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						AllowHTTPKey:     "false",
						IngressClassKey:  "gce",
						PreSharedCertKey: "shared-cert-key",
						StaticIPNameKey:  "1.2.3.4",
					},
				},
			},
			allowHTTP:    false,
			useNamedTLS:  "shared-cert-key",
			staticIPName: "1.2.3.4",
			ingressClass: "gce",
		},
		{
			ing: &v1beta1.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						ReserveGlobalStaticIPNameKey: "1.2.3.4-managed",
					},
				},
			},
			allowHTTP:                 true,
			reserveGlobalStaticIPName: "1.2.3.4-managed",
		},
	} {
		ing := FromIngress(tc.ing)
		if x := ing.AllowHTTP(); x != tc.allowHTTP {
			t.Errorf("ingress %+v; AllowHTTP() = %v, want %v", tc.ing, x, tc.allowHTTP)
		}
		if x := ing.UseNamedTLS(); x != tc.useNamedTLS {
			t.Errorf("ingress %+v; UseNamedTLS() = %v, want %v", tc.ing, x, tc.useNamedTLS)
		}
		if x := ing.StaticIPName(); x != tc.staticIPName {
			t.Errorf("ingress %+v; StaticIPName() = %v, want %v", tc.ing, x, tc.staticIPName)
		}
		if x := ing.IngressClass(); x != tc.ingressClass {
			t.Errorf("ingress %+v; IngressClass() = %v, want %v", tc.ing, x, tc.ingressClass)
		}
		if x := ing.ReserveGlobalStaticIPName(); x != tc.reserveGlobalStaticIPName {
			t.Errorf("ingress %+v; ReserveGlobalStaticIPName() = %v, want %v", tc.ing, x, tc.reserveGlobalStaticIPName)
		}
	}
}
