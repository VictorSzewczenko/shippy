#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

echoBold(){
   echo "${bold}$1${normal}"
}

echoBold "Creating Fast storage class..."
kubectl create -f storage-class.yaml --namespace micro