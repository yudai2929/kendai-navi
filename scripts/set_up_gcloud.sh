#!/bin/bash

# Set up the project ID
PROJECT_ID="xxx"

# Set up the region
REGION="asia-northeast1"

# Set up the zone
ZONE="asia-northeast1"


# Set up gcloud
gcloud init

# Set up the project
gcloud config set project $PROJECT_ID

# Set up the region
gcloud config set compute/region $REGION

# Set up the zone
gcloud config set compute/zone $ZONE

