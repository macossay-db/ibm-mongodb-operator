//
// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package controllers

const installConfigMap = `
---
# Source: icp-mongodb/templates/mongodb-install-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  labels:
    app.kubernetes.io/name: icp-mongodb
    app.kubernetes.io/instance: icp-mongodb
    app.kubernetes.io/version: 4.0.12-build.3
    app.kubernetes.io/component: database
    app.kubernetes.io/part-of: common-services-cloud-pak
    app.kubernetes.io/managed-by: operator
    release: mongodb
  name: icp-mongodb-install
data:
  install.sh: |
    #!/bin/bash

    # Copyright 2016 The Kubernetes Authors. All rights reserved.
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

    # This volume is assumed to exist and is shared with the peer-finder
    # init container. It contains on-start/change configuration scripts.
    WORKDIR_VOLUME="/work-dir"
    CONFIGDIR_VOLUME="/data/configdb"

    for i in "$@"
    do
    case $i in
        -c=*|--config-dir=*)
        CONFIGDIR_VOLUME="${i#*=}"
        shift
        ;;
        -w=*|--work-dir=*)
        WORKDIR_VOLUME="${i#*=}"
        shift
        ;;
        *)
        # unknown option
        ;;
    esac
    done

    echo installing config scripts into "${WORKDIR_VOLUME}"
    mkdir -p "${WORKDIR_VOLUME}"
    cp /peer-finder "${WORKDIR_VOLUME}"/
    echo "I am running as " $(whoami)

    cp /configdb-readonly/mongod.conf "${CONFIGDIR_VOLUME}"/mongod.conf
    cp /keydir-readonly/key.txt "${CONFIGDIR_VOLUME}"/
    cp /ca-readonly/tls.key "${CONFIGDIR_VOLUME}"/tls.key
    cp /ca-readonly/tls.crt "${CONFIGDIR_VOLUME}"/tls.crt

    chmod 600 "${CONFIGDIR_VOLUME}"/key.txt
    # chown -R 999:999 /work-dir
    # chown -R 999:999 /data

    # Root file system is readonly but still need write and execute access to tmp
    # chmod -R 777 /tmp`
