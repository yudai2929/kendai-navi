resource "google_cloud_run_service" "default" {
  name     = "kenda-navi-dev"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

output "url" {
  value = google_cloud_run_service.default.status[0].url
}