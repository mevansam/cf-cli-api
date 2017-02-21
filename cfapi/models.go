package cfapi

// Model structs not present in CF CLI API

// ServiceBindingDetail -
type ServiceBindingDetail struct {
	Entity struct {
		AppGUID             string                 `json:"app_guid,omitempty"`
		ServiceInstanceGUID string                 `json:"service_instance_guid,omitempty"`
		Credentials         map[string]interface{} `json:"credentials,omitempty"`
	} `json:"entity,omitempty"`
}
