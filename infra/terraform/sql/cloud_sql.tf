variable "credentials" {}
variable "project" {}
variable "region" {}
variable "password" {}

resource "google_sql_database_instance" "master" {
  name             = "shs-web-sql"
  database_version = "MYSQL_8_0"
  region = var.region

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_user" "root" {
    name = "root"
    instance = google_sql_database_instance.master.name
    password = var.password
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.master.name
}
