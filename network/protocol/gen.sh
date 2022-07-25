#!/usr/bin/env bash
echo protocol generate
echo current dir: "$PWD"

cur_path=$(pwd)
src_path=$cur_path
dst_path="$cur_path/gen/"

echo "SRC_DIR: $src_path"
echo "DST_DIR: $dst_path"

# go_path=${GOPATH}
# protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/address_book.proto
# protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto
mkdir gen

protoc --proto_path=./proto/ --go_out=./gen/ proto/*.proto

mkdir doc

protoc --validate_out="lang=go:./gen" --go_out=./gen/ --doc_out=./doc --doc_opt=html,index.html proto/*.proto


