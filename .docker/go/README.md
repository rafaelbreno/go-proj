### _.docker/go/_
#### _Dockerfile_
- [Docker Hub](https://hub.docker.com/_/golang)
- `FROM golang:alpine`
    - Selecting image
    - _{image}:{version}_
- `WORKDIR /app`
    - This will define in which directoy the RUN/CMD will run
_ `COPY go.mod go.sum ./`
    - Copying files to install my dependencies
- `RUN go mod download`
    - _RUN_ will execute a command, in this case:
    - > $ go mod download
- `RUN go build -o main`
    - This command will compile my app into a _binary_
    - _go build_ -> build the app
    - _-o_ -> force to overwrite any already existent binary
    - _main_ -> the package that will be built
- `EXPOSE 8080`
    - Exposing this image to a port
    - In this case will be the same as I've configured in my _nginx.conf_ file
- `CMD ["./main"]`
    - Run some shell commands
    - in this case:
    - > $ ./main
