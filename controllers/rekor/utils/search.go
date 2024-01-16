package utils

import (
	"context"
	"errors"

	rhtasv1alpha1 "github.com/securesign/operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func FindRekor(ctx context.Context, cli client.Client, namespace string, labels map[string]string) (*rhtasv1alpha1.Rekor, error) {
	list := &rhtasv1alpha1.RekorList{}
	err := cli.List(ctx, list, client.InNamespace(namespace), client.MatchingLabels(labels), client.Limit(1))
	if err != nil {
		return nil, err
	}
	if len(list.Items) == 1 {
		return &list.Items[0], nil
	}
	// try to find any resource in namespace
	err = cli.List(ctx, list, client.InNamespace(namespace), client.Limit(1))
	if err != nil {
		return nil, err
	}

	if len(list.Items) == 1 {
		return &list.Items[0], nil
	}
	return nil, errors.New("component not found")
}
