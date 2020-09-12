package utils

func SendResponse(success bool, obj map[string]interface{}) map[string]interface{} {
	newObj := make(map[string]interface{})
	for key, value := range obj {
		newObj[key] = value
	}
	newObj["success"] = success
	return newObj
}
