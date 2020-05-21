FROM ubuntu:18.04
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV DEBIAN_FRONTEND=noninteractive
ENV PGVER 10
ENV POSTGRES_HOST /var/run/postgresql/
ENV POSTGRES_PORT 5432
ENV POSTGRES_DB pinterest
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password

EXPOSE $POSTGRES_PORT

RUN apt-get update && apt-get install -y postgresql-$PGVER

USER postgres

COPY build/scripts.sql scripts.sql

RUN service postgresql start &&\
    psql -U postgres -c "ALTER USER postgres PASSWORD 'password';" &&\
    psql -U postgres -c 'CREATE DATABASE "pinterest";' &&\
    psql -U postgres -d pinterest -a -f scripts.sql &&\
    service postgresql stop

#COPY config/pg_hba.conf /etc/postgresql/$PGVER/main/pg_hba.conf
RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf &&\
    echo "shared_buffers=256MB" >> /etc/postgresql/$PGVER/main/postgresql.conf &&\
    echo "full_page_writes=off" >> /etc/postgresql/$PGVER/main/postgresql.conf &&\
    echo "unix_socket_directories = '/var/run/postgresql'" >> /etc/postgresql/$PGVER/main/postgresql.conf

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD ["/usr/lib/postgresql/10/bin/postgres", "-D", "/var/lib/postgresql/10/main", "-c", "config_file=/etc/postgresql/10/main/postgresql.conf"]