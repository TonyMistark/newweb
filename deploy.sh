#!/bin/bash

kill -9 $(pgrep webserver)
cd /home/omar/projects/newweb/
git pull git@github.com:TonyMistark/newweb.git
cd webserver
./webserver &
