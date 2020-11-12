#!/bin/bash
set -v

sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl \
    -a fetch-config \
    -m ec2 \
    -c file:/opt/cloudwatch-config.json \
    -s

sudo /var/lib/webapps/webapp > /dev/null 2> /dev/null < /dev/null &