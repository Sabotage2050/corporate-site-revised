# infrastructure/terraform/common/circleci-oidc/variables.tf
variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "revised-corporate-site"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
}

variable "circleci_org_id" {
  description = "CircleCI Organization ID"
  type        = string
  default = "9ddd9dc1-d266-4467-9171-8acadd3261a4"
}

variable "role_name" {
  description = "Name of the IAM role for CircleCI"
  type        = string
  default     = "circleci-oidc-role"
}

variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
  default     = "circleci-oidc-bucket"
}