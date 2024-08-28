# Terraform Provider for Gravicore

The Terraform provider for Gravicore allows you to manage AppSync Merged resources using Terraform. This document provides an overview of the provider's configuration and resources.

## Overview

The `gravicore` provider is currently used to interact with AWS AppSync resources. This documentation will cover:

- Provider Configuration
- Resources

## Provider Configuration

To use the `gravicore` provider, you need to specify the required configuration in your Terraform configuration file. The provider uses the same access key and secret key from the environment that the `aws` provider uses.

## Usage

- [provider](docs/index.md)
- [resource_gravicore_aws_appsync_graphql_api](docs/resources/aws_appsync_graphql_api.md)
- [resource_gravicore_aws_appsync_merged_api_association](docs/resources/aws_appsync_merged_api_association.md)
- [resource_gravicore_aws_appsync_start_schema_merge](docs/resources/aws_appsync_start_schema_merge.md)
