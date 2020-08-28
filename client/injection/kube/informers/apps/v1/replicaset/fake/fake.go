/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	replicaset "knative.dev/pkg/client/injection/kube/informers/apps/v1/replicaset"
	fake "knative.dev/pkg/client/injection/kube/informers/factory/fake"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = replicaset.Get

func init() {
	injection.Fake.RegisterInformer(
		withInformer,
		metav1.GroupVersionResource{
			Group:    "apps",
			Version:  "v1",
			Resource: "replicasets",
		})
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Apps().V1().ReplicaSets()
	return context.WithValue(ctx, replicaset.Key{}, inf), inf.Informer()
}
