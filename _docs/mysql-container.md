# MySQL Container Documentation

## Overview

This project uses MySQL 8.4 LTS running in a Docker container for database services. The database data is persisted in a local volume at `.data/mysql` which is excluded from version control.

## Container Configuration

The MySQL container is configured via `docker-compose.yml` with the following settings:

- **MySQL Version**: 8.4 (LTS)
- **Container Name**: snippetbox-mysql
- **Database Name**: snippetbox
- **Port**: 3306 (mapped to host)
- **Data Volume**: `./.data/mysql:/var/lib/mysql`
- **SQL Scripts Volume**: `./sql:/opt/sql:ro` (read-only)

### Credentials

- **Root User**: root / rootpass123
- **Application User**: web / pass123
- **Database**: snippetbox

## Starting the MySQL Container Stack

### Start the container in detached mode (recommended)
```bash
docker compose up -d
```

### Start with logs visible
```bash
docker compose up
```

### Verify the container is running
```bash
docker compose ps
```

## Connecting to MySQL

### Connect via Docker Exec

#### As root user
```bash
docker exec -it snippetbox-mysql mysql -uroot -prootpass123 --default-character-set=utf8mb4
```

#### As application user
```bash
docker exec -it snippetbox-mysql mysql -uweb -ppass123 snippetbox --default-character-set=utf8mb4
```

### Connect from Host Machine

If you have MySQL client installed locally:

```bash
# As root
mysql -h 127.0.0.1 -P 3306 -uroot -prootpass123

# As application user to snippetbox database
mysql -h 127.0.0.1 -P 3306 -uweb -ppass123 snippetbox
```

### Connect from Go Application

```go
// Connection string format
dsn := "web:pass123@tcp(localhost:3306)/snippetbox?parseTime=true"

// Example connection
db, err := sql.Open("mysql", dsn)
```

## Managing the MySQL Container Stack

### Container Lifecycle Commands

#### Stop the container
```bash
docker compose stop
```

#### Stop and remove the container
```bash
docker compose down
```

#### Stop, remove container AND delete volume data (WARNING: Data loss!)
```bash
docker compose down -v
rm -rf .data/mysql
```

#### Restart the container
```bash
docker compose restart
```

### View Container Logs

#### View all logs
```bash
docker compose logs mysql
```

#### Follow logs in real-time
```bash
docker compose logs -f mysql
```

#### View last 100 lines
```bash
docker compose logs --tail=100 mysql
```

### Container Resource Management

#### Check container resource usage
```bash
docker stats snippetbox-mysql
```

#### Inspect container details
```bash
docker inspect snippetbox-mysql
```

## Working with SQL Scripts

### Executing SQL Scripts from MySQL Prompt

SQL scripts are mounted read-only at `/opt/sql` inside the container. Once connected to MySQL, execute scripts using:

```sql
-- Execute a script file
source /opt/sql/create-database.sql;

-- Alternative syntax
\. /opt/sql/create-database.sql
```

### List Available Scripts

From the container shell:
```bash
docker exec -it snippetbox-mysql ls -la /opt/sql/
```

### Important Notes About SQL Scripts

- Scripts are mounted **read-only** and persist even with `docker compose down -v`
- Scripts are **NOT** automatically executed on container startup
- All files in the local `./sql/` directory are available at `/opt/sql/` in the container
- Use the `source` command or `\.` shortcut to execute scripts manually

## Automatic Script Execution

> Note: this feature is not implemented and is only provided for reference.

### About /docker-entrypoint-initdb.d

The MySQL Docker image provides a special directory `/docker-entrypoint-initdb.d` that automatically executes scripts during **initial container creation** (not on restarts). This feature is **not currently configured** in our setup to maintain manual control over script execution.

### How to Enable Automatic Execution

If you wanted to enable automatic script execution, you would modify the volume mount in `docker-compose.yml`:

```yaml
volumes:
  - ./.data/mysql:/var/lib/mysql
  - ./sql:/docker-entrypoint-initdb.d:ro  # Auto-execute on creation
```

### Automatic Execution Behavior

When using `/docker-entrypoint-initdb.d`:

- **Execution Order**: Files are executed in alphabetical order
- **File Types Supported**:
  - `.sql` files - Executed by MySQL
  - `.sh` files - Executed as shell scripts
  - `.sql.gz` files - Decompressed and executed
