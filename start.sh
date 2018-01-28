#!/usr/bin/env bash
echo "Installing NPM dependencies with yarn"
(cd app && yarn install)

echo "Running go server"
(cd server && go run main.go) &
echo "Running frontend application"
(cd app && yarn start) &
