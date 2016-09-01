#!/bin/bash

set -e -x

find . -name '*.yaml'

which go || echo 'No go!'

