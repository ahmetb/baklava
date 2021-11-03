# baklava

## Set up

Re-deploy the Cloud Run app:

```
gcloud run deploy baklava \
    --project=ahmet-personal-api \
    --platform=managed --region=us-central1 \
    --image=$(KO_DOCKER_REPO=gcr.io/ahmet-personal-api/baklava ko publish .)
```

Re-deploy the scheduler setup:

```
terraform -chdir=./tf apply
```
