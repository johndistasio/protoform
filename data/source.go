package data

// A Data is a generic map.
type Data map[string]interface{}

// A Source is a generic source of data to use when populating a template.
type Source interface {
	GetData() (Data, error)
}
