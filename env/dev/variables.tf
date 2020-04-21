variable "region" {
  type    = string
  default = "us-east-1"
}
variable "cluster_name" {
  type    = string
  default = "test_cluster"
}
variable "create_cloudwatch_log_group" {
  type    = bool
  default = true
}
