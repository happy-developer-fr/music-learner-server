#!/bin/bash
source ./image.sh
echo version : $IMAGETAG
sudo docker stop $IMAGETAG
sudo docker build --tag ${IMAGETAG} ../
