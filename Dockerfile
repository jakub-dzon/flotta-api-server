# Copyright 2017 The Kubernetes Authors.
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

FROM golang:1.16 as build
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor  -a -o flotta-apiserver main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=build /workspace/flotta-apiserver /
USER 65532:65532

ENTRYPOINT ["/flotta-apiserver"]
