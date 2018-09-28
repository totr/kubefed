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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedServiceAccounts implements FederatedServiceAccountInterface
type FakeFederatedServiceAccounts struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var federatedserviceaccountsResource = schema.GroupVersionResource{Group: "core.federation.k8s.io", Version: "v1alpha1", Resource: "federatedserviceaccounts"}

var federatedserviceaccountsKind = schema.GroupVersionKind{Group: "core.federation.k8s.io", Version: "v1alpha1", Kind: "FederatedServiceAccount"}

// Get takes name of the federatedServiceAccount, and returns the corresponding federatedServiceAccount object, and an error if there is any.
func (c *FakeFederatedServiceAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedserviceaccountsResource, c.ns, name), &v1alpha1.FederatedServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccount), err
}

// List takes label and field selectors, and returns the list of FederatedServiceAccounts that match those selectors.
func (c *FakeFederatedServiceAccounts) List(opts v1.ListOptions) (result *v1alpha1.FederatedServiceAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedserviceaccountsResource, federatedserviceaccountsKind, c.ns, opts), &v1alpha1.FederatedServiceAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedServiceAccountList{ListMeta: obj.(*v1alpha1.FederatedServiceAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.FederatedServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedServiceAccounts.
func (c *FakeFederatedServiceAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedserviceaccountsResource, c.ns, opts))

}

// Create takes the representation of a federatedServiceAccount and creates it.  Returns the server's representation of the federatedServiceAccount, and an error, if there is any.
func (c *FakeFederatedServiceAccounts) Create(federatedServiceAccount *v1alpha1.FederatedServiceAccount) (result *v1alpha1.FederatedServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedserviceaccountsResource, c.ns, federatedServiceAccount), &v1alpha1.FederatedServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccount), err
}

// Update takes the representation of a federatedServiceAccount and updates it. Returns the server's representation of the federatedServiceAccount, and an error, if there is any.
func (c *FakeFederatedServiceAccounts) Update(federatedServiceAccount *v1alpha1.FederatedServiceAccount) (result *v1alpha1.FederatedServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedserviceaccountsResource, c.ns, federatedServiceAccount), &v1alpha1.FederatedServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedServiceAccounts) UpdateStatus(federatedServiceAccount *v1alpha1.FederatedServiceAccount) (*v1alpha1.FederatedServiceAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedserviceaccountsResource, "status", c.ns, federatedServiceAccount), &v1alpha1.FederatedServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccount), err
}

// Delete takes name of the federatedServiceAccount and deletes it. Returns an error if one occurs.
func (c *FakeFederatedServiceAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedserviceaccountsResource, c.ns, name), &v1alpha1.FederatedServiceAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedServiceAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedserviceaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedServiceAccountList{})
	return err
}

// Patch applies the patch and returns the patched federatedServiceAccount.
func (c *FakeFederatedServiceAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedserviceaccountsResource, c.ns, name, data, subresources...), &v1alpha1.FederatedServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccount), err
}