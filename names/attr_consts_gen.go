// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by internal/generate/attrconsts/main.go; DO NOT EDIT.

package names

import (
	"fmt"
)

const (
	AttrARN                        = "arn"
	AttrARNs                       = "arns"
	AttrAccessKey                  = "access_key"
	AttrAccountID                  = "account_id"
	AttrAction                     = "action"
	AttrAddress                    = "address"
	AttrAlias                      = "alias"
	AttrApplyImmediately           = "apply_immediately"
	AttrAttributes                 = "attributes"
	AttrAutoMinorVersionUpgrade    = "auto_minor_version_upgrade"
	AttrAvailabilityZone           = "availability_zone"
	AttrAvailabilityZones          = "availability_zones"
	AttrBucket                     = "bucket"
	AttrBucketName                 = "bucket_name"
	AttrBucketPrefix               = "bucket_prefix"
	AttrCatalogID                  = "catalog_id"
	AttrCertificate                = "certificate"
	AttrCertificateARN             = "certificate_arn"
	AttrCertificateChain           = "certificate_chain"
	AttrClientID                   = "client_id"
	AttrClientSecret               = "client_secret"
	AttrClusterIdentifier          = "cluster_identifier"
	AttrClusterName                = "cluster_name"
	AttrComment                    = "comment"
	AttrCondition                  = "condition"
	AttrConfiguration              = "configuration"
	AttrContent                    = "content"
	AttrContentType                = "content_type"
	AttrCreateTime                 = "create_time"
	AttrCreatedAt                  = "created_at"
	AttrCreatedDate                = "created_date"
	AttrCreatedTime                = "created_time"
	AttrCreationDate               = "creation_date"
	AttrCreationTime               = "creation_time"
	AttrDNSName                    = "dns_name"
	AttrDatabase                   = "database"
	AttrDatabaseName               = "database_name"
	AttrDefaultValue               = "default_value"
	AttrDeleteOnTermination        = "delete_on_termination"
	AttrDescription                = "description"
	AttrDestination                = "destination"
	AttrDestinationARN             = "destination_arn"
	AttrDeviceName                 = "device_name"
	AttrDisplayName                = "display_name"
	AttrDomain                     = "domain"
	AttrDomainName                 = "domain_name"
	AttrDuration                   = "duration"
	AttrEmail                      = "email"
	AttrEnabled                    = "enabled"
	AttrEncrypted                  = "encrypted"
	AttrEncryptionConfiguration    = "encryption_configuration"
	AttrEndpoint                   = "endpoint"
	AttrEndpointType               = "endpoint_type"
	AttrEndpoints                  = "endpoints"
	AttrEngineVersion              = "engine_version"
	AttrEnvironment                = "environment"
	AttrExecutionRoleARN           = "execution_role_arn"
	AttrExpression                 = "expression"
	AttrFamily                     = "family"
	AttrField                      = "field"
	AttrFileSystemID               = "file_system_id"
	AttrFilter                     = "filter"
	AttrForceDestroy               = "force_destroy"
	AttrFormat                     = "format"
	AttrFunctionARN                = "function_arn"
	AttrGroupName                  = "group_name"
	AttrHeader                     = "header"
	AttrHealthCheck                = "health_check"
	AttrHostedZoneID               = "hosted_zone_id"
	AttrIAMRoleARN                 = "iam_role_arn"
	AttrID                         = "id"
	AttrIDs                        = "ids"
	AttrIOPS                       = "iops"
	AttrIPAddress                  = "ip_address"
	AttrIPAddressType              = "ip_address_type"
	AttrIdentifier                 = "identifier"
	AttrInstanceCount              = "instance_count"
	AttrInstanceID                 = "instance_id"
	AttrInstanceType               = "instance_type"
	AttrInterval                   = "interval"
	AttrIssuer                     = "issuer"
	AttrJSON                       = "json"
	AttrKMSKey                     = "kms_key"
	AttrKMSKeyARN                  = "kms_key_arn"
	AttrKMSKeyID                   = "kms_key_id"
	AttrKey                        = "key"
	AttrKeyID                      = "key_id"
	AttrLastUpdatedDate            = "last_updated_date"
	AttrLaunchTemplate             = "launch_template"
	AttrLogGroupName               = "log_group_name"
	AttrLoggingConfiguration       = "logging_configuration"
	AttrMax                        = "max"
	AttrMessage                    = "message"
	AttrMetricName                 = "metric_name"
	AttrMin                        = "min"
	AttrMode                       = "mode"
	AttrMostRecent                 = "most_recent"
	AttrName                       = "name"
	AttrNamePrefix                 = "name_prefix"
	AttrNames                      = "names"
	AttrNamespace                  = "namespace"
	AttrNetworkConfiguration       = "network_configuration"
	AttrNetworkInterfaceID         = "network_interface_id"
	AttrOwner                      = "owner"
	AttrOwnerID                    = "owner_id"
	AttrParameter                  = "parameter"
	AttrParameters                 = "parameters"
	AttrPassword                   = "password"
	AttrPath                       = "path"
	AttrPermissions                = "permissions"
	AttrPolicy                     = "policy"
	AttrPort                       = "port"
	AttrPreferredMaintenanceWindow = "preferred_maintenance_window"
	AttrPrefix                     = "prefix"
	AttrPrincipal                  = "principal"
	AttrPriority                   = "priority"
	AttrPrivateKey                 = "private_key"
	AttrProfile                    = "profile"
	AttrProperties                 = "properties"
	AttrProtocol                   = "protocol"
	AttrPublicKey                  = "public_key"
	AttrPubliclyAccessible         = "publicly_accessible"
	AttrRegion                     = "region"
	AttrRepositoryName             = "repository_name"
	AttrResourceARN                = "resource_arn"
	AttrResourceID                 = "resource_id"
	AttrResourceTags               = "resource_tags"
	AttrResourceType               = "resource_type"
	AttrResources                  = "resources"
	AttrRetentionPeriod            = "retention_period"
	AttrRole                       = "role"
	AttrRoleARN                    = "role_arn"
	AttrRule                       = "rule"
	AttrS3Bucket                   = "s3_bucket"
	AttrS3BucketName               = "s3_bucket_name"
	AttrS3KeyPrefix                = "s3_key_prefix"
	AttrSNSTopicARN                = "sns_topic_arn"
	AttrSchedule                   = "schedule"
	AttrScheduleExpression         = "schedule_expression"
	AttrSchema                     = "schema"
	AttrScope                      = "scope"
	AttrSecretKey                  = "secret_key"
	AttrSecurityGroupIDs           = "security_group_ids"
	AttrSecurityGroups             = "security_groups"
	AttrServiceName                = "service_name"
	AttrServiceRoleARN             = "service_role_arn"
	AttrSession                    = "session"
	AttrSharedConfigFiles          = "shared_config_files"
	AttrSize                       = "size"
	AttrSkipCredentialsValidation  = "skip_credentials_validation"
	AttrSkipDestroy                = "skip_destroy"
	AttrSkipRequestingAccountID    = "skip_requesting_account_id"
	AttrSource                     = "source"
	AttrSourceType                 = "source_type"
	AttrStage                      = "stage"
	AttrStartTime                  = "start_time"
	AttrState                      = "state"
	AttrStatus                     = "status"
	AttrStatusMessage              = "status_message"
	AttrStorageType                = "storage_type"
	AttrStreamARN                  = "stream_arn"
	AttrSubnetID                   = "subnet_id"
	AttrSubnetIDs                  = "subnet_ids"
	AttrSubnets                    = "subnets"
	AttrTableName                  = "table_name"
	AttrTags                       = "tags"
	AttrTagsAll                    = "tags_all"
	AttrTarget                     = "target"
	AttrTargetARN                  = "target_arn"
	AttrTimeout                    = "timeout"
	AttrTimeouts                   = "timeouts"
	AttrTopicARN                   = "topic_arn"
	AttrTransitGatewayAttachmentID = "transit_gateway_attachment_id"
	AttrTransitGatewayID           = "transit_gateway_id"
	AttrTriggers                   = "triggers"
	AttrType                       = "type"
	AttrURI                        = "uri"
	AttrURL                        = "url"
	AttrUnit                       = "unit"
	AttrUserName                   = "user_name"
	AttrUsername                   = "username"
	AttrVPCConfig                  = "vpc_config"
	AttrVPCConfiguration           = "vpc_configuration"
	AttrVPCEndpointID              = "vpc_endpoint_id"
	AttrVPCID                      = "vpc_id"
	AttrVPCSecurityGroupIDs        = "vpc_security_group_ids"
	AttrValue                      = "value"
	AttrValues                     = "values"
	AttrVersion                    = "version"
	AttrVolumeSize                 = "volume_size"
	AttrVolumeType                 = "volume_type"
	AttrWeight                     = "weight"
)

