package templates

import (
	values "github.com/lllamnyp/typedhelm/values"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	Deployment = func() appsv1.Deployment {
		d := appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      values.Name,
				Namespace: values.Namespace,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &values.Replicas,
				Selector: &metav1.LabelSelector{
					MatchLabels: values.Labels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: values.Labels,
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  values.Name,
								Image: values.Image,
							},
						},
					},
				},
			},
		}

		d.Spec.Replicas = &values.Replicas
		d.Spec.Selector.MatchLabels = values.Labels
		d.Spec.Template.Labels = d.Spec.Selector.MatchLabels

		return d
	}()
	_ = Register(&Deployment)
)
