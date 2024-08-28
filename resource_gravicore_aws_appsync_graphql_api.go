package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGravicoreAwsAppsyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceGravicoreAwsAppsyncGraphQLApiCreate,
		Read:   resourceGravicoreAwsAppsyncGraphQLApiRead,
		Update: resourceGravicoreAwsAppsyncGraphQLApiUpdate,
		Delete: resourceGravicoreAwsAppsyncGraphQLApiDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"additional_authentication_providers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of additional authentication providers for the GraphqlApi API.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The authentication type for the additional authentication provider.",
						},
						"lambda_authorizer_config": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The configuration for the Lambda authorizer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorizer_result_ttl_seconds": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "The TTL for the authorizer result in seconds.",
									},
									"authorizer_uri": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The URI of the authorizer Lambda function.",
									},
									"identity_validation_expression": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The validation expression for the identity.",
									},
								},
							},
						},
						"openid_connect_config": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The configuration for the OpenID Connect provider.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The TTL for the authentication token in seconds.",
									},
									"client_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The client ID for the OpenID Connect provider.",
									},
									"iat_ttl": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The TTL for the IAT token in seconds.",
									},
									"issuer": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The issuer URL for the OpenID Connect provider.",
									},
								},
							},
						},
						"user_pool_config": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The configuration for the user pool.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_id_client_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regular expression for validating the app client ID.",
									},
									"aws_region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The AWS region for the user pool.",
									},
									"user_pool_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The ID of the user pool.",
									},
								},
							},
						},
					},
				},
			},
			"api_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "MERGED",
				Description: "The type of the GraphQL API.",
			},
			"authentication_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The authentication type for the GraphQL API.",
			},
			"enhanced_metrics_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The configuration for the enhanced metrics.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_level_metrics_behavior": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The behavior for data source level metrics.",
						},
						"operation_level_metrics_config": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The configuration for operation level metrics.",
						},
						"resolver_level_metrics_behavior": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The behavior for resolver level metrics.",
						},
					},
				},
			},
			"introspection_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ENABLED",
				Description: "The configuration for the introspection.",
			},
			"lambda_authorizer_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The configuration for the Lambda authorizer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorizer_result_ttl_seconds": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The TTL for the authorizer result in seconds.",
						},
						"authorizer_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URI of the authorizer Lambda function.",
						},
						"identity_validation_expression": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The validation expression for the identity.",
						},
					},
				},
			},
			"log_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The configuration for the log.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_logs_role_arn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ARN of the CloudWatch Logs role.",
						},
						"exclude_verbose_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether to exclude verbose content from the log.",
						},
						"field_log_level": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The log level for the field.",
						},
					},
				},
			},
			"merged_api_execution_role_arn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ARN of the merged API execution role.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the GraphQL API.",
			},
			"openid_connect_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The configuration for the OpenID Connect provider.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The TTL for the authentication token in seconds.",
						},
						"client_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The client ID for the OpenID Connect provider.",
						},
						"iat_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The TTL for the IAT token in seconds.",
						},
						"issuer": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The issuer URL for the OpenID Connect provider.",
						},
					},
				},
			},
			"owner_contact": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The contact information for the owner.",
			},
			"query_depth_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The depth limit for the query.",
			},
			"resolver_count_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The count limit for the resolver.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the GraphQL API.",
			},
			"user_pool_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The configuration for the user pool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_id_client_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The regular expression for validating the app client ID.",
						},
						"aws_region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The AWS region for the user pool.",
						},
						"default_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The default action for the user pool.",
						},
						"user_pool_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the user pool.",
						},
					},
				},
			},
			"visibility": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "GLOBAL",
				ForceNew:    true,
				Description: "The visibility of the GraphQL API.",
			},
			"xray_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether X-Ray tracing is enabled for the GraphQL API.",
			},
			"schema": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The schema definition for the GraphQL API.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the GraphQL API.",
			},
			"arn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ARN of the GraphQL API.",
			},
		},
	}
}

func resourceGravicoreAwsAppsyncGraphQLApiCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.CreateGraphqlApiInput{
		AdditionalAuthenticationProviders: expandAdditionalAuthenticationProviders(d.Get("additional_authentication_providers").([]interface{})),
		ApiType:                           aws.String(d.Get("api_type").(string)),
		AuthenticationType:                aws.String(d.Get("authentication_type").(string)),
		EnhancedMetricsConfig:             expandEnhancedMetricsConfig(d.Get("enhanced_metrics_config").([]interface{})),
		IntrospectionConfig:               aws.String(d.Get("introspection_config").(string)),
		LambdaAuthorizerConfig:            expandLambdaAuthorizerConfig(d.Get("lambda_authorizer_config").([]interface{})),
		LogConfig:                         expandLogConfig(d.Get("log_config").([]interface{})),
		MergedApiExecutionRoleArn:         aws.String(d.Get("merged_api_execution_role_arn").(string)),
		Name:                              aws.String(d.Get("name").(string)),
		OpenIDConnectConfig:               expandOpenIDConnectConfig(d.Get("openid_connect_config").([]interface{})),
		OwnerContact:                      aws.String(d.Get("owner_contact").(string)),
		QueryDepthLimit:                   aws.Int64(int64(d.Get("query_depth_limit").(int))),
		ResolverCountLimit:                aws.Int64(int64(d.Get("resolver_count_limit").(int))),
		Tags:                              expandTags(d.Get("tags").(map[string]interface{})),
		UserPoolConfig:                    expandUserPoolConfig(d.Get("user_pool_config").([]interface{})),
		Visibility:                        aws.String(d.Get("visibility").(string)),
		XrayEnabled:                       aws.Bool(d.Get("xray_enabled").(bool)),
	}

	result, err := client.AppSync.CreateGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId(aws.StringValue(result.GraphqlApi.ApiId))

	return resourceGravicoreAwsAppsyncGraphQLApiRead(d, meta)
}

