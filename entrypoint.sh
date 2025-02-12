#!/bin/sh

echo "Congiring AWS CLI profile"
aws configure set aws_access_key_id ${AWS_CLI_ACCOUNT}
aws configure set region "us-east-1"
aws configure set aws_secret_access_key ${AWS_CLI_TOKEN}
echo "Check if doc sxist"
FILE=/doc/swagger.json
if [ -f "$FILE" ]; then
    echo "$FILE exists."
else 
    echo "$FILE does not exist."
fi


echo "Download parameters APP $APPNAME"
mkdir configs;
aws ssm get-parameter --name "${APPNAME}" --query "Parameter.Value" --output text   --profile default > configs/service.json ; 
cat configs/service.json | yq -P > configs/service.yaml;

echo  "Starting APP $APPNAME"

echo "..######..########....###....########..########";
echo ".##....##....##......##.##...##.....##....##...";
echo ".##..........##.....##...##..##.....##....##...";
echo "..######.....##....##.....##.########.....##...";
echo ".......##....##....#########.##...##......##...";
echo ".##....##....##....##.....##.##....##.....##...";
echo "..######.....##....##.....##.##.....##....##...";


./$APPNAME