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
pysql -d alurachallenge7
```

