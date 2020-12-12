# GoProject
## Releases
- [Alpha](https://github.com/rafaelbreno/go-proj/tree/release/alpha-1.0)
    - First version
    - Simple Todo List
- [Master](#)
    - Banking System
## About
- Random project for learning purposes
- Expects a lot of comments
## Running this App
- > $ git clone https://github.com/rafaelbreno/go-proj.git
- > $ cd go-proj
- > $ sudo ./start.sh
- Update the generated .env
- > $ sudo ./start.sh
-   ```
        Choose an option:
        1- Run Tests
        2- Run Prod
        3- Run Tests and Prod
        4- Purge All
    ```
- > $ 3
- Hit enter and Enjoy 
## Packages
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/go-gorm/gorm)
- [Godotenv](github.com/joho/godotenv)
- [Zap](https://github.com/uber-go/zap)
## TODO:

### Study: 
- [ ] Understanding of Design Patterns
- [ ] SSO systems
- [ ] Microservice architecture
- [ ] Communication protocols
- [ ] Virtualisation
- [ ] Continuous integration (CI/CD)
- [ ] Unit testing

### Application
#### Structure
- [x] __FOCUS ON GOROUTINES__
- [x] WebServer
    - [x] Nginx
- [ ] Database
    - [x] Choose one:
        - ~MySQL~
        - [x] PostgreSQL
        - ~MongoDB~
    - [ ] Redis
    - [ ] Connection
        - [ ] Postgres
        - [ ] Redis
    - [ ] Migrations
    - [ ] Seeding

#### API
- [ ] Database Connection
- [ ] Banking App
    - [ ] TODO List
    - [ ] TODO Items
- [ ] User CRUD
    - [ ] Sign Up (_registration_)
    - [ ] Sign In (_login_)
    - [ ] Sign off
    - [ ] Update
    - [ ] Delete
- [ ] JWT
    - [ ] Creating Token
    - [ ] Deleting Token
    - [ ] Refreshing Token
- [ ] Explode Functionalities into Microservices

#### Development Pattern
- [x] Implement some design pattern
    - [x] SOLID
    - [x] TDD
    - [x] [Hexagonal Architecture](https://www.qwan.eu/2020/08/20/hexagonal-architecture.html)
### DevOps
#### CI/CD
- [ ] Choose one
    - ~Jenkins~
    - ~Gitlab CI~
    - [x] Github Actions
    - ~CircleCI~
- [_/.github_](https://github.com/rafaelbreno/go-proj/tree/master/.github)
#### Container
- [ ] Choose One
    - [x] Docker
- [_/.docker_](https://github.com/rafaelbreno/go-proj/tree/master/.docker)

- [ ] Configuration Management
    - [ ] Ansible

- [ ] Container Orchestration
    - [ ] Kubernetes

- [ ] Infraestructure Provisioning
    - [ ] Terraform

- [ ] Service Mesh
    - [ ] Istio 

### Monitoring Tools
- [ ] Logs
    - [ ] Elastic Stack
- [ ] Application Monitoring
    - [ ] Jaeger
    - [ ] New Relic
- [ ] Implement monitoring system
    - [ ] Prometheus
    - [ ] Grafana
- [ ] Cloud Providers
    - [ ] AWS
    - [ ] Digital Ocean
    - [ ] Linode

-----

## Project Folder Structure
### _.docker/_
- Here will be all Dockerfiles
- _go/_
    - Go image
- _nginx/_
    - Server image
- _postgresql_
    - Database image
- _redis/_
    - Key/Value Database image

-----

### _app/_
- Here will be all Dockerfiles

-----

### _cmd/_
- Functionalities globally accessible
- E.G.
    - Database Connection
    - Helpers
    - Errors Handlers
    - Logging
    - etc.

-----

### _domain/_
- Stores all App's Domain
- I came from a MVC background, so I like to think Domain as a Model

#### _domain.go_
- Retrieve DB connections
- Redis, PostgreSQL
- Run Migrations

#### _user.go_
- Here is the User Domain("Model")
- Here'll be the User main struct and Repository

#### _userRepositoryDB.go_
- Almost like and extension for the User's Domain
- Here'll be all methods/actions that this Domain will need to have

-----

### _routes/_
- Here will be the routing file

-----

### _service/_
- Here will be all Services from the App
- Connecting Handler and Domain
- Handler <-> Service <-> Domain

-----
-----

Liked what I'm doing? And want to support?

<a href="https://www.buymeacoffee.com/rafiusky" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

