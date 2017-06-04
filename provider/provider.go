package provider

type TemplateData map[string]interface{}

type Provider interface {
	GetData() (TemplateData, error)
}
