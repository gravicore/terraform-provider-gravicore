package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGravicoreAwsAppsyncMergedApiAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceGravicoreAwsAppsyncMergedApiAssociationCreate,
		Read:   resourceGravicoreAwsAppsyncMergedApiAssociationRead,
		Update: resourceGravicoreAwsAppsyncMergedApiAssociationUpdate,
		Delete: resourceGravicoreAwsAppsyncMergedApiAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"merged_api_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_api_association_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merge_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "AUTO_MERGE",
						},
					},
				},
			},
			"source_api_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGravicoreAwsAppsyncMergedApiAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	mergedApiIdentifier := d.Get("merged_api_id").(string)
	input := &appsync.AssociateMergedGraphqlApiInput{
		Description:                aws.String(d.Get("description").(string)),
		MergedApiIdentifier:        aws.String(mergedApiIdentifier),
		SourceApiAssociationConfig: expandSourceApiAssociationConfig(d.Get("source_api_association_config").([]interface{})),
		SourceApiIdentifier:        aws.String(d.Get("source_api_id").(string)),
	}

	result, err := client.AppSync.AssociateMergedGraphqlApi(input)
	if err != nil {
		return err
	}

	associationId := aws.StringValue(result.SourceApiAssociation.AssociationId)
	d.SetId(fmt.Sprintf("%s_%s", mergedApiIdentifier, associationId))

	return resourceGravicoreAwsAppsyncMergedApiAssociationRead(d, meta)
}

func resourceGravicoreAwsAppsyncMergedApiAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.GetSourceApiAssociationInput{
		AssociationId:       aws.String(strings.Split(d.Id(), "_")[1]),
		MergedApiIdentifier: aws.String(strings.Split(d.Id(), "_")[0]),
	}
	result, err := client.AppSync.GetSourceApiAssociation(input)
	if err != nil {
		return err
	}

	d.Set("description", result.SourceApiAssociation.Description)
	d.Set("merged_api_id", result.SourceApiAssociation.MergedApiId)
	d.Set("source_api_association_config", flattenSourceApiAssociationConfig(result.SourceApiAssociation.SourceApiAssociationConfig))
	d.Set("source_api_id", result.SourceApiAssociation.SourceApiId)
	d.Set("arn", result.SourceApiAssociation.AssociationArn)

	return nil
}

func resourceGravicoreAwsAppsyncMergedApiAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.UpdateSourceApiAssociationInput{
		AssociationId:       aws.String(d.Id()),
		MergedApiIdentifier: aws.String(d.Get("merged_api_id").(string)),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("source_api_association_config"); ok {
		input.SourceApiAssociationConfig = expandSourceApiAssociationConfig(v.([]interface{}))
	}

	_, err := client.AppSync.UpdateSourceApiAssociation(input)
	if err != nil {
		return err
	}

	return resourceGravicoreAwsAppsyncMergedApiAssociationRead(d, meta)
}

func resourceGravicoreAwsAppsyncMergedApiAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.DisassociateMergedGraphqlApiInput{
		AssociationId:       aws.String(d.Id()),
		SourceApiIdentifier: aws.String(d.Get("source_api_id").(string)),
	}

	_, err := client.AppSync.DisassociateMergedGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
