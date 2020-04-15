package apis

import (
	"github.com/thefirstofthe300/mass-relay/pkg/apis/massrelay/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
