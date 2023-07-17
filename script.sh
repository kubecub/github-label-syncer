#!/bin/bash
# Copyright © 2023 KubeCub open source community. All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.


# 执行 make build
make build

# 定义函数执行命令
execute_command() {
    echo "Executing command: $@"
    "$@"
}

# 定义要执行的命令列表
commands=(
    "./_output/bin/platforms/linux/amd64/exporter kubernetes kubernetes --yaml -f ./labels-templates/kubernetes-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes kubernetes --table -f ./labels-templates/kubernetes-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes kubernetes --json -f ./labels-templates/kubernetes-json.json"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes kubernetes --xml -f ./labels-templates/kubernetes-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter OpenIMSDK Open-IM-Server --yaml -f ./labels-templates/openim-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter OpenIMSDK Open-IM-Server --table -f ./labels-templates/openim-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter OpenIMSDK Open-IM-Server --json -f ./labels-templates/openim-json.json"
    "./_output/bin/platforms/linux/amd64/exporter OpenIMSDK Open-IM-Server --xml -f ./labels-templates/openim-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter labring sealos --yaml -f ./labels-templates/sealos-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter labring sealos --table -f ./labels-templates/sealos-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter labring sealos --json -f ./labels-templates/sealos-json.json"
    "./_output/bin/platforms/linux/amd64/exporter labring sealos --xml -f ./labels-templates/sealos-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter microsoft vscode --yaml -f ./labels-templates/vscode-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter microsoft vscode --table -f ./labels-templates/vscode-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter microsoft vscode --json -f ./labels-templates/vscode-json.json"
    "./_output/bin/platforms/linux/amd64/exporter microsoft vscode --xml -f ./labels-templates/vscode-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes-sigs kustomize --yaml -f ./labels-templates/kustomize-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes-sigs kustomize --table -f ./labels-templates/kustomize-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes-sigs kustomize --json -f ./labels-templates/kustomize-json.json"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes-sigs kustomize --xml -f ./labels-templates/kustomize-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter prometheus prometheus --yaml -f ./labels-templates/prometheus-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter prometheus prometheus --table -f ./labels-templates/prometheus-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter prometheus prometheus --json -f ./labels-templates/prometheus-json.json"
    "./_output/bin/platforms/linux/amd64/exporter prometheus prometheus --xml -f ./labels-templates/prometheus-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes test-infra --yaml -f ./labels-templates/test-infra-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes test-infra --table -f ./labels-templates/test-infra-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes test-infra --json -f ./labels-templates/test-infra-json.json"
    "./_output/bin/platforms/linux/amd64/exporter kubernetes test-infra --xml -f ./labels-templates/test-infra-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter redis redis --yaml -f ./labels-templates/redis-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter redis redis --table -f ./labels-templates/redis-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter redis redis --json -f ./labels-templates/redis-json.json"
    "./_output/bin/platforms/linux/amd64/exporter redis redis --xml -f ./labels-templates/redis-xml.xml"
    "./_output/bin/platforms/linux/amd64/exporter kubecub go-project-layout --yaml -f ./labels-templates/go-project-layout-yaml.yaml"
    "./_output/bin/platforms/linux/amd64/exporter kubecub go-project-layout --table -f ./labels-templates/go-project-layout-table.txt"
    "./_output/bin/platforms/linux/amd64/exporter kubecub go-project-layout --json -f ./labels-templates/go-project-layout-json.json"
    "./_output/bin/platforms/linux/amd64/exporter kubecub go-project-layout --xml -f ./labels-templates/go-project-layout-xml.xml"
)

# 依次执行命令列表中的命令
for command in "${commands[@]}"; do
    execute_command $command
done
