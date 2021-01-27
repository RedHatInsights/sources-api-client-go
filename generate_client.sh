#!/usr/bin/env bash

### Run like so:
# FILE=openapi.json ./generate_client.sh
# the openapi schema needs to be in this directory for the container to see it.
# it would have to be something like /workdir/output

docker run -it --rm \
    -v /tmp:/tmp \
    -v $PWD:/workdir \
    openapitools/openapi-generator-cli:v5.0.0 \
    generate -i /workdir/$FILE -g go -o /workdir/ -c /workdir/config.json

echo -e "=======
Client generated from $FILE
======="
