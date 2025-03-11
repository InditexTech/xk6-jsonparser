package json

import (
	"bytes"

	"github.com/bytedance/sonic"
	"github.com/grafana/sobek"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/modules"
)

var logger = logrus.New()

type JSON struct {
	Vu modules.VU
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (json *JSON) Exports() modules.Exports {
	return modules.Exports{Default: json}
}

func (json *JSON) Unmarshal(s string) sobek.Value {
	return json.Vu.Runtime().ToValue(unmarshal(s))
}

func (json *JSON) Marshal(object any) sobek.Value {
	return json.Vu.Runtime().ToValue(marshal(object))
}

func unmarshal(s string) any {
	var v any
	err := sonic.Unmarshal([]byte(s), &v)

	if err != nil {
		logger.Errorf("Error parsing JSON: %v", err)
		return nil
	}

	return v
}

func marshal(object any) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := sonic.ConfigDefault.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(object)

	if err != nil {
		logger.Errorf("Error marshall object to JSON: %v", err)
		return ""
	}
	return bf.String()
}
