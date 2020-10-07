# webapp
Instructions:
1. Prerequisites for building and deploying your application locally.
    - Install [Golang](https://golang.org/dl/)
    - Place the codebase in `GOPATH/src/`
    - Install the dependencies listed in go.mod(optional)
2. Build and Deploy instructions for web application.
    - Build: `go build`
    - Test: `go test ./...`
    - Run: `go run main.go`

api spec:
- [hw2](https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-02)
- [hw3](https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-03)

Changes from a2 to a3:
- Added more APIs
- Table `user` changes to `users`
- Endpoint get: `/v1/user/self` changes to `v1/userself`
- Endpoint get: `v1/user` changes to `v1/users`

a3 demo procedure:
- add two users, a and b
- add a question for a user
- add an answer for a user
- user b can't edit the answer
- user b can't delete the question
- user a can't delete the question due to answer exists
- user a delete the answer
- user a delete the question
