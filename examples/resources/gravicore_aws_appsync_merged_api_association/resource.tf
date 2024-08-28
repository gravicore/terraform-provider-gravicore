resource "gravicore_aws_appsync_graphql_api" "default" {
  name                = "provider-test-1"
  authentication_type = "AMAZON_COGNITO_USER_POOLS"
  api_type            = "MERGED"

  user_pool_config {
    user_pool_id   = "us-east-1_test"
    aws_region     = "us-east-1"
    default_action = "ALLOW"
  }

  merged_api_execution_role_arn = "arn:aws:iam::123456:role/test"
  xray_enabled                  = true

  enhanced_metrics_config {
    resolver_level_metrics_behavior    = "PER_RESOLVER_METRICS"
    data_source_level_metrics_behavior = "PER_DATA_SOURCE_METRICS"
    operation_level_metrics_config     = "ENABLED"
  }

  log_config {
    field_log_level          = "NONE"
    cloudwatch_logs_role_arn = "arn:aws:iam::123456:role/test"
    exclude_verbose_content  = true
  }

  tags = {
    Name  = "provider-test-1"
    Test1 = "test1"
    Test2 = "test2"
    Test3 = "test3"
  }
}


resource "aws_appsync_graphql_api" "default" {
  authentication_type = "AMAZON_COGNITO_USER_POOLS"
  name                = "provider-test-2"

  user_pool_config {
    user_pool_id   = "us-east-1_test"
    aws_region     = "us-east-1"
    default_action = "ALLOW"
  }

  schema = <<EOF
schema {
    query: Query
}
type Query {
  test1: Int
}
EOF

  tags = {
    Name = "example"
    Test = "test"
  }
}

resource "gravicore_aws_appsync_merged_api_association" "default" {
  description   = "This is the default association"
  merged_api_id = gravicore_aws_appsync_graphql_api.default.id
  source_api_id = aws_appsync_graphql_api.default.id
  source_api_association_config {
    merge_type = "MANUAL_MERGE"
  }
}
