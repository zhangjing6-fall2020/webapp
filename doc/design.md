Webapp is a blog web application, which includes:
- create/update/delete users with user information(name, email, password, etc)
- authorized users are able to create/update/delete questions with images, and answers with images
- users will receive emails if the questions they asked are answered, updated or deleted
- use DynamoDB to store the message published on SNS and sent to users by SES to avoid sending duplicate emails
- auto scaling groups to scale out/in according to CPU utilization
- connect to rds using ssl
- load balancer uses ssl to confirm security
- use https to visit domain(dev/prod.bh7cw.me)
- creates log/metric in cloudwatch

The webapp can:
- trigger ci to build and test the application
- deploy webapp automatically on aws using codedeploy
- do load test using JMeter

Endpoints:
- https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-02#/
- https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-03
- https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-05

Related repo:
- each assignment](https://github.com/bh7cw/demo?organization=bh7cw&organization=bh7cw)

secrets setting:
- `cicd` user in prod or dev account

webapp design logic:
- question with any answers cannot be deleted
- the categories won't be deleted if the question is deleted
- delete question will delete all the files attached to it
- delete answer will delete all the files attached to it