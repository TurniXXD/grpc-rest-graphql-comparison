#!/bin/bash
export PATH="$PATH:$HOME/go/bin"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/proto/gallery.proto grpc/proto/user.proto
