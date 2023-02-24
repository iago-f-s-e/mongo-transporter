## Mongo Transporter
[![Docker Build and Push](https://github.com/iago-f-s-e/mongo-transporter/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/iago-f-s-e/mongo-transporter/actions/workflows/docker-publish.yml)

<details>
  <summary>Language</summary>
  
  * [Portuguese](https://github.com/iago-f-s-e/mongo-transporter/blob/main/docs/pt-br/README.md)

</details>

---

This is a project for transferring data from one MongoDB database to another. It supports the following databases to act as the *Receiver* (database that will receive the transfer):

- DynamoDB
- MongoDB

With this project, you can easily transfer data from one MongoDB database to another compatible database with the desired settings.

### Summary

- [Prerequisites](#pré-requisitos)
- [Installation with Docker](#instalação-com-docker)
- [Installation with Repository](#instalação-com-repositório)
- [Configuration](#configuração)
- [Execution](#execução)
- [Stop](#parada)
- [Contribution](#contribuição)

### <p name="pré-requisitos">Prerequisites</p>
 - Docker
 - Docker Compose

### <p name="instalação-com-docker">Instalação com Docker</p>
1. Download the project image by running the command:
  ```bash
  docker pull iagofse/mongo-transporter:latest
  ```

2. Create the configuration file with the name `config.toml`:
```toml
batch-size=1000
database-name="app-db"
transfer-collections=["users", "products"]
watch-collections=["users"]

[receiver]
  connection="mongodb://localhost:27017/"
  type="mongodb"
  region="local"

[sender]
  connection="mongodb://remote-server:27017/?directConnection=true"

[[mapping]]
  collection-name="products"
  collection-map="production_products"

[[mapping]]
  collection-name="users"
  collection-map="REMOTE_USERS"
```
Learn more about the configurations [_here_](#Configuration)

3. Start the container in the same directory where you created the configuration file:
  ```bash
  docker run -v $(pwd):/app/config mongo-transporter-test-build    
  ```

### <p name="instalação-com-repositório">Installation with Repository</p>
1. Clone this repository:
  ```bash
  git clone https://github.com/iago-f-s-e/mongo-transporter.git
  ```

2. Enter the project folder:
  ```bash
  cd mongo-transporter
  ```

3. Make a copy of the `config.example.toml` file with the name `config.toml` and fill in the necessary information:
  ```bash
  cp config.example.toml config/config.toml
  ```

---
4. Start the database container (_optional_):
    1. Make a copy the `.env.example` file with the name `.env` and fill in the necessary information:
    ```bash
    cp .drivers/[mongo, dynamo]/.env.example .drivers/[mongo, dynamo]/.env
    ```

    2. Start the container:
    ```bash
    docker-compose -f .drivers/[mongo, dynamo]/docker-compose.yml up -d
    ```

    3. Wait for the container to start, you can check it on the endpoint below:
    ```url
    http://localhost:{UI_PORT}
    ```

    4. To stop the container, just run the following command:
    ```bash
    docker-compose -f .drivers/[mongo, dynamo]/docker-compose.yml down
    ```
---

5. Start the container:
  ```bash
  docker-compose up -d
  ```

### <p name="configuração">Configuration</p>
The `config.toml` configuration file allows you to configure the source and destination information for the data. Additionally, it is possible to configure other options such as the batch size tha will be transferred. Learn mode [_here._](https://github.com/iago-f-s-e/mongo-transporter/blob/main/docs/en/config.md)

### <p name="execução">Execution</p>
When running the application, the container will be started and the project will begin transporting the data. It is possible to check the progress in the container log.

### <p name="parada">Stop</p>
To stop the container, just run the following command:
  ```bash
  docker-compose down
  ```

### <p name="contribuição">Contribution</p>
Contributions are welcome! Feel free to submit a pull request with your changes.