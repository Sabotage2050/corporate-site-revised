# RDSモジュール (modules/rds/variables.tf)
variable "project_name" {
  description = "The name of the project"
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs for the RDS instance"
  type        = list(string)
}

variable "db_instance_class" {
  description = "The instance type of the RDS instance"
  type        = string
}

variable "db_username" {
  description = "Username for the database"
  type        = string
}

variable "db_password" {
  description = "Password for the database"
  type        = string
}

variable "db_name" {
  description = "Database Name"
  type = string
}

variable "rds_security_group_id" {
  description = "ID of the RDS security group"
  type        = string
}