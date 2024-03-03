package model

type IModel interface {
	ToMap(model interface{}) (map[string]interface{}, error)
	Fields() []string
}
