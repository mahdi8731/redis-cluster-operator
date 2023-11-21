//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisCluster":       schema_pkg_apis_redis_v1alpha1_DistributedRedisCluster(ref),
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterSpec":   schema_pkg_apis_redis_v1alpha1_DistributedRedisClusterSpec(ref),
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterStatus": schema_pkg_apis_redis_v1alpha1_DistributedRedisClusterStatus(ref),
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackup":            schema_pkg_apis_redis_v1alpha1_RedisClusterBackup(ref),
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupSpec":        schema_pkg_apis_redis_v1alpha1_RedisClusterBackupSpec(ref),
		"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupStatus":      schema_pkg_apis_redis_v1alpha1_RedisClusterBackupStatus(ref),
	}
}

func schema_pkg_apis_redis_v1alpha1_DistributedRedisCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DistributedRedisCluster is the Schema for the distributedredisclusters API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterSpec", "github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.DistributedRedisClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redis_v1alpha1_DistributedRedisClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DistributedRedisClusterSpec defines the desired state of DistributedRedisCluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"command": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"masterSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"clusterReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"serviceName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"toleRations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PodSecurityContext"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisStorage"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"passwordSecret": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisStorage", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.PodSecurityContext", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_redis_v1alpha1_DistributedRedisClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DistributedRedisClusterStatus defines the observed state of DistributedRedisCluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"numberOfMaster": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterNode"),
									},
								},
							},
						},
					},
				},
				Required: []string{"status", "nodes"},
			},
		},
		Dependencies: []string{
			"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterNode"},
	}
}

func schema_pkg_apis_redis_v1alpha1_RedisClusterBackup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterBackup is the Schema for the redisclusterbackups API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupSpec", "github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1.RedisClusterBackupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redis_v1alpha1_RedisClusterBackupSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterBackupSpec defines the desired state of RedisClusterBackup",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_redis_v1alpha1_RedisClusterBackupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterBackupStatus defines the observed state of RedisClusterBackup",
				Type:        []string{"object"},
			},
		},
	}
}
