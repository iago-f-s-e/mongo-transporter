## Mongo Transporter

Este é um projeto para transferência de dados de um banco de dados MongoDB para outro. Ele suporta os seguintes bancos de dados para atuar como *Receiver* (Banco que receberá a transferência):

- MongoDB

Com este projeto, você pode facilmente transferir dados de um banco de dados MongoDB para outro banco compatível com as configurações desejadas.

### Pré-requisitos
 - Docker
 - Docker Compose

### Instalação
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
    cp .env.example .env
    ```

    2. Inicie o container:
    ```bash
    docker-compose -f dev-db.docker-compose.yml up -d
    ```

    3. Aguarde o container subir, você pode verificar no endpoint abaixo:
    ```url
    http://localhost:{MONGO_EXPRESS_PORT}
    ```

    4. Para parar o container, basta executar o seguinte comando:
    ```bash
    docker-compose -f dev-db.docker-compose.yml down
    ```
---


5. Inicie o container:
  ```bash
  docker-compose up -d
  ```

### Configuração 
O arquivo de configuração `config.toml` permite configurar as informações de origem e destino dos dados. Além disso, é possível configurar outras opções, como o tamanho do lote que será transferido. Veja mais [_aqui._](./docs/config.md)


### Execução
Ao executar a aplicação, o container será iniciado e o projeto começará a transportar os dados. É possível verificar o progresso no log do container.

### Parada
Para parar o container, basta executar o seguinte comando:
  ```bash
  docker-compose down
  ```

### Contribuição
Contribuições são bem-vindas! Sinta-se livre para enviar um pull request com suas alterações.