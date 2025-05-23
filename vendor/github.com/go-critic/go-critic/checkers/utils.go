package checkers

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/go-critic/go-critic/linter"
)

// goStdlib contains `go list std` command output list.
// `internal` and `vendor` packages are excluded.
// Used to detect packages that belong to standard Go packages distribution.
var goStdlib = map[string]bool{
	"archive/tar":          true,
	"archive/zip":          true,
	"bufio":                true,
	"bytes":                true,
	"cmp":                  true,
	"compress/bzip2":       true,
	"compress/flate":       true,
	"compress/gzip":        true,
	"compress/lzw":         true,
	"compress/zlib":        true,
	"container/heap":       true,
	"container/list":       true,
	"container/ring":       true,
	"context":              true,
	"crypto":               true,
	"crypto/aes":           true,
	"crypto/cipher":        true,
	"crypto/des":           true,
	"crypto/dsa":           true,
	"crypto/ecdh":          true,
	"crypto/ecdsa":         true,
	"crypto/ed25519":       true,
	"crypto/elliptic":      true,
	"crypto/hmac":          true,
	"crypto/md5":           true,
	"crypto/rand":          true,
	"crypto/rc4":           true,
	"crypto/rsa":           true,
	"crypto/sha1":          true,
	"crypto/sha256":        true,
	"crypto/sha512":        true,
	"crypto/subtle":        true,
	"crypto/tls":           true,
	"crypto/x509":          true,
	"crypto/x509/pkix":     true,
	"database/sql":         true,
	"database/sql/driver":  true,
	"debug/buildinfo":      true,
	"debug/dwarf":          true,
	"debug/elf":            true,
	"debug/gosym":          true,
	"debug/macho":          true,
	"debug/pe":             true,
	"debug/plan9obj":       true,
	"embed":                true,
	"encoding":             true,
	"encoding/ascii85":     true,
	"encoding/asn1":        true,
	"encoding/base32":      true,
	"encoding/base64":      true,
	"encoding/binary":      true,
	"encoding/csv":         true,
	"encoding/gob":         true,
	"encoding/hex":         true,
	"encoding/json":        true,
	"encoding/pem":         true,
	"encoding/xml":         true,
	"errors":               true,
	"expvar":               true,
	"flag":                 true,
	"fmt":                  true,
	"go/ast":               true,
	"go/build":             true,
	"go/build/constraint":  true,
	"go/constant":          true,
	"go/doc":               true,
	"go/doc/comment":       true,
	"go/format":            true,
	"go/importer":          true,
	"go/parser":            true,
	"go/printer":           true,
	"go/scanner":           true,
	"go/token":             true,
	"go/types":             true,
	"go/version":           true,
	"hash":                 true,
	"hash/adler32":         true,
	"hash/crc32":           true,
	"hash/crc64":           true,
	"hash/fnv":             true,
	"hash/maphash":         true,
	"html":                 true,
	"html/template":        true,
	"image":                true,
	"image/color":          true,
	"image/color/palette":  true,
	"image/draw":           true,
	"image/gif":            true,
	"image/jpeg":           true,
	"image/png":            true,
	"index/suffixarray":    true,
	"io":                   true,
	"io/fs":                true,
	"io/ioutil":            true,
	"iter":                 true,
	"log":                  true,
	"log/slog":             true,
	"log/syslog":           true,
	"maps":                 true,
	"math":                 true,
	"math/big":             true,
	"math/bits":            true,
	"math/cmplx":           true,
	"math/rand":            true,
	"math/rand/v2":         true,
	"mime":                 true,
	"mime/multipart":       true,
	"mime/quotedprintable": true,
	"net":                  true,
	"net/http":             true,
	"net/http/cgi":         true,
	"net/http/cookiejar":   true,
	"net/http/fcgi":        true,
	"net/http/httptest":    true,
	"net/http/httptrace":   true,
	"net/http/httputil":    true,
	"net/http/pprof":       true,
	"net/mail":             true,
	"net/netip":            true,
	"net/rpc":              true,
	"net/rpc/jsonrpc":      true,
	"net/smtp":             true,
	"net/textproto":        true,
	"net/url":              true,
	"os":                   true,
	"os/exec":              true,
	"os/signal":            true,
	"os/user":              true,
	"path":                 true,
	"path/filepath":        true,
	"plugin":               true,
	"reflect":              true,
	"regexp":               true,
	"regexp/syntax":        true,
	"runtime":              true,
	"runtime/cgo":          true,
	"runtime/coverage":     true,
	"runtime/debug":        true,
	"runtime/metrics":      true,
	"runtime/pprof":        true,
	"runtime/race":         true,
	"runtime/trace":        true,
	"slices":               true,
	"sort":                 true,
	"strconv":              true,
	"strings":              true,
	"structs":              true,
	"sync":                 true,
	"sync/atomic":          true,
	"syscall":              true,
	"testing":              true,
	"testing/fstest":       true,
	"testing/iotest":       true,
	"testing/quick":        true,
	"testing/slogtest":     true,
	"text/scanner":         true,
	"text/tabwriter":       true,
	"text/template":        true,
	"text/template/parse":  true,
	"time":                 true,
	"time/tzdata":          true,
	"unicode":              true,
	"unicode/utf16":        true,
	"unicode/utf8":         true,
	"unique":               true,
	"unsafe":               true,
}

