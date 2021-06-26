package config

import (
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
	"gopkg.in/yaml.v3"
)

var Schemaconfig = []byte(
`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "template": {
      "type": "object",
      "properties":{
        "type":{
          "type":"string",
          "enum": [ "file", "http", "s3" ]
        },
        "path": {
          "type": "string"
        }
      }
    },
    "parameters" : {
      "type": ["object", "null"]
    },
    "timeout":{
      "type": ["integer", "null"]
    },
    "stackname":{
      "type": ["string", "null"]
    },
    "tags":{
      "type": ["object", "null"]
    },
    "rolearn":{
      "type": ["string", "null"]
    },
    "servicerolearn":{
      "type": ["string", "null"]
    },
    "externalID":{
      "type": ["string", "null"]
    },
    "duration":{
      "type": ["string", "null"]
    },
    "acctnum":{
      "type": ["string", "null"]
    },
    "expirywindow":{
      "type": ["string", "null"]
    },
    "env":{
      "type": ["string", "null"]
    },
    "region":{
      "type": ["string", "null"]
    },
    "syncwave":{
      "type": ["integer", "null"]
    }
  },
  "additionalProperties": false,
  "required": [
    "template"
  ]
}`)

func ValidateJSONSchema(schemaData []byte) (*jsonschema.RootSchema, error) {
	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		return nil, fmt.Errorf( "schema unmarshalling error %s", err)
	}
	return rs, nil
}


func ConfigValidate(inputSchemaData, schemaData []byte) (bool, error) {
	rs, err := ValidateJSONSchema(schemaData)
	if err != nil{
		return false, fmt.Errorf("failed ValidateJSONSchema due to %s", err)
	}
	var body interface{}
	if err := yaml.Unmarshal(inputSchemaData, &body); err != nil{
		return false, fmt.Errorf("yaml unmarshal failed inputSchemaData due to: %s", err)
	}
	inputContentJson, err := json.Marshal(body)
	if err != nil {
		return false, fmt.Errorf("json marshal failed inputContentJson due to: %s", err)
	}

	if errs, _ := rs.ValidateBytes(inputContentJson); len(errs) > 0 {
		return false, fmt.Errorf("validate failed due to: %s", errs)
	}

	return true, nil
}