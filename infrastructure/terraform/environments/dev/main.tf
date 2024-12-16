# Backend configuration for state management
terraform {
  backend "s3" {
    bucket = "terraform-state-revised-corporate-site"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}
# Provider configuration
provider "aws" {
  region = "ap-northeast-1"
}

# アプリケーション用S3バケット
module "s3_bucket" {
  source        = "../../modules/s3"
  bucket_name   = "${var.project_name}-bucket-dev"
  versioning    = true
  force_destroy = true
  tags = {
    Environment = "development"
    Project     = var.project_name
  }
}

