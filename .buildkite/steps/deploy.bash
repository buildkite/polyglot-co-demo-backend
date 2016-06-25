#!/bin/bash

set -euo pipefail

echo "--- :elasticbeanstalk: Creating zip for Elastic Beanstalk"

make eb.zip

echo "+++ :elasticbeanstalk: Deploying to Elastic Beanstalk"

s3_upload_key="deploys/pipeline-co-demo-backend-${BUILDKITE_BUILD_NUMBER}.zip"

aws s3 cp --region "${EB_REGION}" eb.zip "s3://${S3_EB_APP_BUCKET_NAME}/${s3_upload_key}"

aws elasticbeanstalk create-application-version \
  --region "${EB_REGION}" \
  --application-name "${EB_APP_NAME}" \
  --version-label "${BUILDKITE_BUILD_NUMBER}" \
  --source-bundle S3Bucket="${S3_EB_APP_BUCKET_NAME}",S3Key="${s3_upload_key}"

aws elasticbeanstalk update-environment \
  --region "${EB_REGION}" \
  --environment-name "${EB_ENVIRONMENT_NAME}" \
  --version-label "${BUILDKITE_BUILD_NUMBER}"
