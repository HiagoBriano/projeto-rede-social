FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=Devbook

COPY ./sql/sql.sql /docker-entrypoint-initdb.d/

EXPOSE 3306

# docker build -t devbook_mysql .
# docker run -d -p 3306:3306 devbook_mysql