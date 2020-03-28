// Copyright 2020 bmcstdio
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

package cilium

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/bmcstdio/kubectl-cilium/internal/constants"
)

func DiscoverCiliumNamespace(kubeClient kubernetes.Interface) (string, error) {
	pp, err := kubeClient.CoreV1().Pods(corev1.NamespaceAll).List(metav1.ListOptions{
		LabelSelector: constants.CiliumLabelSelector,
		Limit:         1,
	})
	if err != nil {
		return "", fmt.Errorf("failed to discover Cilium namespace: %v", err)
	}
	if len(pp.Items) == 0 {
		return "", fmt.Errorf("failed to discover Cilium namespace: no Cilium pods seem to be running in the cluster")
	}
	return pp.Items[0].Namespace, nil
}
