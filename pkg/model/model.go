package model

var Models = make([]interface{}, 0)

// AddSyncModels 同步mysql表
func AddSyncModels(m ...interface{}) {
	Models = append(Models, m...)
}
