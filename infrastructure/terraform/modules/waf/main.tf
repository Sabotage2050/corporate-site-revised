# IPv4用のIPセット
resource "aws_wafv2_ip_set" "allowed_ips_v4" {
  name               = "${var.project_name}-allowed-ips-v4"
  description        = "Allowed IPv4 addresses"
  scope              = "REGIONAL"
  ip_address_version = "IPV4"
  addresses          = var.allowed_ipv4_addresses
}

# IPv6用のIPセット
resource "aws_wafv2_ip_set" "allowed_ips_v6" {
  name               = "${var.project_name}-allowed-ips-v6"
  description        = "Allowed IPv6 addresses"
  scope              = "REGIONAL"
  ip_address_version = "IPV6"
  addresses          = var.allowed_ipv6_addresses
}

# WAF Web ACLの定義
resource "aws_wafv2_web_acl" "main" {
  name        = "${var.project_name}-web-acl"
  description = "WAF rules for ${var.project_name}"
  scope       = "REGIONAL"
  
  default_action {
    block {}
  }

  # レート制限ルール
  rule {
    name     = "IPRateLimit"
    priority = 1
    
    action {
      block {}
    }
    
    statement {
      rate_based_statement {
        limit              = var.waf_rate_limit
        aggregate_key_type = "IP"
      }
    }
    
    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "IPRateLimit"
      sampled_requests_enabled   = true
    }
  }

  # IP制限ルール
  rule {
    name     = "IPAllowList"
    priority = 2
    
    action {
      allow {}
    }
    
    statement {
      or_statement {
        statement {
          ip_set_reference_statement {
            arn = aws_wafv2_ip_set.allowed_ips_v4.arn
          }
        }
        statement {
          ip_set_reference_statement {
            arn = aws_wafv2_ip_set.allowed_ips_v6.arn
          }
        }
      }
    }
    
    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "IPAllowList"
      sampled_requests_enabled   = true
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = true
    metric_name                = "${var.project_name}-web-acl"
    sampled_requests_enabled   = true
  }
}

# ALBとの関連付け
resource "aws_wafv2_web_acl_association" "main" {
  resource_arn = var.alb_arn
  web_acl_arn  = aws_wafv2_web_acl.main.arn
}

# ログ設定
# CloudWatchロググループ
resource "aws_cloudwatch_log_group" "waf_logs" {
  name              = "aws-waf-logs-${var.project_name}" # aws-waf-logs- プレフィックスが必要
  retention_in_days = var.waf_log_retention_days
}

# WAF Logging設定
resource "aws_wafv2_web_acl_logging_configuration" "main" {
  log_destination_configs = ["${aws_cloudwatch_log_group.waf_logs.arn}:*"]
  resource_arn           = aws_wafv2_web_acl.main.arn
}