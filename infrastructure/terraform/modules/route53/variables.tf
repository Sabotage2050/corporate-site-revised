# terraform/modules/route53/variables.tf
variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "domain_name" {
  description = "Domain name"
  type        = string
}

# バックエンド用の変数
variable "backend_subdomain" {
  description = "Subdomain for backend (e.g., api, staging)"
  type        = string
}

variable "alb_dns_name" {
  description = "ALB DNS name"
  type        = string
}

variable "alb_zone_id" {
  description = "ALB Zone ID"
  type        = string
}

# フロントエンド用の変数
variable "frontend_subdomain" {
  description = "Subdomain for frontend (e.g., www)"
  type        = string
}

variable "cloudfront_dns_name" {
  description = "CloudFront distribution DNS name"
  type        = string
  default     = ""  # フロントエンドを使用しない場合のデフォルト値
}

variable "cloudfront_zone_id" {
  description = "CloudFront distribution hosted zone ID"
  type        = string
  default     = ""  # フロントエンドを使用しない場合のデフォルト値
}