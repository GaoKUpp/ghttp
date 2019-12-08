#! /bin/bash

BIN_DIR=$(go env | grep -i gobin | awk -v FS="\"" '{print $2}')

rm -rf "$BIN_DIR/ghttp" && go install
