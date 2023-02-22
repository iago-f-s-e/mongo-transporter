### Configurações


#### Exemplo

```toml
batch-size=1000
database-name="app-db"
transfer-collections=["users", "products"]
watch-collections=["users"]

[receiver]
  connection="mongodb://localhost:27017/"
  type="mongodb"
  region="local"
  disable-ssl=true

[sender]
  connection="mongodb://remote-server:27017/?directConnection=true"

[[mapping]]
  collection-name="products"
  collection-map="production_products"

[[mapping]]
  collection-name="users"
  collection-map="REMOTE_USERS"
```


#### Propriedades

- **batch-size _(opcional)_:** Define o tamanho do lote de transferência. Se informado um valor inferior a 1000, o valor será sobrescrito como 1000.

---

- **database-name _(obrigatório)_:** Nome do banco de dados que será transferido.

---

- **transfer-collections _(obrigatório)_:** Lista de coleções que serão transferidas.

---

- **watch-collections _(opcional)_:** Lista de coleções que serão monitoradas para refletir as atualizações em outra coleção.

---

- **receiver _(obrigatório)_:** Seção que armazena as informações de conexão do banco de dados destinatário.

  - **connection _(opcional somente para cloud dynamodb)_:** URL de conexão do banco de dados destinatário

    ---

  - **type _(opcional)_:** Tipo do banco de dados destinatário. Lista de tipos compatíveis:
    - mongodb _(default)_
    - dynamodb

    ---

  - **region _(opcional)_:** Região onde está hospedado o banco de dados, exemplos:
    - local _(default)_
    - us-east-1
    - us-east-2
    - us-west-1
    - us-west-2

    ---

  - **access-key-id _(opcional)_:** Chave de acesso (AWS)

    ---

  - **secret-access-key _(opcional)_:** Chave de acesso secret (AWS)

    ---

  - **session-token _(opcional)_:** Token da sessão (AWS)
---

- **sender _(obrigatório)_:** Seção que armazena as informações de conexão do banco de dados remetente.
  - **connection _(obrigatório)_:** URL de conexão do banco de dados remetente.

---

- **mapping _(opcional)_:** Seção que informa as configurações de mapeamento.

  - **collection-name _(obrigatório somente se informado `collection-map`)_:** Coleção do _remetente_ a ser mapeada no _destinatário_

    ---

  - **collection-map _(obrigatório somente se informado `collection-name`)_:** Coleção mapeada no _destinatário_