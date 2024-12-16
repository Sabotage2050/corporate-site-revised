# modules/waf/variables.tf

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "allowed_ipv4_addresses" {
  description = "List of allowed IPv4 addresses in CIDR notation"
  type        = list(string)
  default     = [
    "49.109.155.217/32"  # 現在のIPv4アドレス
  ]
}

variable "allowed_ipv6_addresses" {
  description = "List of allowed IPv6 addresses in CIDR notation"
  type        = list(string)
  default     = []  # IPv6アドレスが必要な場合はここに追加
}

variable "alb_arn" {
  description = "ARN of the Application Load Balancer"
  type        = string
}

variable "waf_rate_limit" {
  description = "Maximum number of requests allowed from an IP in 5 minutes"
  type        = number
  default     = 2000
}

variable "waf_log_retention_days" {
  description = "Number of days to retain WAF logs"
  type        = number
  default     = 30
}
