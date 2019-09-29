#!/bin/sh
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=. -I. $1