# NGINX Image
FROM nginx:alpine

# Copying the conf file
COPY /.docker/nginx/conf/nginx.conf /etc/nginx/nginx.conf

# Exposing a port, where this container will
# be accessible to
EXPOSE 80 443

# Defining an entrypoint
ENTRYPOINT ["nginx"]

# This be executed as:
# $ nginx -g
# $ nginx daemon off
CMD ["-g", "daemon off;"]
