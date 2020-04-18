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
	"fmt"

	errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNodeNameFromPod(kubeClient kubernetes.Interface, namespace, name string) (string, error) {
	n, err := kubeClient.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return "", fmt.Errorf("pod %q does not exist in namespace %q", name, namespace)
		}
		return "", nil
	}
	if n.Spec.NodeName == "" {
		return "", fmt.Errorf("pod %q in namespace %q hasn't been assigned to a node", name, namespace)
	}
	return n.Spec.NodeName, nil
}
