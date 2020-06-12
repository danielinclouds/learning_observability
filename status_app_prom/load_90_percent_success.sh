#!/bin/bash

HOST="$1"

while true;
do 
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
    http http://${HOST}/status/200 --print h
    sleep 0.5
done