#!/bin/bash

protoc --micro_out=../pb/ --go_out=../pb/ *.proto

protoc-go-inject-tag -input=../pb/users.pb.go