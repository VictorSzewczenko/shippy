#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Un-Deploying User Service..."

echoBold "Deleting deployment..."
kubectl delete -f deployment.yaml --namespace micro

echoBold "Deleting services..."
kubectl delete -f service.yaml --namespace micro