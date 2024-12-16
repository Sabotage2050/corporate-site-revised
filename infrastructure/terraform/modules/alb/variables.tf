# ALBモジュール (modules/alb/variables.tf)
variable "project_name" {
  description = "The name of the project"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "public_subnet_ids" {
  description = "List of public subnet IDs"
  type        = list(string)
}

variable "alb_security_group_id" {
  description = "ID of the ALB security group"
  type        = string
}

variable "health_check_path" {
  description = "Health check path for the default target group"
  type        = string
  default     = "/"
}

variable "certificate_arn" {
  description = "The ARN of the ACM certificate"
  type        = string
}

# terraform/modules/alb/variables.tf に追加
variable "certificate_validation_completion" {
  description = "Certificate validation completion dependency"
  type        = any
  default     = null
}