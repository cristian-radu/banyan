package kubernetes

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"
	"github.com/cristian-radu/banyan/pkg/ownership"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	v1listers "k8s.io/client-go/listers/core/v1"
)

type Store struct {
	client          kubernetes.Interface
	configMapLister v1listers.ConfigMapLister
}

func NewStore(client kubernetes.Interface) (s Store) {
	sharedInformerFactory := informers.NewSharedInformerFactoryWithOptions(client, 30*time.Second, informers.WithNamespace(v1alpha1.BanyanNamespace))
	return Store{
		client:          client,
		configMapLister: sharedInformerFactory.Core().V1().ConfigMaps().Lister(),
	}
}

func (s Store) Check(item ownership.Item) bool {
	switch item.GetKind() {
	case v1alpha1.DomainKind:
		return s.CheckDomainOwnership(item.GetName())
	default:
		return false
	}
}

func (s Store) Set(item ownership.Item) error {
	switch kind := item.GetKind(); {
	case kind == v1alpha1.DomainKind:
		return s.SetDomainOwnership(item.GetName())
	default:
		err := fmt.Errorf("unable to set ownership; unrecognized item kind: %s", kind)
		return err
	}
}

func (s Store) CheckDomainOwnership(name string) bool {
	_, err := s.configMapLister.ConfigMaps(v1alpha1.BanyanNamespace).Get(name)
	if err != nil {
		log.Printf("failed to check domain ownership: %s", err.Error())
		return false
	}
	return true
}

func (s Store) SetDomainOwnership(name string) error {
	domainConfigMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "banyan-" + name,
		},
		Data: map[string]string{"Domain": "Owned"},
	}
	_, err := s.client.CoreV1().ConfigMaps(v1alpha1.BanyanNamespace).Create(domainConfigMap)
	if err != nil {
		return err
	}
	return nil
}
