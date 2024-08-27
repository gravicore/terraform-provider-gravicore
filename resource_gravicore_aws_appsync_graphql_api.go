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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"lambda_authorizer_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorizer_result_ttl_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"authorizer_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"identity_validation_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"openid_connect_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_ttl": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"client_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"iat_ttl": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"issuer": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"user_pool_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_id_client_regex": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"aws_region": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"user_pool_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"api_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "MERGED",
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enhanced_metrics_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_level_metrics_behavior": {
							Type:     schema.TypeString,
							Required: true,
						},
						"operation_level_metrics_config": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resolver_level_metrics_behavior": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"introspection_config": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"lambda_authorizer_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorizer_result_ttl_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"authorizer_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"identity_validation_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"log_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_logs_role_arn": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude_verbose_content": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"field_log_level": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"merged_api_execution_role_arn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"openid_connect_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iat_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"owner_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_depth_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resolver_count_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"user_pool_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_id_client_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"aws_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_pool_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"visibility": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GLOBAL",
				ForceNew: true,
			},
			"xray_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"schema": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
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
