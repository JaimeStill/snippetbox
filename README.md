# SnippetBox

A web application for storing and retrieving code snippets. Companion app for the book [Let's Go](https://lets-go.alexedwards.net/) by [Alex Edwards](https://github.com/alexedwards).

## Quickstart

Assuming [Initialization](#initialization) has been completed:

1. Run the web server:

```sh
go run ./cmd/web
```

2. Connect to MySQL:

```sh
docker exec -it snippetbox-mysql mysql -uweb -ppass snippetbox
```

Navigate to http://localhost:5000

Test POST:

```sh
curl -iL -d "" http://localhost:5000/snippet/create
```

## Generate TLS Cert

```sh
# executed in ./tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

## MySQL Integration

Instead of directly installing MySQL, I've opted to use a Docker container for the MySQL database. This approach simplifies the setup process and ensures consistency across different development environments.

> For comprehensive details on working with the MySQL container and MySQL in general, reference [MySQL Container Documentation](./_docs/mysql-container.md)

The [container configuration](./docker-compose.yml) is setup to persist data in a volume in a `.data` directory at the repository root. It also exposes read-only access to the [`sql`](./sql) directory on `/opt/sql` so that the SQL infrastructure can be initialized directly from the container.

> You could provide a script to automatically initialize the SQL infrastructure on initialization, but I opted not to take this approach in favor of explicit execution for learning purposes. See [Automatic Script Execution](./_docs/mysql-container.md#automatic-script-execution).

### Initialization

Start the container in detached mode:

```bash
docker compose up -d
```

Verify the container is running:

```bash
docker ps
```

Connect via Docker Exec:

```sh
# as root
docker exec -it snippetbox-mysql mysql -uroot -prootpass123
```

Initialize SQL infrastructure:

```sh
# sequentially execute SQL scripts
source /opt/sql/db-init.sql
```

Exit MySQL and test the infrastructure:

```sh
# connect to mysql as user web
docker exec -it snippetbox-mysql mysql -uweb -ppass snippetbox
```

```sql
-- test select
SELECT id, title, expires FROM snippets;

-- test drop permissions
DROP TABLE snippets;
```

You should see the following outputs:

```
mysql> SELECT id, title, expires FROM snippets;
+----+------------------------+---------------------+
| id | title                  | expires             |
+----+------------------------+---------------------+
|  1 | An old silent pond     | 2026-08-15 14:45:07 |
|  2 | Over the wintry forest | 2026-08-15 14:45:07 |
|  3 | First autumn morning   | 2025-08-22 14:45:07 |
+----+------------------------+---------------------+
3 rows in set (0.00 sec)

mysql> DROP TABLE snippets;
ERROR 1142 (42000): DROP command denied to user 'web'@'localhost' for table 'snippets'
```

### Shutdown

Stop the container:

```bash
docker compose stop
```

Stop and remove the container:

```bash
docker compose down
```

Stop and remove the container and delete the volume:

```bash
docker compose down -v
sudo rm -rf .data/
```

## References

Helpful links discovered throughout this book.

- [Alex Edwards Blog](https://www.alexedwards.net/blog)
- [Tour of Go](https://go.dev/tour/welcome/1)
- [Go Modules](https://go.dev/wiki/Modules)
- [Go Standard Library](https://pkg.go.dev/std)
- [Go Project Layout](https://go.dev/doc/modules/layout)
- [Go SQL Drivers](https://go.dev/wiki/SQLDrivers)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [Go sql.NullString](https://gist.github.com/alexedwards/dc3145c8e2e6d2fd6cd9)
- [Go templates](https://pkg.go.dev/text/template)
- [OWASP Secure Headers Project](https://owasp.org/www-project-secure-headers/)
- [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/CSP)
- [Embedding in Go](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)
- [Go Generics](https://go.dev/doc/tutorial/generics)
- [GopherCon 2021 Generics!](https://www.youtube.com/watch?v=Pa_e9EeCdy8)
- [Let's Encrypt](https://letsencrypt.org/)
- [mkcert](https://github.com/FiloSottile/mkcert)
- [HTTP/2 in Go](https://www.youtube.com/watch?v=FARQMJndUn0)
