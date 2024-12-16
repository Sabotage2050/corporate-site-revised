# modules/ssm/outputs.tf
output "parameter_arns" {
  description = "The ARNs of the SSM parameters"
  value = [for param in aws_ssm_parameter.parameters : param.arn]
}

output "parameter_names" {
  description = "The names of the SSM parameters"
  value = [for param in aws_ssm_parameter.parameters : param.name]
}