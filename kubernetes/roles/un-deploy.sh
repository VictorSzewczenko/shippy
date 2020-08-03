#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Un-Deploying RBAC rules"

echoBold "Deleting cluster role..."
kubectl delete -f cluster-role.yaml --namespace micro

echoBold "Deleting cluster role binding..."
kubectl delete -f cluster-role-binding.yaml --namespace micro