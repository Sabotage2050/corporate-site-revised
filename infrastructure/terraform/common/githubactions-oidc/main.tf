provider "aws" {
  region = var.aws_region
}

# Terraform状態用のS3バケット
module "terraform_state_bucket" {
  source       = "../../modules/s3"
  bucket_name  = "terraform-state-${var.project-name}"
  acl          = "private"
  versioning   = true
  force_destroy = true
  tags = {
    Environment = var.environment
    Project     = var.project-name
  }
}

module "github_oidc" {
  source = "../../modules/github-oidc"
}

module "github_iam_role" {
  source           = "../../modules/githubactions-iam-role"
  oidc_provider_arn = module.github_oidc.oidc_provider_arn
  role_name        = var.role_name
  github_org       = var.github_org
  github_repo      = var.github_repo
  environment      = var.environment
}


