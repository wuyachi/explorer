#!/bin/bash

base=build

if [ ! -d "./$base" ]; then
  mkdir -p "./$base"
fi

if [ ! -d "./$base/palette_server" ]; then
  mkdir -p "./$base/palette_server"
fi

if [ ! -d "./$base/palette_http" ]; then
  mkdir -p "./$base/palette_http"
fi

if [ ! -d "./$base/palette_http/conf" ]; then
  mkdir -p "./$base/palette_http/conf"
fi

cd ./cmd
go build -o palette_server main.go
mv palette_server ./../$base/palette_server
cp ./../conf/config.json ./../$base/palette_server

cd ./..
go build -o palette_http main.go
mv palette_http ./$base/palette_http
cp ./conf/app.conf ./$base/palette_http/conf