- **Timing**: Only runs during **initial container creation**
- **One-Time Only**: Scripts do not re-run on container restarts
- **Database Context**: SQL scripts run against the database specified in `MYSQL_DATABASE`

### Example Automatic Setup

```yaml
# docker-compose.yml - Auto-execution example (NOT current config)
services:
  mysql:
    image: mysql:8.4
    environment:
      MYSQL_DATABASE: snippetbox
      # ... other env vars
    volumes:
      - ./.data/mysql:/var/lib/mysql
      - ./sql:/docker-entrypoint-initdb.d:ro
```

With this setup, `create-database.sql` would run automatically when the container is first created.

### Why We Use Manual Execution Instead

Our current configuration uses `/opt/sql` for manual execution because:

- **Explicit Control**: You decide when to run scripts
- **Repeatability**: Scripts can be re-run as needed
- **Development Flexibility**: No need to recreate containers to re-run scripts
- **Debugging**: Easier to troubleshoot script issues manually
- **Selective Execution**: Run only specific scripts when needed

### Switching Between Approaches

To switch from manual to automatic execution:
1. Change volume mount from `./sql:/opt/sql:ro` to `./sql:/docker-entrypoint-initdb.d:ro`
2. Remove existing container and volume: `docker compose down -v`
3. Recreate container: `docker compose up -d`
4. Scripts will execute automatically during container creation

## Working with MySQL CLI

### Basic MySQL Commands

Once connected to MySQL, you can use these commands:

```sql
-- Show all databases
SHOW DATABASES;

-- Use a specific database
USE snippetbox;

-- Show all tables
SHOW TABLES;

-- Describe table structure
DESCRIBE table_name;

-- Show current user
SELECT USER();

-- Show current database
SELECT DATABASE();

-- Show MySQL version
SELECT VERSION();

-- Exit MySQL CLI
EXIT;
-- or
QUIT;
```

### Database Management

```sql
-- Create a new database
CREATE DATABASE IF NOT EXISTS dbname;

-- Drop a database (WARNING: Permanent!)
DROP DATABASE IF EXISTS dbname;

-- Show database size
SELECT
    table_schema AS 'Database',
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)'
FROM information_schema.tables
WHERE table_schema = 'snippetbox'
GROUP BY table_schema;
```

### User Management

```sql
-- Create a new user
CREATE USER 'username'@'%' IDENTIFIED BY 'password';

-- Grant privileges
GRANT ALL PRIVILEGES ON snippetbox.* TO 'username'@'%';

-- Show user privileges
SHOW GRANTS FOR 'web'@'%';

-- Flush privileges after changes
FLUSH PRIVILEGES;
```

## MySQL Best Practices

### 1. Connection Pooling

When connecting from your Go application, use connection pooling:

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

### 2. Query Optimization

- Always use indexes on columns used in WHERE, JOIN, and ORDER BY clauses
- Use EXPLAIN to analyze query performance:
  ```sql
  EXPLAIN SELECT * FROM snippets WHERE expires > UTC_TIMESTAMP();
  ```

### 3. Security Guidelines

- Never store credentials in code; use environment variables
- Always use parameterized queries to prevent SQL injection
- Regularly update passwords and use strong passwords
- Limit user privileges to only what's necessary
- Consider using SSL/TLS for connections in production

### 4. Backup Strategies

#### Manual backup
```bash
# Backup database to file
docker exec snippetbox-mysql mysqldump -uroot -prootpass123 snippetbox > backup_$(date +%Y%m%d).sql

# Backup with compression
docker exec snippetbox-mysql mysqldump -uroot -prootpass123 snippetbox | gzip > backup_$(date +%Y%m%d).sql.gz
```

#### Restore from backup
```bash
# Restore from SQL file
docker exec -i snippetbox-mysql mysql -uroot -prootpass123 snippetbox < backup.sql

# Restore from compressed file
gunzip < backup.sql.gz | docker exec -i snippetbox-mysql mysql -uroot -prootpass123 snippetbox
```

### 5. Performance Monitoring

```sql
-- Show running processes
SHOW PROCESSLIST;

-- Show slow queries
SHOW VARIABLES LIKE 'slow_query%';

-- Check table statistics
SHOW TABLE STATUS FROM snippetbox;

-- Monitor connections
SHOW STATUS LIKE 'Threads%';
```

### 6. Data Integrity

