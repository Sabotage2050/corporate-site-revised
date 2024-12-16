# Provider設定
provider "aws" {
  region = var.aws_region
}

provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

terraform {
  backend "s3" {
    bucket = "terraform-state-revised-corporate-site"
    key    = "staging/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

locals {
  table_prefix = "${var.project_name}-${var.environment}-"
}

# DynamoDB テーブル
module "dynamodb" {
  source       = "../../modules/dynamodb"
  table_prefix = local.table_prefix
  table_name   = var.dynamodb_table_name
  environment  = var.environment
  project_name = var.project_name
}

# VPC
module "vpc" {
  source             = "../../modules/vpc"
  project_name       = var.project_name
  vpc_cidr          = var.vpc_cidr
  subnet_cidrs      = var.subnet_cidrs
  availability_zones = var.availability_zones
}

# Security Groups
module "security_groups" {
  source               = "../../modules/security_group"
  project_name         = var.project_name
  vpc_id               = module.vpc.vpc_id
  allowed_ip_addresses = ["0.0.0.0/0"]
}

module "route53" {
  source = "../../modules/route53"
  providers = {
    aws = aws
  }
  project_name         = var.project_name
  domain_name          = var.domain_name
  backend_subdomain    = var.subdomain
  alb_dns_name         = module.alb.alb_dns_name
  alb_zone_id          = module.alb.alb_zone_id
  frontend_subdomain   = var.frontend_subdomain
  cloudfront_dns_name  = module.cloudfront.cloudfront_domain_name
  cloudfront_zone_id   = module.cloudfront.cloudfront_hosted_zone_id
}
# 
module "acm_main" {
  source = "../../modules/acm"
  providers = {
    aws = aws
  }
  domain_name      = "${var.subdomain}.${var.domain_name}"
  route53_zone_id  = module.route53.zone_id
  project_name     = var.project_name
  environment      = var.environment
}
# 
module "acm_cloudfront" {
  source = "../../modules/acm"
  providers = {
    aws = aws.us_east_1
  }
  domain_name      = "${var.frontend_subdomain}.${var.domain_name}"
  route53_zone_id  = module.route53.zone_id
  project_name     = var.project_name
  environment      = var.environment
}
# 
module "alb" {
  source                = "../../modules/alb"
  project_name          = var.project_name
  vpc_id                = module.vpc.vpc_id
  public_subnet_ids     = module.vpc.public_subnet_ids
  alb_security_group_id = module.security_groups.alb_security_group_id
  health_check_path     = var.health_check_path
  certificate_arn       = module.acm_main.certificate_arn
  certificate_validation_completion = module.acm_main.certificate_validation_completion
}
# 
module "artifact_bucket" {
  source = "../../modules/s3"
  providers = {
    aws = aws
  }
  bucket_name   = "${var.project_name}-${var.environment}-artifacts"
  versioning    = true
  force_destroy = true

  cors_rules = [
    {
      allowed_headers = ["*"]
      allowed_methods = ["GET", "PUT", "POST"]
      allowed_origins = ["https://${var.subdomain}.${var.domain_name}"]
      expose_headers  = ["ETag"]
      max_age_seconds = 3000
    }
  ]

  tags = {
    Environment = var.environment
    Project     = var.project_name
    Type        = "artifacts"
  }
}
# 
module "log_bucket" {
  source = "../../modules/s3"
  providers = {
    aws = aws
  }
  bucket_name   = "${var.project_name}-${var.environment}-logs"
  versioning    = false
  force_destroy = true

  bucket_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowLogDelivery"
        Effect = "Allow"
        Principal = {
          Service = "logging.s3.amazonaws.com"
        }
        Action   = "s3:PutObject"
        Resource = "arn:aws:s3:::${var.project_name}-${var.environment}-logs/*"
      }
    ]
  })
# 
  tags = {
    Environment = var.environment
    Project     = var.project_name
    Type        = "logs"
  }
}
# 
module "waf" {
  source = "../../modules/waf"
  # 
  project_name           = var.project_name
  alb_arn               = module.alb.alb_arn
  allowed_ipv4_addresses = var.allowed_ipv4_addresses
  allowed_ipv6_addresses = var.allowed_ipv6_addresses
  waf_rate_limit        = var.waf_rate_limit
  waf_log_retention_days = var.waf_log_retention_days
}
# 
module "iam" {
  source       = "../../modules/iam"
  project_name = var.project_name
}

module "cloudwatch" {
  source       = "../../modules/cloudwatch"
  project_name = var.project_name
}
# 
module "ecs" {
  source = "../../modules/ecs"
  
  project_name = var.project_name
  environment  = var.environment
  aws_region   = var.aws_region
  
  subnet_id              = module.vpc.public_subnet_ids[0]
  ecs_security_group_id  = module.security_groups.ecs_security_group_id
  
  task_cpu              = var.ecs_task_cpu
  task_memory           = var.ecs_task_memory
  service_desired_count = var.ecs_service_desired_count
  
  ecr_repository_url           = var.ecr_repository_arn
  image_tag                    = var.image_tag
  ecs_task_execution_role_arn  = module.iam.ecs_task_execution_role_arn
  ecs_task_role_arn           = module.iam.ecs_task_role_arn
  alb_target_group_arn        = module.alb.alb_target_group_arn
}
# 
module "frontend_bucket" {
  source = "../../modules/s3"
  providers = {
    aws = aws
  }
  bucket_name     = "${var.project_name}-${var.environment}-frontend"
  versioning      = true
  force_destroy   = true
  website_enabled = true
  tags = {
    Environment = var.environment
    Project     = var.project_name
    Type        = "frontend"
  }
}

module "cloudfront" {
  source = "../../modules/cloudfront"
  providers = {
    aws = aws.us_east_1
  }
  domain_name     = "${var.frontend_subdomain}.${var.domain_name}"
  s3_bucket_regional_domain_name = module.frontend_bucket.bucket_regional_domain_name
  certificate_arn = module.acm_cloudfront.certificate_arn
  project_name    = var.project_name
  environment     = var.environment
}
# 
resource "aws_s3_bucket_policy" "frontend" {
  bucket = module.frontend_bucket.bucket_id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowCloudFrontServicePrincipal"
        Effect    = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "${module.frontend_bucket.bucket_arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = module.cloudfront.distribution_arn
          }
        }
      }
    ]
  })

  depends_on = [module.cloudfront]
}