# terraform/modules/route53/outputs.tf
output "zone_id" {
  description = "The ID of the Route53 zone"
  value       = data.aws_route53_zone.selected.zone_id
}

output "backend_url" {
  description = "The full URL for the backend"
  value       = "${var.backend_subdomain}.${var.domain_name}"
}

output "frontend_url" {
  description = "The full URL for the frontend"
  value       = var.cloudfront_dns_name != "" ? "${var.frontend_subdomain}.${var.domain_name}" : null
}