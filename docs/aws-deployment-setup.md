# Setting up Polyglot Co Backend Demo

The following steps are what’s required to setup the backend pipeline so it can automatically deploy after it passes its tests.

## Create an Elastic Beanstalk application

Create a new web app in Elastic Beanstalk called `Polyglot Co`, select `Platform: Docker` and `App code: Sample Application`.

Once it’s created, record the application name (e.g. `polyglot-co`) which you'll need to set in the S3 environment hook below.

## Update the EB App’s Software Configuration

In your EB application’s `Configure` then `Software Configuration` menu, set the following environment variable to point to the API gateway you created during the Lambda setup, and hit `Apply`:

`WEATHER_SERVICE_URL=<url to the API Gateway endpoint>`

For example:

`WEATHER_SERVICE_URL=https://cbz123.execute-api.ap-southeast-2.amazonaws.com/prod/polyglot-co-weather_fetchWeather`

## Allow the agent to deploy the EB app

We need to give permission for the agent to upload a new version of the backend to the EB application’s S3 bucket, and create new versions. To do this, find the bucket name (such as `elasticbeanstalk-ap-southeast-2-534940912648`) and the ARN of the application, and then add an Inline Policy to your agent role that allows it to write to the S3 bucket under the `/deploys/*` key, and create versions and update the environment.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": ["arn:aws:s3:::elasticbeanstalk-ap-southeast-2-534940912648/deploys/*"]
    },
    {
      "Effect": "Allow",
      "Action": [
        "elasticbeanstalk:CreateApplicationVersion",
        "elasticbeanstalk:UpdateEnvironment"
      ],
      "Resource": [
        "arn:aws:elasticbeanstalk:ap-southeast-2:534940912648:*"
      ]
    }
  ]
}
```

## Upload the pipeline environment hook to S3

For the deploy pipeline to work it needs some secrets exposed using the pipeline’s environment hook (which as standard in Buildkite AWS Elastic Stack is autoloaded from `s3://the-stack-secret-bucket/pipeline-slug/env`).

For the deploy pipeline you need to expose some config and secrets in the environment hook. You can do this by uploading the following file to the secrets bucket at `/backend-deploy/env`, just like you did with `lambda-deploy`, substituting in real values for the keys:

```bash
#!/bin/bash

export S3_EB_APP_BUCKET_NAME="<elastic-beanstalk-app-s3-bucket>"
export EB_REGION="<elastic-beanstalk-app-region>"
export EB_APP_NAME="<elastic-beanstalk-app-name>"
export EB_ENVIRONMENT_NAME="<elastic-beanstalk-environment-name>"
```

for example:

```bash
#!/bin/bash

export S3_EB_APP_BUCKET_NAME="elasticbeanstalk-ap-southeast-2-534940912648"
export EB_REGION="ap-southeast-2"
export EB_APP_NAME="polyglot-co"
export EB_ENVIRONMENT_NAME="polyglotco"
```

## Test the Backend pipeline

Create a build on the Backend pipeline and watch the CI and deploy happen :tada:

## All done!

The backend side is all done—you've now got a continuously tested and deployed backend.
