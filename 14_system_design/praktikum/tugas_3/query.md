### REDIS

Untuk mendapatkan semua data dari tabel users di Redis:

```shell
KEYS users
```

Gunakan perintah:
```shell
GET
```

Hasil:
```
123456
789012
```

Perintah ini akan mengembalikan value dari kunci tertentu

### Neo4j

Untuk mendapatkan semua data dari tabel users di Neo4j:

```shell
MATCH (u:User)
RETURN u
```

Hasil:
```
{
  "id": "JJKFIG",
  "name": "John Mack Doe",
  "email": "johnmack@example.com"
}
{
  "id": "789012",
  "name": "Liana Dew",
  "email": "lianadew@example.com"
}
```

### Cassandra

Untuk mendapatkan semua data dari tabel users di Cassandra:

```shell
SELECT *
FROM users;
```

Hasil:
```
id | name | email
-- | -- | --
123456 | John Doe | johndoe@example.com
789012 | Jane Doe | janedoe@example.com
```