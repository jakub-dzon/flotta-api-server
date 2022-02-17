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

package edgedeployment

import (
	"github.com/jakub-dzon/flotta-apiserver/pkg/apis/v1alpha1"
	"github.com/jakub-dzon/flotta-apiserver/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.Storage, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &v1alpha1.EdgeDeployment{} },
		NewListFunc:              func() runtime.Object { return &v1alpha1.EdgeDeploymentList{} },
		PredicateFunc:            MatchEdgeDeployment,
		DefaultQualifiedResource: v1alpha1.Resource("edgedeployments"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(v1alpha1.Resource("edgedeployments")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	main := &registry.REST{Store: store}

	return &registry.Storage{Main: main}, nil
}
