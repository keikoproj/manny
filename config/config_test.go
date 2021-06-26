package config

import (
	"github.com/onsi/gomega"
	"testing"
)
//Test_Validate_Success
func Test_Validate_Success(t *testing.T) {
	var schemaData = []byte(`{
		 "title": "Person",
      "type": "object",
      "properties": {
          "firstName": {
              "type": "string"
          },
          "lastName": {
              "type": "string"
          },
          "age": {
              "description": "Age in years",
              "type": "integer",
              "minimum": 0
          }
      },
      "required": ["firstName", "lastName"]
    }`)
	g := gomega.NewGomegaWithT(t)
	var valid = []byte(`{
		"firstName":	"foo",
		"lastName":		"bar",
		"age":			24
	}`)
	ret, _ := ConfigValidate(valid, schemaData)
	g.Expect(ret).To(gomega.BeTrue())
}

//Test_Validate_Failed
func Test_Validate_Failed(t *testing.T) {
	var schemaData = []byte(`{
		 "title": "Person",
      "type": "object",
      "properties": {
          "firstName": {
              "type": "string"
          },
          "lastName": {
              "type": "string"
          },
          "age": {
              "description": "Age in years",
              "type": "integer",
              "minimum": 0
          },
      },
      "required": ["firstName", "lastName"]
    }`)
	g := gomega.NewGomegaWithT(t)
	var valid = []byte(`{
		"firstName":	"foo",
		"lastName":		45
	}`)
	ret, _ := ConfigValidate(valid, schemaData)
	g.Expect(ret).To(gomega.BeFalse())
}

func TestGetConfigSchemaError(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	_, err := ValidateJSONSchema([]byte("invalid.schema"))
	g.Expect(err).To(gomega.HaveOccurred())
}