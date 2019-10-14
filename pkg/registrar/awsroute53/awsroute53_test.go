package awsroute53

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
)

type mockedDomainAvailability struct {
	route53domainsiface.Route53DomainsAPI
	Response *route53domains.CheckDomainAvailabilityOutput
}

func (m mockedDomainAvailability) CheckDomainAvailability(input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
	return m.Response, nil
}

func TestIsDomainAvailable(t *testing.T) {
	cases := []struct {
		DomainName     string
		Response       *route53domains.CheckDomainAvailabilityOutput
		ExpectedResult bool
	}{
		{
			"testdomain1",
			&domainAvailable,
			true,
		},
		{
			"testdomain2",
			&domainUnavailable,
			false,
		},
	}

	for _, c := range cases {
		r53 := AWSRoute53{
			client: mockedDomainAvailability{Response: c.Response},
		}
		check, err := r53.IsNameAvailable(c.DomainName)
		if err != nil {
			t.Errorf("Test failed with error: %s", err.Error())
		}
		if check != c.ExpectedResult {
			t.Errorf("Domain: %s does not match expected value: %t", c.DomainName, c.ExpectedResult)
		}
	}

}
