# セキュリティグループモジュール (modules/security_group/variables.tf)
variable "project_name" {
  description = "The name of the project"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "allowed_ip_addresses" {
  description = "List of allowed IP addresses in CIDR notation"
  type        = list(string)
}

