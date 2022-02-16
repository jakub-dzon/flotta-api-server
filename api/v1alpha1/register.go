package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var AddToScheme = func(scheme *runtime.Scheme) error {
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{
		Group:   "management.project-flotta.io",
		Version: "v1alpha1",
	})
	// +kubebuilder:scaffold:install
	scheme.AddKnownTypes(schema.GroupVersion{
		Group:   "management.project-flotta.io",
		Version: "v1alpha1",
	}, &EdgeDevice{}, &EdgeDeviceList{})

	scheme.AddKnownTypes(schema.GroupVersion{
		Group:   "management.project-flotta.io",
		Version: "v1alpha1",
	}, &EdgeDeployment{}, &EdgeDeploymentList{})
	return nil
}