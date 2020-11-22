# https://hub.docker.com/_/redis
FROM redis:6-alpine

# Copying our conf file
COPY /.docker/redis/conf/redis.conf /usr/local/etc/redis/redis.conf

# Specifying the conf file dir when starting running Redis
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

#CMD ["redis-server"]
# Exposing Redis port
EXPOSE 6379 
