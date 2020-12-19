# webapp
webapp running at port: 8080

Instructions:
1. Prerequisites for building and deploying your application locally.
    - Install [Golang](https://golang.org/dl/)
    - Place the codebase in `GOPATH/src/`
    - Install the dependencies listed in go.mod(optional)
2. Build and Deploy instructions for web application.
    - Build: `go build`
    - Build for ubuntu: `env GOOS=linux GOARCH=amd64 go build`
    - Test: `go test ./...`
    - Run: `go run main.go`
3. JMeter do load test.
   - install [JMeter](https://jmeter.apache.org/)
   - run JMeter: `jmeter`
   - open the jmx file
   - run the load test

Stack used:
- golang
- Gorm
- Gin

Components on AWS used:
- EC2 Instances, Security Groups, AMI, Auto Scaling, Load balaner
- Rds
- DynamoDB
- S3
- CloudWatch
- VPC
- CodeDeploy
- Route53
- Lambda
- SNS
- SES
- Certificate Manager

Tools:
- Github actions
- JMeter

See more details:
- Design:  [doc/design.md](https://github.com/zhangjing6-fall2020/webapp/blob/master/doc/design.md)
- Version: [doc/version.md](https://github.com/zhangjing6-fall2020/webapp/blob/master/doc/version.md)
- Command: [doc/command.md](https://github.com/zhangjing6-fall2020/webapp/blob/master/doc/command.md)