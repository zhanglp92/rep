#!/bin/bash

set -e

PROTOC=protoc
PB_DIR=github.com/zhanglp92/rep/api/pb

cd ${GOPATH}/src

${PROTOC} --go_out=. ./github.com/zhanglp92/rep/api/pb/user/*.proto
${PROTOC} --go_out=. ./github.com/zhanglp92/rep/api/pb/item/*.proto

echo "Done."
