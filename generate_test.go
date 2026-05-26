package hyperdbgsdk

import (
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:       "clone/SDK",
		OutputDir:        ".",
		PackageName:      "hyperdbgsdk",
		HeaderOrder:      []string{"HyperDbgSdk.h"},
		ExtraIncludeDirs: []string{"clone", "clone/SDK/headers"},
		Predefined: `
#include <windows.h>
`,
	}})
}
