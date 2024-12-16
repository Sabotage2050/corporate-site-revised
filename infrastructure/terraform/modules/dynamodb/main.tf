resource "aws_dynamodb_table" "example" {
  name         = "${var.table_prefix}${var.table_name}"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "Enginetype"
    type = "S"
  }

  attribute {
    name = "SerialNo"
    type = "S"
  }

  hash_key  = "Enginetype"
  range_key = "SerialNo"

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    Environment = var.environment
    Project     = var.project_name
  }
}