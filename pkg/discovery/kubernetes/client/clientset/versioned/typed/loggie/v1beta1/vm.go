/*
Copyright The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	scheme "github.com/loggie-io/loggie/pkg/discovery/kubernetes/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VmsGetter has a method to return a VmInterface.
// A group's client should implement this interface.
type VmsGetter interface {
	Vms() VmInterface
}

// VmInterface has methods to work with Vm resources.
type VmInterface interface {
	Create(ctx context.Context, vm *v1beta1.Vm, opts v1.CreateOptions) (*v1beta1.Vm, error)
	Update(ctx context.Context, vm *v1beta1.Vm, opts v1.UpdateOptions) (*v1beta1.Vm, error)
	UpdateStatus(ctx context.Context, vm *v1beta1.Vm, opts v1.UpdateOptions) (*v1beta1.Vm, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Vm, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VmList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Vm, err error)
	VmExpansion
}

// vms implements VmInterface
type vms struct {
	client rest.Interface
}

// newVms returns a Vms
func newVms(c *LoggieV1beta1Client) *vms {
	return &vms{
		client: c.RESTClient(),
	}
}

// Get takes name of the vm, and returns the corresponding vm object, and an error if there is any.
func (c *vms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Vm, err error) {
	result = &v1beta1.Vm{}
	err = c.client.Get().
		Resource("vms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Vms that match those selectors.
func (c *vms) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VmList{}
	err = c.client.Get().
		Resource("vms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vms.
func (c *vms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("vms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vm and creates it.  Returns the server's representation of the vm, and an error, if there is any.
func (c *vms) Create(ctx context.Context, vm *v1beta1.Vm, opts v1.CreateOptions) (result *v1beta1.Vm, err error) {
	result = &v1beta1.Vm{}
	err = c.client.Post().
		Resource("vms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vm).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vm and updates it. Returns the server's representation of the vm, and an error, if there is any.
func (c *vms) Update(ctx context.Context, vm *v1beta1.Vm, opts v1.UpdateOptions) (result *v1beta1.Vm, err error) {
	result = &v1beta1.Vm{}
	err = c.client.Put().
		Resource("vms").
		Name(vm.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vm).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vms) UpdateStatus(ctx context.Context, vm *v1beta1.Vm, opts v1.UpdateOptions) (result *v1beta1.Vm, err error) {
	result = &v1beta1.Vm{}
	err = c.client.Put().
		Resource("vms").
		Name(vm.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vm).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vm and deletes it. Returns an error if one occurs.
func (c *vms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("vms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vm.
func (c *vms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Vm, err error) {
	result = &v1beta1.Vm{}
	err = c.client.Patch(pt).
		Resource("vms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}