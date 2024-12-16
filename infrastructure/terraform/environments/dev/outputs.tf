output "s3_bucket_name" {
  description = "The name of the S3 bucket for the development environment"
  value       = module.s3_bucket.bucket_name
}

output "s3_bucket_arn" {
  description = "The ARN of the S3 bucket for the development environment"
  value       = module.s3_bucket.bucket_arn
}
