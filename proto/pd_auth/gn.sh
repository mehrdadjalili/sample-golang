#!/bin/bash
protoc --go_out=. --go-grpc_out=. pd_auth_client.proto
protoc --go_out=. --go-grpc_out=. pd_auth_manager.proto
