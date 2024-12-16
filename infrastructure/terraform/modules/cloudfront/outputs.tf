# modules/cloudfront/outputs.tf
output "distribution_arn" {
  description = "ARN of the CloudFront distribution"
  value       = aws_cloudfront_distribution.s3_distribution.arn
}

output "distribution_id" {
  description = "ID of the CloudFront distribution"
  value       = aws_cloudfront_distribution.s3_distribution.id
}

output "cloudfront_domain_name" {
  description = "Domain name of the CloudFront distribution"
  value       = aws_cloudfront_distribution.s3_distribution.domain_name
}

output "cloudfront_hosted_zone_id" {
  description = "Route 53 zone ID of the CloudFront distribution"
  value       = aws_cloudfront_distribution.s3_distribution.hosted_zone_id
}

output "origin_access_control_id" {
  description = "ID of the CloudFront origin access control"
  value       = aws_cloudfront_origin_access_control.this.id
}
