#!/bin/bash
set -v

ls /var/lib/webapps/webapp
ls -l /var/lib/webapps

/var/lib/webapps/webapp > /dev/null 2> /dev/null < /dev/null &
jobs

ls /var/lib/webapps/webapp
ls -l /var/lib/webapps

sleep 10
ls /var/lib/webapps/webapp
ls -l /var/lib/webapps
jobs
jobs -l