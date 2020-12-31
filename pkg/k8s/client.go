package k8s

import (
	"apascualco.com/type/k8s"
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strconv"
	"strings"
	"time"
)

type KClient struct {
	Client *kubernetes.Clientset
}

func New(masterUrl string, kubeConfig string) *KClient {
	if _, err := os.Stat(kubeConfig); os.IsNotExist(err) {
		fmt.Println("NOT EXIST")
	}
	config, _ := clientcmd.BuildConfigFromFlags(masterUrl, kubeConfig)
	client, _ := kubernetes.NewForConfig(config)
	return &KClient{
		Client: client,
	}
}

func (k *KClient) Nodes() []*k8s.Node {
	nodes, _ := k.Client.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	items := nodes.Items
	return k8sNodeToNode(items)
}

func (k *KClient) Namespaces() []*k8s.Namespace {
	namespaces, _ := k.Client.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	items := namespaces.Items
	return k8sNamespaceToNamespace(items)
}

func k8sNodeToNode(nodes []v12.Node) []*k8s.Node {
	var k8sNodes []*k8s.Node

	for _, node := range nodes {
		uptime := strconv.Itoa(int(time.Since(node.ObjectMeta.CreationTimestamp.Time).Hours()/24)) + "d"
		k8sNode := &k8s.Node{
			Name:             node.Name,
			Date:             uptime,
			Version:          node.APIVersion,
			Os:               node.Status.NodeInfo.OSImage,
			KernelVersion:    node.Status.NodeInfo.KernelVersion,
			ContainerRuntime: node.Status.NodeInfo.ContainerRuntimeVersion,
			Architecture:     node.Status.NodeInfo.Architecture,
		}
		for _, c := range node.Status.Conditions {
			if "True" == c.Status {
				k8sNode.Status = string(c.Type)
			}
		}
		var roles []string
		for _, r := range node.Spec.Taints {
			if r.Key != "" {
				roles = append(roles, strings.Split(r.Key, "/")[1])
			}
		}
		if len(roles) > 0 {
			k8sNode.Roles = strings.Join(roles, ",")
		}
		k8sNodes = append(k8sNodes, k8sNode)
	}
	return k8sNodes
}

func k8sNamespaceToNamespace(namespaces []v12.Namespace) []*k8s.Namespace {
	var k8sNamespaces []*k8s.Namespace
	for _, namespace := range namespaces {
		uptime := strconv.Itoa(int(time.Since(namespace.CreationTimestamp.Time).Hours()/24)) + "d"
		k8sNamespace := &k8s.Namespace{
			Name: namespace.Name,
			Date: uptime,
		}
		k8sNamespaces = append(k8sNamespaces, k8sNamespace)

	}
	return k8sNamespaces
}
