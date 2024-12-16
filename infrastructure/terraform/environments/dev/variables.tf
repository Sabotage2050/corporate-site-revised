# environments/dev/variables.tf

variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "revised-corporate-site"
}

variable "circleci_org_id" {
  description = "CircleCI Organization ID"
  type        = string
  default     = "9ddd9dc1-d266-4467-9171-8acadd3261a4"
}