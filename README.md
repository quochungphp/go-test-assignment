### Setup database
1. Pull & run docker iamge

```bash
docker run -it -d -p 5432:5432 --name postgres-local -e POSTGRES_PASSWORD=password postgres
```

2. Go into docker container

```bash
docker exec -it postgres-local bash
```

3. Login pg

```bash
psql -h localhost -U postgres
```

4. Create user

```bash
CREATE USER "user_login" WITH PASSWORD 'password';
ALTER ROLE "user_login" WITH SUPERUSER;
```

5. Create Database

```bash
CREATE DATABASE dbtest;
```

6. Export ENV

```bash
export PG_HOST=127.0.0.1
export PG_PORT=5432
export PG_USER=user_login
export PG_PASS=password
export PG_DB=dbtest
export PORT=9090
```

6. Insert test data, please take a look for db.sql

### Install go packages & run app

1. Install go packages

```bash
go get -u ./...
```

2. Init go vendor for go modules management

```bash
go mod vendor
```
