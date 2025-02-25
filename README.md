# Task Manager DevOps Project

A modern, scalable Task Manager application built as a microservices project to demonstrate DevOps best practices. It features a Go backend with SQLite for lightweight, persistent task management and a React frontend with Vite and Tailwind CSS for a sleek, responsive user interface. This project is designed for containerization, CI/CD pipelines, and cloud deployment, making it an ideal showcase for DevOps skills.

## Overview

- **Backend**: Written in Go (Golang) using the Gin framework, with SQLite for efficient, file-based task storage. Supports CRUD operations (Create, Read, Update, Delete) with rate limiting, error handling, and CORS support.
- **Frontend**: Built with React, Vite for fast development and builds, and Tailwind CSS for a modern, utility-first UI. Provides an interactive task list with add, toggle completion, and delete functionality.
- **DevOps Features**: Ready for Docker containerization, GitHub Actions for CI/CD, Terraform for infrastructure as code, and AWS deployment (EKS, ECS, or EC2). Includes monitoring readiness with Prometheus/Grafana and GitOps with ArgoCD.

## Structure

- `Taks-Manager-Backend/`: Go backend code, including SQLite database logic, Gin API endpoints, and configuration.
- `Task-Manager-Frontend/`: React frontend code, styled with Tailwind CSS and powered by Vite for development.

## Tools & Technologies

- **Languages**: Go, JavaScript (React)
- **Frameworks/Libraries**: Gin, React, Vite, Tailwind CSS
- **Database**: SQLite (embedded, file-based for low-memory systems)
- **DevOps Tools**: Docker, GitHub Actions, SonarQube (for code quality), Nexus Artifactory (for artifacts), Terraform, ArgoCD, Prometheus, Grafana, NGINX, AWS

## Features

- **Task Management**: Add, view, toggle completion, and delete tasks via a RESTful API.
- **Robustness**: Includes error handling, rate limiting, and CORS configuration for cross-origin requests.
- **Scalability**: Designed for containerization and cloud deployment, with monitoring and GitOps readiness.
- **Performance**: Lightweight Go backend (compiled binary) and fast React frontend (Vite) for efficient operation.

## Getting Started

### Prerequisites

- Go (1.21 or higher)
- Node.js and npm (for frontend)
- Docker (optional, for containerization)
- Git

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/username/Task_Manager_DevOps.git
   cd Task_Manager_DevOps
   ```
2. **Backend Setup**
    - Navigate to the backend directory:
        ```bash
        cd Task-Manager-Backend
        ```
    - Ensure Go is installed and the `modernc.org/sqlite` package is available (handled by `go.mod`).
    - Run the backend:
        ```bash
        go run cmd/main.go
        ```
    - The API will be available at `http://localhost:8080`. Verify it’s running by testing with `curl http://localhost:8080/tasks` or using the frontend.

2. **Frontend Setup**
    - Navigate to the frontend directory:
        ```bash
        cd Task-Manager-Frontend
        ```
    - Install Dependencies:
        ```bash
        npm install
        ```
    - Run frontend
        ```bash
        npm run dev
        ```
    - The app will be available at `http://localhost:5173`. Open it in a browser to interact with the Task Manager.

### Usage
- Add a task by typing in the input field and clicking "Add."
- Toggle task completion with the checkbox (strikes through the text).
- Delete tasks with the "Delete" button.
- Use the browser’s dev tools to inspect API calls to `http://localhost:8080/tasks`.

## DevOps Integration
- Containerization: Dockerfiles for both backend and frontend are planned (see Dockerfile examples in future commits or documentation).
- CI/CD: GitHub Actions workflows will automate builds, tests, and deployments of Docker containers.
- Monitoring: Integrate Prometheus for metrics collection and Grafana for visualization, with NGINX as a reverse proxy or load balancer.
- Infrastructure: Use Terraform to define AWS infrastructure (e.g., EKS, ECS, or EC2), and ArgoCD for GitOps-based deployments.
- Code Quality & Artifacts: Plan to integrate SonarQube for code analysis and Nexus Artifactory for managing Docker images and build artifacts.

## Future Enhancements
- Add task editing, sorting, and filtering in the frontend.
- Implement Helm charts for Kubernetes deployment on AWS EKS.
- Enhance monitoring with custom Prometheus alerts and Grafana dashboards.
- Integrate additional DevOps tools like Kubernetes monitoring (e.g., kube-state-metrics) and security scanning (e.g., Trivy).

## Contributing
Contributions are welcome! Please follow these steps:

1. Fork this repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request on GitHub.

## License
This project is open-source under the MIT License. See the `LICENSE` file for details.

## Contact
- Author: [Shubham Pawar/ShubhamPawar-3333]
- Email: [shuubham.pawar.368@gmail.com]
- GitHub: [https://github.com/ShubhamPawar-3333]