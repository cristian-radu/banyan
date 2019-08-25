package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"
	"github.com/cristian-radu/banyan/pkg/controller/domain"
	clientset "github.com/cristian-radu/banyan/pkg/generated/clientset/versioned"
	informers "github.com/cristian-radu/banyan/pkg/generated/informers/externalversions"

	log "github.com/sirupsen/logrus"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	banyanClient, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error building domain clientset: %s", err.Error())
	}

	apiextensionsClient, err := apiextensionsclientset.NewForConfig(config)

	_, err = createCRD(apiextensionsClient, v1alpha1.DomainCRD)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL)

	domainInformerFactory := informers.NewSharedInformerFactory(banyanClient, time.Second*30)

	domainController, err := domain.NewController(
		kubeClient,
		banyanClient,
		domainInformerFactory.Banyan().V1alpha1().Domains(),
	)

	if err != nil {
		log.Fatalf("Error creating the domain controller: %s", err.Error())
	}

	domainInformerFactory.Start(stopCh)
	domainInformerFactory.WaitForCacheSync(stopCh)

	// ToDo: configurable number of workers
	if err = domainController.Run(2, stopCh); err != nil {
		log.Fatalf("Error running domain controller: %s", err.Error())
	}

	sig := <-sigs
	log.Infof("Received signal %v, exiting.", sig)
	close(stopCh)
}

func createCRD(client *apiextensionsclientset.Clientset, CRDIn *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	CRDOut, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(CRDIn)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return CRDOut, nil
	}
	return nil, err
}
