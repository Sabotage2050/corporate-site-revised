resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}-cluster"
}

resource "aws_ecs_task_definition" "app" {
  family                   = "${var.project_name}-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.task_cpu
  memory                   = var.task_memory
  execution_role_arn       = var.ecs_task_execution_role_arn
  task_role_arn            = var.ecs_task_role_arn

 container_definitions = jsonencode([{
  name  = "${var.project_name}-container"
  image = "${var.ecr_repository_url}:${var.image_tag}"
  portMappings = [{
    containerPort = 8080
    hostPort      = 8080
  }]
  environment = [
    {
      name  = "APP_NAME"
      value = var.project_name
    },
    {
      name  = "RUN_ENV"
      value = var.environment
    },
    {
      name = "APP_PORT"
      value = "8080"
    },
    {
      name = "DYNAMODB_TABLE_NAME"
      value = "Forklift"
    }
  ]
  secrets = [
    {
      name      = "SES_FROM_ADDRESS"
      valueFrom = "/${var.project_name}/${var.environment}/ses_from_address"
    },
    {
      name      = "SES_TO_ADDRESS"
      valueFrom = "/${var.project_name}/${var.environment}/ses_to_address"
    },
  ]
  logConfiguration = {
    logDriver = "awslogs"
    options = {
      "awslogs-group"         = "/ecs/${var.project_name}"
      "awslogs-region"        = var.aws_region
      "awslogs-stream-prefix" = "ecs"
    }
  }
}])
}

resource "aws_ecs_service" "main" {
  name            = "${var.project_name}-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = var.service_desired_count
  launch_type     = "FARGATE"
force_new_deployment = var.force_new_deployment  # これを追加
  network_configuration {
    subnets          = [var.subnet_id]
    security_groups  = [var.ecs_security_group_id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = var.alb_target_group_arn
    container_name   = "${var.project_name}-container"
    container_port   = 8080
  }
}