- Always use transactions for related operations
- Define foreign key constraints where appropriate
- Use appropriate data types (don't use VARCHAR(255) for everything)
- Set NOT NULL constraints where data is required

### 7. Development vs Production

**Development Settings** (current docker-compose.yml):
- Simple passwords for convenience
- Direct port exposure (3306)
- Single container setup

**Production Recommendations**:
- Use strong, randomized passwords
- Store credentials in secrets management
- Use private networks, not exposed ports
- Implement regular automated backups
- Enable binary logging for point-in-time recovery
- Configure replication for high availability
- Use SSL/TLS for all connections
- Monitor performance metrics
- Set up alerting for issues

## Troubleshooting

### Container won't start
```bash
# Check logs for errors
docker compose logs mysql

# Ensure port 3306 is not already in use
lsof -i :3306
```

### Can't connect to MySQL
```bash
# Verify container is running
docker compose ps

# Test connection from within container
docker exec -it snippetbox-mysql mysql -uroot -prootpass123 -e "SELECT 1"

# Check firewall rules
sudo iptables -L -n | grep 3306
```

### Data persistence issues
```bash
# Verify volume is mounted
docker inspect snippetbox-mysql | grep -A 5 Mounts

# Check permissions on .data directory
ls -la .data/
```

### Reset everything (WARNING: Data loss!)
```bash
docker compose down -v
rm -rf .data/
docker compose up -d
```

## Additional Resources

- [MySQL 8.4 Documentation](https://dev.mysql.com/doc/refman/8.4/en/)
- [Docker MySQL Official Image](https://hub.docker.com/_/mysql)
- [MySQL Performance Tuning](https://dev.mysql.com/doc/refman/8.4/en/optimization.html)
- [MySQL Security Best Practices](https://dev.mysql.com/doc/refman/8.4/en/security.html)

## Appendix: UTF-8 Character Encoding Configuration

### The Problem

When using MySQL in a Docker container, you may encounter character encoding issues where special characters (em dashes, letters with diacritical marks like ō in "Bashō", etc.) display incorrectly. This happens because:

1. The MySQL Docker image defaults to `latin1` character set for client connections
2. When data is inserted with `latin1` encoding but contains UTF-8 characters, it becomes corrupted ("mojibake")
3. Even if the database and tables are created with `utf8mb4`, the client connection character set matters during data insertion and retrieval

### Symptoms

Data that looks correct in the database:
```
– Matsuo Bashō
```

Appears corrupted when retrieved:
```
â€" Matsuo BashÅ
```

### Solution: Configure UTF-8 as Default

Create a custom MySQL configuration file to set UTF-8 as the default for all connections:

1. Create `mysql-utf8.cnf` in your project root:
```ini
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init-connect = 'SET NAMES utf8mb4'
```

2. Mount this configuration in `docker-compose.yml`:
```yaml
services:
  mysql:
    image: mysql:8.4
    container_name: snippetbox-mysql
    environment:
      MYSQL_ROOT_PASSWORD: rootpass123
    ports:
      - "3306:3306"
    volumes:
      - ./.data/mysql:/var/lib/mysql
      - ./sql:/opt/sql:ro
      - ./mysql-utf8.cnf:/etc/mysql/conf.d/utf8.cnf:ro  # Add this line
    restart: unless-stopped
```

3. Rebuild the container:
```bash
docker compose down
docker compose up -d
```

### Alternative: Command-Line Flags

If you prefer not to use a configuration file, you can specify the character set when connecting:

```bash
# Always add --default-character-set=utf8mb4 when connecting
docker exec -it snippetbox-mysql mysql -uroot -prootpass123 --default-character-set=utf8mb4
```

However, this must be done for every connection, including when running initialization scripts.

### Verifying Character Set Configuration

After implementing the solution, verify the configuration:

```sql
-- Check server and client character sets
SHOW VARIABLES WHERE Variable_name LIKE 'character_set_%' OR Variable_name LIKE 'collation%';
```

You should see `utf8mb4` for all relevant variables:
- character_set_client: utf8mb4
- character_set_connection: utf8mb4
- character_set_database: utf8mb4
- character_set_results: utf8mb4
- character_set_server: utf8mb4

### Note on MySQL Version Differences

The `--skip-character-set-client-handshake` option, which was commonly used in older MySQL versions to force server character set on all clients, was deprecated in MySQL 8.0.35 and **removed in MySQL 8.3.0**. The configuration file approach described above is the recommended solution for MySQL 8.4+.
