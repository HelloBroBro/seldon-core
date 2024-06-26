/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	"github.com/seldonio/seldon-core/operator/client/machinelearning.seldon.io/v1alpha2/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type MachinelearningV1alpha2Interface interface {
	RESTClient() rest.Interface
	SeldonDeploymentsGetter
}

// MachinelearningV1alpha2Client is used to interact with features provided by the machinelearning.seldon.io group.
type MachinelearningV1alpha2Client struct {
	restClient rest.Interface
}

func (c *MachinelearningV1alpha2Client) SeldonDeployments(namespace string) SeldonDeploymentInterface {
	return newSeldonDeployments(c, namespace)
}

// NewForConfig creates a new MachinelearningV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*MachinelearningV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MachinelearningV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new MachinelearningV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MachinelearningV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MachinelearningV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *MachinelearningV1alpha2Client {
	return &MachinelearningV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
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
func (c *MachinelearningV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
