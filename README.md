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
        --interactive \
        --publish 9174:8080 \
        --rm \
        --tty \
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
       --publish 9171:80 \
       --publish 9172:443 \
       --rm \
       dpage/pgadmin4:latest
    ```
