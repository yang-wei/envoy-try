#!/bin/sh
go run /code/service.go &
envoy -c /etc/service-envoy.yaml --service-cluster service${SERVICE_NAME}

