package kubernetes

import (
	"reflect"
	"testing"

	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestSetDomainOwnership(t *testing.T) {
	client := makeTestClient()
	store := NewStore(client)
	err := store.SetDomainOwnership("test")
	if err != nil {
		t.Errorf("Error %s setting domain ownership", err.Error())
	}

	var expected v1.ConfigMap
	expected.ObjectMeta = metav1.ObjectMeta{
		Name:      "banyan-" + "test",
		Namespace: v1alpha1.BanyanNamespace,
	}
	expected.Data = map[string]string{"Domain": "Owned"}

	actual, err := client.CoreV1().ConfigMaps(v1alpha1.BanyanNamespace).Get("banyan-test", metav1.GetOptions{})

	if !reflect.DeepEqual(expected.ObjectMeta, actual.ObjectMeta) {
		t.Errorf("Expected configmap metadata %v, but got %v instead", expected.ObjectMeta, actual.ObjectMeta)
	}
	if !reflect.DeepEqual(expected.Data, actual.Data) {
		t.Errorf("Expected configmap data %v, but got %v instead", expected.Data, actual.Data)
	}
}

func makeTestClient() *fake.Clientset {
	client := fake.NewSimpleClientset()
	_, _ = client.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: v1alpha1.BanyanNamespace,
		},
	})
	return client
}
