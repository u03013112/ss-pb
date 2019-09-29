#!/bin/sh
docker run --rm -v $(pwd):$(pwd) -v $(pwd):$(pwd)/github.com/u03013112/ss-pb -w $(pwd) znly/protoc --grpc-gateway_out=logtostderr=true:. -I. $1