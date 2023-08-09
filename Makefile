#
# Copyright 2023 QKP Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: build-simulate-memory-binary

all: build-simulate-memory-binary

build-simulate-memory-binary: ;$(info $(M)...Begin to build simulate-memory binary.)  @ ## Generate https certificate
	hack/memory/build_binary.sh



help:
	@echo "-----------------------------------------------------------------------------------"
	@echo "make build-simulate-memory-binary                - build the simulate-memory binary"
	@echo "-----------------------------------------------------------------------------------"