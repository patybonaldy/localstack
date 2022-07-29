#!/usr/bin/env bash

set -euo pipefail

echo "configuring aws infrastructure"
echo "==================="
export LOCALSTACK_HOST=localhost
export AWS_DEFAULT_REGION=eu-central-1
export AWS_ACCESS_KEY_ID=irrelevant
export AWS_SECRET_ACCESS_KEY=irrelevant


create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE} --attributes VisibilityTimeout=30
}

create_bucket() {
    local BUCKET_NAME_TO_CREATE=$1
    aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 s3 mb s3://${BUCKET_NAME_TO_CREATE}
}

create_secret(){
  aws --endpoint http://localhost:4566 secretsmanager create-secret --name bd-test --description "Some secret" --secret-string '{"user":"adm","password":"adm","decryptionkey":"123"}' --tags '[{"Key":"user", "Value":"adm"},{"Key":"password","Value":"adm"},{"Key":"decryptionkey","Value":"123"}]' --region us-east-1
}

# create queues
for queue in "queue1"
do
  create_queue $queue
done


# create buckets
for bucket in "bucket1"
do
  create_bucket $bucket
done

echo "finished running localstack bootstrap"
