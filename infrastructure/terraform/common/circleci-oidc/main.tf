# infrastructure/terraform/common/circleci-oidc/main.tf
provider "aws" {
  region = var.aws_region
}

module "circleci_oidc" {
  source          = "../../modules/circleci-oidc"
  role_name       = var.role_name
  circleci_org_id = var.circleci_org_id
  project_name    = var.project_name
  environment     = var.environment
}