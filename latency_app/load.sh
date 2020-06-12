#!/bin/bash

HOST="$1"

while true;
do 
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/1 --print h
    sleep 0.5
    http http://${HOST}/latency/400 --print h
    sleep 0.5
done