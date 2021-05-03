package templates

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var Templates = make([]client.Object, 0)

func Register(obj client.Object) (err error) {
	Templates = append(Templates, obj)
	return nil
}
