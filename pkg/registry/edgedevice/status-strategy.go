package edgedevice

import (
	"context"
	"github.com/jakub-dzon/flotta-apiserver/pkg/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.RESTUpdateStrategy = &statusSubResourceStrategy{}

// StatusSubResourceStrategy defines a default Strategy for the status subresource.
type statusSubResourceStrategy struct {
	rest.RESTUpdateStrategy
}

// PrepareForUpdate calls the PrepareForUpdate function on obj if supported, otherwise does nothing.
func (s *statusSubResourceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	// should panic/fail-fast upon casting failure
	newEdgeDevice := obj.(*v1alpha1.EdgeDevice)
	oldEdgeDevice := old.(*v1alpha1.EdgeDevice)
	// only modifies status
	oldEdgeDevice.Status = newEdgeDevice.Status
	oldEdgeDevice.DeepCopyInto(newEdgeDevice)
}
