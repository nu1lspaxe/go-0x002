#!/bin/sh

# run in wsl
docker run -v `pwd`:/defs namely/gen-grpc-gateway:latest -f /defs/unary/schema.proto -s Service

# run in powershell
# docker run -v ${pwd}:/defs namely/gen-grpc-gateway:latest -f /defs/unary/schema.proto -s Service