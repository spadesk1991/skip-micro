#!/bin/bash
cd ../app/services/proto
protoc --micro_out=../pb/ --go_out=../pb/ users.proto
protoc-go-inject-tag -input=../pb/users.pb.go