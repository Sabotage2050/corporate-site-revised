resource "aws_ses_domain_identity" "example" {
  domain = var.domain_name
}

resource "aws_ses_domain_dkim" "example" {
  domain = aws_ses_domain_identity.example.domain
}

resource "aws_ses_email_identity" "example" {
  email = var.email_address
}
