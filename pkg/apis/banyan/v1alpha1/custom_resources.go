package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//DomainCRD holds the representation and validation options for the Domain CustomResourceDefinition
var DomainCRD = &apiextensionsv1beta1.CustomResourceDefinition{
	TypeMeta: metav1.TypeMeta{
		Kind:       DomainKind,
		APIVersion: BanyanVersion,
	},
	ObjectMeta: metav1.ObjectMeta{
		Name: DomainFullName,
	},
	Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
		Group:   BanyanGroup,
		Version: BanyanVersion,
		Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
			Singular:   DomainSingular,
			Plural:     DomainPlural,
			Kind:       DomainKind,
			ListKind:   DomainListKind,
			Categories: []string{"banyan"},
		},
		Scope: apiextensionsv1beta1.NamespaceScoped,
		Validation: &apiextensionsv1beta1.CustomResourceValidation{
			OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"apiVersion": apiextensionsv1beta1.JSONSchemaProps{
						Type: "string",
					},
					"kind": apiextensionsv1beta1.JSONSchemaProps{
						Type: "string",
					},
					"metadata": apiextensionsv1beta1.JSONSchemaProps{
						Type: "object",
					},
					"spec": apiextensionsv1beta1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
							"Name": {
								Type: "string",
							},
						},
					},
					"status": apiextensionsv1beta1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
							"Registration": {
								Type: "string",
							},
						},
					},
				},
			},
		},
		Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
			Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
		},
		PreserveUnknownFields: &domainPreserveUnknownFields,
	},
}

var domainPreserveUnknownFields = false
