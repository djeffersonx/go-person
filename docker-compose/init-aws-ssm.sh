#!/bin/bash
set -e
set -u

endpoint=http://localhost:4566
region=us-east-1

function create_ssm_PARAMETER() {
  echo "creating: $1"
  NAME=$1
  VALUE=$2
  aws --endpoint-url=$endpoint ssm put-PARAMETER --region "$region" \
                                    --name "$NAME" \
                                    --description "Parameter auto created on script init-aws-ssm.sh $NAME" \
                                    --VALUE "$VALUE"
}

echo "starting: create aws ssm PARAMETERs"

$CREATE_PARAMETERS

PARAMETERS=( "/go-person/development/database.url:url-postgresql"
        "/go-person/development/database.username:user-postgresql" )

for PARAMETER in "${PARAMETERS[@]}"; do
    KEY=${PARAMETER%%:*}
    VALUE=${PARAMETER#*:}
    create_ssm_PARAMETER "$KEY" "$VALUE"
done

echo "finished: create aws ssm parameters"