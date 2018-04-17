package converter

import (
	"github.com/seibert-media/k8s-ingress/model"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Converter struct {
	Domains []model.Domain
	Ingress []v1beta1.Ingress
}

func (c *Converter) Convert() ([]v1beta1.Ingress, error) {
	return []v1beta1.Ingress{v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "traefik",
			},
		},
		Spec: v1beta1.IngressSpec{
			Rules: []v1beta1.IngressRule{
				{
					Host: string("www.example.com"),
					IngressRuleValue: v1beta1.IngressRuleValue{
					},
				},
			},
		},
	}}, nil
}