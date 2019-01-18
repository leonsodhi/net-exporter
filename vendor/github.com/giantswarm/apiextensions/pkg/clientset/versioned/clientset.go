/*
Copyright 2019 Giant Swarm GmbH.

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

package versioned

import (
	applicationv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/application/v1alpha1"
	corev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/core/v1alpha1"
	examplev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/example/v1alpha1"
	providerv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/provider/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Application() applicationv1alpha1.ApplicationV1alpha1Interface
	CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Core() corev1alpha1.CoreV1alpha1Interface
	ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Example() examplev1alpha1.ExampleV1alpha1Interface
	ProviderV1alpha1() providerv1alpha1.ProviderV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Provider() providerv1alpha1.ProviderV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	applicationV1alpha1 *applicationv1alpha1.ApplicationV1alpha1Client
	coreV1alpha1        *corev1alpha1.CoreV1alpha1Client
	exampleV1alpha1     *examplev1alpha1.ExampleV1alpha1Client
	providerV1alpha1    *providerv1alpha1.ProviderV1alpha1Client
}

// ApplicationV1alpha1 retrieves the ApplicationV1alpha1Client
func (c *Clientset) ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface {
	return c.applicationV1alpha1
}

// Deprecated: Application retrieves the default version of ApplicationClient.
// Please explicitly pick a version.
func (c *Clientset) Application() applicationv1alpha1.ApplicationV1alpha1Interface {
	return c.applicationV1alpha1
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// Deprecated: Core retrieves the default version of CoreClient.
// Please explicitly pick a version.
func (c *Clientset) Core() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1Client
func (c *Clientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface {
	return c.exampleV1alpha1
}

// Deprecated: Example retrieves the default version of ExampleClient.
// Please explicitly pick a version.
func (c *Clientset) Example() examplev1alpha1.ExampleV1alpha1Interface {
	return c.exampleV1alpha1
}

// ProviderV1alpha1 retrieves the ProviderV1alpha1Client
func (c *Clientset) ProviderV1alpha1() providerv1alpha1.ProviderV1alpha1Interface {
	return c.providerV1alpha1
}

// Deprecated: Provider retrieves the default version of ProviderClient.
// Please explicitly pick a version.
func (c *Clientset) Provider() providerv1alpha1.ProviderV1alpha1Interface {
	return c.providerV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.applicationV1alpha1, err = applicationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1alpha1, err = corev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.exampleV1alpha1, err = examplev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.providerV1alpha1, err = providerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.applicationV1alpha1 = applicationv1alpha1.NewForConfigOrDie(c)
	cs.coreV1alpha1 = corev1alpha1.NewForConfigOrDie(c)
	cs.exampleV1alpha1 = examplev1alpha1.NewForConfigOrDie(c)
	cs.providerV1alpha1 = providerv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.applicationV1alpha1 = applicationv1alpha1.New(c)
	cs.coreV1alpha1 = corev1alpha1.New(c)
	cs.exampleV1alpha1 = examplev1alpha1.New(c)
	cs.providerV1alpha1 = providerv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}