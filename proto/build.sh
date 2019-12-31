#!/usr/bin/env bash
TARGET_FOLDER="api/v1"
SERVICE_PROTO="./service.proto"

rm -rf ../${TARGET_FOLDER} && mkdir -p ../${TARGET_FOLDER}

protoc -I=. -I=./third_party --go_out=plugins=grpc:../${TARGET_FOLDER} ./*.proto
echo "-> go stubs generated"

protoc -I=. -I=./third_party --grpc-gateway_out=logtostderr=true:../${TARGET_FOLDER} ./*.proto
echo "-> grpc-gateway stubs generated"

protoc -I=. -I=./third_party --swagger_out=logtostderr=true:../${TARGET_FOLDER} "$SERVICE_PROTO"
echo "-> swagger spec generated"
