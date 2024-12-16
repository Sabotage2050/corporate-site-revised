output "alb_arn" {
  description = "The ARN of the application load balancer"
  value       = aws_lb.main.arn
}

output "alb_dns_name" {
  description = "The DNS name of the application load balancer"
  value       = aws_lb.main.dns_name
}

output "alb_zone_id" {
  description = "The zone ID of the application load balancer"
  value       = aws_lb.main.zone_id
}

output "alb_https_listener_arn" {
  description = "The ARN of the HTTPS listener"
  value       = aws_lb_listener.https.arn
}
output "alb_target_group_arn" {
  value = aws_lb_target_group.app.arn
}
