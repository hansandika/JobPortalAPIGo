# JobPortal API

This repository contains the backend code for the JobPortal API, which provides endpoints for managing users, jobs, and applications within a job portal platform.

## Overview

The JobPortal API allows users to perform various actions, including user authentication (login, register, logout), job management (create, list), and application management (create, update, list). The API is designed to be user-friendly and easy to integrate into job portal applications.

## Features

- User authentication: Login, register, and logout functionalities.
- Job management: Create new job listings and retrieve a list of available jobs.
- Application management: Submit job applications, update application status, and retrieve application lists.

## Technologies Used

- Programming Language: Go
- Database: [Specify the database you're using, e.g., PostgreSQL, MySQL]
- ORM: GORM
- Web Framework: [Specify the web framework you're using, e.g., Gin, Echo]

## Getting Started

To get started with the JobPortal API, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/jobportal-api.git
   ```

2. Set up the environment variables:

- Copy the .env.example file to .env:

  ```bash
  cp .env.example .env
  ```

- Update the `.env` file with your database connection details, such as the database host, username, password, and database name.

3. Install dependencies:

   ```bash
   cd jobportal-api
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```
