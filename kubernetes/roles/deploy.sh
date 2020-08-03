#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Deploying RBAC rules"

echoBold "Creating cluster role..."
kubectl create -f cluster-role.yaml --namespace micro

echoBold "Creating cluster role binding..."
kubectl create -f cluster-role-binding.yaml --namespace micro