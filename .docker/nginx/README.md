### _.docker/nginx/_
#### _Dockerfile_
- [Hub](https://hub.docker.com/_/nginx)
- `FROM nginx:alpine`
    - Nginx Image Alpine version
- `COPY /.docker/nginx/conf/nginx.conf /etc/nginx/nginx.conf`
    - Copying _nginx.conf_ into nginx folder inside of my service/container
- `EXPOSE 80 443`
    - [_EXPOSE_ Docs](https://docs.docker.com/engine/reference/builder/#expose)
    - Exposing port 80 into 443
    - or 443 into 80??? ~I don't know~
- `ENTRYPOINT ["nginx"]`
    - [_ENTRYPOINT_ Docs](https://docs.docker.com/engine/reference/builder/#entrypoint)
    - This allows we configure a container as an _executable_
- `CMD ["-g", "daemon off;"]`
    - From [Nginx Hub](https://hub.docker.com/_/nginx)
        - "If you add a custom CMD in the Dockerfile, be sure to include -g daemon off; in the CMD in order for nginx to stay in the foreground, so that Docker can track the process properly (otherwise your container will stop immediately after starting)!"

#### _conf/nginx.conf_
- [NGINX Doc](https://nginx.org/en/docs/)
##### `worker_processes auto;`
    - [Doc](http://nginx.org/en/docs/ngx_core_module.html#worker_processes)
    - This define the number of worker processes ~obviously~
    - The number of workers to be optimal depends on(but not limited to):
        - Number of CPU cores
        - Number of data HD
        - Load Pattern
##### `events`
- [Doc](https://nginx.org/en/docs/ngx_core_module.html#events)
- _"Provides the configuration file context in which the directives that affect connection processing are specified."_
- `worker_connections 1024;`    
    - [Doc](http://nginx.org/en/docs/ngx_core_module.html#worker_connections)
    - Set the maximum number of simultaneous connections that can be opened by a worker process
##### `http`
- [Doc](https://nginx.org/en/docs/http/ngx_http_core_module.html#http)
- _"Provides the configuration file context in which the HTTP server directives are specified."_
- `upstream app`
    - [Doc](http://nginx.org/en/docs/http/ngx_http_upstream_module.html)
    - _"Is used to define groups of servers that can be referenced"_
    - `least_conn`
        - [Doc](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#least_conn)
        - _"Specifies that a group should use a load balancing method where a request is passed to the server with the least number of active connections"_
    - `server`
        - [Doc](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#server)
        - _"Defines the address and other parameters of a server. The address can be specified as a domain name or IP address, with an optional port"_
- `server`
    - [Doc](https://nginx.org/en/docs/http/configuring_https_servers.html)
    - `listen 80;`
        - [Doc](https://nginx.org/en/docs/http/ngx_http_core_module.html#listen)
        - Set an address and port for IP 
    - `server_name 0.0.0.0;`
        - [Doc](https://nginx.org/en/docs/http/ngx_http_core_module.html#server_name)
        - _"Sets names of a virtual server"_
    - `location /`
        - [Doc](https://nginx.org/en/docs/http/ngx_http_core_module.html#location)
        - This context will allow us to write different configurations for different files, folder, and regex patterns
        - [proxy_pass](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass)
        - [proxy_http_version](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_http_version)
        - [proxy_set_header](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_set_header)
        - [proxy_cache_bypass](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache_bypass)
