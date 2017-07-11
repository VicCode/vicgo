package deepcopy

func Copy(value interface{}) interface{} {
    if valueMap, ok := value.(map[string]interface{}); ok {
        newMap := make(map[string]interface{})
        for k, v := range valueMap {
            newMap[k] = Copy(v)
        }

        return newMap
    } else if valueSlice, ok := value.([]interface{}); ok {
        newSlice := make([]interface{}, len(valueSlice))
        for k, v := range valueSlice {
            newSlice[k] = Copy(v)
        }

        return newSlice
    }

    return value
}

