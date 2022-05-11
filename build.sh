#!/bin/bash
set -x

LIB_NAME=xirr_lib
DEST_FOLDER=../Portal.Domcap/app/lib/finance/go/
MAIN_FILE=main/main.go
BUILD_DIR=build
# MAC OS Intel 64BIT:
ARCH=amd64

rm -fr $BUILD_DIR
rm -fr $DEST_FOLDER

for OS_TARGET in darwin linux
do
    if [ "$OS_TARGET" == "darwin" ]
    then
      LIB_NAME=$LIB_NAME.so
    else
      LIB_NAME=xirr_lib
    fi
    GOOS=$OS_TARGET 
    GOARCH=$ARCH 
    CGO_ENABLED=1
    go build -o $BUILD_DIR/$ARCH/$OS_TARGET/$LIB_NAME -buildmode=c-shared
done
cp -a $BUILD_DIR/. $DEST_FOLDER
