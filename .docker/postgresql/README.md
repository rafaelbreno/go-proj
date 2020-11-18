### _.docker/postgresql/_
#### _Dockerfile_
- `FROM postgres:alpine`
    - [Hub](https://hub.docker.com/_/postgres)
- `COPY /.docker/postgresql/conf/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf`
    - This is our custom conf file
- `CMD ["postgres"]`
    - Running the database
- `EXPOSE 5432`
    - Exposing the port
#### _conf/pg_hba.conf_
- [pg_hba Doc](https://www.postgresql.org/docs/9.1/auth-pg-hba-conf.html)
- ``` 
      #       DB     user   address       auth-method
      host	  all	 all	0.0.0.0/0     md5
  ```
