/*
Copyright 2019 Banzai Cloud.

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
	v1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	scheme "github.com/banzaicloud/istio-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RemoteIstiosGetter has a method to return a RemoteIstioInterface.
// A group's client should implement this interface.
type RemoteIstiosGetter interface {
	RemoteIstios(namespace string) RemoteIstioInterface
}

// RemoteIstioInterface has methods to work with RemoteIstio resources.
type RemoteIstioInterface interface {
	Create(*v1beta1.RemoteIstio) (*v1beta1.RemoteIstio, error)
	Update(*v1beta1.RemoteIstio) (*v1beta1.RemoteIstio, error)
	UpdateStatus(*v1beta1.RemoteIstio) (*v1beta1.RemoteIstio, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.RemoteIstio, error)
	List(opts v1.ListOptions) (*v1beta1.RemoteIstioList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RemoteIstio, err error)
	RemoteIstioExpansion
}

// remoteIstios implements RemoteIstioInterface
type remoteIstios struct {
	client rest.Interface
	ns     string
}

// newRemoteIstios returns a RemoteIstios
func newRemoteIstios(c *IstioV1beta1Client, namespace string) *remoteIstios {
	return &remoteIstios{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the remoteIstio, and returns the corresponding remoteIstio object, and an error if there is any.
func (c *remoteIstios) Get(name string, options v1.GetOptions) (result *v1beta1.RemoteIstio, err error) {
	result = &v1beta1.RemoteIstio{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("remoteistios").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RemoteIstios that match those selectors.
func (c *remoteIstios) List(opts v1.ListOptions) (result *v1beta1.RemoteIstioList, err error) {
	result = &v1beta1.RemoteIstioList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("remoteistios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested remoteIstios.
func (c *remoteIstios) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("remoteistios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a remoteIstio and creates it.  Returns the server's representation of the remoteIstio, and an error, if there is any.
func (c *remoteIstios) Create(remoteIstio *v1beta1.RemoteIstio) (result *v1beta1.RemoteIstio, err error) {
	result = &v1beta1.RemoteIstio{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("remoteistios").
		Body(remoteIstio).
		Do().
		Into(result)
	return
}

// Update takes the representation of a remoteIstio and updates it. Returns the server's representation of the remoteIstio, and an error, if there is any.
func (c *remoteIstios) Update(remoteIstio *v1beta1.RemoteIstio) (result *v1beta1.RemoteIstio, err error) {
	result = &v1beta1.RemoteIstio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("remoteistios").
		Name(remoteIstio.Name).
		Body(remoteIstio).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *remoteIstios) UpdateStatus(remoteIstio *v1beta1.RemoteIstio) (result *v1beta1.RemoteIstio, err error) {
	result = &v1beta1.RemoteIstio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("remoteistios").
		Name(remoteIstio.Name).
		SubResource("status").
		Body(remoteIstio).
		Do().
		Into(result)
	return
}

// Delete takes name of the remoteIstio and deletes it. Returns an error if one occurs.
func (c *remoteIstios) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("remoteistios").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *remoteIstios) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("remoteistios").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched remoteIstio.
func (c *remoteIstios) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RemoteIstio, err error) {
	result = &v1beta1.RemoteIstio{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("remoteistios").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
