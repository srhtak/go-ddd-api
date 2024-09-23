# Go DDD API with Docker and Kubernetes

This project demonstrates a simple Go API built with Domain-Driven Design (DDD) principles, containerized with Docker, and deployed on Kubernetes using Minikube. It also includes monitoring setup with Prometheus and Grafana.

## Prerequisites

- Go 1.16+
- Docker
- Minikube
- kubectl
- Helm 

## Project Structure

```
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── domain
│   ├── application
│   ├── infrastructure
│   └── interfaces
├── k8s
│   ├── deployment.yaml
│   └── service.yaml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-ddd-api.git
   cd go-ddd-api
   ```

2. Build the Docker image:
   ```
   docker build -t go-ddd-api:latest .
   ```

3. Start Minikube:
   ```
   minikube start
   ```

4. Load the Docker image into Minikube:
   ```
   minikube image load go-ddd-api:latest
   ```

5. Apply Kubernetes manifests:
   ```
   kubectl apply -f k8s/
   ```

6. Verify the deployment:
   ```
   kubectl get pods
   kubectl get services
   ```

## Accessing the API

1. Get the Minikube IP:
   ```
   minikube ip
   ```

2. Get the NodePort of the service:
   ```
   kubectl get service go-ddd-api -o jsonpath='{.spec.ports[0].nodePort}'
   ```

3. Access the API at `http://<minikube-ip>:<node-port>`

## Monitoring with Prometheus and Grafana

1. Install Prometheus and Grafana using Helm:
   ```
   helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
   helm repo update
   helm install monitoring prometheus-community/kube-prometheus-stack
   ```

2. Expose Grafana:
   ```
   kubectl apply -f k8s/grafana-nodeport.yaml
   ```

3. Access Grafana:
   - URL: `http://<minikube-ip>:30300`
   - Default username: `admin`
   - Get password: 
     ```
     kubectl get secret monitoring-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
     ```

4. Configure Prometheus data source in Grafana:
   - URL: `http://monitoring-kube-prometheus-prometheus:9090`

5. Import dashboards or create your own to visualize API metrics.

## Troubleshooting

- If pods are not starting, check logs:
  ```
  kubectl logs <pod-name>
  ```

- For detailed pod information:
  ```
  kubectl describe pod <pod-name>
  ```

- If the API is not accessible, ensure the service is running:
  ```
  kubectl get services
  ```


