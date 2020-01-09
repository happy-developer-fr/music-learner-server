#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
outputPath="build"
app="music-learner-server"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"

printf "\nBuilding: %s\n" "$app"
time go build -o $output $src
printf "\nBuilt: %s size:" "$app"
ls -lah $output | awk '{print $5}'
printf "\nDone building: %s\n\n" "$app"
