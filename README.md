# IOHub Story API

Your real-time story application built with Golang, MongoDB, and WebSockets

## Introduction

IOHub Story API is a fast and lightweight chat WebSocket service built with Golang, MongoDB, and WebSockets. It aims to provide a scalable and efficient platform for real-time story applications.

## Features

### Completed Features

- **User Authentication**
  - JWT-based authentication.
  - RSA keys for JWT signing.

- **Real-time Communication**
  - WebSocket support for real-time updates.
  - Real-time "typing..." status display.

- **Rooms**
  - Create, join, and leave rooms.

- **Message Types**
  - Support for text messages.

- **Distributed Systems**
  - NATS Streaming for event-driven architecture.
  
- **Logging**
  - Comprehensive logging with Uber's Zap logger.

- **API and Web Server**
  - API built on top of Fiber (FastHTTP).

- **Database**
  - MongoDB for persistent storage.

## API Documentation with Swagger

You can access the API documentation through Swagger at the following endpoint after starting the application:

[http://localhost:8080/api/v1/swagger/](http://localhost:8080/api/v1/swagger/)

Tip: Before delving into the documentation, ensure to execute `make swagger` at the project root. This step regenerates essential Swagger components, keeping the docs updated.
