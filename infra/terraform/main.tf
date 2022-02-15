terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.6.0"
    }
  }
}

module "vpc" {
  source = "terraform-google-modules/network/google"
  version = "4.1.0"
  project_id = var.project
  network_name = "shs-web-test001-vpc"
  routing_mode = "GLOBAL"

  subnets = [
    {
      subnet_name   = "shs-subnet"
      subnet_ip     = "10.10.10.0/24"
      subnet_region = var.region
    },
  ]

  routes = [
        {
        name                   = "egress-internet"
        description            = "route through IGW to access internet"
        destination_range      = "0.0.0.0/0"
        tags                   = "egress-inet"
        next_hop_internet      = "true"
    },
  ]
}

module "gke" {
  source = "./gke"
  zone = var.zone
}

module "mysql" {
  source = "./sql"

  credentials = var.credentials_file
  project     = var.project
  region      = var.region
  password    = var.password
}
