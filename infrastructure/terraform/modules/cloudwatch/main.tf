# CloudWatch Logsモジュール (新規作成: modules/cloudwatch/main.tf)
resource "aws_cloudwatch_log_group" "ecs_logs" {
  name              = "/ecs/${var.project_name}"
  retention_in_days = 30
}