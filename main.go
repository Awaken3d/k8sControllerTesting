package main

import (
//"flag"
"fmt"
//"os"
//"path/filepath"

metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
"k8s.io/client-go/kubernetes"
"k8s.io/client-go/tools/clientcmd"
"k8s.io/client-go/rest"
)

func oocConfig() (*rest.Config, error) {
var kubeconfig string
kubeconfig = "/root/.kube/config"
/*if home := homeDir(); home != "" {
kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
} else {
kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
}
flag.Parse()
*/
// use the current context in kubeconfig
return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func main() {
// creates the in-cluster config
//config, err := rest.InClusterConfig()

config, err := oocConfig()

if err != nil {
panic(err.Error())
}
// creates the clientset
clientset, err := kubernetes.NewForConfig(config)
if err != nil {
panic(err.Error())
}
eps, err := clientset.CoreV1().Endpoints("default").List(metav1.ListOptions{FieldSelector:"metadata.name=kubernetes"})
if err != nil {
panic(err.Error())
}
fmt.Printf("%v\n", eps)
}

