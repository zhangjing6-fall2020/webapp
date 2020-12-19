api spec:
- [hw2](https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-02)
- [hw3](https://app.swaggerhub.com/apis-docs/csye6225/fall2020-csye6225/assignment-03)

Useful link:
- https://toolbox.googleapps.com/apps/dig/#A/
- https://www.whatsmydns.net/#NS/bh7cw.me
- https://docs.aws.amazon.com/codedeploy/latest/userguide/troubleshooting-deployments.html#troubleshooting-long-running-processes

Debug codedeploy:
- log: /opt/codedeploy-agent/deployment-root/deployment-logs/codedeploy-agent-deployments.log

Check port:
- lsof -i:8080
- sudo netstat -pna | grep 8080

Test statsD:
- `netcat -ulzp 8125`
- `echo "foo:1|c" | nc -u 127.0.0.1 8125`
- `echo -n 'some.metric.namespace:1|c' | nc -u -q0 localhost 8125`
- `while true; do curl http://localhost:8080/v1/users;sleep 1;done;`

port:
//https://blog.csdn.net/ws379374000/article/details/74218530
- `sudo netstat -ntulp |grep 8125`

Debug cloud watch:
//https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/install-CloudWatch-Agent-on-EC2-Instance-fleet.html#start-CloudWatch-Agent-EC2-fleet
- stop: `sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -m ec2 -a stop`

Debug code deploy:
//https://docs.aws.amazon.com/codedeploy/latest/userguide/deployments-view-logs.html
- `less /var/log/aws/codedeploy-agent/codedeploy-agent.log`

Related ssl links:
- [connect to rds commands](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToInstance.html)
- [check ssl connection](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ssl-certificate-rotation-mysql.html#ssl-certificate-rotation-mysql.determining-server)
- [check ssl param in mysql](https://dev.mysql.com/doc/refman/8.0/en/performance-schema-quick-start.html)
- [gorm](https://gorm.io/docs/sql_builder.html)

About demo ssl:
- open ssh 22 port
- ssh ec2
- `sudo apt install mysql-client`
- `mysql --host=csye6225-f20.cvefazkbfyp3.us-east-1.rds.amazonaws.com --user=csye6225fall2020 --password=Znt9yjNTp5NR`
- `SELECT id, user, host, connection_type  FROM performance_schema.threads pst  INNER JOIN information_schema.processlist isp  ON pst.processlist_id = isp.id;`
shows:
+----+------------------+------------------+-----------------+
| id | user             | host             | connection_type |
+----+------------------+------------------+-----------------+
|  8 | csye6225fall2020 | 10.0.2.150:58430 | SSL/TLS         |
|  9 | csye6225fall2020 | 10.0.1.36:47372  | SSL/TLS         |
| 10 | csye6225fall2020 | 10.0.3.163:42500 | SSL/TLS         |
| 11 | csye6225fall2020 | 10.0.2.150:58440 | SSL/TLS         |
|  5 | event_scheduler  | localhost        | NULL            |
|  7 | rdsadmin         | localhost:24828  | TCP/IP          |
+----+------------------+------------------+-----------------+