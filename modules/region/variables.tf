variable "region" {
  default = ""
}

# Cloudwatch module
variable "create_cloudwatch_log_group" {
  type    = bool
  default = true
}
