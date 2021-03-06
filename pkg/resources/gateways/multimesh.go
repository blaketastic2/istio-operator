/*
Copyright 2019 Banzai Cloud.

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

package gateways

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/banzaicloud/istio-operator/pkg/k8sutil"
	"github.com/banzaicloud/istio-operator/pkg/util"
)

const (
	multimeshResourceNamePrefix = "istio-multicluster"
)

func (r *Reconciler) multimeshEgressGateway() *k8sutil.DynamicObject {
	return &k8sutil.DynamicObject{
		Gvr: schema.GroupVersionResource{
			Group:    "networking.istio.io",
			Version:  "v1alpha3",
			Resource: "gateways",
		},
		Kind:      "Gateway",
		Name:      multimeshResourceNamePrefix + "-egressgateway",
		Namespace: r.Config.Namespace,
		Spec: map[string]interface{}{
			"servers": []map[string]interface{}{
				{
					"hosts": util.EmptyTypedStrSlice("*.global"),
					"port": map[string]interface{}{
						"name":     "tls",
						"protocol": "TLS",
						"number":   15443,
					},
					"tls": map[string]interface{}{
						"mode": "AUTO_PASSTHROUGH",
					},
				},
			},
			"selector": map[string]interface{}{
				"istio": "egressgateway",
			},
		},
		Owner: r.Config,
	}
}

func (r *Reconciler) multimeshIngressGateway() *k8sutil.DynamicObject {
	return &k8sutil.DynamicObject{
		Gvr: schema.GroupVersionResource{
			Group:    "networking.istio.io",
			Version:  "v1alpha3",
			Resource: "gateways",
		},
		Kind:      "Gateway",
		Name:      multimeshResourceNamePrefix + "-ingressgateway",
		Namespace: r.Config.Namespace,
		Spec: map[string]interface{}{
			"servers": []map[string]interface{}{
				{
					"hosts": util.EmptyTypedStrSlice("*.global"),
					"port": map[string]interface{}{
						"name":     "tls",
						"protocol": "TLS",
						"number":   15443,
					},
					"tls": map[string]interface{}{
						"mode": "AUTO_PASSTHROUGH",
					},
				},
			},
			"selector": map[string]interface{}{
				"istio": "ingressgateway",
			},
		},
		Owner: r.Config,
	}
}

func (r *Reconciler) multimeshEnvoyFilter() *k8sutil.DynamicObject {
	return &k8sutil.DynamicObject{
		Gvr: schema.GroupVersionResource{
			Group:    "networking.istio.io",
			Version:  "v1alpha3",
			Resource: "envoyfilters",
		},
		Kind:      "EnvoyFilter",
		Name:      multimeshResourceNamePrefix + "-ingressgateway",
		Namespace: r.Config.Namespace,
		Spec: map[string]interface{}{
			"workloadLabels": map[string]interface{}{
				"istio": "ingressgateway",
			},
			"filters": []map[string]interface{}{
				{
					"listenerMatch": map[string]interface{}{
						"portNumber":   15443,
						"listenerType": "GATEWAY",
					},
					"insertPosition": map[string]interface{}{
						"index":      "AFTER",
						"relativeTo": "envoy.filters.network.sni_cluster",
					},
					"filterName": "envoy.filters.network.tcp_cluster_rewrite",
					"filterType": "NETWORK",
					"filterConfig": map[string]interface{}{
						"cluster_pattern":     "\\.global$",
						"cluster_replacement": ".svc." + r.Config.Spec.Proxy.ClusterDomain,
					},
				},
			},
		},
		Owner: r.Config,
	}
}

func (r *Reconciler) multimeshDestinationRule() *k8sutil.DynamicObject {
	return &k8sutil.DynamicObject{
		Gvr: schema.GroupVersionResource{
			Group:    "networking.istio.io",
			Version:  "v1alpha3",
			Resource: "destinationrules",
		},
		Kind:      "DestinationRule",
		Name:      multimeshResourceNamePrefix + "-destinationrule",
		Namespace: r.Config.Namespace,
		Spec: map[string]interface{}{
			"host": "*.global",
			"trafficPolicy": map[string]interface{}{
				"tls": map[string]interface{}{
					"mode": "ISTIO_MUTUAL",
				},
			},
		},
		Owner: r.Config,
	}
}
