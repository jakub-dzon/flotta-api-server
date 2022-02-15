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

package registry

import (
	"context"
	"fmt"
	"github.com/rancher/kine/pkg/endpoint"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/rest"
)

type KineRESTOptionsGetter struct {
	delegate generic.RESTOptionsGetter
	host     string
	port     int32
	username string
	password string
}

func NewKineRESTOptionsGetter(originalGetter generic.RESTOptionsGetter, host string, port int32, username, password string) generic.RESTOptionsGetter {
	return &KineRESTOptionsGetter{
		delegate: originalGetter,
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

// GetRESTOptions implements RESTOptionsGetter interface.
func (k *KineRESTOptionsGetter) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	restOptions, err := k.delegate.GetRESTOptions(resource)
	if err != nil {
		return generic.RESTOptions{}, err
	}
	etcdConfig, err := endpoint.Listen(context.TODO(), endpoint.Config{
		Endpoint: fmt.Sprintf("postgres://%s:%d/kine?user=%s&password=%s&sslmode=disable", k.host, k.port, k.username, k.password),
	})

	restOptions.StorageConfig.Transport.ServerList = etcdConfig.Endpoints
	restOptions.StorageConfig.Transport.TrustedCAFile = etcdConfig.TLSConfig.CAFile
	restOptions.StorageConfig.Transport.CertFile = etcdConfig.TLSConfig.CertFile
	restOptions.StorageConfig.Transport.KeyFile = etcdConfig.TLSConfig.KeyFile
	return restOptions, nil
}

func NewStore(host string, port int32, username, password string) rest.StoreFn {
	return func(_ *genericregistry.Store, options *generic.StoreOptions) {
		options.RESTOptions = NewKineRESTOptionsGetter(options.RESTOptions, host, port, username, password)
	}
}
