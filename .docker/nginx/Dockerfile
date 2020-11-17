# NGINX Image
FROM nginx:alpine

# Copying the conf file
COPY /.docker/nginx/conf/nginx.conf /etc/nginx/nginx.conf

# Exposing a port, where this container will
# be accessible to
EXPOSE 80 443

# From https://hub.docker.com/_/nginx
# "If you add a custom CMD in the Dockerfile, 
# be sure to include -g daemon off; 
# in the CMD in order for nginx to stay in the foreground, 
# so that Docker can track the process properly 
# (otherwise your container will stop immediately after starting)!"

# Defining an entrypoint
ENTRYPOINT ["nginx"]

# This be executed as:
# $ nginx -g
# $ nginx daemon off
CMD ["-g", "daemon off"]
