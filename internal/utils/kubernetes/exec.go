// Copyright 2020 bmcustodio
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pod

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

func Exec(client kubernetes.Interface, config *restclient.Config, streams genericclioptions.IOStreams, podNamespace, podName, containerName string, stdin bool, tty bool, command ...string) error {
	req := client.CoreV1().RESTClient().Post().Resource("pods").Namespace(podNamespace).Name(podName).SubResource("exec")
	opt := &v1.PodExecOptions{
		Command:   command,
		Container: containerName,
		Stdin:     stdin,
		Stdout:    true,
		Stderr:    true,
		TTY:       tty,
	}
	req.VersionedParams(
		opt,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		Stdin:  streams.In,
		Stdout: streams.Out,
		Stderr: streams.ErrOut,
	})
}
