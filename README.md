docker run -it -d -p 5432:5432 --name postgres-local -e POSTGRES_PASSWORD=password postgres

2. Go into docker container
```sh
docker exec -it postgres-local bash
```

3. Login pg
```sh
psql -h localhost -U postgres
```

4. Create user
```bash
CREATE USER "user_login" WITH PASSWORD 'password';
ALTER ROLE "user_login" WITH SUPERUSER;
```

5. Create Database

```
CREATE DATABASE dbtest;
```

6. Export ENV

```
export PG_HOST=127.0.0.1
export PG_PORT=5432
export PG_USER=user_login
export PG_PASS=password
export PG_DB=dbtest
```

6. Insert test data
```
INSERT INTO public.users (id, "password", max_todo) VALUES('userId-1', 'password', 5);
INSERT INTO public.users (id, "password", max_todo) VALUES('userId-2', 'password', 5);

INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('1635608919171', 'Tasks - 1', 'userId-1', '2021-10-30');
INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('1635608980210', 'Tasks - 2', 'userId-1', '2021-10-30');
```
