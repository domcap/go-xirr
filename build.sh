#!/bin/bash
set -x

DEST_FOLDER=../Portal.Domcap/app/lib/finance/go/
MAIN_FILE=main/main.go
BUILD_DIR=build
# MAC OS Intel 64BIT:
ARCH=amd64
GOARCH=$ARCH 

case "$OSTYPE" in
  # solaris*) echo "SOLARIS" ;;
  darwin*)
    OS_TARGET=darwin
    LIB_NAME=xirr_lib.so
    ;;
  linux*)
    OS_TARGET=linux
    LIB_NAME=xirr_lib
    ;;
  # bsd*)     echo "BSD" ;;
  # msys*)    echo "WINDOWS" ;;
  # cygwin*)  echo "ALSO WINDOWS" ;;
  *)        
    echo "unknown: $OSTYPE"
    exit 1
    ;;
esac
GOOS=$OS_TARGET 

rm $BUILD_DIR/$ARCH/$OS_TARGET/$LIB_NAME
go build -o $BUILD_DIR/$ARCH/$OS_TARGET/$LIB_NAME -buildmode=c-shared

# cp -a $BUILD_DIR/. $DEST_FOLDER
