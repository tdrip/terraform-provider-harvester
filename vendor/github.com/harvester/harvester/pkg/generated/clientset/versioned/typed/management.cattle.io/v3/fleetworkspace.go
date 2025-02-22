/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FleetWorkspacesGetter has a method to return a FleetWorkspaceInterface.
// A group's client should implement this interface.
type FleetWorkspacesGetter interface {
	FleetWorkspaces() FleetWorkspaceInterface
}

// FleetWorkspaceInterface has methods to work with FleetWorkspace resources.
type FleetWorkspaceInterface interface {
	Create(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.CreateOptions) (*v3.FleetWorkspace, error)
	Update(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (*v3.FleetWorkspace, error)
	UpdateStatus(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (*v3.FleetWorkspace, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.FleetWorkspace, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.FleetWorkspaceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FleetWorkspace, err error)
	FleetWorkspaceExpansion
}

// fleetWorkspaces implements FleetWorkspaceInterface
type fleetWorkspaces struct {
	client rest.Interface
}

// newFleetWorkspaces returns a FleetWorkspaces
func newFleetWorkspaces(c *ManagementV3Client) *fleetWorkspaces {
	return &fleetWorkspaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the fleetWorkspace, and returns the corresponding fleetWorkspace object, and an error if there is any.
func (c *fleetWorkspaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.FleetWorkspace, err error) {
	result = &v3.FleetWorkspace{}
	err = c.client.Get().
		Resource("fleetworkspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FleetWorkspaces that match those selectors.
func (c *fleetWorkspaces) List(ctx context.Context, opts v1.ListOptions) (result *v3.FleetWorkspaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.FleetWorkspaceList{}
	err = c.client.Get().
		Resource("fleetworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fleetWorkspaces.
func (c *fleetWorkspaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("fleetworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fleetWorkspace and creates it.  Returns the server's representation of the fleetWorkspace, and an error, if there is any.
func (c *fleetWorkspaces) Create(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.CreateOptions) (result *v3.FleetWorkspace, err error) {
	result = &v3.FleetWorkspace{}
	err = c.client.Post().
		Resource("fleetworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fleetWorkspace).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fleetWorkspace and updates it. Returns the server's representation of the fleetWorkspace, and an error, if there is any.
func (c *fleetWorkspaces) Update(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (result *v3.FleetWorkspace, err error) {
	result = &v3.FleetWorkspace{}
	err = c.client.Put().
		Resource("fleetworkspaces").
		Name(fleetWorkspace.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fleetWorkspace).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fleetWorkspaces) UpdateStatus(ctx context.Context, fleetWorkspace *v3.FleetWorkspace, opts v1.UpdateOptions) (result *v3.FleetWorkspace, err error) {
	result = &v3.FleetWorkspace{}
	err = c.client.Put().
		Resource("fleetworkspaces").
		Name(fleetWorkspace.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fleetWorkspace).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fleetWorkspace and deletes it. Returns an error if one occurs.
func (c *fleetWorkspaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("fleetworkspaces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fleetWorkspaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("fleetworkspaces").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fleetWorkspace.
func (c *fleetWorkspaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FleetWorkspace, err error) {
	result = &v3.FleetWorkspace{}
	err = c.client.Patch(pt).
		Resource("fleetworkspaces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
