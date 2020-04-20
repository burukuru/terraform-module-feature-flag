provider "aws" {
  region = var.region
}

module "vpc" {
  source = "./vpc"
}

module "ecs-app" {
  source = "./ecs-app"

  cluster_name = var.cluster_name
}

module "cloudwatch" {
  source = "./cloudwatch"

  create_cloudwatch_log_group = var.create_cloudwatch_log_group
}