// ConstOrQuote returns the constant name for the given attribute if it exists.
// Otherwise, it returns the attribute quoted. This is intended for use in
// generated code and templates.
func ConstOrQuote(constant string) string {
	allConstants := map[string]string{
		"arn":                           "AttrARN",
		"arns":                          "AttrARNs",
		"access_key":                    "AttrAccessKey",
		"account_id":                    "AttrAccountID",
		"action":                        "AttrAction",
		"address":                       "AttrAddress",
		"alias":                         "AttrAlias",
		"apply_immediately":             "AttrApplyImmediately",
		"attributes":                    "AttrAttributes",
		"auto_minor_version_upgrade":    "AttrAutoMinorVersionUpgrade",
		"availability_zone":             "AttrAvailabilityZone",
		"availability_zones":            "AttrAvailabilityZones",
		"bucket":                        "AttrBucket",
		"bucket_name":                   "AttrBucketName",
		"bucket_prefix":                 "AttrBucketPrefix",
		"catalog_id":                    "AttrCatalogID",
		"certificate":                   "AttrCertificate",
		"certificate_arn":               "AttrCertificateARN",
		"certificate_chain":             "AttrCertificateChain",
		"client_id":                     "AttrClientID",
		"client_secret":                 "AttrClientSecret",
		"cluster_identifier":            "AttrClusterIdentifier",
		"cluster_name":                  "AttrClusterName",
		"comment":                       "AttrComment",
		"condition":                     "AttrCondition",
		"configuration":                 "AttrConfiguration",
		"content":                       "AttrContent",
		"content_type":                  "AttrContentType",
		"create_time":                   "AttrCreateTime",
		"created_at":                    "AttrCreatedAt",
		"created_date":                  "AttrCreatedDate",
		"created_time":                  "AttrCreatedTime",
		"creation_date":                 "AttrCreationDate",
		"creation_time":                 "AttrCreationTime",
		"dns_name":                      "AttrDNSName",
		"database":                      "AttrDatabase",
		"database_name":                 "AttrDatabaseName",
		"default_value":                 "AttrDefaultValue",
		"delete_on_termination":         "AttrDeleteOnTermination",
		"description":                   "AttrDescription",
		"destination":                   "AttrDestination",
		"destination_arn":               "AttrDestinationARN",
		"device_name":                   "AttrDeviceName",
		"display_name":                  "AttrDisplayName",
		"domain":                        "AttrDomain",
		"domain_name":                   "AttrDomainName",
		"duration":                      "AttrDuration",
		"email":                         "AttrEmail",
		"enabled":                       "AttrEnabled",
		"encrypted":                     "AttrEncrypted",
		"encryption_configuration":      "AttrEncryptionConfiguration",
		"endpoint":                      "AttrEndpoint",
		"endpoint_type":                 "AttrEndpointType",
		"endpoints":                     "AttrEndpoints",
		"engine_version":                "AttrEngineVersion",
		"environment":                   "AttrEnvironment",
		"execution_role_arn":            "AttrExecutionRoleARN",
		"expression":                    "AttrExpression",
		"family":                        "AttrFamily",
		"field":                         "AttrField",
		"file_system_id":                "AttrFileSystemID",
		"filter":                        "AttrFilter",
		"force_destroy":                 "AttrForceDestroy",
		"format":                        "AttrFormat",
		"function_arn":                  "AttrFunctionARN",
		"group_name":                    "AttrGroupName",
		"header":                        "AttrHeader",
		"health_check":                  "AttrHealthCheck",
		"hosted_zone_id":                "AttrHostedZoneID",
		"iam_role_arn":                  "AttrIAMRoleARN",
		"id":                            "AttrID",
		"ids":                           "AttrIDs",
		"iops":                          "AttrIOPS",
		"ip_address":                    "AttrIPAddress",
		"ip_address_type":               "AttrIPAddressType",
		"identifier":                    "AttrIdentifier",
		"instance_count":                "AttrInstanceCount",
		"instance_id":                   "AttrInstanceID",
		"instance_type":                 "AttrInstanceType",
		"interval":                      "AttrInterval",
		"issuer":                        "AttrIssuer",
		"json":                          "AttrJSON",
		"kms_key":                       "AttrKMSKey",
		"kms_key_arn":                   "AttrKMSKeyARN",
		"kms_key_id":                    "AttrKMSKeyID",
		"key":                           "AttrKey",
		"key_id":                        "AttrKeyID",
		"last_updated_date":             "AttrLastUpdatedDate",
		"launch_template":               "AttrLaunchTemplate",
		"log_group_name":                "AttrLogGroupName",
		"logging_configuration":         "AttrLoggingConfiguration",
		"max":                           "AttrMax",
		"message":                       "AttrMessage",
		"metric_name":                   "AttrMetricName",
		"min":                           "AttrMin",
		"mode":                          "AttrMode",
		"most_recent":                   "AttrMostRecent",
		"name":                          "AttrName",
		"name_prefix":                   "AttrNamePrefix",
		"names":                         "AttrNames",
		"namespace":                     "AttrNamespace",
		"network_configuration":         "AttrNetworkConfiguration",
		"network_interface_id":          "AttrNetworkInterfaceID",
		"owner":                         "AttrOwner",
		"owner_id":                      "AttrOwnerID",
		"parameter":                     "AttrParameter",
		"parameters":                    "AttrParameters",
		"password":                      "AttrPassword",
		"path":                          "AttrPath",
		"permissions":                   "AttrPermissions",
		"policy":                        "AttrPolicy",
		"port":                          "AttrPort",
		"preferred_maintenance_window":  "AttrPreferredMaintenanceWindow",
		"prefix":                        "AttrPrefix",
		"principal":                     "AttrPrincipal",
		"priority":                      "AttrPriority",
		"private_key":                   "AttrPrivateKey",
		"profile":                       "AttrProfile",
		"properties":                    "AttrProperties",
		"protocol":                      "AttrProtocol",
		"public_key":                    "AttrPublicKey",
		"publicly_accessible":           "AttrPubliclyAccessible",
		"region":                        "AttrRegion",
		"repository_name":               "AttrRepositoryName",
		"resource_arn":                  "AttrResourceARN",
		"resource_id":                   "AttrResourceID",
		"resource_tags":                 "AttrResourceTags",
		"resource_type":                 "AttrResourceType",
		"resources":                     "AttrResources",
		"retention_period":              "AttrRetentionPeriod",
		"role":                          "AttrRole",
		"role_arn":                      "AttrRoleARN",
		"rule":                          "AttrRule",
		"s3_bucket":                     "AttrS3Bucket",
		"s3_bucket_name":                "AttrS3BucketName",
		"s3_key_prefix":                 "AttrS3KeyPrefix",
		"sns_topic_arn":                 "AttrSNSTopicARN",
		"schedule":                      "AttrSchedule",
		"schedule_expression":           "AttrScheduleExpression",
		"schema":                        "AttrSchema",
		"scope":                         "AttrScope",
		"secret_key":                    "AttrSecretKey",
		"security_group_ids":            "AttrSecurityGroupIDs",
		"security_groups":               "AttrSecurityGroups",
		"service_name":                  "AttrServiceName",
		"service_role_arn":              "AttrServiceRoleARN",
		"session":                       "AttrSession",
		"shared_config_files":           "AttrSharedConfigFiles",
		"size":                          "AttrSize",
		"skip_credentials_validation":   "AttrSkipCredentialsValidation",
		"skip_destroy":                  "AttrSkipDestroy",
		"skip_requesting_account_id":    "AttrSkipRequestingAccountID",
		"source":                        "AttrSource",
		"source_type":                   "AttrSourceType",
		"stage":                         "AttrStage",
		"start_time":                    "AttrStartTime",
		"state":                         "AttrState",
		"status":                        "AttrStatus",
		"status_message":                "AttrStatusMessage",
		"storage_type":                  "AttrStorageType",
		"stream_arn":                    "AttrStreamARN",
		"subnet_id":                     "AttrSubnetID",
		"subnet_ids":                    "AttrSubnetIDs",
		"subnets":                       "AttrSubnets",
		"table_name":                    "AttrTableName",
		"tags":                          "AttrTags",
		"tags_all":                      "AttrTagsAll",
		"target":                        "AttrTarget",
		"target_arn":                    "AttrTargetARN",
		"timeout":                       "AttrTimeout",
		"timeouts":                      "AttrTimeouts",
		"topic_arn":                     "AttrTopicARN",
		"transit_gateway_attachment_id": "AttrTransitGatewayAttachmentID",
		"transit_gateway_id":            "AttrTransitGatewayID",
		"triggers":                      "AttrTriggers",
		"type":                          "AttrType",
		"uri":                           "AttrURI",
		"url":                           "AttrURL",
		"unit":                          "AttrUnit",
		"user_name":                     "AttrUserName",
		"username":                      "AttrUsername",
		"vpc_config":                    "AttrVPCConfig",
		"vpc_configuration":             "AttrVPCConfiguration",
		"vpc_endpoint_id":               "AttrVPCEndpointID",
		"vpc_id":                        "AttrVPCID",
		"vpc_security_group_ids":        "AttrVPCSecurityGroupIDs",
		"value":                         "AttrValue",
		"values":                        "AttrValues",
		"version":                       "AttrVersion",
		"volume_size":                   "AttrVolumeSize",
		"volume_type":                   "AttrVolumeType",
		"weight":                        "AttrWeight",
	}

	if v, ok := allConstants[constant]; ok {
		return fmt.Sprintf("names.%s", v)
	}
	return fmt.Sprintf("%q", constant)
}
