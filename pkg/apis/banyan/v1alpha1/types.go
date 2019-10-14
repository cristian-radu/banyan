package v1alpha1

import (
	"time"

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
	Name                            string        `json:"name"`
	AutoRenew                       bool          `json:"autoRenew"`
	DurationInYears                 int           `json:"durationInYears"`
	AdminContact                    ContactDetail `json:"adminContact"`
	PrivacyProtectAdminContact      bool          `json:"privacyProtectAdminContact"`
	RegistrantContact               ContactDetail `json:"registrantContact"`
	PrivacyProtectRegistrantContact bool          `json:"privacyProtectRegistrantContact"`
	TechContact                     ContactDetail `json:"techContact"`
	PrivacyProtectTechContact       bool          `json:"privacyProtectTechContact"`
}

// DomainStatus describes the status of a Domain resource
type DomainStatus struct {
	StatusList        []string      `json:"statusList"`
	Nameservers       []Nameserver  `json:"nameservers"`
	RegistrarName     string        `json:"registrarName"`
	RegistrarURL      string        `json:"registrarURL"`
	Reseller          string        `json:"reseller"`
	CreationDate      time.Time     `json:"creationTimestamp"`
	ExpirationDate    time.Time     `json:"expirationTimestamp"`
	UpdatedDate       time.Time     `json:"updatedTimestamp"`
	AutoRenew         bool          `json:"autoRenew"`
	AbuseContactEmail string        `json:"abuseContactEmail"`
	AbuseContactPhone string        `json:"abuseContactPhone"`
	AdminContact      ContactDetail `json:"adminContact"`
	AdminPrivacy      bool          `json:"adminPrivacy"`
	RegistrantContact ContactDetail `json:"registrantContact"`
	RegistrantPrivacy bool          `json:"registrantPrivacy"`
	TechContact       ContactDetail `json:"techContact"`
	TechPrivacy       bool          `json:"techPrivacy"`
	WhoIsServer       string        `json:"whoIsServer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DomainList is a list of Domain resources
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Domain `json:"items"`
}

// ContactDetail holds contact information for domains
type ContactDetail struct {
	// Indicates whether the contact is a person, company, association, or public
	// organization.
	ContactType  string `json:"contactType"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city"`
	CountryCode  string `json:"countryCode"`
	Email        string `json:"email"`
	// ExtraParams []*ExtraParam `type:"list"`
	Fax string `json:"fax"`
	// Name of the organization for contact types other than PERSON.
	OrganizationName string `json:"organizationalName,omitempty"`
	PhoneNumber      string `json:"phoneNumber"`
	State            string `json:"state"`
	ZipCode          string `json:"zipCode"`
}

// Nameserver holds information about name servers
type Nameserver struct {
	Name    string   `json:"name"`
	GlueIps []string `json:"glueIps"`
}

// GetKind returns the object kind for a Domain
func (d Domain) GetKind() string {
	return DomainKind
}

// GetName returns the domain name string
func (d Domain) GetName() string {
	return d.GetObjectMeta().GetName()
}
