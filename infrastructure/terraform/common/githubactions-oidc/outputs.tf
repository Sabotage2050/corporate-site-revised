# outputs.tf
output "terraform_state_bucket_id" {
  description = "The name of the Terraform state bucket"
  value       = module.terraform_state_bucket.bucket_id
}

output "terraform_state_bucket_arn" {
  description = "The ARN of the Terraform state bucket"
  value       = module.terraform_state_bucket.bucket_arn
}

output "github_oidc_provider_arn" {
  description = "The ARN of the GitHub OIDC provider"
  value       = module.github_oidc.oidc_provider_arn
}

output "github_iam_role_arn" {
  description = "The ARN of the IAM role for GitHub Actions"
  value       = module.github_iam_role.role_arn
}
