## Example Go Backend

# API endpoints

Please see [README](./api/README.md) within api folder for list of the endpoints

# Running backend locally

- Make sure you have the .env file with necessary environment variables
- To run open up this repo in vscode, install Golang and within repo root folder run (see debugging if it fails):

```
go run src/main.go
```

# Redeploying backend to google cloud run

- Install google cloud sdk https://cloud.google.com/sdk/docs/install
- Initialize google cloud sdk https://cloud.google.com/sdk/docs/initializing
  At this point only Ryan can run and redeploy service with my cloud account (will look into soon)

Commands to redeploy (must be within src project repo):

```
gcloud config set project chore-project
gcloud run deploy --source .
```

- Specify region to be us-east5
- If asked enable the artifact registry API
- If asked allow for unathenticated invocations

If a 403 error arises an IAM permission is most likely blocking the deploy
(visit IAM permissions and add viewer role onto the @gcloudbuild.gserviceaccount principal)

# Golang setup

Download Golang here: https://go.dev/doc/install

# mysql setup

- This is important for local testing and running of sql commands to google cloud sql
- download mysql community server: https://dev.mysql.com/downloads/mysql/

# Tableplus

- Download tableplus for editing/working with sql database:
- https://tableplus.com/download
- For setup contact Ryan for user/password and host ip

# Debugging

- If you run into a "cannot find package src/main.go" the run the following

```
export GO111MODULE="on"
go mod tidy
```

- If you run into a "Warning: cloudsqlconn.NewDialer:" then it means the database cannot connect due to
  it not finding the default credentails. You will need to run:

```
gcloud auth application-default login
```

This will store the proper credentials within your local machine:

- Linux, macOS: $HOME/.config/gcloud/application_default_credentials.json
- Windows: %APPDATA%\gcloud\application_default_credentials.json
