#!/bin/bash

set -euo pipefail

echo "--- :elasticbeanstalk: Creating zip for Elastic Beanstalk"

make eb.zip

echo "+++ :elasticbeanstalk: Deploying to Elastic Beanstalk"

s3_upload_key="deploys/pipeline-co-demo-backend-${BUILDKITE_BUILD_NUMBER}.zip"

aws s3 cp --region "${EB_REGION}" eb.zip "s3://${S3_EB_APP_BUCKET_NAME}/${s3_upload_key}"

existing_version=$(aws --region "${EB_REGION}" elasticbeanstalk describe-application-versions --application-name "${EB_APP_NAME}" --version-labels ${BUILDKITE_COMMIT} --query "ApplicationVersions[0].Status" --output text)

if [[ $existing_version == "None" ]]; then
  aws elasticbeanstalk create-application-version \
    --region "${EB_REGION}" \
    --application-name "${EB_APP_NAME}" \
    --version-label "${BUILDKITE_COMMIT}" \
    --source-bundle S3Bucket="${S3_EB_APP_BUCKET_NAME}",S3Key="${s3_upload_key}"
fi

aws elasticbeanstalk update-environment \
  --region "${EB_REGION}" \
  --environment-name "${EB_ENVIRONMENT_NAME}" \
  --version-label "${BUILDKITE_COMMIT}"
