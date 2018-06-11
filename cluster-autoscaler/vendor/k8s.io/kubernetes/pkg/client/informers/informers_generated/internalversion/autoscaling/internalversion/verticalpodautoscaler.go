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

package internalversion

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/autoscaling/internalversion"
)

// VerticalPodAutoscalerInformer provides access to a shared informer and lister for
// VerticalPodAutoscalers.
type VerticalPodAutoscalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.VerticalPodAutoscalerLister
}

type verticalPodAutoscalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVerticalPodAutoscalerInformer constructs a new informer for VerticalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVerticalPodAutoscalerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVerticalPodAutoscalerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVerticalPodAutoscalerInformer constructs a new informer for VerticalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVerticalPodAutoscalerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Autoscaling().VerticalPodAutoscalers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Autoscaling().VerticalPodAutoscalers(namespace).Watch(options)
			},
		},
		&autoscaling.VerticalPodAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *verticalPodAutoscalerInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVerticalPodAutoscalerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *verticalPodAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscaling.VerticalPodAutoscaler{}, f.defaultInformer)
}

func (f *verticalPodAutoscalerInformer) Lister() internalversion.VerticalPodAutoscalerLister {
	return internalversion.NewVerticalPodAutoscalerLister(f.Informer().GetIndexer())
}
