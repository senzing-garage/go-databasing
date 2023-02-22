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
   See [bitnami/postgresql](https://hub.docker.com/r/bitnami/postgresql)
