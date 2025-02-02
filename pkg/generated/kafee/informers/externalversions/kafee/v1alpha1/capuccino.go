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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	kafeev1alpha1 "k8s.io/sample-apiserver/pkg/apis/kafee/v1alpha1"
	versioned "k8s.io/sample-apiserver/pkg/generated/kafee/clientset/versioned"
	internalinterfaces "k8s.io/sample-apiserver/pkg/generated/kafee/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/sample-apiserver/pkg/generated/kafee/listers/kafee/v1alpha1"
)

// CapuccinoInformer provides access to a shared informer and lister for
// Capuccinos.
type CapuccinoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CapuccinoLister
}

type capuccinoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCapuccinoInformer constructs a new informer for Capuccino type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCapuccinoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCapuccinoInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCapuccinoInformer constructs a new informer for Capuccino type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCapuccinoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafeeV1alpha1().Capuccinos().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafeeV1alpha1().Capuccinos().Watch(context.TODO(), options)
			},
		},
		&kafeev1alpha1.Capuccino{},
		resyncPeriod,
		indexers,
	)
}

func (f *capuccinoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCapuccinoInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *capuccinoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kafeev1alpha1.Capuccino{}, f.defaultInformer)
}

func (f *capuccinoInformer) Lister() v1alpha1.CapuccinoLister {
	return v1alpha1.NewCapuccinoLister(f.Informer().GetIndexer())
}
