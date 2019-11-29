/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platform "tkestack.io/tke/api/platform"
)

// FakeHelms implements HelmInterface
type FakeHelms struct {
	Fake *FakePlatform
}

var helmsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "", Resource: "helms"}

var helmsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "", Kind: "Helm"}

// Get takes name of the helm, and returns the corresponding helm object, and an error if there is any.
func (c *FakeHelms) Get(name string, options v1.GetOptions) (result *platform.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(helmsResource, name), &platform.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Helm), err
}

// List takes label and field selectors, and returns the list of Helms that match those selectors.
func (c *FakeHelms) List(opts v1.ListOptions) (result *platform.HelmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(helmsResource, helmsKind, opts), &platform.HelmList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platform.HelmList{ListMeta: obj.(*platform.HelmList).ListMeta}
	for _, item := range obj.(*platform.HelmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helms.
func (c *FakeHelms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(helmsResource, opts))
}

// Create takes the representation of a helm and creates it.  Returns the server's representation of the helm, and an error, if there is any.
func (c *FakeHelms) Create(helm *platform.Helm) (result *platform.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(helmsResource, helm), &platform.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Helm), err
}

// Update takes the representation of a helm and updates it. Returns the server's representation of the helm, and an error, if there is any.
func (c *FakeHelms) Update(helm *platform.Helm) (result *platform.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(helmsResource, helm), &platform.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Helm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelms) UpdateStatus(helm *platform.Helm) (*platform.Helm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(helmsResource, "status", helm), &platform.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Helm), err
}

// Delete takes name of the helm and deletes it. Returns an error if one occurs.
func (c *FakeHelms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(helmsResource, name), &platform.Helm{})
	return err
}

// Patch applies the patch and returns the patched helm.
func (c *FakeHelms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platform.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(helmsResource, name, pt, data, subresources...), &platform.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.Helm), err
}