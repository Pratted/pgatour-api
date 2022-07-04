#!/bin/bash

echo "Generating models..."
oapi-codegen -generate types,skip-prune -package common ../openapi/models.yaml > ../pkg/common/models.go
echo "Done."