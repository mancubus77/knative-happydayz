# Description

A simple K-Native go application with GET and POST functions

## Build

func build --registry quay.io/mancubus77

## Deployment

### Remote deployment with Tekton pipeline

kn func deploy --remote

### Create KN Service

kn service create happydayz \
--image quay.io/mancubus77/kn-happydayz:latest \
--annotation serving.knative.openshift.io/disableRoute=true
