package main

import (
	"context"

	"github.com/lllamnyp/typedhelm/templates"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	ctx := context.Background()
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		panic(err)
	}
	for _, obj := range templates.Templates {
		c.Create(ctx, obj)
	}
}
