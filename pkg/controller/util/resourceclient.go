/*
Copyright 2018 The Kubernetes Authors.

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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type ResourceClient interface {
	Resources(namespace string) dynamic.ResourceInterface
}

type resourceClient struct {
	client      dynamic.Interface
	apiResource schema.GroupVersionResource
}

func NewResourceClient(config *rest.Config, apiResource *metav1.APIResource) (ResourceClient, error) {
	resource := schema.GroupVersionResource{
		Group:    apiResource.Group,
		Version:  apiResource.Version,
		Resource: apiResource.Name,
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &resourceClient{
		client:      client,
		apiResource: resource,
	}, nil
}

func (c *resourceClient) Resources(namespace string) dynamic.ResourceInterface {
	return c.client.Resource(c.apiResource).Namespace(namespace)
}
