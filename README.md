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

INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('8dfbbb79-6a58-4ae0-957e-4a377e0d2f22', 'I''m a developer', 'userId-1', '2021-10-31');
INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('6251b656-b334-4948-96c8-ca1a11484ce5', 'I''m a design', 'userId-1', '2021-10-31');
INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('7408f5db-26b1-46e9-9ed9-609f71bc6942', 'I''m a cooker', 'userId-1', '2021-10-31');
INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('31dcd51a-f62e-4268-83a5-8048dfa95c7b', 'I''m a police', 'userId-1', '2021-10-31');
INSERT INTO public.tasks (id, "content", user_id, created_date) VALUES('83b2cea1-c346-464d-92a6-3c7224dd76d3', 'I''m a PR', 'userId-1', '2021-10-31');

```
