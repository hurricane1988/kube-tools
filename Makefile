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

.PHONY: build-binary build-image clean-binary clean-image save-image go-vet

all: build-binary build-image clean-binary clean-image save-image go-vet

build-binary: ;$(info $(M)...Begin to build binary.)  @ ## build binary
	hack/build_binary.sh

build-image: ;$(info $(M)...Begin to build image.)  @ ## build image
	hack/build_image.sh

clean-binary: ;$(info $(M)...Begin to clean binary.)  @ ## clean binary
	hack/clean_binary.sh

clean-image: ;$(info $(M)...Begin to clean image.)  @ ## clean image
	hack/clean_image.sh

save-image: ;$(info $(M)...Begin to save image.)  @ ## save image
	hack/save_image.sh

go-vet: ;$(info $(M)...Begin to go vet.)  @ ## go vet
	hack/go_vet.sh

help:
	@echo "\033[1;33m------------------------------------------------------\033[0m"
	@echo "\033[1;33m          [Kube-Tools Makefile Commands]          \033[0m"
	@echo "\033[1;36mbuild-binary\033[0m           - \033[1;32mbuild the kube-tools binary\033[0m"
	@echo "\033[1;36mbuild-image\033[0m            - \033[1;32mbuild the kube-tools image\033[0m"
	@echo "\033[1;31mclean-binary\033[0m           - \033[1;91mclean the kube-tools binary\033[0m"
	@echo "\033[1;31mclean-image\033[0m            - \033[1;91mclean the kube-tools image\033[0m"
	@echo "\033[1;33msave-image\033[0m             - \033[1;32msave the kube-tools image\033[0m"
	@echo "\033[1;33mgo-vet\033[0m                 - \033[1;32mexecute go vet test.\033[0m"
	@echo "\033[1;33m------------------------------------------------------\033[0m"