var goBuiltins = map[string]bool{
	// Types
	"any":        true,
	"bool":       true,
	"byte":       true,
	"comparable": true,
	"complex64":  true,
	"complex128": true,
	"error":      true,
	"float32":    true,
	"float64":    true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"rune":       true,
	"string":     true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,

	// Constants
	"true":  true,
	"false": true,
	"iota":  true,

	// Zero value
	"nil": true,

	// Functions
	"append":  true,
	"cap":     true,
	"clear":   true,
	"close":   true,
	"complex": true,
	"copy":    true,
	"delete":  true,
	"imag":    true,
	"len":     true,
	"make":    true,
	"min":     true,
	"max":     true,
	"new":     true,
	"panic":   true,
	"print":   true,
	"println": true,
	"real":    true,
	"recover": true,
}

// isBuiltin reports whether sym belongs to a predefined identifier set.
func isBuiltin(sym string) bool {
	return goBuiltins[sym]
}

// isStdlibPkg reports whether pkg is a package from the Go standard library.
func isStdlibPkg(pkg *types.Package) bool {
	return pkg != nil && goStdlib[pkg.Path()]
}

// isExampleTestFunc reports whether FuncDecl looks like a testable example function.
func isExampleTestFunc(fn *ast.FuncDecl) bool {
	return len(fn.Type.Params.List) == 0 && strings.HasPrefix(fn.Name.String(), "Example")
}

// isUnitTestFunc reports whether FuncDecl declares testing function.
func isUnitTestFunc(ctx *linter.CheckerContext, fn *ast.FuncDecl) bool {
	if !strings.HasPrefix(fn.Name.Name, "Test") {
		return false
	}
	typ := ctx.TypesInfo.TypeOf(fn.Name)
	if sig, ok := typ.(*types.Signature); ok {
		return sig.Results().Len() == 0 &&
			sig.Params().Len() == 1 &&
			sig.Params().At(0).Type().String() == "*testing.T"
	}
	return false
}

// qualifiedName returns called expr fully-qualified name.
//
// It works for simple identifiers like f => "f" and identifiers
// from other package like pkg.f => "pkg.f".
//
// For all unexpected expressions returns empty string.
func qualifiedName(x ast.Expr) string {
	switch x := x.(type) {
	case *ast.SelectorExpr:
		pkg, ok := x.X.(*ast.Ident)
		if !ok {
			return ""
		}
		return pkg.Name + "." + x.Sel.Name
	case *ast.Ident:
		return x.Name
	default:
		return ""
	}
}

// identOf returns identifier for x that can be used to obtain associated types.Object.
// Returns nil for expressions that yield temporary results, like `f().field`.
func identOf(x ast.Node) *ast.Ident {
	switch x := x.(type) {
	case *ast.Ident:
		return x
	case *ast.SelectorExpr:
		return identOf(x.Sel)
	case *ast.TypeAssertExpr:
		// x.(type) - x may contain ident.
		return identOf(x.X)
	case *ast.IndexExpr:
		// x[i] - x may contain ident.
		return identOf(x.X)
	case *ast.StarExpr:
		// *x - x may contain ident.
		return identOf(x.X)
	case *ast.SliceExpr:
		// x[:] - x may contain ident.
		return identOf(x.X)

	default:
		// Note that this function is not comprehensive.
		return nil
	}
}
