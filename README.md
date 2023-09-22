# Effective work with SQL in Go
- [Go: ефективна робота з SQL](https://dou.ua/forums/topic/34806/)

### Support Ukraine 🇺🇦
- Help Ukraine via [SaveLife fund](https://savelife.in.ua/en/donate-en/)
- Help Ukraine via [Dignitas fund](https://dignitas.fund/donate/)
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

### Testing
```bash
make env-up
make docker-go-version
make docker-pg-version
make migrate-up
make go-test
make env-down
```
```bash
make env-up
time make test
```
```text
real	0m16.016s
user	0m0.168s
sys	0m0.079s
```

### Examples
```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "./internal/storage/postgres/queries/"
    schema: "./internal/storage/postgres/migrations/"

    gen:
      go:
        package: "dbs"
        sql_package: "database/sql"
        out: "./internal/storage/postgres/dbs"
        emit_prepared_queries: true
```
```sql
-- name: CompanyNewAndGetID :one
INSERT INTO companies (alias, name, created_by, created_at)
VALUES (@alias, @name, @created_by, @created_at)
RETURNING id;
```

### Development
```bash
make create-new-migration-file NAME=migration_name
```
```bash
mkdir -p ./internal/storage/postgres/queries/
```
```bash
make generate-sqlc
```

### Tools
* [github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)
* [github.com/pressly/goose](https://github.com/pressly/goose)
