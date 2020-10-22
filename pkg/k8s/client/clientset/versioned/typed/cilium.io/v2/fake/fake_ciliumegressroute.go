// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumEgressRoutes implements CiliumEgressRouteInterface
type FakeCiliumEgressRoutes struct {
	Fake *FakeCiliumV2
}

var ciliumegressroutesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2", Resource: "ciliumegressroutes"}

var ciliumegressroutesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2", Kind: "CiliumEgressRoute"}

// Get takes name of the ciliumEgressRoute, and returns the corresponding ciliumEgressRoute object, and an error if there is any.
func (c *FakeCiliumEgressRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumEgressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumegressroutesResource, name), &v2.CiliumEgressRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressRoute), err
}

// List takes label and field selectors, and returns the list of CiliumEgressRoutes that match those selectors.
func (c *FakeCiliumEgressRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumEgressRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumegressroutesResource, ciliumegressroutesKind, opts), &v2.CiliumEgressRouteList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CiliumEgressRouteList{ListMeta: obj.(*v2.CiliumEgressRouteList).ListMeta}
	for _, item := range obj.(*v2.CiliumEgressRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumEgressRoutes.
func (c *FakeCiliumEgressRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumegressroutesResource, opts))
}

// Create takes the representation of a ciliumEgressRoute and creates it.  Returns the server's representation of the ciliumEgressRoute, and an error, if there is any.
func (c *FakeCiliumEgressRoutes) Create(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.CreateOptions) (result *v2.CiliumEgressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumegressroutesResource, ciliumEgressRoute), &v2.CiliumEgressRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressRoute), err
}

// Update takes the representation of a ciliumEgressRoute and updates it. Returns the server's representation of the ciliumEgressRoute, and an error, if there is any.
func (c *FakeCiliumEgressRoutes) Update(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (result *v2.CiliumEgressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumegressroutesResource, ciliumEgressRoute), &v2.CiliumEgressRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCiliumEgressRoutes) UpdateStatus(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (*v2.CiliumEgressRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ciliumegressroutesResource, "status", ciliumEgressRoute), &v2.CiliumEgressRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressRoute), err
}

// Delete takes name of the ciliumEgressRoute and deletes it. Returns an error if one occurs.
func (c *FakeCiliumEgressRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ciliumegressroutesResource, name), &v2.CiliumEgressRoute{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumEgressRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumegressroutesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CiliumEgressRouteList{})
	return err
}

// Patch applies the patch and returns the patched ciliumEgressRoute.
func (c *FakeCiliumEgressRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumegressroutesResource, name, pt, data, subresources...), &v2.CiliumEgressRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CiliumEgressRoute), err
}
