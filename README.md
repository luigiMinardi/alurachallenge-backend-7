# How to install and use postgres
## Installing postgres

### Windows
Use your favorite GUI or the WSL to use postgres. [this](https://www.postgresql.org/download/windows/) and [this](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql/) may help you.

### Mac
Install with
```bash
brew install postgresql@15
```
If you want it to turn on by default on reboot run with
```bash
brew services start postgresql 
```
If you want to run just now do
```bash
brew services run postgresql
```

### Linux
Install postgres with the package manager from your distro.
Make sure to have postgres running on your machine.
If you dont have it `start`ed and `enable`d in the systemctl and don't want to do so you can just run this in some terminal to start postgres:

```bash
postgres
```

## How to create your postgres database and tables

In your terminal run:
```bash
createdb alurachallenge7
```

Enter your db with:
```bash
psql -d alurachallenge7
```

Create extension to use images with lo:
```sql
create extension lo;
```

Create table with:
```sql
create table reviews (
id serial primary key,
name varchar,
review varchar,
image lo);
```

Check if everything worked:
```bash
\dt
```
You should see something like that:
```
        List of relations
 Schema |  Name   | Type  | Owner 
--------+---------+-------+-------
 public | reviews | table | your_user
```

Now you can quit the psql with `exit`.

# How to start the project

First you need to connect your db to the project, to do so, copy and rename the `db/db.go.example`
```bash
cp db/db.go.example db/db.go
```
Then in the `dsn` variable change the `your_postgres_user` to your user and `your_postgres_passwd` to your password.
> **The user and password is related to the database and not your OS user**
> The user will most likely be the same as the one printed as `owner` when you ran the `\dt` command.

