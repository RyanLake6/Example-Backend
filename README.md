## Backend for Chore Project

# API endpoints

Please see [README](./api/README.md) within api folder for list of the endpoints

# running backend locally

To run open up this repo in vscode, install Golang and within repo root folder run:

```
go run src/main.go
```

# redeploying backend to google cloud run

- Install google cloud sdk https://cloud.google.com/sdk/docs/install
- Initialize google cloud sdk https://cloud.google.com/sdk/docs/initializing
  At this point only Ryan can run and redeploy service with my cloud account (will look into soon)

Commands to redeploy (must be within src project repo):

```
gcloud config set project chore-project
gcloud run deploy --source .
```

Specify region to be us-east5
If asked enable the artifact registry API
If asked allow for unathenticated invocations

If a 403 error arises an IAM permission is most likely blocking the deploy
(visit IAM permissions and add viewer role onto the @gcloudbuild.gserviceaccount principal)

# Golang setup

Download Golang here: https://go.dev/doc/install

# Debugging

If you run into a "cannot find package src/main.go" the run the following

```
export GO111MODULE="on"
go mod tidy
```
