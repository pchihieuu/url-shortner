The Shortened URL project is a system that helps users shorten long URLs into shorter, more manageable links that are easier to share and manage. 

The system provides features for managing and tracking shortened URLs, along with tools to help users optimize and secure their links.

## Tables of contents

1. [Introduction](#Introduction)
2. [Key Features](#Key-Features)
3. [System Overview](#System-Overview)
4. [Repository Structure](#Repository-Structure)
5. [Installation Guide](#Installation-Guide)
    - [Requirements](#Requirements-📋)
    - [Setup](#Set-up)

## Introduction

-   The [Shortened URL](https://en.wikipedia.org/wiki/URL_shortening) project is a system that allows users to easily shorten long URLs into more concise links, making it easier to share and manage links. By shortening URLs, users can share them via email, social media, or other platforms without worrying about the length of the link.
-   Additionally, the system offers powerful features such as tracking and analyzing click statistics for shortened URLs, managing created links. The system also includes an API that allows external applications to integrate URL shortening functionality into their services.
-   With the goal of optimizing the user experience, the project also enables custom domain options for shortened URLs, making it easier for users to identify and manage their links effectively.

## Key Features

The project focuses on the following main features:

-   URL Shortening: Users can input a long URL and receive a shortened version, making the URL more compact and easy to share on various platforms like social media, email, or messaging apps.
-   Link Management: Users can manage their shortened URLs, view details about them, and perform actions such as deleting or updating the links.
-   Link Expiry: Each shortened URL can have an expiry date, automatically deactivating the link after the specified time.
-   Customizable Shortened URLs: Users can customize the alias or name of their shortened URL, making it more recognizable and meaningful.
-   Rate Limiting: The system includes rate limiting to control the number of requests a user can make, preventing abuse and ensuring fair usage.
-   Link Tagging: Users can assign tags to their shortened URLs for better organization and easy searchability.

## System Overview

The system is designed using a monolithic architecture, with the following technologies:

-   [Go](https://go.dev/): Build backend applications.
-   [Gin Framework](https://gin-gonic.com/): Build backend applications with high performance and support for routing, middleware, and JSON handling.
-   [Redis](https://redis.io/): NoSQL in-memory key-value database.
-   [Docker](https://www.docker.com/): Containerize services.
-   [Docker Compose](https://docs.docker.com/compose/): Manage containers.

<img loading="lazy" src="./docs/images/url_shorter.drawio.svg" alt="Architecture" width="100%" height=600>

## Repository Structure
```plaintext
/url-shortener
├── api/
│   ├── database/
│   │   └── database.go  # Redis client setup
│   ├── models/
│   │   └── models.go    # Models for Request, Response, etc.
│   ├── routes/
│   │   ├── addTag.go    # Add tag to URL
│   │   ├── editUrl.go   # Update URL
│   │   ├── getUrl.go    # Get shortened URL
│   │   └── deleteUrl.go # Delete shortened URL
│   └── utils/
│       └── utils.go     # Utility functions
├── .env                 # Environment variables
├── docker-compose.yml   # Docker Compose configuration
└── main.go              # Application entry point
```
## Installation Guide

### Requirements

To install and run the project, you need to install the following tools. Please follow the installation instructions for your appropriate operating system:

-   [Docker-Installation](https://docs.docker.com/get-docker/)
-   [Docker-Compose-Installation](https://docs.docker.com/compose/install/)
-   [Golang v1.22 Installation](https://go.dev/doc/install)
-   [Git Installation](https://git-scm.com/downloads)

### Setup

1. **Clone the repository**: First, clone the repository to your local machine by running the following command in your terminal

```bash
git clone https://github.com/pchihieuu/url-shortner.git

cd url-shortner
```

2. **Install Dependencies**: Inside the project folder, install the required Go modules by running

```bash
go mod tidy
```

3. **Set Up Redis:**

-   If Redis is not installed, you can either install it locally or run it using Docker.
-   To install Redis locally, follow the [Redis installation guide](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/).
-   If you're using Docker, run the following command to start a Redis container:

```bash
docker run --name redis-container -p 6379:6379 -d redis
```

> **Note:** This will run Redis on the default port 6379.

4. **Configure Environment Variables**: In the project directory, you may need to set environment variables for configurations. For example, create a .env file and define any necessary values, such as Redis host and port:

```bash
DB_ADDR="YOUR_DB_ADDR"
DB_PASS="YOUR_DB_PASS"
APP_PORT="YOUR_PORT"
DOMAIN="YOUR_DOMAIN"
API_QUOTA=YOUR_API_QUOTA
GIN_MODE=YOUR_MODE
```

5.  **Run the Application**: To start the services with a single Docker Compose command, use:

```bash
docker-compose up -d
```

6. **Access the Application**: The server will start running on the default port. You can access the URL shortener API by visiting:

```bash
http://localhost:YOUR_PORT/api/v1
```