# Service account for SQL Proxy
resource "google_service_account" "proxy_account" {
  account_id = "cloud-sql-proxy"
}
resource "google_project_iam_member" "role" {
  role   = "roles/cloudsql.editor"
  member = "serviceAccount:${google_service_account.proxy_account.email}"
}
resource "google_service_account_key" "key" {
  service_account_id = google_service_account.proxy_account.name
}


# SQL Proxy itself
data "google_compute_subnetwork" "regional_subnet" {
  name   = google_compute_network.vpc-db.name
  region = var.gcp_region_1
}

resource "google_compute_instance" "db_proxy" {
  name                      = "db-proxy"
  machine_type              = "f1-micro"
  zone                      = "us-central1-a"
  desired_status            = "RUNNING"
  allow_stopping_for_update = true
  tags = ["ssh-enabled"]
  boot_disk {
    initialize_params {
      image = "cos-cloud/cos-stable"
      size  = 10
      type  = "pd-ssd"
    }
  }
    metadata = {
    enable-oslogin = "TRUE"
  }
  metadata_startup_script = templatefile("${path.module}/run_cloud_sql_proxy.tpl", {
    "db_instance_name"    = "db-proxy",
    "service_account_key" = base64decode(google_service_account_key.key.private_key),
  })
  network_interface {
    network    = var.vpc_name
    subnetwork = data.google_compute_subnetwork.regional_subnet.self_link
    access_config {}
  }
  scheduling {
    on_host_maintenance = "MIGRATE"
  }
  service_account {
    email = module.serviceaccount.email
    scopes = ["cloud-platform"]
  }
}