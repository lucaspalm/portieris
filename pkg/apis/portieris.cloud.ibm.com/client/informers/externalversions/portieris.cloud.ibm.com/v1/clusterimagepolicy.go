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

package v1

import (
	time "time"

	versioned "github.com/IBM/portieris/pkg/apis/portieris.cloud.ibm.com/client/clientset/versioned"
	internalinterfaces "github.com/IBM/portieris/pkg/apis/portieris.cloud.ibm.com/client/informers/externalversions/internalinterfaces"
	v1 "github.com/IBM/portieris/pkg/apis/portieris.cloud.ibm.com/client/listers/portieris.cloud.ibm.com/v1"
	portieriscloudibmcomv1 "github.com/IBM/portieris/pkg/apis/portieris.cloud.ibm.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterImagePolicyInformer provides access to a shared informer and lister for
// ClusterImagePolicies.
type ClusterImagePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterImagePolicyLister
}

type clusterImagePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PortierisV1().ClusterImagePolicies().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PortierisV1().ClusterImagePolicies().Watch(options)
			},
		},
		&portieriscloudibmcomv1.ClusterImagePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterImagePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterImagePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&portieriscloudibmcomv1.ClusterImagePolicy{}, f.defaultInformer)
}

func (f *clusterImagePolicyInformer) Lister() v1.ClusterImagePolicyLister {
	return v1.NewClusterImagePolicyLister(f.Informer().GetIndexer())
}