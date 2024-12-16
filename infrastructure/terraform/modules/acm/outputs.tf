# terraform/modules/acm/outputs.tf
output "certificate_arn" {
  description = "The ARN of the ACM certificate"
  value       = aws_acm_certificate.https_cert.arn
}

output "certificate_id" {
  description = "The ID of the ACM certificate"
  value       = aws_acm_certificate.https_cert.id
}

output "domain_validation_options" {
  description = "Domain validation options for the certificate"
  value       = aws_acm_certificate.https_cert.domain_validation_options
}

# terraform/modules/acm/outputs.tf に追加
output "certificate_validation_completion" {
  description = "A dependency value to ensure certificate validation is complete"
  value       = aws_acm_certificate_validation.cert_validation.id
}