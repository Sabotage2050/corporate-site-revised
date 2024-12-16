# RDSモジュール (modules/rds/main.tf)
resource "aws_db_subnet_group" "default" {
  name       = "${var.project_name}-db-subnet-group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "${var.project_name}-db-subnet-group"
  }
}

resource "aws_db_instance" "default" {
  identifier           = "${var.project_name}-db"
  engine               = "mysql"
  engine_version       = "8.0"
  instance_class       = var.db_instance_class
  allocated_storage    = 20
  storage_type         = "gp2"
  username             = var.db_username
  password             = var.db_password
  db_name              = var.db_name
  db_subnet_group_name = aws_db_subnet_group.default.name
  vpc_security_group_ids = [var.rds_security_group_id]
  skip_final_snapshot  = true

  tags = {
    Name = "${var.project_name}-db"
  }
}