func resourceGravicoreAwsAppsyncGraphQLApiRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.GetGraphqlApiInput{
		ApiId: aws.String(d.Id()),
	}
	result, err := client.AppSync.GetGraphqlApi(input)
	if err != nil {
		return err
	}

	d.Set("additional_authentication_providers", flattenAdditionalAuthenticationProviders(result.GraphqlApi.AdditionalAuthenticationProviders))
	d.Set("api_type", result.GraphqlApi.ApiType)
	d.Set("authentication_type", result.GraphqlApi.AuthenticationType)
	d.Set("enhanced_metrics_config", flattenEnhancedMetricsConfig(result.GraphqlApi.EnhancedMetricsConfig))
	d.Set("introspection_config", result.GraphqlApi.IntrospectionConfig)
	d.Set("lambda_authorizer_config", flattenLambdaAuthorizerConfig(result.GraphqlApi.LambdaAuthorizerConfig))
	d.Set("log_config", flattenLogConfig(result.GraphqlApi.LogConfig))
	d.Set("merged_api_execution_role_arn", result.GraphqlApi.MergedApiExecutionRoleArn)
	d.Set("name", result.GraphqlApi.Name)
	d.Set("open_id_connect_config", flattenOpenIDConnectConfig(result.GraphqlApi.OpenIDConnectConfig))
	d.Set("owner_contact", result.GraphqlApi.OwnerContact)
	d.Set("query_depth_limit", result.GraphqlApi.QueryDepthLimit)
	d.Set("tags", aws.StringValueMap(result.GraphqlApi.Tags))
	d.Set("user_pool_config", flattenUserPoolConfig(result.GraphqlApi.UserPoolConfig))
	d.Set("visibility", result.GraphqlApi.Visibility)
	d.Set("xray_enabled", result.GraphqlApi.XrayEnabled)
	d.Set("arn", aws.StringValue(result.GraphqlApi.Arn))

	return nil
}

func resourceGravicoreAwsAppsyncGraphQLApiUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.UpdateGraphqlApiInput{
		ApiId:              aws.String(d.Id()),
		AuthenticationType: aws.String(d.Get("authentication_type").(string)),
		Name:               aws.String(d.Get("name").(string)),
	}

	if d.HasChangesExcept("tags", "tags_all") {
		if v, ok := d.GetOk("additional_authentication_providers"); ok {
			input.AdditionalAuthenticationProviders = expandAdditionalAuthenticationProviders(v.([]interface{}))
		}

		if v, ok := d.GetOk("enhanced_metrics_config"); ok {
			input.EnhancedMetricsConfig = expandEnhancedMetricsConfig(v.([]interface{}))
		}

		if v, ok := d.GetOk("introspection_config"); ok {
			input.IntrospectionConfig = aws.String(v.(string))
		}

		if v, ok := d.GetOk("lambda_authorizer_config"); ok {
			input.LambdaAuthorizerConfig = expandLambdaAuthorizerConfig(v.([]interface{}))
		}

		if v, ok := d.GetOk("log_config"); ok {
			input.LogConfig = expandLogConfig(v.([]interface{}))
		}

		if v, ok := d.GetOk("merged_api_execution_role_arn"); ok {
			input.MergedApiExecutionRoleArn = aws.String(v.(string))
		}

		if v, ok := d.GetOk("open_id_connect_config"); ok {
			input.OpenIDConnectConfig = expandOpenIDConnectConfig(v.([]interface{}))
		}

		if v, ok := d.GetOk("owner_contact"); ok {
			input.OwnerContact = aws.String(v.(string))
		}

		if v, ok := d.GetOk("query_depth_limit"); ok {
			input.QueryDepthLimit = aws.Int64(int64(v.(int)))
		}

		if v, ok := d.GetOk("resolver_count_limit"); ok {
			input.ResolverCountLimit = aws.Int64(int64(v.(int)))
		}

		if v, ok := d.GetOk("user_pool_config"); ok {
			input.UserPoolConfig = expandUserPoolConfig(v.([]interface{}))
		}

		if v, ok := d.GetOk("xray_enabled"); ok {
			input.XrayEnabled = aws.Bool(v.(bool))
		}

		_, err := client.AppSync.UpdateGraphqlApi(input)
		if err != nil {
			return err
		}
	}

	if d.HasChange("tags") {
		err := updateTags(d, meta)
		if err != nil {
			return err
		}
	}

	return resourceGravicoreAwsAppsyncGraphQLApiRead(d, meta)
}

func resourceGravicoreAwsAppsyncGraphQLApiDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.DeleteGraphqlApiInput{
		ApiId: aws.String(d.Id()),
	}

	_, err := client.AppSync.DeleteGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
