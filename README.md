# go-databasing

## Viewing databases

### SQLite

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
