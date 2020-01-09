#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="music-learner-server"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: %S\n" "$app"
# Set all ENV vars for the server to run
export $(grep -v '^#' .env | xargs) && time go run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: %s\n\n" "$app"