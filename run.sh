#!/bin/bash

docker build -t receipt-processor-challenge .
docker run -p 8080:8080 receipt-processor-challenge