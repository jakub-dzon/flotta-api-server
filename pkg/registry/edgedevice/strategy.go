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

package edgedevice

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

// NewStrategy creates and returns a edgeDeviceStrategy instance
func NewStrategy(typer runtime.ObjectTyper) edgeDeviceStrategy {
	return edgeDeviceStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not an EdgeDevice
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*v1alpha1.EdgeDevice)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not an EdgeDevice")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchEdgeDevice is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchEdgeDevice(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *v1alpha1.EdgeDevice) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type edgeDeviceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s edgeDeviceStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return []string{}
}

func (s edgeDeviceStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

func (edgeDeviceStrategy) NamespaceScoped() bool {
	return true
}

func (edgeDeviceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (edgeDeviceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (edgeDeviceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	edgeDevice := obj.(*v1alpha1.EdgeDevice)
	return ValidateEdgeDevice(edgeDevice)
}

func (edgeDeviceStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (edgeDeviceStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (edgeDeviceStrategy) Canonicalize(obj runtime.Object) {
}

func (edgeDeviceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// ValidateEdgeDevice validates an EdgeDevice.
func ValidateEdgeDevice(_ *v1alpha1.EdgeDevice) field.ErrorList {
	return field.ErrorList{}
}
