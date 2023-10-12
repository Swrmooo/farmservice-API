package lib

import (
	"strings"
	"encoding/json"
	"sort"
)

func ListToMapOfArray(arr []map[string]interface{}, code string) map[string][]map[string]interface{}{
	objects := map[string][]map[string]interface{}{}
	for _, rdata := range arr {
		if arr,ok := objects[T(rdata, code)]; ok {
			objects[T(rdata, code)] = append(arr, rdata)
		}else{
			objects[T(rdata, code)] = []map[string]interface{}{rdata}
		}
	}
	return objects
}

func MapToSortList(olist map[string]map[string]interface{}, key1 string, key2 string, key3 string) []map[string]interface{} {
	keys := []string{}
	resultset := map[string]map[string]interface{}{}
	for _, currMap := range olist {
		text := T(currMap,key1)
		if key2 != "" {
			text += T(currMap,key2)
		}
    if key3 != "" {
			text += T(currMap,key3)
		}
		resultset[text] = currMap
		keys = append(keys, text)
	}
	sort.Strings(keys)
	list := []map[string]interface{}{}
	for _, kv := range keys {
		if rdata, ok := resultset[kv]; ok {
			list = append(list, rdata)
		}
	}
	return list
}

func MapOfArrayToList(olist map[string][]map[string]interface{}) []map[string]interface{} {
	rlist := []map[string]interface{}{}
	for _, list := range olist {
		for _, item := range list {
			rlist = append(rlist, item)
		}
	}
	return rlist
}

func RawStringToPureASCII(src string) string{
  s1 := strings.Replace(strings.Replace(strings.Replace(strings.Replace(src,"\r\n","\n",-1),"\n","#NL#",-1),"'","&#39;",-1),`"`,"#DQ#",-1)
  return s1
}

func MapToString(payload map[string]interface{}) string{
  fields := map[string]interface{}{}
  for key, v := range payload {
    switch value := v.(type) {
     case string:
        fields[key] = RawStringToPureASCII(T(payload,key))
        break
     default:
        fields[key] = value
        break
    }
  }
  fieldsByte, _ := json.Marshal(fields)
  return string(fieldsByte)
}

func ArrayOfMapToString(payload []map[string]interface{}) string{
  fieldValue := ""
  for _, obj := range payload {
    for k, v := range obj {
      switch v.(type) {
       case string:
         obj[k] = RawStringToPureASCII(T(obj,k))
         break
       default:
         break
      }
    }
    jsonString, err := json.Marshal(obj);
    if err != nil {
      panic("error.JsonArrayParsingFailed")
    }
    fieldValue += ","+string(jsonString)
  }

  if fieldValue != "" {
    fieldValue = `[`+fieldValue[1:]+`]`
  }
  return fieldValue
}
