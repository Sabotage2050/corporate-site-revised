# 基本設定
variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "revised-corporate-site"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "staging"
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

# VPC関連
variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_cidrs" {
  description = "CIDR blocks for the subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "availability_zones" {
  description = "Availability zones"
  type        = list(string)
  default     = ["ap-northeast-1a", "ap-northeast-1c"]
}

# ECS関連
variable "ecs_task_cpu" {
  description = "CPU units for the ECS task"
  type        = string
  default     = "256"
}

variable "ecs_task_memory" {
  description = "Memory for the ECS task"
  type        = string
  default     = "512"
}

variable "ecs_service_desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 1
}

variable "image_tag" {
  description = "Docker image tag"
  type        = string
}

variable "ecr_repository_arn" {
  description = "value of the ECR repository ARN"
  type        = string
}

# ALB関連
variable "health_check_path" {
  description = "Health check path for ALB"
  type        = string
  default     = "/health"
}

# DynamoDB関連
variable "dynamodb_table_prefix" {
  description = "Prefix for DynamoDB table names"
  type        = string
  default     = ""
}

variable "dynamodb_table_name" {
  description = "Name of the DynamoDB table"
  type        = string
  default     = "Forklift"
}

# WAF関連
variable "allowed_ipv4_addresses" {
  description = "List of allowed IPv4 addresses in CIDR notation"
  type        = list(string)
  default = [
    "49.109.139.57/32"    # 現在のIPv4アドレス
  ]
}

variable "allowed_ipv6_addresses" {
  description = "List of allowed IPv6 addresses in CIDR notation"
  type        = list(string)
  default = [
    "240a:61:52f6:70a5::/64" # 現在のIPv6範囲
  ]
}

variable "waf_log_retention_days" {
  description = "Number of days to retain WAF logs"
  type        = number
  default     = 30
}

variable "waf_rate_limit" {
  description = "Maximum number of requests allowed from an IP in 5 minutes"
  type        = number
  default     = 2000
}

variable "waf_enabled" {
  description = "Enable or disable WAF protection"
  type        = bool
  default     = true
}

variable "waf_block_mode_enabled" {
  description = "Enable blocking mode (true) or count mode (false)"
  type        = bool
  default     = true
}

# ドメイン関連
variable "domain_name" {
  description = "Base domain name (e.g., example.com)"
  type        = string
  default     = "sanichi-machine.com"
}

variable "subdomain" {
  description = "Subdomain for staging environment"
  type        = string
  default     = "api"
}

variable "frontend_subdomain" {
  description = "Subdomain for frontend (www)"
  type        = string
  default     = "www"
}

# CloudFront関連
variable "cloudfront_price_class" {
  description = "CloudFront distribution price class"
  type        = string
  default     = "PriceClass_All"  # PriceClass_100, PriceClass_200, PriceClass_All
}

variable "cloudfront_ttl" {
  description = "CloudFront cache TTL settings"
  type        = map(number)
  default = {
    min     = 0
    default = 3600    # 1時間
    max     = 86400   # 24時間
  }
}