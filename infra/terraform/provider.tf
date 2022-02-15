provider "google" {
  version = "4.6.0"
  credentials = file(var.credentials_file)
  project = var.project
  region  = var.region
  zone    = var.zone
}
