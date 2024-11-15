#!/usr/bin/env sh

echo "------- goose: start migrations ---------------------------------"
./goose-custom up
echo "------- goose: finish migrations --------------------------------"
./go-app
