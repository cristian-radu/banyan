package v1alpha1

const (
	BanyanNamespace = "banyan"
	BanyanGroup     = "banyan.argonauts.dev"
	BanyanVersion   = "v1alpha1"

	DomainKind     = "Domain"
	DomainSingular = "domain"
	DomainPlural   = "domains"
	DomainListKind = "DomainList"
	DomainFullName = DomainPlural + "." + BanyanGroup
)
