variable "project_id" {
  description = "The GCP project ID"
}
variable "default_region" {
  description = "The default region for resources"
  default     = "asia-southeast1"
}
variable "credentials_path" {
  description = "Path to the service account key file"
}
