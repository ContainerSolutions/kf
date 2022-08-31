// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	kfsystemv1alpha1 "kf-operator/pkg/apis/kfsystem/v1alpha1"
	versioned "kf-operator/pkg/client/clientset/versioned"
	internalinterfaces "kf-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kf-operator/pkg/client/listers/kfsystem/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KfSystemInformer provides access to a shared informer and lister for
// KfSystems.
type KfSystemInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KfSystemLister
}

type kfSystemInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKfSystemInformer constructs a new informer for KfSystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKfSystemInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKfSystemInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKfSystemInformer constructs a new informer for KfSystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKfSystemInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KfV1alpha1().KfSystems().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KfV1alpha1().KfSystems().Watch(context.TODO(), options)
			},
		},
		&kfsystemv1alpha1.KfSystem{},
		resyncPeriod,
		indexers,
	)
}

func (f *kfSystemInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKfSystemInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kfSystemInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kfsystemv1alpha1.KfSystem{}, f.defaultInformer)
}

func (f *kfSystemInformer) Lister() v1alpha1.KfSystemLister {
	return v1alpha1.NewKfSystemLister(f.Informer().GetIndexer())
}