package domain

type AppSyncEvent struct {
	OperationName string                 `mapstructure:"operationName"`
	Arguments     map[string]interface{} `mapstructure:"arguments"`
}
