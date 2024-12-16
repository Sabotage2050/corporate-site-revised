# modules/ssm/main.tf
# modules/ssm/main.tf
resource "aws_ssm_parameter" "parameters" {
  for_each = var.parameters

  name        = "/${var.project_name}/${var.environment}/${each.key}"
  description = coalesce(each.value.description, "${var.project_name} ${var.environment} ${each.key} parameter")
  type        = "SecureString"
  value       = each.value.value

  lifecycle {
    ignore_changes = [
      value
    ]
  }

  tags = {
    Environment = var.environment
    Project     = var.project_name
  }
}