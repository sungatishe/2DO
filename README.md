# 2DO Project

2DO is a microservices-based task management application. The project is organized as a set of independent services, each handling different parts of the system's functionality, including user management, authentication, task management, and notifications.


## Architecture

2DO follows a microservices architecture, where each service is isolated and communicates with others through gRPC. This setup allows for high scalability and maintainability, making it easy to add new features or modify existing ones independently.

## Technologies Used

- **Backend**: Golang, gRPC
- **Frontend**: Vue.js
- **Database**: PostgreSQL
- **Cache**: Redis
- **Queue**: RabbitMQ
- **Containerization**: Docker
- **Build Automation**: Makefile

## Services

Each service is located in a dedicated directory and performs a specific function:

- **api-gateway**: Entry point for client requests, handling routing to various services.
- **auth-service**: Manages user authentication and authorization.
- **user-service**: Manages user profiles and user data.
- **todo-service**: Manages tasks (todos) for users.
- **push-service**: Handles notifications, sending reminders based on task deadlines.
- **frontend**: Vue.js-based frontend that interacts with the `api-gateway`.


