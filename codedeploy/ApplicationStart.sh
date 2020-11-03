#!/bin/bash


# https://docs.aws.amazon.com/codedeploy/latest/userguide/troubleshooting-deployments.html#troubleshooting-long-running-processes
/var/lib/webapps/webapp > /dev/null 2> /dev/null < /dev/null &
jobs