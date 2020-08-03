#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Deleting fast storage class..."
kubectl delete -f storage-class.yaml --namespace micro