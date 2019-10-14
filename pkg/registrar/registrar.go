package registrar

import (
	"github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"
)

// Interface is the interface that must be implemented by a registrar in order to be used with the domain controller
type Interface interface {
	IsNameAvailable(d v1alpha1.Domain) (bool, error)
	RegisterDomain(d v1alpha1.Domain) error
}
