# manny

[![Build Status][BuildStatusImg]][BuildMasterURL]
[![Code Coverage][CodecovImg]][CodecovURL]

Argo CD tool to generate K8s manifests from GitOps repo

## Installation

This process for adding additional custom tools to Argo CD is documented here:

* https://argoproj.github.io/argo-cd/operator-manual/custom_tools/

The ArgoCD repo-server deployment must be updated to include an init container that downloads and installed manny.

``` yaml
spec:
  template:
    spec:
      volumes:
      - name: custom-tools
        emptyDir: {}
      initContainers:
      - name: download-tools
        image: alpine:3.8
        command: [sh, -c]
        args:
        - wget -q -O manny.gz https://github.com/keikoproj/manny/manny-vlatest-linux-amd64.gz &&
          gunzip manny.gz &&
          chmod +x manny &&
          mv manny /custom-tools/manny
        volumeMounts:
        - mountPath: /custom-tools
          name: custom-tools
      containers:
      - name: argocd-repo-server
        volumeMounts:
        - mountPath: /usr/local/bin/manny
          name: custom-tools
          subPath: manny
```

The Argo CD configmap must be updated to install manny as a plugin.

``` yaml
data:
  configManagementPlugins: |
    - name: manny
      generate:
        command: [sh, -c]
        args: ["manny build ."]
```

Also, for each Argo CD app that intends to use manny, the Application must be updated to reference the manny plugin.

``` yaml
spec:
  source:
    plugin:
      name: manny
```