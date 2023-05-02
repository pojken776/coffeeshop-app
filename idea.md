# Purpose

Create a project that I can use to showcase my abilities in:

- CI/CD: Create build pipelines
- Infrastructure as Code: Deploy the infrastructure on which the app runs using Infrastructure as Code
- GitOps: Deploy the application and maintain using ArgoCD
- Kubernetes: Hosts the application containers
- Certificate Management: Use letsencrypt to manage the site certificate
- Database: Store the data in a light sql database
- Golang: Write the application
- Cloud
- Observability: Observe key metrics, configure alerting
- Message queue: demonstrate understanding of decoupling services in a microservices architecture

## Site Idea

### Stage 1: Create a backend API

CRUD Operations

1. fetch data related to the coffee shop from the database
2. add new coffee shops to the database
3. edit an existing coffee shop entry in the database
4. delete a coffee shop in the database

Tech:

- golang
- database
- message queue
- unit testing/test driven development

## Stage 2: Create a frontend

1. main page
   1. Display the coffee shops
   2. allow sorting the shops by attribute
   3. allow filtering
2. about page

## Stage 3: Build Container Images

1. Use github actions to build image
   1. run unit/integration tests
   2. scan for vulnerabilities
2. Use github actions to push image to container repo

> free tier docker

## Stage 4: Deploy Application on Kubernetes

- deploy a free k8s instance locally
- deploy the container instances
- deploy the application via gitops

> minikube

## Stage 5: Observability

- setup observability metrics + alerting
- setup logging monitoring
- setup tracing

## Stage 6: Deploy Application on Cloud Infrastructure

- setup free tier cloud k8s
- purchase a DNS
- deploy the application onto cloud infrastructure
- setup security
  - WAF
  - DDOS prevention
