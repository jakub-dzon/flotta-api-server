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
	"context"
	"fmt"
	"github.com/jakub-dzon/flotta-apiserver/pkg/apis/v1alpha1"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

// NewStrategy creates and returns a edgeDeploymentStrategy instance
func NewStrategy(typer runtime.ObjectTyper) edgeDeploymentStrategy {
	return edgeDeploymentStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not an EdgeDeployment
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*v1alpha1.EdgeDeployment)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not an EdgeDeployment")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchEdgeDeployment is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchEdgeDeployment(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *v1alpha1.EdgeDeployment) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type edgeDeploymentStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s edgeDeploymentStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return []string{}
}

func (s edgeDeploymentStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

func (edgeDeploymentStrategy) NamespaceScoped() bool {
	return true
}

func (edgeDeploymentStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (edgeDeploymentStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (edgeDeploymentStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	edgeDevice := obj.(*v1alpha1.EdgeDeployment)
	return ValidateEdgeDeployment(edgeDevice)
}

func (edgeDeploymentStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (edgeDeploymentStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (edgeDeploymentStrategy) Canonicalize(obj runtime.Object) {
}

func (edgeDeploymentStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// ValidateEdgeDeployment validates an EdgeDeployment.
func ValidateEdgeDeployment(_ *v1alpha1.EdgeDeployment) field.ErrorList {
	return field.ErrorList{}
}
