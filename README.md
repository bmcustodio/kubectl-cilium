# kubectl-cilium

A `kubectl` plugin for interacting with [Cilium](https://cilium.io).

![Release](https://img.shields.io/github/v/release/bmcstdio/kubectl-cilium)
![Downloads](https://img.shields.io/github/downloads/bmcstdio/kubectl-cilium/total?color=green)
[![Build](https://img.shields.io/travis/com/bmcstdio/kubectl-cilium)](https://travis-ci.com/bmcstdio/kubectl-cilium)
![License](https://img.shields.io/github/license/bmcstdio/kubectl-cilium)

## Installation

At the moment, `kubectl-cilium` must be installed by running

```shell
$ go get github.com/bmcstdio/kubectl-cilium/cmd/kubectl-cilium
```

or by cloning this repository, running

```shell
$ make build
```

and copying `./bin/kubectl-cilium` to a directory in your `$PATH`.

## Examples

### `kubectl-cilium exec`

The `exec` command allows you to execute a command on a Cilium agent targeted by either a node name or a `namespace/name` key targeting any pod in the cluster.
For example, assuming that your Kubernetes cluster has nodes

```shell
$ kubectl get node                                                                                                                                            
NAME                               STATUS   ROLES    AGE   VERSION                                                                                      
kind-cilium-mesh-2-control-plane   Ready    master   87m   v1.18.0                                                                           
kind-cilium-mesh-2-worker          Ready    <none>   86m   v1.18.0                                                    
kind-cilium-mesh-2-worker2         Ready    <none>   86m   v1.18.0
```

and pods

```shell
$ kubectl get pod --all-namespaces
NAMESPACE            NAME                                                       READY   STATUS        RESTARTS   AGE   IP            NODE
kube-system          cilium-8dzvt                                               1/1     Running       0          86m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          cilium-8lmlb                                               1/1     Running       0          86m   172.17.0.5    kind-cilium-mesh-2-worker       
kube-system          cilium-etcd-2pphwr5jhg                                     1/1     Running       0          88m   10.20.2.12    kind-cilium-mesh-2-worker       
kube-system          cilium-etcd-operator-59b4987745-n94pb                      1/1     Running       0          90m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          cilium-jcs44                                               1/1     Running       0          86m   172.17.0.6    kind-cilium-mesh-2-worker2      
kube-system          cilium-operator-77dd4b8544-vb2kf                           1/1     Running       0          86m   172.17.0.5    kind-cilium-mesh-2-worker       
kube-system          coredns-5644d7b6d9-fhmff                                   1/1     Running       0          81m   10.20.2.62    kind-cilium-mesh-2-worker       
kube-system          coredns-5644d7b6d9-x9ksr                                   1/1     Running       0          81m   10.20.0.68    kind-cilium-mesh-2-control-plane
kube-system          etcd-kind-cilium-mesh-2-control-plane                      1/1     Running       0          90m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          etcd-operator-59cf4cfb7c-4snqh                             1/1     Running       0          81m   10.20.2.109   kind-cilium-mesh-2-worker       
kube-system          kube-apiserver-kind-cilium-mesh-2-control-plane            1/1     Running       0          90m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          kube-controller-manager-kind-cilium-mesh-2-control-plane   1/1     Running       0          90m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          kube-proxy-tk2xm                                           1/1     Running       0          90m   172.17.0.5    kind-cilium-mesh-2-worker       
kube-system          kube-proxy-wwl2g                                           1/1     Running       0          90m   172.17.0.6    kind-cilium-mesh-2-worker2      
kube-system          kube-proxy-zrz4z                                           1/1     Running       0          91m   172.17.0.7    kind-cilium-mesh-2-control-plane
kube-system          kube-scheduler-kind-cilium-mesh-2-control-plane            1/1     Running       0          90m   172.17.0.7    kind-cilium-mesh-2-control-plane
local-path-storage   local-path-provisioner-7745554f7f-9mmr8                    1/1     Running       1          91m   10.20.0.114   kind-cilium-mesh-2-control-plane
```

running

```shell
$ kubectl-cilium exec local-path-storage/local-path-provisioner-7745554f7f-9mmr8 cilium monitor
```

will start monitoring all Cilium-managed traffic in node `kind-cilium-mesh-2-control-plane`:

```shell
$ kubectl-cilium exec local-path-storage/local-path-provisioner-7745554f7f-9mmr8 cilium monitor
Listening for events on 4 CPUs with 64x4096 of shared memory
Press Ctrl-C to quit
level=info msg="Initializing dissection cache..." subsys=monitor
-> overlay flow 0x883a3a1b identity 4->0 state new ifindex cilium_vxlan orig-ip 0.0.0.0: 10.20.0.234:4240 -> 10.20.2.90:59456 tcp ACK
-> endpoint 1534 flow 0xef026664 identity 6->4 state established ifindex lxc_health orig-ip 10.20.2.90: 10.20.2.90:59456 -> 10.20.0.234:4240 tcp ACK
-> endpoint 1980 flow 0xb0e8433f identity 1->104 state new ifindex lxcb42c26c3bbf8 orig-ip 10.20.0.122: 10.20.0.122:47168 -> 10.20.0.68:8181 tcp SYN
-> host from flow 0x4a33a3f3 identity 104->1 state reply ifindex cilium_net orig-ip 0.0.0.0: 10.20.0.68:8181 -> 10.20.0.122:47168 tcp SYN, ACK
(...)
```

If no command is specified, `/bin/bash` is used by default:

```shell
$ kubectl-cilium exec local-path-storage/local-path-provisioner-7745554f7f-9mmr8
root@kind-cilium-mesh-2-control-plane:~#
```

## License

Copyright 2020 bmcstdio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
