# PubJudge 🧑‍⚖️

PubJudge is a **distributed, scalable online code judging system** inspired by platforms like LeetCode, HackerRank, and Codeforces. It leverages Go (Fiber), MongoDB, Redis, Docker, and Kubernetes to achieve **horizontal scalability**, **real-time code evaluation**, and **microservice isolation**.

---

## 🚀 Project Goals

- ⚙️ Run and evaluate code submissions asynchronously using isolated judge workers  
- 🧵 Event-driven architecture with Redis Pub/Sub and Streams  
- 🛠️ Auto-scalable Kubernetes deployments with HPA (Horizontal Pod Autoscaler)  
- 🐳 Fully containerized microservices for reproducible builds  
- 🔐 Secure JWT-based authentication and authorization  
- 💾 Use MongoDB for users and submissions, Redis for job/event transport  
- 🧱 Clean and modular Go codebase for easier onboarding and collaboration  

---

## 📁 Tech Stack

| Layer         | Technology            |
|---------------|------------------------|
| API Layer     | Go (Fiber)             |
| Worker Engine | Go + Redis Sub/Pub     |
| Message Queue | Redis Streams          |
| Database      | MongoDB                |
| Authentication| JWT                    |
| Container     | Docker                 |
| Orchestration | Kubernetes + HPA       |
| Autoscaling   | Kubernetes HPA         |
| DevOps        | GitHub Actions (planned) |

---

## ☸️ Kubernetes-Native Deployment

PubJudge is cloud-native and optimized for Kubernetes:

### 🧩 Microservices as Independent Pods

- **Gateway Service**  
  - REST API server exposed via Ingress or LoadBalancer  
  - Auto-scalable via CPU/memory metrics  

- **Judge Worker Pods**  
  - Stateless services subscribed to job queue  
  - Can scale up to meet submission demand  

- **Redis & MongoDB**  
  - Support both managed (e.g., Redis Cloud, MongoDB Atlas) or Helm-deployed local versions  

### 🔐 Secrets & Config

- Kubernetes **Secrets** for credentials  
- **ConfigMaps** for service-wide settings  

### 📈 Autoscaling

- HPA is configured for each pod group  
- Easily customizable for memory or CPU utilization thresholds  

### 📦 Containerization

- Dockerfiles for all services with multi-stage builds  
- Clean separation of app and runtime for optimized image size  

---

## 🧭 Architecture Overview

- **Gateway** → Accepts API calls → Validates & enqueues jobs  
- **Worker** → Subscribed to Redis → Runs user code in sandboxed environments  
- **DB** → Stores user info, submissions, results  

---

## 🧪 Running Locally

```bash
git clone https://github.com/pnaskardev/pubjudge.git
cd pubjudge
go run main.go
