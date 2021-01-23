
# create database
resource "google_sql_database" "programservicedb" {
  name = var.db_name
  project = var.project_id
  instance = google_sql_database_instance.postgresql.name
  charset = var.db_charset
  collation = var.db_collation
}


resource "google_sql_database_instance" "postgresql" {
  name = join ("", [var.app_name, "-db1"])
  project = var.project_id
  region = var.gcp_region_1
  database_version = var.db_version
  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = var.db_tier
    activation_policy = var.db_activation_policy
    disk_autoresize = var.db_disk_autoresize
    disk_size = var.db_disk_size
    disk_type = var.db_disk_type
    pricing_plan = var.db_pricing_plan

    location_preference {
      zone = var.gcp_zone_1
    }

    maintenance_window {
      day  = "7"  # sunday
      hour = "3" # 3am
    }

    backup_configuration {
      enabled = true
      start_time = "00:00"
    }

    ip_configuration {
      ipv4_enabled = "false"
      private_network = google_compute_network.vpc-db.self_link
    }
  }
}

# create user
resource "random_id" "user_password" {
  byte_length = 8
}
resource "google_sql_user" "postgresql_user" {
  name = var.db_user_name
  project  = var.app_project
  instance = google_sql_database_instance.postgresql.name
  host = var.db_user_host
  password = var.db_user_password == "" ?random_id.user_password.hex : var.db_user_password
}
