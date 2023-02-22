### Configurations


#### Example

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


#### Properties

- **batch-size _(optional)_:** This setting defines the transfer batch size. If a value lower than 1000 is provided, it will be overwritten with a value of 1000.

---

- **database-name _(required)_:** Name the database that will be transferred.

---

- **transfer-collections _(required)_:** List of collection to be transferred.

---

- **watch-collections _(optional)_:** List of collection that will be monitored to reflect updates to another collection.

---

- **receiver _(required)_:** Section that stores the destination database connection information.

  - **connection _(optional only AWS DynamoDB)_:** Destination database connection URL.

    ---

  - **type _(optional)_:** Destination database type. List of compatible types:
    - mongodb _(default)_
    - dynamodb

    ---

  - **region _(optional)_:** Region where the database is hosted. Examples:
    - local _(default)_
    - us-east-1
    - us-east-2
    - us-west-1
    - us-west-2

    ---

  - **access-key-id _(optional)_:** IAM user access key.

    ---

  - **secret-access-key _(optional)_:** IAM user secret key.

    ---

  - **session-token _(optional)_:** IAM user session token.
---

- **sender _(required)_:** This section stores the connection information of the source database.
  - **connection _(required)_:** Source database connection URL.

---

- **mapping _(optional)_:** This section informs the mapping configuration settings.

  - **collection-name _(required only if `collection-map` is informed)_:** Collection of the _sender_ to be mapped in the _receiver_.

    ---

  - **collection-map _(required only if `collection-name` is informed)_:** Collection mapped on the _receiver_ database.