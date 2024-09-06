# go-databasing development

The following instructions are useful during development.

**Note:** This has been tested on Linux and Darwin/macOS.
It has not been tested on Windows.

## Prerequisites for development

:thinking: The following tasks need to be complete before proceeding.
These are "one-time tasks" which may already have been completed.

1. The following software programs need to be installed:
    1. [git]
    1. [make]
    1. [docker]
    1. [go]
    1. [Oracle instant client]

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing-garage
    export GIT_REPOSITORY=go-databasing
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow
   steps in [clone-repository] to install the Git repository.

## Dependencies

1. A one-time command to install dependencies needed for `make` targets.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies-for-development

    ```

1. Install dependencies needed for [Go] code.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies

    ```

## Environment variables

## Lint

1. Run linting.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make lint

    ```

## Test

To run the tests successfully, Sqlite, PostgreSQL, MySQL, and MsSQL databases need to be accessable.

1. :thinking: Add to `LD_LIBRARY_PATH` for Oracle database.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/er/lib:/opt/oracle/instantclient_23_5:${LD_LIBRARY_PATH}
    ```

1. Run tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup test

    ```

## Coverage

Create a code coverage map.

1. Run Go tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean setup coverage

    ```

   A web-browser will show the results of the coverage.
   The goal is to have over 80% coverage.
   Anything less needs to be reflected in [testcoverage.yaml].

## Documentation

1. View documentation.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean documentation

    ```

1. If a web page doesn't appear, visit [localhost:6060].
1. Senzing documentation will be in the "Third party" section.
   `github.com` > `senzing-garage` > `go-databasing`

1. When a versioned release is published with a `v0.0.0` format tag,
the reference can be found by clicking on the following badge at the top of the README.md page.
Example:

    [![Go Reference Badge]][Go Reference]

1. To stop the `godoc` server, run

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## Clean

1. Remove files and bring down docker-compose formation.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## Viewing databases

### View SQLite

1. View the SQLite database.
   Example:

    ```console
    docker run \
        --env SQLITE_DATABASE=G2C.db \
        --interactive \
        --name SqliteWeb \
        --publish 9174:8080 \
        --rm \
        --tty \
        --volume /tmp/sqlite:/data \
        coleifer/sqlite-web

    ```

   Visit [localhost:9174](http://localhost:9174).

### View PostgreSQL

1. Visit [localhost:9171](http://localhost:9171).
    1. **Username:** <postgres@postgres.com>
    1. **Password:** password
1. "Servers" > "senzing"
    1. **Password:** postgres

### View MySQL

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

1. Visit [localhost:9173](http://localhost:9173).
    1. **Username:** root
    1. **Password:** root

### View MsSQL

1. View the MsSql database.
   _Caveat:_ The setting of `ADMINER_DEFAULT_SERVER` may not work in all cases.
   Example:

    ```console
    export ADMINER_DEFAULT_SERVER=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)
    echo ${ADMINER_DEFAULT_SERVER}

    docker run \
        --env ADMINER_DEFAULT_SERVER \
        --name adminer \
        --publish 9177:8080 \
        --rm \
        senzing/adminer:latest

    ```

1. Visit [localhost:9177](http://localhost:9177).
    1. **System:** MS SQL (beta)
    1. **Server:** [Value of ADMINER_DEFAULT_SERVER]
    1. **Username:** sa
    1. **Password:** Passw0rd
    1. **Database:** G2

## Misc

### Misc MsSQL

1. Create the `G2` database instance.
   _Caveat:_ The setting of `DATABASE_HOST` may not work in all cases.
   Example:

    ```console
    export DATABASE_HOST=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)

    docker run \
        --env ACCEPT_EULA=Y \
        --env MSSQL_PID=Developer \
        --env MSSQL_SA_PASSWORD=Passw0rd \
        --name mssql-tools \
        --rm \
        mcr.microsoft.com/mssql-tools:latest /opt/mssql-tools/bin/sqlcmd \
            -P Passw0rd \
            -Q "CREATE DATABASE G2" \
            -S ${DATABASE_HOST},1433 \
            -U sa

    ```

### Misc Db2

1. Create empty database.
   See [ibmcom/db2](https://hub.docker.com/r/ibmcom/db2).
   Example:

    ```console
    docker run \
        --env DB2INST1_PASSWORD=db2inst1 \
        --env LICENSE=accept \
        --interactive \
        --name db2 \
        --privileged \
        --publish 50000:50000 \
        --rm \
        --tty \
        --volume ${GIT_REPOSITORY_DIR}/testdata/db2:/var/custom \
        ibmcom/db2:latest

    ```

## Cleanup

1. Clean up.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean

    ```

## Alternatives

1. Create an empty **PostgreSQL** database.
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
        bitnami/postgresql

    ```

1. Create an empty **MySQL** database.
   See [bitnami/mysql](https://hub.docker.com/r/bitnami/mysql).
   Example:

    ```console
    docker run \
        --env MYSQL_DATABASE=G2 \
        --env MYSQL_PASSWORD=mysql \
        --env MYSQL_ROOT_PASSWORD=root \
        --env MYSQL_USER=mysql \
        --interactive \
        --name mysql \
        --publish 3306:3306 \
        --rm \
        --tty \
        bitnami/mysql

    ```

1. Create an empty **MsSQL** database.
   See [Configure SQL Server settings with environment variables on Linux](https://docs.microsoft.com/en-us/sql/linux/sql-server-linux-configure-environment-variables).
   Example:

    ```console
    docker run \
        --env ACCEPT_EULA=Y \
        --env MSSQL_PID=Developer \
        --env MSSQL_SA_PASSWORD=Passw0rd \
        --name mssql \
        --publish 1433:1433 \
        --rm \
        --tty \
        mcr.microsoft.com/mssql/server

    ```

## References

[clone-repository]: https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/clone-repository.md
[docker]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/docker.md
[git]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/git.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/template-go.svg
[Go Reference]: https://pkg.go.dev/github.com/senzing-garage/template-go
[go]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/go.md
[localhost:6060]: http://localhost:6060/pkg/github.com/senzing-garage/template-go/
[make]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/make.md
[Oracle instant client]: https://github.com/senzing-garage/knowledge-base/blob/main/WHATIS/oracle-instant-client.md
[testcoverage.yaml]: ../.github/coverage/testcoverage.yaml
