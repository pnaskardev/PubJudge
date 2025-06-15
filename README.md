
# PubJudge 🧠📦

PubJudge is a scalable, modular, and clean-code-based backend system inspired by platforms like LeetCode, designed to run, evaluate, and manage code submissions. It follows a layered Clean Architecture pattern with the aim of achieving:

- High scalability and testability

- Clear separation of concerns

- Easy maintenance and extension

- Support for multiple data sources (MongoDB, Redis, PostgreSQL)

- Support horizontal scalability using Kubernetes and autoscaling.



## 🏗️ Project Structure

This project uses a Go monorepo with workspaces. The two main components are:

### 1.Gateway (/gateway)
 Responsible for:

- Exposing REST APIs via Fiber

- Handling routes, services, and request validation

- Hydrating routes with necessary dependencies (DB, cache, config)

### 2. Worker (/worker)
Responsible for:

- Async job processing (code execution, test evaluation)

- Pub/Sub communication using Redis queues
## 🎯 Goals of PubJudge
- 🧪 Accept user code submissions

- ⚙️ Queue and evaluate code against test cases in an async worker system

- 🔐 Secure user auth with hashed passwords and JWT

- 🔍 Use MongoDB for users and submissions storage

- 💨 Use Redis for in-memory pub-sub queues and cache

- 📈 Scalable & Kubernetes-ready deployment

- 📦 Clean layered architecture for clear separation:

    - Entities: Core business models (e.g., User, Submission)
    
    -  Repository: Interfaces to interact with the database

    - Service: Application logic that uses the repositories

    - Handlers: Fiber-compatible HTTP request handlers

    - Presenter: Unified response formatting layer





