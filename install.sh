#!/bin/bash

CLINAME=jackpass
FINAL_CLI_FOLDER_PATH="/usr/local/bin/jackpass/"

go build -ldflags="-s -w" -o $CLINAME

chmod +x $CLINAME

if [ ! -d "$FINAL_CLI_FOLDER_PATH" ]; then
    echo "Directory $FINAL_CLI_FOLDER_PATH not exist. Enter your password for create it."
    sudo mkdir -p "$FINAL_CLI_FOLDER_PATH"
fi

sudo mv "$CLINAME" "$FINAL_CLI_FOLDER_PATH"