# go-databasing

## Specific databases

### SQLite

1. Create empty database.
   Example:

    ```console
    rm -rf /tmp/sqlite
    mkdir  /tmp/sqlite
    touch /tmp/sqlite/G2C.db

    ```

1. View the SQLite database.
   Example:

    ```console
    docker run \
        --env SQLITE_DATABASE=G2C.db \
        --name SqliteWeb \
        --publish 9174:8080 \
        --rm \
        --volume /tmp/sqlite:/data \
        coleifer/sqlite-web

    ```

   Visit [localhost:9174](http://localhost:9174).

### PostgreSQL

1. Create empty database.
   See [bitnami/postgresql](https://hub.docker.com/r/bitnami/postgresql).
   Example:

    ```console
    docker run \
        --env POSTGRESQL_DATABASE=G2 \
        --env POSTGRESQL_PASSWORD=senzing \
        --env POSTGRESQL_POSTGRES_PASSWORD=postgres \
        --env POSTGRESQL_USERNAME=senzing \
        --name postgresql \
        --publish 5432:5432 \
        --rm \
        bitnami/postgresql:latest

    ```

1. View the PostgreSql database.
   Example:

    ```console
    docker run \
        --env PGADMIN_CONFIG_DEFAULT_SERVER='"0.0.0.0"' \
        --env PGADMIN_DEFAULT_EMAIL=postgres@postgres.com \
        --env PGADMIN_DEFAULT_PASSWORD=password \
        --name pgAdmin \
        --publish 9171:80 \
        --publish 9172:443 \
        --rm \
        dpage/pgadmin4:latest

    ```

   Visit [localhost:9171](http://localhost:9171).

### MySQL

1. Create empty database.
   See [bitnami/mysql](https://hub.docker.com/r/bitnami/mysql).
   Example:

    ```console
    docker run \
        --env MYSQL_DATABASE=G2 \
        --env MYSQL_PASSWORD=mysql \
        --env MYSQL_ROOT_PASSWORD=root \
        --env MYSQL_USER=mysql \
        --name mysql \
        --publish 3306:3306 \
        --rm \
        --tty \
        bitnami/mysql:latest

    ```

1. View the MySql database.
   _Caveat:_ The setting of `DATABASE_HOST` may not work in all cases.
   Example:

    ```console
    export DATABASE_HOST=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)

    docker run \
        --env DATABASE_HOST \
        --name phpMyAdmin \
        --publish 9173:8080 \
        --rm \
        bitnami/phpmyadmin:latest

    ```

   Visit [localhost:9173](http://localhost:9173).
