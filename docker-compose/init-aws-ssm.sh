#!/bin/bash
set -e
set -u

endpoint=http://localhost:4566
region=us-east-1

function create_ssm_parameter() {
  echo "creating: $1"
  name=$1
  value=$2
  aws --endpoint-url=$endpoint ssm put-parameter --region "$region" \
                                    --name "$name" \
                                    --description "Parameter auto created on script init-aws-ssm.sh $name" \
                                    --value "$value"
}

echo "starting: create aws ssm parameters"

parameters=( "/go-person/dev/database.url:url-postgresql"
        "/go-person/dev/database.username:user-postgresql" )

for parameter in "${parameters[@]}"; do
    key=${parameter%%:*}
    value=${parameter#*:}
    create_ssm_parameter "$key" "$value"
done

echo "finished: create aws ssm parameters"