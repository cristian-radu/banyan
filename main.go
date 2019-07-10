package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/rest"
)

func main() {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	apiextensionsClient, err := apiextensionsclientset.NewForConfig(config)

	_, err = createCRD(apiextensionsClient, v1alpha1.DomainCRD)
	if err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL)

	sig := <-sigs
	log.Printf("Received signal %v, exiting.", sig)
	os.Exit(0)

}

func createCRD(client *apiextensionsclientset.Clientset, CRDIn *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	CRDOut, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(CRDIn)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return CRDOut, nil
	}
	return nil, err
}
