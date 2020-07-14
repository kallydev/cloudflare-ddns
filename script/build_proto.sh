#!/bin/bash

function listFiles() {
  for file in "$1"/*; do
    if test -d "$file"; then
      cd "$file" || continue
      listFiles .
    elif test -e "$file" && [[ ${file##*.} == "proto" ]]; then
      build "$file"
    fi
  done
  cd ..
}

function build() {
  echo "$(pwd)$1" building...
  protoc -I . "$1" --go_out=plugins=grpc:.
}

listFiles .
