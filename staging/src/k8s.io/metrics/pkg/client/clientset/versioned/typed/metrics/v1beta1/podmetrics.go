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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	scheme "k8s.io/metrics/pkg/client/clientset/versioned/scheme"
)

// PodMetricsesGetter has a method to return a PodMetricsInterface.
// A group's client should implement this interface.
type PodMetricsesGetter interface {
	PodMetricses(namespace string) PodMetricsInterface
}

// PodMetricsInterface has methods to work with PodMetrics resources.
type PodMetricsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*metricsv1beta1.PodMetrics, error)
	List(ctx context.Context, opts v1.ListOptions) (*metricsv1beta1.PodMetricsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	PodMetricsExpansion
}

// podMetricses implements PodMetricsInterface
type podMetricses struct {
	*gentype.ClientWithList[*metricsv1beta1.PodMetrics, *metricsv1beta1.PodMetricsList]
}

// newPodMetricses returns a PodMetricses
func newPodMetricses(c *MetricsV1beta1Client, namespace string) *podMetricses {
	return &podMetricses{
		gentype.NewClientWithList[*metricsv1beta1.PodMetrics, *metricsv1beta1.PodMetricsList](
			"pods",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *metricsv1beta1.PodMetrics { return &metricsv1beta1.PodMetrics{} },
			func() *metricsv1beta1.PodMetricsList { return &metricsv1beta1.PodMetricsList{} },
			gentype.PrefersProtobuf[*metricsv1beta1.PodMetrics](),
		),
	}
}
