package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGravicoreAwsAppsyncStartSchemaMerge() *schema.Resource {
	return &schema.Resource{
		Create: resourceGravicoreAwsAppsyncStartSchemaMergeCreate,
		Read:   schema.Noop,
		Update: resourceGravicoreAwsAppsyncStartSchemaMergeUpdate,
		Delete: schema.Noop,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identifier of the schema association.",
			},
			"merged_api_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identifier of the merged API.",
			},
			"timeout_seconds": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "The number of seconds to wait for the schema merge to complete. ",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the schema merge.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the schema merge.",
			},
		},
	}
}

func resourceGravicoreAwsAppsyncStartSchemaMergeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	associationId := d.Get("association_id").(string)
	mergedApiId := d.Get("merged_api_id").(string)

	input := &appsync.StartSchemaMergeInput{
		AssociationId:       aws.String(associationId),
		MergedApiIdentifier: aws.String(mergedApiId),
	}

	_, err := client.AppSync.StartSchemaMerge(input)
	if err != nil {
		return err
	}

	status, err := wait("MERGE_SUCCESS", d, meta)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", associationId, mergedApiId))
	d.Set("status", status)

	return nil
}

func resourceGravicoreAwsAppsyncStartSchemaMergeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceGravicoreAwsAppsyncStartSchemaMergeCreate(d, meta)
}

func wait(expectedStatus string, d *schema.ResourceData, meta interface{}) (string, error) {
	associationId := d.Get("association_id").(string)
	mergedApiId := d.Get("merged_api_id").(string)
	timeoutSeconds := d.Get("timeout_seconds").(int)

	client := meta.(*AWSClient)
	for i := 0; i < timeoutSeconds; i++ {
		input := &appsync.GetSourceApiAssociationInput{
			AssociationId:       aws.String(associationId),
			MergedApiIdentifier: aws.String(mergedApiId),
		}

		result, err := client.AppSync.GetSourceApiAssociation(input)
		if err != nil {
			return "", err
		}

		status := aws.StringValue(result.SourceApiAssociation.SourceApiAssociationStatus)
		if status == expectedStatus {
			return status, nil
		}

		time.Sleep(1 * time.Second)
	}

	return "", fmt.Errorf("schema merge did not succeed after %d attempts", timeoutSeconds)
}
