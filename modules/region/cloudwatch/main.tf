resource "aws_cloudwatch_log_group" "test_cluster" {
  count = var.create_cloudwatch_log_group ? 1 : 0
  name  = "/ecs/test-cluster"
}
