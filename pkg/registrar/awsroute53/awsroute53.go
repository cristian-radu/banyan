package awsroute53

import (
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
func (r53 *AWSRoute53) IsNameAvailable(name string) (bool, error) {
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
