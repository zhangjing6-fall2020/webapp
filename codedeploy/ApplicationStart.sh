#!/bin/bash
set -v

/var/lib/webapps/webapp > /dev/null 2> /dev/null < /dev/null &
sleep 60
jobs