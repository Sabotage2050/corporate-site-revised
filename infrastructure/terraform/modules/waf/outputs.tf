# outputs.tf
output "web_acl_id" {
  description = "The ID of the WAF Web ACL"
  value       = aws_wafv2_web_acl.main.id
}

output "web_acl_arn" {
  description = "The ARN of the WAF Web ACL"
  value       = aws_wafv2_web_acl.main.arn
}

output "ip_rate_limit_rule_metric" {
  description = "Metric name for the IP rate limit rule"
  value       = "IPRateLimit"
}
