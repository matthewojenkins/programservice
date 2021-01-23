# define the GCP authentication file
variable "gcp_auth_file" {
  type = string
  description = "GCP authentication file"
}
# define GCP project name
variable "project_id" {
  type = string
  description = "GCP project id"
}
# define GCP project name
variable "app_project" {
  type = string
  description = "GCP project name"
}
# define GCP project name
variable "app_name" {
  type = string
  description = "GCP application name"
}
# define GCP region
variable "gcp_region_1" {
  type = string
  description = "GCP region"
}
# define GCP zone
variable "gcp_zone_1" {
  type = string
  description = "GCP zone"
}