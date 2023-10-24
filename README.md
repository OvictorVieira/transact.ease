# Transact Ease API Documentation

## Introduction

An intuitive transaction management system, designed for seamless card operations and account handling.

## Production URL

The live application is deployed on [Heroku](https://www.heroku.com/) and can be accessed at [https://transact-ease-605ea8e7ab48.herokuapp.com/](https://transact-ease-605ea8e7ab48.herokuapp.com/).

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Project Structure](#project-structure)
3. [Postman Documentation](#postman-documentation)
4. [Postman Collection](#postman-collection)
5. [Installation](#installation)
6. [Database Structure](#database-structure)
7. [Usage](#usage)
8. [Endpoints](#endpoints)
9. [Testing](#testing)
10. [Release and Deployment](#release-and-deployment)

## Prerequisites

To run local:

1. **Export the `LOCAL` environment variable:**
   ```bash
   export LOCAL=LOCAL
   ```

1. **Create an `.env` file on `internal/config`:**
   ```bash
   cp internal/config/.env.example internal/config/.env
   ```

### Docker

Ensure you have Docker and Docker Compose installed. If not, follow the installation guides:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Project Structure

Below is the directory structure of the Transact Ease API:

1. **cmd:**
    - **api:** Contains the entry point (`main.go`) and server setup (`server.go`) for your application.
    - **helpers:** Houses common helper function(s) (`common_helper.go`) that could be used across different parts of your application.
    - **migration:** Manages database migrations with SQL scripts for creating, indexing, and dropping tables, and a `main.go` for executing these migrations.
    - **seed:** Holds SQL seeding scripts along with a `main.go` file to manage the seeding process.

2. **infra:**
    - **local:** Contains Docker Compose and initial SQL setup for local development, as well as the data directory for your Postgres database.

3. **internal:**
    - **config:** Encapsulates configuration loading and database setup.
    - **constants:** Defines various constant values used across your application.
    - **controllers:** Contains controller logic for handling HTTP requests and responses.
    - **domains:** Holds domain models, interfaces, and mappers for managing data and business logic.
    - **dto (Data Transfer Objects):** Contains request and response structures for handling data transfer between layers.
    - **mocks:** Houses mock implementations for testing purposes.
    - **repositories:** Encapsulates data access logic.
    - **routes:** Manages routing of HTTP requests to the appropriate handlers.
    - **usecases:** Contains use case implementations which orchestrate the flow of data to and from the entities.

4. **pkg:**
    - **helpers:** Contains utility helper function(s).
    - **logger:** Manages logging across your application.
    - **validators:** Manages payload validation.

5. **scripts:** This directory could be used for storing various script files that assist in development, build or deployment processes.

This structure adheres to a clean architecture, separating concerns into different layers and following a logical structure that groups related functionalities together. Each directory serves a clear purpose, aiding in maintaining a clean and well-organized codebase.

## Postman Documentation

To access the documentation click [here](https://documenter.getpostman.com/view/10569183/2s9YRDyqKZ)

## Postman Collection

To test the API endpoints, you can use the provided Postman collection. [Download the collection here](https://github.com/OvictorVieira/transact.ease/files/13126655/API.Documentation.postman_collection.json) and import it into your Postman application.

To the **local** environment use [localhost:8080](http://localhost:3000) as the `url` variable value.

To the **production** environment use [https://transact-ease-605ea8e7ab48.herokuapp.com/](https://transact-ease-605ea8e7ab48.herokuapp.com/) as the `url` variable value.

## Installation

1. **Clone the Repository:**
   ```bash
   git clone git@github.com:OvictorVieira/transact.ease.git
   cd transact.ease
   ```

2. **Use `go mod download` to download the dependencies:**
   ```bash
   go mod download
   ```

3. **Build and Start the Docker Containers:**
   ```bash
   make docker-up
   ```

## Database Structure

### Accounts Table

| Column Name     | Data Type | Description                       |
|-----------------|-----------|-----------------------------------|
| id              | Integer   | Primary Key                       |
| document_number | String    | Document number                   |
| created_at      | Tiemstamp | Date that the account was created |
| updated_at      | Tiemstamp | Date that the account was updated |

### Operation Types Table

| Column Name     | Data Type | Description                                                            |
|-----------------|-----------|------------------------------------------------------------------------|
| id              | Integer   | Primary Key                                                            |
| description     | String    | Operation description (e.g., CASH PURCHASE, INSTALLMENT PURCHASE, etc) |
| password_digest | String    | Encrypted password                                                     |

### Transactions Table

| Column Name       | Data Type | Description                           |
|-------------------|-----------|---------------------------------------|
| id                | Integer   | Primary Key                           |
| account_id        | Integer   | Id of an account                      |
| operation_type_id | Integer   | Id of a operation type                |
| amount            | Float     | Value of a transaction                |
| event_date        | Tiemstamp | Date that the transaction occurred    |
| created_at        | Tiemstamp | Date that the transaction was created |
| updated_at        | Tiemstamp | Date that the transaction was updated |

## Usage

- **Configurations:** Ensure all configurations are set correctly in the `config` directory.
- **Starting the Server:** Use `run-server` to start the server.
- **Using the API:** Use tools like `curl` or Postman to interact with the API.

## Endpoints

- **Accounts:**
    - `GET /api/accounts/:accountId`: Fetches a specific account.
    - `POST /api/accounts`: Creates a new account.

- **Transactions:**
    - `POST /api/transactions`: Creates a new transaction.

## Testing

- **Running Tests:** Use `make run-tests` to run the tests.

## Test Coverage

At Transact Ease API, we prioritize the quality and reliability of our codebase. To ensure that our application is thoroughly tested, we utilize the `coverage` gem to monitor our test coverage.

### Checking Test Coverage

- **Running tests with coverage:** 
   ```bash
   make run-tests-with-cover
   ```

## Release and Deployment

### Publishing via Tags

At Transact Ease API, we believe in structured and organized releases. To ensure consistency and traceability, we use a tag-based release system. This means that every significant change or release of the application is marked with a specific tag, following the [Semantic Versioning](https://semver.org/) convention.

#### How it Works:

1. **Development & Testing:** Once a feature is developed and thoroughly tested, it's merged into the main branch.
2. **Tagging:** Before a release, a new tag is created. This tag represents the version of the application. For example, `v1.0.0`.
3. **Deployment:** Our CI/CD pipeline is triggered by the creation of tags. Once a new tag is pushed, the pipeline automates the deployment process, ensuring that the tagged version of the application is deployed.
4. **Accessing Releases:** You can view all the tags and releases in the "Releases" section of our GitHub repository. Each release will have detailed notes and changes that were introduced in that version.

#### Benefits:

- **Traceability:** Each release is traceable, and you can easily revert to a previous version if needed.
- **Stability:** By releasing through tags, we ensure that only stable and tested features make it to production.
- **Clarity:** Each tag comes with release notes, providing a clear understanding of the changes.
