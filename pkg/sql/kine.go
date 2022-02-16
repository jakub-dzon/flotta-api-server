package sql

import (
	"context"
	"github.com/rancher/kine/pkg/endpoint"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
)

type KineProxiedRESTOptionsGetter struct {}

// GetRESTOptions implements RESTOptionsGetter interface.
func (k *KineProxiedRESTOptionsGetter) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	restOptions := generic.RESTOptions{}

	etcdConfig, err := endpoint.Listen(context.TODO(), endpoint.Config{
		Endpoint: "",
	})
	if err != nil {
		return generic.RESTOptions{}, err
	}

	restOptions.StorageConfig.Transport.ServerList = etcdConfig.Endpoints
	restOptions.StorageConfig.Transport.TrustedCAFile = etcdConfig.TLSConfig.CAFile
	restOptions.StorageConfig.Transport.CertFile = etcdConfig.TLSConfig.CertFile
	restOptions.StorageConfig.Transport.KeyFile = etcdConfig.TLSConfig.KeyFile
	return restOptions, nil
}
