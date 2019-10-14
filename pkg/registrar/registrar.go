package registrar

type Interface interface {
	isDomainAvailable(name string) (bool, error)
}
