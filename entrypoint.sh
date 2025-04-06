#!/bin/sh

echo "Configuring AWS CLI profile"
aws configure set aws_access_key_id ${AWS_CLI_ACCOUNT}
aws configure set region "us-east-1"
aws configure set aws_secret_access_key ${AWS_CLI_TOKEN}


# aws ssm get-parameter --name "${APPNAME}" --query "Parameter.Value"  --profile default


echo "Check if doc exist"
FILE=/docs/swagger.json
if [ -f "$FILE" ]; then
    echo "$FILE exists."
else 
    echo "$FILE does not exist."
fi

term_handler() {
  echo "SIGTERM received, shutting down gracefully..."
  # Aquí puedes colocar cualquier comando de limpieza que necesites
  if [ -n "$APP_PID" ]; then
    kill -TERM "$APP_PID"
    wait "$APP_PID"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}

# Atrapamos SIGTERM y SIGINT (por si usas Ctrl+C)
trap 'term_handler' SIGTERM SIGINT

# Iniciar la aplicación en segundo plano y guardar el PID
./$APPNAME &
APP_PID=$!

# Esperar indefinidamente para que el script no termine
wait "$APP_PID"