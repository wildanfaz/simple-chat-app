# simple-chat-app

simple-chat-app is a Golang-based chat application using WebSocket for real-time communication. Integrated with Supabase for message storage, its main feature is peer-to-peer communication, allowing direct interaction between users for a private chat experience.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' (optional). 

    * On Debian/Ubuntu, you can use:

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/):

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/):

    ```bash
    choco install make
    ```

3. Install Docker by following the instructions provided on the [Docker Official Website](https://www.docker.com/products/docker-desktop/)

4. Install Git on [Git](https://www.git-scm.com/downloads) to clone the repositories

## Usage

1. Open new terminal, then setup supabase locally by following the instructions below or on the [Supabase Self-Hosting](https://supabase.com/docs/guides/self-hosting/docker)

    ```
    # Get the code
    git clone --depth 1 https://github.com/supabase/supabase
    
    # Go to the docker folder
    cd supabase/docker
    
    # Copy the fake env vars
    cp .env.example .env
    
    # Pull the latest images
    docker compose pull
    
    # Start the services (in detached mode)
    docker compose up -d

    ```

2. Open new terminal, then clone this repository:

    ```bash
    git clone https://github.com/wildanfaz/simple-chat-app.git
    ```

3. Change to the project directory:

    ```bash
    cd simple-chat-app
    ```

4. Start the application using docker:

    ```bash
    docker compose up
    ```

5. Test the http endpoint using this postman documentation

    * [simple-chat-app-http](https://documenter.getpostman.com/view/22978251/2s9YsNdq7n)

6. Test the ws endpoint

    * Example User 1:
    ```
    ws://127.0.0.1:3000/ws?sender_id=1&receiver_id=2
    Type : Text
    Example Value : Hello
    Example Response :
    {
        "error": false,
        "message": "success",
        "data": {
            "chat_id": 9,
            "room_id": "1_2",
            "sender_id": 1,
            "receiver_id": 2,
            "message": "Hello",
            "created_at": "2024-01-14T07:09:31.928276138Z"
        }
    }
    ```

    * Example User 2:
    ```
    ws://127.0.0.1:3000/ws?sender_id=2&receiver_id=1
    Type : Text
    Example Value : World
    Example Response :
    {
        "error": false,
        "message": "success",
        "data": {
            "chat_id": 10,
            "room_id": "1_2",
            "sender_id": 2,
            "receiver_id": 1,
            "message": "World",
            "created_at": "2024-01-14T07:10:25.193007406Z"
        }
    }
    ```

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```