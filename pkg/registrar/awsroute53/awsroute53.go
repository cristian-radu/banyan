package awsroute53

import (
	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
)

// AWSRoute53 is a registrar implemented using the AWS Route53Domains service
type AWSRoute53 struct {
	client route53domainsiface.Route53DomainsAPI
}

var (
	available         = "AVAILABLE"
	domainAvailable   = route53domains.CheckDomainAvailabilityOutput{Availability: &available}
	unavailable       = "UNAVAILABLE"
	domainUnavailable = route53domains.CheckDomainAvailabilityOutput{Availability: &unavailable}
)

// New returns an initialized AWS Route53Domains registrar
func New(s *session.Session) AWSRoute53 {
	return AWSRoute53{
		client: route53domains.New(s),
	}
}

// IsNameAvailable checks if a given domain name is available for registering
func (r53 *AWSRoute53) IsNameAvailable(d v1alpha1.Domain) (bool, error) {
	name := d.GetName()
	input := route53domains.CheckDomainAvailabilityInput{
		DomainName: &name,
	}

	output, err := r53.client.CheckDomainAvailability(&input)
	if err != nil {
		return false, err
	}

	if output == &domainAvailable {
		return true, nil
	}
	return false, nil
}

// RegisterDomain registers a given domain name
func (r53 *AWSRoute53) RegisterDomain(d v1alpha1.Domain) error {
	input := route53domains.RegisterDomainInput{
		AdminContact: &route53domains.ContactDetail{
			AddressLine1:     &d.Spec.AdminContact.AddressLine1,
			AddressLine2:     &d.Spec.AdminContact.AddressLine2,
			City:             &d.Spec.AdminContact.City,
			ContactType:      &d.Spec.AdminContact.ContactType,
			CountryCode:      &d.Spec.AdminContact.CountryCode,
			Email:            &d.Spec.AdminContact.Email,
			Fax:              &d.Spec.AdminContact.Fax,
			FirstName:        &d.Spec.AdminContact.FirstName,
			LastName:         &d.Spec.AdminContact.LastName,
			OrganizationName: &d.Spec.AdminContact.OrganizationName,
			PhoneNumber:      &d.Spec.AdminContact.PhoneNumber,
			State:            &d.Spec.AdminContact.State,
			ZipCode:          &d.Spec.AdminContact.ZipCode,
		},
		RegistrantContact: &route53domains.ContactDetail{
			AddressLine1:     &d.Spec.RegistrantContact.AddressLine1,
			AddressLine2:     &d.Spec.RegistrantContact.AddressLine2,
			City:             &d.Spec.RegistrantContact.City,
			ContactType:      &d.Spec.RegistrantContact.ContactType,
			CountryCode:      &d.Spec.RegistrantContact.CountryCode,
			Email:            &d.Spec.RegistrantContact.Email,
			Fax:              &d.Spec.RegistrantContact.Fax,
			FirstName:        &d.Spec.RegistrantContact.FirstName,
			LastName:         &d.Spec.RegistrantContact.LastName,
			OrganizationName: &d.Spec.RegistrantContact.OrganizationName,
			PhoneNumber:      &d.Spec.RegistrantContact.PhoneNumber,
			State:            &d.Spec.RegistrantContact.State,
			ZipCode:          &d.Spec.RegistrantContact.ZipCode,
		},
		TechContact: &route53domains.ContactDetail{
			AddressLine1:     &d.Spec.TechContact.AddressLine1,
			AddressLine2:     &d.Spec.TechContact.AddressLine2,
			City:             &d.Spec.TechContact.City,
			ContactType:      &d.Spec.TechContact.ContactType,
			CountryCode:      &d.Spec.TechContact.CountryCode,
			Email:            &d.Spec.TechContact.Email,
			Fax:              &d.Spec.TechContact.Fax,
			FirstName:        &d.Spec.TechContact.FirstName,
			LastName:         &d.Spec.TechContact.LastName,
			OrganizationName: &d.Spec.TechContact.OrganizationName,
			PhoneNumber:      &d.Spec.TechContact.PhoneNumber,
			State:            &d.Spec.TechContact.State,
			ZipCode:          &d.Spec.TechContact.ZipCode,
		},
		AutoRenew:                       &d.Spec.AutoRenew,
		DomainName:                      &d.Spec.Name,
		DurationInYears:                 &d.Spec.DurationInYears,
		PrivacyProtectAdminContact:      &d.Spec.PrivacyProtectAdminContact,
		PrivacyProtectRegistrantContact: &d.Spec.PrivacyProtectRegistrantContact,
		PrivacyProtectTechContact:       &d.Spec.PrivacyProtectTechContact,
	}

	_, err := r53.client.RegisterDomain(&input)
	if err != nil {
		return err
	}

	return nil
}
