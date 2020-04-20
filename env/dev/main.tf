module "global" {
  source = "../../modules/global"
}

module "us_east_1" {
  source = "../../modules/region"
  region = "us-east-1"

  cluster_name                = var.cluster_name
  create_cloudwatch_log_group = var.create_cloudwatch_log_group
}

# module "us_west_1" {
#   source = "../../modules/region"
#   region = "us-west-1"
# }
