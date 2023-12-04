#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <nama module>"
    exit 1
fi

module_name=$1

mkdir -p "./module/$module_name/dto"
mkdir -p "./module/$module_name/handler"
mkdir -p "./module/$module_name/repository"
mkdir -p "./module/$module_name/service"

# Membuat file-filenya
touch "./module/$module_name/dto/req.go"
touch "./module/$module_name/dto/res.go"
touch "./module/$module_name/handler/index.go"
touch "./module/$module_name/index.go"
touch "./module/$module_name/repository/index.go"
touch "./module/$module_name/service/index.go"

echo "Module '$module_name' berhasil dibuat!"
