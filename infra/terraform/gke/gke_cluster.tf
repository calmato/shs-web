variable "zone" {}

resource "google_service_account" "default" {
  account_id = "gke-account"
  display_name = "gke-account"
}

resource "google_container_cluster" "microservices" {
  name = "shs-web-gke-cluster"
  location = var.zone
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "microservices_nodes" {
  name = "shs-web-node-pool"
  location = var.zone
  cluster = google_container_cluster.microservices.name
  node_count = 3

  node_config {
    preemptible = true
    machine_type = "e2-small"
  }
}
