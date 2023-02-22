## Mongo Transporter
[![Docker Build and Push](https://github.com/iago-f-s-e/mongo-transporter/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/iago-f-s-e/mongo-transporter/actions/workflows/docker-publish.yml)

Este é um projeto para transferência de dados de um banco de dados MongoDB para outro. Ele suporta os seguintes bancos de dados para atuar como *Receiver* (Banco que receberá a transferência):

- DynamoDB
- MongoDB

Com este projeto, você pode facilmente transferir dados de um banco de dados MongoDB para outro banco compatível com as configurações desejadas.


### Sumário

- [Pré-requisitos](#pré-requisitos)
- [Instalação com Docker](#instalação-com-docker)
- [Instalação com Repositório](#instalação-com-repositório)
- [Configuração](#configuração)
- [Execução](#execução)
- [Parada](#parada)
- [Contribuição](#contribuição)


### <p name="pré-requisitos">Pré-requisitos</p>
 - Docker
 - Docker Compose

### <p name="instalação-com-docker">Instalação com Docker</p>
1. Baixe a imagem do projeto executando o comando:
  ```bash
  docker pull iagofse/mongo-transporter:latest
  ```

1. Crie o arquivo de configuração com o nome `config.toml`:
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

Veja mais sobre as configurações [_aqui_](#Configuração)

3. Inicie o container no mesmo diretório que você criou o arquivo de configuração:
  ```bash
  docker run -v $(pwd):/app/config mongo-transporter-test-build    
  ```

### <p name="instalação-com-repositório">Instalação com Repositório</p>
1. Clone este repositório:
  ```bash
  git clone https://github.com/iago-f-s-e/mongo-transporter.git
  ```

2. Entre na pasta do projeto:
  ```bash
  cd mongo-transporter
  ```

3. Crie uma cópia do arquivo `config.example.toml` com o nome `config.toml` e preencha as informações necessárias:
  ```bash
  cp config.example.toml config/config.toml
  ```


---
4. Inicie o container do banco de dados *(opcional)*:
    1. Crie uma cópia do arquivo `.env.example` com o nome `.env` e preencha as informações necessárias:
    ```bash
    cp .drivers/[mongo, dynamo]/.env.example .drivers/[mongo, dynamo]/.env
    ```

    2. Inicie o container:
    ```bash
    docker-compose -f .drivers/[mongo, dynamo]/docker-compose.yml up -d
    ```

    3. Aguarde o container subir, você pode verificar no endpoint abaixo:
    ```url
    http://localhost:{UI_PORT}
    ```

    4. Para parar o container, basta executar o seguinte comando:
    ```bash
    docker-compose -f .drivers/[mongo, dynamo]/docker-compose.yml down
    ```
---


5. Inicie o container:
  ```bash
  docker-compose up -d
  ```

### <p name="configuração">Configuração</p>
O arquivo de configuração `config.toml` permite configurar as informações de origem e destino dos dados. Além disso, é possível configurar outras opções, como o tamanho do lote que será transferido. Veja mais [_aqui._](https://github.com/iago-f-s-e/mongo-transporter/blob/main/docs/config.md)

### <p name="execução">Execução</p>
Ao executar a aplicação, o container será iniciado e o projeto começará a transportar os dados. É possível verificar o progresso no log do container.

### <p name="parada">Parada</p>
Para parar o container, basta executar o seguinte comando:
  ```bash
  docker-compose down
  ```

### <p name="contribuição">Contribuição</p>
Contribuições são bem-vindas! Sinta-se livre para enviar um pull request com suas alterações.