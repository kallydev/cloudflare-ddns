#!/usr/bin/env bash

project="cloudflare-ddns"
release=$1

platforms=(
  "linux_amd64"
)

for(( i=0;i<${#platforms[@]};i++)) do
  array=($(echo "${platforms[i]}" | tr '_' ' '))
  name=${project}_${platforms[i]}
  if [[ $release == "--release" ]]; then
    mkdir -p release
    GOOS=${array[0]} GOARCH=${array[1]} /snap/go/current/bin/go build -ldflags "-s -w" -o release/"${name}" cmd/main.go
    strip -s release/"$name"
    upx -9 release/"$name"
    7z a release/"$name".7z release/"$name" README.md LICENSE config.json public.pem private.pem
    sha1sum release/"$name".7z > release/"$name".7z.sha1
  else
    mkdir -p build
    GOOS=${array[0]} GOARCH=${array[1]} /snap/go/current/bin/go build -o build/"${name}" cmd/main.go
  fi
done;
