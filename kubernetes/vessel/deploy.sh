#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Deploying Vessel Service..."

echoBold "Creating deployment..."
kubectl create -f deployment.yaml --namespace micro

echoBold "Creating services..."
kubectl create -f service.yaml --namespace micro