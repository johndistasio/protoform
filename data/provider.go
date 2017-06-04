package data

type Data map[string]interface{}

// A Provider is a generic source of data to use when populating a template.
type Provider interface {
	GetData() (Data, error)
}
