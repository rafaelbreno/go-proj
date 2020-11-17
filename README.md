# GoProject
- Random project for learning purposes
- Expects a lot of comments
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
- [ ] WebServer
    - [ ] Nginx
- [ ] Database
    - [ ] Choose one:
        - [ ] MySQL
        - [ ] PostgreSQL
        - [ ] MongoDB
    - [ ] Redis
    - [ ] Connection
- [ ] Build an API from Scratch
    - [ ] Database Connection
    - [ ] OAuth2
- [ ] Explode Functionalities into Microservices

### Development Pattern
- [ ] Implement some design pattern
    - [ ] SOLID
    - [ ] TDD
    - Choose one:
        - ~DDD~
        - ~BDD~
        - [x] MVC Laravel-ish
    - I didn't find a great and concrete sample/example of a Golang project(or any other language) in DDD
    - So I'm gonna stick with the _MVC-ish_ pattern
### DevOps
- [ ] Choose one CI/CD
    - [ ] Jenkins
    - [ ] Gitlab CI
    - [ ] Github Actions
- [ ] Container
    - [ ] Docker
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