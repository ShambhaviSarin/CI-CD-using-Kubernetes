# CI-CD-using-Kubernetes
This repository demonstrates a full CI/CD pipeline for a simple Golang application deployed to Kubernetes using GitHub Actions and Minikube.
It includes:
A Golang HTTP application (helloworld)
Dockerfile to containerize the app
Kubernetes manifests (application.yaml) for Deployment, Service, and HPA
GitHub Actions workflow to automate build, test, and deploy
Features
Automated CI/CD pipeline triggered on every push to main
Unit testing of Golang code
Containerized app with Docker
Kubernetes Deployment with 2 replicas
Horizontal Pod Autoscaling (HPA) based on CPU usage
Minikube NodePort service for testing the app inside the CI environment
