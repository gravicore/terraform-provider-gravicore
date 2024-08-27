package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					"region": {
						Type:        schema.TypeString,
						Required:    true,
						DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", nil),
						Description: "The region where AWS operations will take place.",
					},
				},
				ResourcesMap: map[string]*schema.Resource{
					"gravicore_aws_appsync_graphql_api":            resourceGravicoreAwsAppsyncGraphQLApi(),
					"gravicore_aws_appsync_merged_api_association": resourceGravicoreAwsAppsyncMergedApiAssociation(),
					"gravicore_aws_appsync_start_schema_merge":     resourceGravicoreAwsAppsyncStartSchemaMerge(),
				},
				ConfigureFunc: configureFunc,
			}
		},
	})
}
