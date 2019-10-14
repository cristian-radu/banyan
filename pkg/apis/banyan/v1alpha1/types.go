package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Domain describes a DNS domain and is used when interacting with public registrars
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DomainSpec   `json:"spec"`
	Status DomainStatus `json:"status"`
}

// DomainSpec is the spec for a Domain resource
type DomainSpec struct {
	Name string `json:"name"`
}

// DomainStatus describes the status of a Domain resource
type DomainStatus struct {
	Registration string `json:"registration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DomainList is a list of Domain resources
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Domain `json:"items"`
}

func (d Domain) GetKind() string {
	return DomainKind
}

func (d Domain) GetName() string {
	return d.GetObjectMeta().GetName()
}
