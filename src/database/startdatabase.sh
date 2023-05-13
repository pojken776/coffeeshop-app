#!/bin/bash
set -x

password=$1

docker run --name postgres-0 -e POSTGRES_PASSWORD=$password -d -p 5432:5432 -v ./data:/var/lib/postgresql/data postgres:alpine
