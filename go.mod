module github.com/jakub-dzon/flotta-apiserver

go 1.16

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rancher/kine v0.4.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/apiserver v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/klog v1.0.0
	sigs.k8s.io/apiserver-runtime v1.0.2
)


replace sigs.k8s.io/apiserver-runtime => github.com/jakub-dzon/apiserver-runtime v1.0.3-0.20220216125922-1657cdadd1a0
