output "db_instance_endpoint" {
  description = "The connection endpoint for the RDS instance"
  value       = aws_db_instance.default.endpoint
}

output "db_instance_arn" {
  description = "The ARN of the RDS instance"
  value       = aws_db_instance.default.arn
}

output "db_instance_name" {
  description = "The database name"
  value       = aws_db_instance.default.db_name
}

output "db_instance_port" {
  description = "The database port"
  value       = aws_db_instance.default.port
}

output "db_instance_username" {
  description = "The master username for the database"
  value       = aws_db_instance.default.username
}


# 注意: パスワードは機密情報のため、通常は出力しません
output "db_instance_password" {
  description = "The database password"
  value       = aws_db_instance.default.password
}