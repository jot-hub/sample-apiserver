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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/kafee/v1alpha1"
)

// EspressoLister helps list Espressos.
// All objects returned here must be treated as read-only.
type EspressoLister interface {
	// List lists all Espressos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Espresso, err error)
	// Espressos returns an object that can list and get Espressos.
	Espressos(namespace string) EspressoNamespaceLister
	EspressoListerExpansion
}

// espressoLister implements the EspressoLister interface.
type espressoLister struct {
	indexer cache.Indexer
}

// NewEspressoLister returns a new EspressoLister.
func NewEspressoLister(indexer cache.Indexer) EspressoLister {
	return &espressoLister{indexer: indexer}
}

// List lists all Espressos in the indexer.
func (s *espressoLister) List(selector labels.Selector) (ret []*v1alpha1.Espresso, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Espresso))
	})
	return ret, err
}

// Espressos returns an object that can list and get Espressos.
func (s *espressoLister) Espressos(namespace string) EspressoNamespaceLister {
	return espressoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EspressoNamespaceLister helps list and get Espressos.
// All objects returned here must be treated as read-only.
type EspressoNamespaceLister interface {
	// List lists all Espressos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Espresso, err error)
	// Get retrieves the Espresso from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Espresso, error)
	EspressoNamespaceListerExpansion
}

// espressoNamespaceLister implements the EspressoNamespaceLister
// interface.
type espressoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Espressos in the indexer for a given namespace.
func (s espressoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Espresso, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Espresso))
	})
	return ret, err
}

// Get retrieves the Espresso from the indexer for a given namespace and name.
func (s espressoNamespaceLister) Get(name string) (*v1alpha1.Espresso, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("espresso"), name)
	}
	return obj.(*v1alpha1.Espresso), nil
}
