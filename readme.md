## User service
This service will handle data login, avatar, followers & following and this service has own database.
There are 2 service on this project, 1 is this service (user-service) and the other 1 is in different repo called (comment-service).
User service will connect with Comment service through inter-communication using GRPC

### Pre-requisite
- Minikube
- Docker
- Kubectl

### Technologies
- Golang
- GRPC
- Gorm
- Postgre
- Docker
- Kubernetes

### How to deploy
``$ sh deploy.sh``
- In this deployment command, will execute build docker image, cleanup existing deployment & pod service, and deploy it to minikube

```$ kubectl port-forward service/user-svc 8081:8081```
- After pod is deployed on minikube, need to forward the port using above command because by default service port is not accesible by public. Here I am using clusterIp, to make it similar like in production microservices. After that service will serve and can access using ``localhost:8081``

### Api Doc
- HOST: `localhost:8081`
- Get all comments
    - description: in this endpoint, user-service will get related data by organization to comment-service using GRPC
    - method: `GET`
    - organization_name: `xendit`
    - endpoint: `{HOST}/orgs/{organization_name}/members`
    - example response:
    `
        {
            "status": 200,
            "data": [
                {
                    "login": "robert",
                    "avatar": "https://google.com",
                    "followers": 1500,
                    "following": 500
                },
                {
                    "login": "maulana",
                    "avatar": "https://google.com1",
                    "followers": 324,
                    "following": 44
                }
            ],
            "message": "success"
        }
    `
### DB design
https://dbdiagram.io/d/5ef929340425da461f03fa2a

### How to run testing
- `$ cd x-user-service/test/users`
- `$ go users_test.go`