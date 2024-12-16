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


variable "parameters" {
  description = "Map of parameter names and their default values"
  type = map(object({
    value = string
    description = optional(string)
  }))
  default = {
    "ses_from_address" = {
      value = "removable4463@gmail.com"
      description = "SES sender email address"
    }
    "ses_to_address" = {
      value = "akanatamaya80@gmail.com"
      description = "SES recipient email address"
    }
  }
}