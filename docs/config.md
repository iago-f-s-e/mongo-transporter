### Configurações

#### Propriedades

- **batch-size _(opcional)_:** Define o tamanho do lote de transferência. Se informado um valor inferior a 1000, o valor será sobrescrito como 1000.
- **database-name _(obrigatório)_:** Nome do banco de dados que será transferido.
- **transfer-collections _(obrigatório)_:** Lista de coleções que serão transferidas.
- **watch-collections _(opcional)_:** Lista de coleções que serão monitoradas para refletir as atualizações em outra coleção.
- **receiver _(obrigatório)_:** Seção que armazena as informações de conexão do banco de dados destinatário.
  - **connection _(obrigatório)_:** URL de conexão do banco de dados destinatário
  - **type _(opcional)_:** Tipo do banco de dados destinatário. Lista de tipos compatíveis:
    - mongodb _(default)_
- **sender _(obrigatório)_:** Seção que armazena as informações de conexão do banco de dados remetente.
  - **connection _(obrigatório)_:** URL de conexão do banco de dados remetente.

#### Exemplo

```toml
batch-size=1000
database-name="app-db"
transfer-collections=["users", "products"]
watch-collections=["users"]

[receiver]
  connection="mongodb://localhost:27017/"
  type="mongodb"

[sender]
  connection="mongodb://remote-server:27017/?directConnection=true"
```
