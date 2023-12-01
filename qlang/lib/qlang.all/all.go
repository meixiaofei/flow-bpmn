package qall

import (
	"github.com/meixiaofei/flow-bpmn/qlang/lib/bufio"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/bytes"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/crypto/md5"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/encoding/hex"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/encoding/json"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/eqlang"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/errors"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/io"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/io/ioutil"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/math"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/meta"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/net/http"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/os"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/path"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/reflect"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/runtime"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/strconv"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/strings"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/sync"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/terminal"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/tpl/extractor"
	"github.com/meixiaofei/flow-bpmn/qlang/lib/version"
	qlang "github.com/meixiaofei/flow-bpmn/qlang/spec"

	// qlang builtin modules
	_ "github.com/meixiaofei/flow-bpmn/qlang/lib/builtin"
	_ "github.com/meixiaofei/flow-bpmn/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
