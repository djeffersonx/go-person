#!/bin/bash
set -e
set -u

endpoint=http://localhost:4566
region=us-east-1

function create_ssm_parameter() {
  echo "creating: $1"
  NAME=$1
  value=$2
  aws --endpoint-url=$endpoint ssm put-parameter --region "$region" \
                                    --name "$NAME" \
                                    --description "Parameter auto created on script init-aws-ssm.sh $NAME" \
                                    --value "$value"
}

echo "starting: create aws ssm parameters"

parameters=( "/go-person/development/database.url:url-postgresql"
        "/go-person/development/database.username:user-postgresql" )

for parameter in "${parameters[@]}"; do
    key=${parameter%%:*}
    value=${parameter#*:}
    create_ssm_parameter "$key" "$value"
done

echo "finished: create aws ssm parameters"