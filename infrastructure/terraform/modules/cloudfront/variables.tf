# modules/cloudfront/variables.tf
variable "domain_name" {
  description = "Domain name for the CloudFront distribution"
  type        = string
}

variable "s3_bucket_regional_domain_name" {
  description = "Regional domain name of the S3 bucket"
  type        = string
}

variable "certificate_arn" {
  description = "ARN of the ACM certificate in us-east-1"
  type        = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "environment" {
  description = "Name of the environment"
  type        = string
}

variable "price_class" {
  description = "CloudFront distribution price class"
  type        = string
  default     = "PriceClass_All"
}

variable "default_ttl" {
  description = "Default TTL for cache behavior"
  type        = number
  default     = 3600
}

variable "min_ttl" {
  description = "Minimum TTL for cache behavior"
  type        = number
  default     = 0
}

variable "max_ttl" {
  description = "Maximum TTL for cache behavior"
  type        = number
  default     = 86400
}