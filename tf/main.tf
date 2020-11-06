provider google {
  project = "ahmet-personal-api"
}


data "google_cloud_run_service" "default" {
  # managed by Google Cloud Build + Cloud Run integration; not TF
  # although there's a bit of cyclic dependency since it uses the
  # service account provisioned below.
  name     = "baklava"
  location = "us-central1"
}

resource "google_service_account" "default" {
  account_id = "baklava"
}

resource "google_cloud_scheduler_job" "job" {
  name        = "baklava"
  description = "updates baklava prices"
  schedule    = "6 8 * * *" # every morning
  time_zone   = "Europe/Istanbul"
  region      = "us-east1" # because that's where GAE zone (hence scheduler) is :(

  attempt_deadline = "120s"
  retry_config {
    retry_count = 0
  }
  http_target {
    http_method = "GET"
    uri = "${element(data.google_cloud_run_service.default.status, 0).url}/run"

    oidc_token {
      service_account_email = google_service_account.default.email
    }
  }
}

data "google_iam_policy" "invoker" {
  binding {
    role = "roles/run.invoker"
    members = [
      "serviceAccount:${google_service_account.default.email}",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "policy" {
  location    = data.google_cloud_run_service.default.location
  project     = data.google_cloud_run_service.default.project
  service     = data.google_cloud_run_service.default.name
  policy_data = data.google_iam_policy.invoker.policy_data
}

output "scheduler_job" {
  value = google_cloud_scheduler_job.job.id
}

output "service_account_email" {
  value = google_service_account.default.email
}


output "cloud_run_url" {
  value = element(data.google_cloud_run_service.default.status, 0).url
}
