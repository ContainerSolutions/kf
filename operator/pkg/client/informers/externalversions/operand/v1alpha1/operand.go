// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	operandv1alpha1 "kf-operator/pkg/apis/operand/v1alpha1"
	versioned "kf-operator/pkg/client/clientset/versioned"
	internalinterfaces "kf-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kf-operator/pkg/client/listers/operand/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OperandInformer provides access to a shared informer and lister for
// Operands.
type OperandInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OperandLister
}

type operandInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOperandInformer constructs a new informer for Operand type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOperandInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOperandInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOperandInformer constructs a new informer for Operand type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOperandInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperandV1alpha1().Operands().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperandV1alpha1().Operands().Watch(context.TODO(), options)
			},
		},
		&operandv1alpha1.Operand{},
		resyncPeriod,
		indexers,
	)
}

func (f *operandInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOperandInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *operandInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operandv1alpha1.Operand{}, f.defaultInformer)
}

func (f *operandInformer) Lister() v1alpha1.OperandLister {
	return v1alpha1.NewOperandLister(f.Informer().GetIndexer())
}