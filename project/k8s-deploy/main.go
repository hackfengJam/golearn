package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	var service *string
	var tag *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	service = flag.String("service", "", "service to upgrade")
	tag = flag.String("tag", "latest", "tag to upgrade")

	flag.Parse()
	if *service == "" {
		fmt.Println("service name not provided. ")
		flag.Usage()
		return
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// for {
	// pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	namespace := "default"

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(*service, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Get deployment %v/%v error %v\n", namespace, *service, err)
	}
	newDeployement := deployment.DeepCopy()

	srcImage := newDeployement.Spec.Template.Spec.Containers[0].Image
	srcImageTag := strings.Split(srcImage, ":")
	srcImageName := strings.Join(srcImageTag[:len(srcImageTag)-1], ":")
	newImage := fmt.Sprintf("%v:%v", srcImageName, *tag)
	newDeployement.Spec.Template.Spec.Containers[0].Image = newImage
	fmt.Printf("Trying upgrade service %v, %v -> %v\n", *service, srcImage, newImage)
	_, err = clientset.AppsV1().Deployments(namespace).Update(newDeployement)
	if err != nil {
		fmt.Printf("Upgrade service failure %v \n", err)
	}
	fmt.Println("Upgrade finished.")

	// for i := 1; i < 4; i++ {
	// 	deployment, err = clientset.AppsV1().Deployments(namespace).Get(*service, metav1.GetOptions{})
	// 	if err != nil {
	// 		fmt.Printf("Get deployment %v/%v error %v\n", namespace, *service, err)
	// 	}
	// 	fmt.Printf("Ready: %v UP-TO-DATE: %v AVAILABLE: %v \n", deployment.Status.ReadyReplicas, deployment.Status.UpdatedReplicas, deployment.Status.AvailableReplicas)
	// 	time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	// }

	// 	time.Sleep(10 * time.Second)
	// }
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
