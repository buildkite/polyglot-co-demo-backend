#!/bin/bash

set -euo pipefail

echo "--- :elasticbeanstalk: Creating zip for Elastic Beanstalk"

make eb.zip

echo "--- :elasticbeanstalk: Deploying to Elastic Beanstalk"

s3_upload_key="pipeline-co-demo-backend-${BUILDKITE_BUILD_NUMBER}.zip"

aws s3 cp eb.zip "s3://${S3_EB_APP_BUCKET_NAME}/${s3_upload_key}"

aws elasticbeanstalk create-application-version \
  --application-name "${EB_APP_NAME}" \
  --version-label "${BUILDKITE_COMMIT}" \
  --source-bundle S3Bucket="${S3_EB_APP_BUCKET_NAME}",S3Key="${s3_upload_key}"

aws elasticbeanstalk update-environment \
  --environment-name "${EB_ENVIRONMENT_NAME}" \
  --version-label "${BUILDKITE_COMMIT}"
