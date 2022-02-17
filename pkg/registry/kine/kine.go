package kine

import (
	"context"
	"fmt"
	"github.com/rancher/kine/pkg/endpoint"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
)

type RESTOptionsGetter struct {
	delegate   generic.RESTOptionsGetter
	etcdConfig endpoint.ETCDConfig
}

func NewRESTOptionsGetter(delegate generic.RESTOptionsGetter, host string, port int32, username, password string) (generic.RESTOptionsGetter, error) {
	etcdConfig, err := endpoint.Listen(context.TODO(), endpoint.Config{
		Endpoint: fmt.Sprintf("postgres://%s:%d/kine?user=%s&password=%s&sslmode=disable", host, port, username, password),
	})
	if err != nil {
		return nil, err
	}
	return &RESTOptionsGetter{
		delegate:   delegate,
		etcdConfig: etcdConfig,
	}, nil
}

// GetRESTOptions implements RESTOptionsGetter interface.
func (k *RESTOptionsGetter) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	restOptions, err := k.delegate.GetRESTOptions(resource)
	if err != nil {
		return restOptions, err
	}

	if err != nil {
		return restOptions, err
	}
	restOptions.StorageConfig.Transport.ServerList = k.etcdConfig.Endpoints
	restOptions.StorageConfig.Transport.TrustedCAFile = k.etcdConfig.TLSConfig.CAFile
	restOptions.StorageConfig.Transport.CertFile = k.etcdConfig.TLSConfig.CertFile
	restOptions.StorageConfig.Transport.KeyFile = k.etcdConfig.TLSConfig.KeyFile

	return restOptions, nil
}
