# GoProject
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
- [x] __FOCUS ON GOROUTINES__
- [x] WebServer
    - [x] Nginx
- [ ] Database
    - [x] Choose one:
        - ~MySQL~
        - [x] PostgreSQL
        - ~MongoDB~
    - [x] Redis
    - [x] Connection
        - [x] Postgres
        - [x] Redis
    - [x] Migrations
    - [ ] Seeding
- [ ] Build an API
    - [x] Database Connection
    - [x] Simple TODO
        - [x] TODO List
        - [x] TODO Items
    - [ ] User CRUD
        - [x] Sign Up (_registration_)
        - [x] Sign In (_login_)
        - [x] Sign off
        - [ ] Update
        - [ ] Delete
    - [x] JWT
        - [x] Creating Token
        - [x] Deleting Token
        - [x] Refreshing Token
- [ ] Explode Functionalities into Microservices

### Development Pattern
- [x] Implement some design pattern
    - [x] SOLID
    - [x] TDD
    - Choose one:
        - ~DDD~
        - ~BDD~
        - [x] MVC Laravel-ish
    - I didn't find a great and concrete sample/example of a Golang project(or any other language) in DDD
    - So I'm gonna stick with the _MVC-ish_ pattern
### DevOps
- [x] Choose one CI/CD
    - [x] Jenkins
    - ~Gitlab CI~
    - ~Github Actions~
    - ~CircleCI~
- [x] Container
    - [x] Docker
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

## Project Folder Structure
### _.docker/_
- Here will be all Dockerfiles
- _nginx/_
- _postgresql/_
- _golang/_
### _app/_
- _helpers/_
- _http/_
    - _controllers/_
    - _middlewares/_ 
- _models/_
- _repositories/_
### _cmd/_
### _config/_
### _database/_
- _factories/_
- _migrations/_
- _seeders/_
### _routes/_
### _tests/_
