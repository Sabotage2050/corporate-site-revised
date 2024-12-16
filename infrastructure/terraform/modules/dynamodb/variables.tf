variable "table_prefix" {
  description = "Prefix for the DynamoDB table name"
  type        = string
}

variable "table_name" {
  description = "Base name of the DynamoDB table"
  type        = string
}

variable "environment" {
  description = "Environment name (e.g., staging, production)"
  type        = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}