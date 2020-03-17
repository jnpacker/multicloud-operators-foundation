// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/open-cluster-management/multicloud-operators-foundation/pkg/client/cluster_clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
)

type ClusterregistryV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
}

// ClusterregistryV1alpha1Client is used to interact with features provided by the clusterregistry.k8s.io group.
type ClusterregistryV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ClusterregistryV1alpha1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

// NewForConfig creates a new ClusterregistryV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ClusterregistryV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterregistryV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterregistryV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterregistryV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterregistryV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ClusterregistryV1alpha1Client {
	return &ClusterregistryV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ClusterregistryV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
