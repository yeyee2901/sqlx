FROM mysql:8.0

COPY ./db/*.sql /docker-entrypoint-initdb.d/
