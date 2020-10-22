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

package v2

import (
	"context"
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumEgressRoutesGetter has a method to return a CiliumEgressRouteInterface.
// A group's client should implement this interface.
type CiliumEgressRoutesGetter interface {
	CiliumEgressRoutes() CiliumEgressRouteInterface
}

// CiliumEgressRouteInterface has methods to work with CiliumEgressRoute resources.
type CiliumEgressRouteInterface interface {
	Create(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.CreateOptions) (*v2.CiliumEgressRoute, error)
	Update(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (*v2.CiliumEgressRoute, error)
	UpdateStatus(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (*v2.CiliumEgressRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.CiliumEgressRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.CiliumEgressRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressRoute, err error)
	CiliumEgressRouteExpansion
}

// ciliumEgressRoutes implements CiliumEgressRouteInterface
type ciliumEgressRoutes struct {
	client rest.Interface
}

// newCiliumEgressRoutes returns a CiliumEgressRoutes
func newCiliumEgressRoutes(c *CiliumV2Client) *ciliumEgressRoutes {
	return &ciliumEgressRoutes{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumEgressRoute, and returns the corresponding ciliumEgressRoute object, and an error if there is any.
func (c *ciliumEgressRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumEgressRoute, err error) {
	result = &v2.CiliumEgressRoute{}
	err = c.client.Get().
		Resource("ciliumegressroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumEgressRoutes that match those selectors.
func (c *ciliumEgressRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumEgressRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumEgressRouteList{}
	err = c.client.Get().
		Resource("ciliumegressroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumEgressRoutes.
func (c *ciliumEgressRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumegressroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumEgressRoute and creates it.  Returns the server's representation of the ciliumEgressRoute, and an error, if there is any.
func (c *ciliumEgressRoutes) Create(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.CreateOptions) (result *v2.CiliumEgressRoute, err error) {
	result = &v2.CiliumEgressRoute{}
	err = c.client.Post().
		Resource("ciliumegressroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressRoute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumEgressRoute and updates it. Returns the server's representation of the ciliumEgressRoute, and an error, if there is any.
func (c *ciliumEgressRoutes) Update(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (result *v2.CiliumEgressRoute, err error) {
	result = &v2.CiliumEgressRoute{}
	err = c.client.Put().
		Resource("ciliumegressroutes").
		Name(ciliumEgressRoute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressRoute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ciliumEgressRoutes) UpdateStatus(ctx context.Context, ciliumEgressRoute *v2.CiliumEgressRoute, opts v1.UpdateOptions) (result *v2.CiliumEgressRoute, err error) {
	result = &v2.CiliumEgressRoute{}
	err = c.client.Put().
		Resource("ciliumegressroutes").
		Name(ciliumEgressRoute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumEgressRoute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumEgressRoute and deletes it. Returns an error if one occurs.
func (c *ciliumEgressRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumegressroutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumEgressRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumegressroutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumEgressRoute.
func (c *ciliumEgressRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumEgressRoute, err error) {
	result = &v2.CiliumEgressRoute{}
	err = c.client.Patch(pt).
		Resource("ciliumegressroutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
