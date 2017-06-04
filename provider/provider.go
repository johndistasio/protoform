package provider

type Data map[string]interface{}

type Provider interface {
	GetData() (Data, error)
}
