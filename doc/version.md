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

Changes from a3 to a4:
- No development, just change category to unique

Changes from a4 to a5:
- add file, answer file, question file
- post file: upload the file to S3 and database
- delete file: delete the file from S3 and database

Changes from a5 to a6:
- use codedeploy to deploy webapp on ec2 instance

Changes from a6 to a7:
- add logs using logrus and metrics using statsD
- send to watch cloud show the logs and metrics

Changes from a7 to a8:
- add auto scaling instead of ec2
- add load-balancer
- add jmeter to do load test

Changes from a8 to a9:
- add create/update/delete answer will publish message in sns on aws

Changes from a9 to a10:
- add ssl secure database connection 