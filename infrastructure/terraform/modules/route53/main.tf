# terraform/modules/route53/main.tf
data "aws_route53_zone" "selected" {
  name         = var.domain_name
  private_zone = false
}

# バックエンド用のレコード
resource "aws_route53_record" "backend" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = "${var.backend_subdomain}.${var.domain_name}"
  type    = "A"

  alias {
    name                   = var.alb_dns_name
    zone_id                = var.alb_zone_id
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "frontend" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = "${var.frontend_subdomain}.${var.domain_name}"
  type    = "A"

  alias {
    name                   = coalesce(var.cloudfront_dns_name, "example.com") # デフォルト値を設定
    zone_id                = coalesce(var.cloudfront_zone_id, "Z2FDTNDATAQYW2") # デフォルト値を設定
    evaluate_target_health = var.cloudfront_dns_name != "" ? false : true
  }
}
