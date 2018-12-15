#!/bin/bash

kill -g $(pgrep webserver)
cd ~/projects/newweb/
git pull git@github.com:TonyMistark/newweb.git
cd webserver
./webserver &
