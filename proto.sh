#!/bin/bash

protoc -Iproto --go_out=. --go-grpc_out=. proto/user.proto
