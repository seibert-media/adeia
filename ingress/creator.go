// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"github.com/seibert-media/adeia/domain"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Creator for transform domain to ingress
type Creator struct {
	Ingressname string
	Servicename string
	Serviceport string
	Namespace string
}

// Create ingress for the given domains.
func (c *Creator) Create(domains []domain.Domain) *v1beta1.Ingress {
	return &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "extensions/v1beta1",
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "traefik",
			},
			Name:      c.Ingressname,
			Namespace: c.Namespace,
		},
		Spec: v1beta1.IngressSpec{
			Rules: c.buildRuleSet(domains),
		},
	}
}

func (c *Creator) buildRuleSet(domains []domain.Domain) []v1beta1.IngressRule {
	var ingressRules []v1beta1.IngressRule
	for _, domain := range domains {
		ingressRule := v1beta1.IngressRule{
			Host: string(domain),
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: &v1beta1.HTTPIngressRuleValue{
					Paths: []v1beta1.HTTPIngressPath{
						{
							Path: "/",
							Backend: v1beta1.IngressBackend{
								ServiceName: c.Servicename,
								ServicePort: intstr.Parse(c.Serviceport),
							},
						},
					},
				},
			},
		}
		ingressRules = append(ingressRules, ingressRule)
	}
	return ingressRules
}
