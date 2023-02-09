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

4. Inicie o container:
  ```bash
  docker-compose up -d
  ```

### Configuração 
O arquivo de configuração `config.toml` permite configurar as informações de origem e destino dos dados. Além disso, é possível configurar outras opções, como a quantidade de goroutine utilizadas para a transferências.

### Execução
Ao executar a aplicação, ela irá aguardar 30 segundos antes de iniciar a transferência de dados. Isto é necessário para garantir que a imagem do MongoDB tenha tempo suficiente para subir.

### Parada
Para parar o container, basta executar o seguinte comando:
  ```bash
  docker-compose down
  ```

### Contribuição
Contribuições são bem-vindas! Sinta-se livre para enviar um pull request com suas alterações.