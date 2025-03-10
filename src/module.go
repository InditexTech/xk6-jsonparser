package code

import (
	"github.com/inditex/xk6-jsonparser/pkg/jsonparser"
	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/json", new(RootModule))
}

// RootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/json` module instances for each VU.
type RootModule struct{}

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &json.JSON{}
)

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &json.JSON{Vu: vu}
}
