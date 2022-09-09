// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubean.io/api/generated/kubeancluster/clientset/versioned/typed/kubeancluster/v1alpha1"
)

type FakeKubeanclusterV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKubeanclusterV1alpha1) KuBeanClusters() v1alpha1.KuBeanClusterInterface {
	return &FakeKuBeanClusters{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubeanclusterV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
