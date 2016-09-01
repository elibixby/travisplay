#!/bin/bash

set -e -x

find . -name '*.yaml'

which go || echo 'No go!'

curl https://sdk.cloud.google.com | bash
source /root/.bashrc

gsutil ls gs://cbro-scratch
