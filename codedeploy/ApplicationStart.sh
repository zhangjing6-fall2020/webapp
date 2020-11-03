#!/bin/bash
set -v

whoami
echo ${HOSTNAME}
echo ${DBHOSTNAME}

#/opt/webapp > /dev/null 2> /dev/null < /dev/null &
sudo /var/lib/webapps/webapp