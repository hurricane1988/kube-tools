#!/usr/bin/env bash

# Copyright 2023 QKP Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e

TAG=${TAG:-latest}
OUTPUT=${OUTPUT:-_output}
TYPE=${TYPE:-tar}
IMAGE=${IMAGE:-kube-tools}

# support other container tools. e.g. podman
CONTAINER_CLI=${CONTAINER_CLI:-docker}
CONTAINER_BUILDER=${CONTAINER_BUILDER:-build}
CONTAINER_SAVE=${CONTAINER_SAVE:-save}

mkdir -p "${OUTPUT}"
# shellcheck disable=SC2086
${CONTAINER_CLI} "${CONTAINER_SAVE}" ${IMAGE}:${TAG} --output "${OUTPUT}"/"${IMAGE}".${TYPE}