/*
Copyright 2016 The Kubernetes Authors.

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

package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/jakub-dzon/flotta-apiserver/api/v1alpha1"
	"github.com/jakub-dzon/flotta-apiserver/pkg/registry"
	"k8s.io/klog"
	// load all auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

var DBConfig struct {
	Host     string `envconfig:"HOST" default:"postgres-postgresql.default"`
	Port     int32  `envconfig:"PORT" default:"5432"`
	Username string `envconfig:"USERNAME" default:"postgres"`
	Password string `envconfig:"PASSWORD"`
}

func main() {
	err := envconfig.Process("DB", &DBConfig)
	if err != nil {
		klog.Fatal(err, "unable to process configuration values")
	}
	store := registry.NewStore(DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password)
	err = builder.APIServer.
		WithResourceAndStorage(&v1alpha1.EdgeDevice{}, store).
		WithLocalDebugExtension().
		WithOptionsFns(func(options *builder.ServerOptions) *builder.ServerOptions {
			options.RecommendedOptions.CoreAPI = nil
			options.RecommendedOptions.Admission = nil
			return options
		}).
		Execute()

	if err != nil {
		klog.Fatal(err)
	}
}
