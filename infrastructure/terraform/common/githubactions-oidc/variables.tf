variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "project-name" {
  description = "value of the project name"
  type         = string
  default      = "revised-corporate-site"
}

variable "github_org" {
  description = "The GitHub organization name"
  type        = string
  default     = "Sabotage2050"
}

variable "github_repo" {
  description = "The GitHub repository name"
  type        = string
  default     = "revised-corporate-site"
}

variable "role_name" {
  description = "The role name to be used for the instance"
  type        = string
  default     = "github-actions-role"
}

variable "environment" {
  description = "value of environment variable"
  type        = string
  default     = "dev"
}

