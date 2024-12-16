variable "project_name" {
  description = "The name of the project"
  type        = string
}

variable "environment" {
  description = "The environment to deploy to"
  type        = string
}

variable "aws_region" {
  description = "The AWS region"
  type        = string
}

variable "subnet_id" {
  description = "Subnet ID for the ECS tasks"
  type        = string
}

variable "ecs_security_group_id" {
  description = "Security group ID for ECS tasks"
  type        = string
}

variable "task_cpu" {
  description = "CPU units for the task"
  type        = string
}

variable "task_memory" {
  description = "Memory for the task"
  type        = string
}

variable "service_desired_count" {
  description = "Desired number of tasks"
  type        = number
}

variable "ecr_repository_url" {
  description = "URL of the ECR repository"
  type        = string
}

variable "image_tag" {
  description = "Docker image tag"
  type        = string
}

variable "alb_target_group_arn" {
  description = "ARN of the ALB target group"
  type        = string
}

variable "ecs_task_execution_role_arn" {
  description = "ARN of the ECS task execution role"
  type        = string
}

variable "ecs_task_role_arn" {
  description = "ARN of the ECS task role"
  type        = string
}

variable "force_new_deployment" {
  description = "Force a new deployment of the service"
  type        = bool
  default     = true
}
