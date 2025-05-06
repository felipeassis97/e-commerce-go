FROM postgres:14.17

COPY config/docker/db/migrations/*.sql /docker-entrypoint-initdb.d/

EXPOSE 5432