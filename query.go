package ksqldbx

type Properties map[string]string

func (prop Properties) Set(key, val string) {
	prop[key] = val
}

func (prop Properties) Delete(key string) {
	delete(prop, key)
}

type Variables map[string]any

func (v Variables) Set(key string, val any) {
	v[key] = val
}

func (v Variables) Delete(key string) {
	delete(v, key)
}

type QuerySQL struct {
	QueryID    string     `json:"queryId,omitempty"`
	SQL        string     `json:"sql,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Variables  Variables  `json:"sessionVariables,omitempty"`
}
