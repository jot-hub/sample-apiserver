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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/kafee/v1alpha1"
)

// FakeEspressos implements EspressoInterface
type FakeEspressos struct {
	Fake *FakeKafeeV1alpha1
	ns   string
}

var espressosResource = schema.GroupVersionResource{Group: "kafee.example.com", Version: "v1alpha1", Resource: "espressos"}

var espressosKind = schema.GroupVersionKind{Group: "kafee.example.com", Version: "v1alpha1", Kind: "Espresso"}

// Get takes name of the espresso, and returns the corresponding espresso object, and an error if there is any.
func (c *FakeEspressos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Espresso, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(espressosResource, c.ns, name), &v1alpha1.Espresso{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Espresso), err
}

// List takes label and field selectors, and returns the list of Espressos that match those selectors.
func (c *FakeEspressos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EspressoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(espressosResource, espressosKind, c.ns, opts), &v1alpha1.EspressoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EspressoList{ListMeta: obj.(*v1alpha1.EspressoList).ListMeta}
	for _, item := range obj.(*v1alpha1.EspressoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested espressos.
func (c *FakeEspressos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(espressosResource, c.ns, opts))

}

// Create takes the representation of a espresso and creates it.  Returns the server's representation of the espresso, and an error, if there is any.
func (c *FakeEspressos) Create(ctx context.Context, espresso *v1alpha1.Espresso, opts v1.CreateOptions) (result *v1alpha1.Espresso, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(espressosResource, c.ns, espresso), &v1alpha1.Espresso{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Espresso), err
}

// Update takes the representation of a espresso and updates it. Returns the server's representation of the espresso, and an error, if there is any.
func (c *FakeEspressos) Update(ctx context.Context, espresso *v1alpha1.Espresso, opts v1.UpdateOptions) (result *v1alpha1.Espresso, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(espressosResource, c.ns, espresso), &v1alpha1.Espresso{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Espresso), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEspressos) UpdateStatus(ctx context.Context, espresso *v1alpha1.Espresso, opts v1.UpdateOptions) (*v1alpha1.Espresso, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(espressosResource, "status", c.ns, espresso), &v1alpha1.Espresso{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Espresso), err
}

// Delete takes name of the espresso and deletes it. Returns an error if one occurs.
func (c *FakeEspressos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(espressosResource, c.ns, name, opts), &v1alpha1.Espresso{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEspressos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(espressosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EspressoList{})
	return err
}

// Patch applies the patch and returns the patched espresso.
func (c *FakeEspressos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Espresso, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(espressosResource, c.ns, name, pt, data, subresources...), &v1alpha1.Espresso{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Espresso), err
}
