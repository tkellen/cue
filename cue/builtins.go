// Code generated by go generate. DO NOT EDIT.

package cue

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"math"
	"math/big"
	"math/bits"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode"

	"cuelang.org/go/cue/literal"
	"cuelang.org/go/cue/parser"
	"cuelang.org/go/cue/token"
	"cuelang.org/go/internal/third_party/yaml"
	goyaml "github.com/ghodss/yaml"
)

func init() {
	initBuiltins(builtinPackages)
}

var _ io.Reader

var split = path.Split

var pathClean = path.Clean

var pathExt = path.Ext

var pathBase = path.Base

var pathIsAbs = path.IsAbs

var pathDir = path.Dir

var builtinPackages = map[string][]*builtin{
	"": []*builtin{{}},
	"crypto/md5": []*builtin{{
		Name:  "Size",
		Const: intFromGo("16"),
	}, {
		Name:  "BlockSize",
		Const: intFromGo("64"),
	}, {
		Name:   "Sum",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return md5.Sum(data)
			}()
		},
	}},
	"crypto/sha1": []*builtin{{
		Name:  "Size",
		Const: intFromGo("20"),
	}, {
		Name:  "BlockSize",
		Const: intFromGo("64"),
	}, {
		Name:   "Sum",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha1.Sum(data)
			}()
		},
	}},
	"crypto/sha256": []*builtin{{
		Name:  "Size",
		Const: intFromGo("32"),
	}, {
		Name:  "Size224",
		Const: intFromGo("28"),
	}, {
		Name:  "BlockSize",
		Const: intFromGo("64"),
	}, {
		Name:   "Sum256",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha256.Sum256(data)
			}()
		},
	}, {
		Name:   "Sum224",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha256.Sum224(data)
			}()
		},
	}},
	"crypto/sha512": []*builtin{{
		Name:  "Size",
		Const: intFromGo("64"),
	}, {
		Name:  "Size224",
		Const: intFromGo("28"),
	}, {
		Name:  "Size256",
		Const: intFromGo("32"),
	}, {
		Name:  "Size384",
		Const: intFromGo("48"),
	}, {
		Name:  "BlockSize",
		Const: intFromGo("128"),
	}, {
		Name:   "Sum512",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha512.Sum512(data)
			}()
		},
	}, {
		Name:   "Sum384",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha512.Sum384(data)
			}()
		},
	}, {
		Name:   "Sum512_224",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha512.Sum512_224(data)
			}()
		},
	}, {
		Name:   "Sum512_256",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return sha512.Sum512_256(data)
			}()
		},
	}},
	"encoding/csv": []*builtin{{
		Name:   "Encode",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			x := c.value(0)
			c.ret, c.err = func() (interface{}, error) {
				buf := &bytes.Buffer{}
				w := csv.NewWriter(buf)
				iter, err := x.List()
				if err != nil {
					return "", err
				}
				for iter.Next() {
					row, err := iter.Value().List()
					if err != nil {
						return "", err
					}
					a := []string{}
					for row.Next() {
						col := row.Value()
						if str, err := col.String(); err == nil {
							a = append(a, str)
						} else {
							b, err := col.MarshalJSON()
							if err != nil {
								return "", err
							}
							a = append(a, string(b))
						}
					}
					w.Write(a)
				}
				w.Flush()
				return buf.String(), nil
			}()
		},
	}, {
		Name:   "Decode",
		Params: []kind{stringKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			r := c.reader(0)
			c.ret, c.err = func() (interface{}, error) {
				return csv.NewReader(r).ReadAll()
			}()
		},
	}},
	"encoding/hex": []*builtin{{
		Name:   "EncodedLen",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			n := c.int(0)
			c.ret = func() interface{} {
				return hex.EncodedLen(n)
			}()
		},
	}, {
		Name:   "DecodedLen",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.int(0)
			c.ret = func() interface{} {
				return hex.DecodedLen(x)
			}()
		},
	}, {
		Name:   "Decode",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret, c.err = func() (interface{}, error) {
				return hex.DecodeString(s)
			}()
		},
	}, {
		Name:   "Dump",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return hex.Dump(data)
			}()
		},
	}, {
		Name:   "Encode",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			src := c.bytes(0)
			c.ret = func() interface{} {
				return hex.EncodeToString(src)
			}()
		},
	}},
	"encoding/json": []*builtin{{
		Name:   "Valid",
		Params: []kind{stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret = func() interface{} {
				return json.Valid(data)
			}()
		},
	}, {
		Name:   "Compact",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			src := c.bytes(0)
			c.ret, c.err = func() (interface{}, error) {
				dst := bytes.Buffer{}
				if err := json.Compact(&dst, src); err != nil {
					return "", err
				}
				return dst.String(), nil
			}()
		},
	}, {
		Name:   "Indent",
		Params: []kind{stringKind, stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			src, prefix, indent := c.bytes(0), c.string(1), c.string(2)
			c.ret, c.err = func() (interface{}, error) {
				dst := bytes.Buffer{}
				if err := json.Indent(&dst, src, prefix, indent); err != nil {
					return "", err
				}
				return dst.String(), nil
			}()
		},
	}, {
		Name:   "HTMLEscape",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			src := c.bytes(0)
			c.ret = func() interface{} {
				dst := &bytes.Buffer{}
				json.HTMLEscape(dst, src)
				return dst.String()
			}()
		},
	}, {
		Name:   "Marshal",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			v := c.value(0)
			c.ret, c.err = func() (interface{}, error) {
				b, err := json.Marshal(v)
				return string(b), err
			}()
		},
	}, {
		Name:   "MarshalStream",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			v := c.value(0)
			c.ret, c.err = func() (interface{}, error) {

				iter, err := v.List()
				if err != nil {
					return "", err
				}
				buf := &bytes.Buffer{}
				for iter.Next() {
					b, err := json.Marshal(iter.Value())
					if err != nil {
						return "", err
					}
					buf.Write(b)
					buf.WriteByte('\n')
				}
				return buf.String(), nil
			}()
		},
	}, {
		Name:   "Unmarshal",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			b := c.bytes(0)
			c.ret, c.err = func() (interface{}, error) {
				if !json.Valid(b) {
					return nil, fmt.Errorf("json: invalid JSON")
				}
				fset := token.NewFileSet()
				expr, err := parser.ParseExpr(fset, "json", b)
				if err != nil {

					return nil, fmt.Errorf("json: could not parse JSON: %v", err)
				}
				return expr, nil
			}()
		},
	}},
	"encoding/yaml": []*builtin{{
		Name:   "Marshal",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			v := c.value(0)
			c.ret, c.err = func() (interface{}, error) {
				b, err := goyaml.Marshal(v)
				return string(b), err
			}()
		},
	}, {
		Name:   "MarshalStream",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			v := c.value(0)
			c.ret, c.err = func() (interface{}, error) {

				iter, err := v.List()
				if err != nil {
					return "", err
				}
				buf := &bytes.Buffer{}
				for i := 0; iter.Next(); i++ {
					if i > 0 {
						buf.WriteString("---\n")
					}
					b, err := goyaml.Marshal(iter.Value())
					if err != nil {
						return "", err
					}
					buf.Write(b)
				}
				return buf.String(), nil
			}()
		},
	}, {
		Name:   "Unmarshal",
		Params: []kind{stringKind},
		Result: topKind,
		Func: func(c *callCtxt) {
			data := c.bytes(0)
			c.ret, c.err = func() (interface{}, error) {
				fset := token.NewFileSet()
				return yaml.Unmarshal(fset, "", data)
			}()
		},
	}},
	"html": []*builtin{{
		Name:   "Escape",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return html.EscapeString(s)
			}()
		},
	}, {
		Name:   "Unescape",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return html.UnescapeString(s)
			}()
		},
	}},
	"list": []*builtin{{}},
	"math": []*builtin{{
		Name:  "MaxExp",
		Const: intFromGo("2147483647"),
	}, {
		Name:  "MinExp",
		Const: intFromGo("-2147483648"),
	}, {
		Name:  "MaxPrec",
		Const: intFromGo("4294967295"),
	}, {
		Name:  "ToNearestEven",
		Const: intFromGo("0"),
	}, {
		Name:  "ToNearestAway",
		Const: intFromGo("1"),
	}, {
		Name:  "ToZero",
		Const: intFromGo("2"),
	}, {
		Name:  "AwayFromZero",
		Const: intFromGo("3"),
	}, {
		Name:  "ToNegativeInf",
		Const: intFromGo("4"),
	}, {
		Name:  "ToPositiveInf",
		Const: intFromGo("5"),
	}, {
		Name:  "Below",
		Const: intFromGo("-1"),
	}, {
		Name:  "Exact",
		Const: intFromGo("0"),
	}, {
		Name:  "Above",
		Const: intFromGo("1"),
	}, {
		Name:   "Jacobi",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x, y := c.bigInt(0), c.bigInt(1)
			c.ret = func() interface{} {
				return big.Jacobi(x, y)
			}()
		},
	}, {
		Name:  "MaxBase",
		Const: intFromGo("62"),
	}, {
		Name:   "Floor",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Floor(x)
			}()
		},
	}, {
		Name:   "Ceil",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Ceil(x)
			}()
		},
	}, {
		Name:   "Trunc",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Trunc(x)
			}()
		},
	}, {
		Name:   "Round",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Round(x)
			}()
		},
	}, {
		Name:   "RoundToEven",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.RoundToEven(x)
			}()
		},
	}, {
		Name:   "Abs",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Abs(x)
			}()
		},
	}, {
		Name:   "Acosh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Acosh(x)
			}()
		},
	}, {
		Name:   "Asin",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Asin(x)
			}()
		},
	}, {
		Name:   "Acos",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Acos(x)
			}()
		},
	}, {
		Name:   "Asinh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Asinh(x)
			}()
		},
	}, {
		Name:   "Atan",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Atan(x)
			}()
		},
	}, {
		Name:   "Atan2",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			y, x := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Atan2(y, x)
			}()
		},
	}, {
		Name:   "Atanh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Atanh(x)
			}()
		},
	}, {
		Name:   "Cbrt",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Cbrt(x)
			}()
		},
	}, {
		Name:  "E",
		Const: floatFromGo("2.71828182845904523536028747135266249775724709369995957496696763"),
	}, {
		Name:  "Pi",
		Const: floatFromGo("3.14159265358979323846264338327950288419716939937510582097494459"),
	}, {
		Name:  "Phi",
		Const: floatFromGo("1.61803398874989484820458683436563811772030917980576286213544861"),
	}, {
		Name:  "Sqrt2",
		Const: floatFromGo("1.41421356237309504880168872420969807856967187537694807317667974"),
	}, {
		Name:  "SqrtE",
		Const: floatFromGo("1.64872127070012814684865078781416357165377610071014801157507931"),
	}, {
		Name:  "SqrtPi",
		Const: floatFromGo("1.77245385090551602729816748334114518279754945612238712821380779"),
	}, {
		Name:  "SqrtPhi",
		Const: floatFromGo("1.27201964951406896425242246173749149171560804184009624861664038"),
	}, {
		Name:  "Ln2",
		Const: floatFromGo("0.693147180559945309417232121458176568075500134360255254120680009"),
	}, {
		Name:  "Log2E",
		Const: floatFromGo("1.442695040888963407359924681001892137426645954152985934135449408"),
	}, {
		Name:  "Ln10",
		Const: floatFromGo("2.3025850929940456840179914546843642076011014886287729760333278"),
	}, {
		Name:  "Log10E",
		Const: floatFromGo("0.43429448190325182765112891891660508229439700580366656611445378"),
	}, {
		Name:   "Copysign",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x, y := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Copysign(x, y)
			}()
		},
	}, {
		Name:   "Dim",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x, y := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Dim(x, y)
			}()
		},
	}, {
		Name:   "Erf",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Erf(x)
			}()
		},
	}, {
		Name:   "Erfc",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Erfc(x)
			}()
		},
	}, {
		Name:   "Erfinv",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Erfinv(x)
			}()
		},
	}, {
		Name:   "Erfcinv",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Erfcinv(x)
			}()
		},
	}, {
		Name:   "Exp",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Exp(x)
			}()
		},
	}, {
		Name:   "Exp2",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Exp2(x)
			}()
		},
	}, {
		Name:   "Expm1",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Expm1(x)
			}()
		},
	}, {
		Name:   "Gamma",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Gamma(x)
			}()
		},
	}, {
		Name:   "Hypot",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			p, q := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Hypot(p, q)
			}()
		},
	}, {
		Name:   "J0",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.J0(x)
			}()
		},
	}, {
		Name:   "Y0",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Y0(x)
			}()
		},
	}, {
		Name:   "J1",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.J1(x)
			}()
		},
	}, {
		Name:   "Y1",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Y1(x)
			}()
		},
	}, {
		Name:   "Jn",
		Params: []kind{intKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			n, x := c.int(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Jn(n, x)
			}()
		},
	}, {
		Name:   "Yn",
		Params: []kind{intKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			n, x := c.int(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Yn(n, x)
			}()
		},
	}, {
		Name:   "Ldexp",
		Params: []kind{numKind, intKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			frac, exp := c.float64(0), c.int(1)
			c.ret = func() interface{} {
				return math.Ldexp(frac, exp)
			}()
		},
	}, {
		Name:   "Log",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Log(x)
			}()
		},
	}, {
		Name:   "Log10",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Log10(x)
			}()
		},
	}, {
		Name:   "Log2",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Log2(x)
			}()
		},
	}, {
		Name:   "Log1p",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Log1p(x)
			}()
		},
	}, {
		Name:   "Logb",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Logb(x)
			}()
		},
	}, {
		Name:   "Ilogb",
		Params: []kind{numKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Ilogb(x)
			}()
		},
	}, {
		Name:   "Mod",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x, y := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Mod(x, y)
			}()
		},
	}, {
		Name:   "Pow",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x, y := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Pow(x, y)
			}()
		},
	}, {
		Name:   "Pow10",
		Params: []kind{intKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			n := c.int(0)
			c.ret = func() interface{} {
				return math.Pow10(n)
			}()
		},
	}, {
		Name:   "Remainder",
		Params: []kind{numKind, numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x, y := c.float64(0), c.float64(1)
			c.ret = func() interface{} {
				return math.Remainder(x, y)
			}()
		},
	}, {
		Name:   "Signbit",
		Params: []kind{numKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Signbit(x)
			}()
		},
	}, {
		Name:   "Cos",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Cos(x)
			}()
		},
	}, {
		Name:   "Sin",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Sin(x)
			}()
		},
	}, {
		Name:   "Sinh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Sinh(x)
			}()
		},
	}, {
		Name:   "Cosh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Cosh(x)
			}()
		},
	}, {
		Name:   "Sqrt",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Sqrt(x)
			}()
		},
	}, {
		Name:   "Tan",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Tan(x)
			}()
		},
	}, {
		Name:   "Tanh",
		Params: []kind{numKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			x := c.float64(0)
			c.ret = func() interface{} {
				return math.Tanh(x)
			}()
		},
	}},
	"math/bits": []*builtin{{
		Name:   "And",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			a, b := c.bigInt(0), c.bigInt(1)
			c.ret = func() interface{} {
				wa := a.Bits()
				wb := b.Bits()
				n := len(wa)
				if len(wb) < n {
					n = len(wb)
				}
				w := make([]big.Word, n)
				for i := range w {
					w[i] = wa[i] & wb[i]
				}
				i := &big.Int{}
				i.SetBits(w)
				return i
			}()
		},
	}, {
		Name:   "Or",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			a, b := c.bigInt(0), c.bigInt(1)
			c.ret = func() interface{} {
				wa := a.Bits()
				wb := b.Bits()
				var w []big.Word
				n := len(wa)
				if len(wa) > len(wb) {
					w = append(w, wa...)
					n = len(wb)
				} else {
					w = append(w, wb...)
				}
				for i := 0; i < n; i++ {
					w[i] = wa[i] | wb[i]
				}
				i := &big.Int{}
				i.SetBits(w)
				return i
			}()
		},
	}, {
		Name:   "Xor",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			a, b := c.bigInt(0), c.bigInt(1)
			c.ret = func() interface{} {
				wa := a.Bits()
				wb := b.Bits()
				var w []big.Word
				n := len(wa)
				if len(wa) > len(wb) {
					w = append(w, wa...)
					n = len(wb)
				} else {
					w = append(w, wb...)
				}
				for i := 0; i < n; i++ {
					w[i] = wa[i] ^ wb[i]
				}
				i := &big.Int{}
				i.SetBits(w)
				return i
			}()
		},
	}, {
		Name:   "Clear",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			a, b := c.bigInt(0), c.bigInt(1)
			c.ret = func() interface{} {
				wa := a.Bits()
				wb := b.Bits()
				w := append([]big.Word(nil), wa...)
				for i, m := range wb {
					if i >= len(w) {
						break
					}
					w[i] = wa[i] &^ m
				}
				i := &big.Int{}
				i.SetBits(w)
				return i
			}()
		},
	}, {
		Name:   "OnesCount",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.uint64(0)
			c.ret = func() interface{} {
				return bits.OnesCount64(x)
			}()
		},
	}, {
		Name:   "RotateLeft",
		Params: []kind{intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x, k := c.uint64(0), c.int(1)
			c.ret = func() interface{} {
				return bits.RotateLeft64(x, k)
			}()
		},
	}, {
		Name:   "Reverse",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.uint64(0)
			c.ret = func() interface{} {
				return bits.Reverse64(x)
			}()
		},
	}, {
		Name:   "ReverseBytes",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.uint64(0)
			c.ret = func() interface{} {
				return bits.ReverseBytes64(x)
			}()
		},
	}, {
		Name:   "Len",
		Params: []kind{intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			x := c.uint64(0)
			c.ret = func() interface{} {
				return bits.Len64(x)
			}()
		},
	}},
	"path": []*builtin{{
		Name:   "Split",
		Params: []kind{stringKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				file, dir := split(path)
				return []string{file, dir}
			}()
		},
	}, {
		Name:   "Match",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			pattern, name := c.string(0), c.string(1)
			c.ret, c.err = func() (interface{}, error) {
				return path.Match(pattern, name)
			}()
		},
	}, {
		Name:   "Clean",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				return pathClean(path)
			}()
		},
	}, {
		Name:   "Ext",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				return pathExt(path)
			}()
		},
	}, {
		Name:   "Base",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				return pathBase(path)
			}()
		},
	}, {
		Name:   "IsAbs",
		Params: []kind{stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				return pathIsAbs(path)
			}()
		},
	}, {
		Name:   "Dir",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			path := c.string(0)
			c.ret = func() interface{} {
				return pathDir(path)
			}()
		},
	}},
	"regexp": []*builtin{{
		Name:   "Match",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			pattern, s := c.string(0), c.string(1)
			c.ret, c.err = func() (interface{}, error) {
				return regexp.MatchString(pattern, s)
			}()
		},
	}, {
		Name:   "QuoteMeta",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return regexp.QuoteMeta(s)
			}()
		},
	}},
	"runtime": []*builtin{{
		Name:   "Path",
		Params: []kind{},
		Result: stringKind,
		Func: func(c *callCtxt) {
			c.ret = func() interface{} {
				return ""
			}()
		},
	}},
	"strconv": []*builtin{{
		Name:   "Unquote",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret, c.err = func() (interface{}, error) {
				return literal.Unquote(s)
			}()
		},
	}, {
		Name:   "ParseBool",
		Params: []kind{stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			str := c.string(0)
			c.ret, c.err = func() (interface{}, error) {
				return strconv.ParseBool(str)
			}()
		},
	}, {
		Name:   "FormatBool",
		Params: []kind{boolKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			b := c.bool(0)
			c.ret = func() interface{} {
				return strconv.FormatBool(b)
			}()
		},
	}, {
		Name:   "ParseFloat",
		Params: []kind{stringKind, intKind},
		Result: numKind,
		Func: func(c *callCtxt) {
			s, bitSize := c.string(0), c.int(1)
			c.ret, c.err = func() (interface{}, error) {
				return strconv.ParseFloat(s, bitSize)
			}()
		},
	}, {
		Name:  "IntSize",
		Const: intFromGo("64"),
	}, {
		Name:   "ParseUint",
		Params: []kind{stringKind, intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, base, bitSize := c.string(0), c.int(1), c.int(2)
			c.ret, c.err = func() (interface{}, error) {
				return strconv.ParseUint(s, base, bitSize)
			}()
		},
	}, {
		Name:   "ParseInt",
		Params: []kind{stringKind, intKind, intKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, base, bitSize := c.string(0), c.int(1), c.int(2)
			c.ret, c.err = func() (interface{}, error) {
				return strconv.ParseInt(s, base, bitSize)
			}()
		},
	}, {
		Name:   "Atoi",
		Params: []kind{stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret, c.err = func() (interface{}, error) {
				return strconv.Atoi(s)
			}()
		},
	}, {
		Name:   "FormatFloat",
		Params: []kind{numKind, intKind, intKind, intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			f, fmt, prec, bitSize := c.float64(0), c.byte(1), c.int(2), c.int(3)
			c.ret = func() interface{} {
				return strconv.FormatFloat(f, fmt, prec, bitSize)
			}()
		},
	}, {
		Name:   "FormatUint",
		Params: []kind{intKind, intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			i, base := c.uint64(0), c.int(1)
			c.ret = func() interface{} {
				return strconv.FormatUint(i, base)
			}()
		},
	}, {
		Name:   "FormatInt",
		Params: []kind{intKind, intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			i, base := c.int64(0), c.int(1)
			c.ret = func() interface{} {
				return strconv.FormatInt(i, base)
			}()
		},
	}, {
		Name:   "Quote",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strconv.Quote(s)
			}()
		},
	}, {
		Name:   "QuoteToASCII",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strconv.QuoteToASCII(s)
			}()
		},
	}, {
		Name:   "QuoteToGraphic",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strconv.QuoteToGraphic(s)
			}()
		},
	}, {
		Name:   "QuoteRune",
		Params: []kind{intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			r := c.rune(0)
			c.ret = func() interface{} {
				return strconv.QuoteRune(r)
			}()
		},
	}, {
		Name:   "QuoteRuneToASCII",
		Params: []kind{intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			r := c.rune(0)
			c.ret = func() interface{} {
				return strconv.QuoteRuneToASCII(r)
			}()
		},
	}, {
		Name:   "QuoteRuneToGraphic",
		Params: []kind{intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			r := c.rune(0)
			c.ret = func() interface{} {
				return strconv.QuoteRuneToGraphic(r)
			}()
		},
	}, {
		Name:   "CanBackquote",
		Params: []kind{stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strconv.CanBackquote(s)
			}()
		},
	}, {
		Name:   "IsPrint",
		Params: []kind{intKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			r := c.rune(0)
			c.ret = func() interface{} {
				return strconv.IsPrint(r)
			}()
		},
	}, {
		Name:   "IsGraphic",
		Params: []kind{intKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			r := c.rune(0)
			c.ret = func() interface{} {
				return strconv.IsGraphic(r)
			}()
		},
	}},
	"strings": []*builtin{{
		Name:   "ToTitle",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {

				prev := ' '
				return strings.Map(
					func(r rune) rune {
						if unicode.IsSpace(prev) {
							prev = r
							return unicode.ToTitle(r)
						}
						prev = r
						return r
					},
					s)
			}()
		},
	}, {
		Name:   "ToCamel",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {

				prev := ' '
				return strings.Map(
					func(r rune) rune {
						if unicode.IsSpace(prev) {
							prev = r
							return unicode.ToLower(r)
						}
						prev = r
						return r
					},
					s)
			}()
		},
	}, {
		Name:   "Compare",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			a, b := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Compare(a, b)
			}()
		},
	}, {
		Name:   "Count",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, substr := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Count(s, substr)
			}()
		},
	}, {
		Name:   "Contains",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			s, substr := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Contains(s, substr)
			}()
		},
	}, {
		Name:   "ContainsAny",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			s, chars := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.ContainsAny(s, chars)
			}()
		},
	}, {
		Name:   "LastIndex",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, substr := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.LastIndex(s, substr)
			}()
		},
	}, {
		Name:   "IndexAny",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, chars := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.IndexAny(s, chars)
			}()
		},
	}, {
		Name:   "LastIndexAny",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, chars := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.LastIndexAny(s, chars)
			}()
		},
	}, {
		Name:   "SplitN",
		Params: []kind{stringKind, stringKind, intKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			s, sep, n := c.string(0), c.string(1), c.int(2)
			c.ret = func() interface{} {
				return strings.SplitN(s, sep, n)
			}()
		},
	}, {
		Name:   "SplitAfterN",
		Params: []kind{stringKind, stringKind, intKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			s, sep, n := c.string(0), c.string(1), c.int(2)
			c.ret = func() interface{} {
				return strings.SplitAfterN(s, sep, n)
			}()
		},
	}, {
		Name:   "Split",
		Params: []kind{stringKind, stringKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			s, sep := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Split(s, sep)
			}()
		},
	}, {
		Name:   "SplitAfter",
		Params: []kind{stringKind, stringKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			s, sep := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.SplitAfter(s, sep)
			}()
		},
	}, {
		Name:   "Fields",
		Params: []kind{stringKind},
		Result: listKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strings.Fields(s)
			}()
		},
	}, {
		Name:   "Join",
		Params: []kind{listKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			a, sep := c.strList(0), c.string(1)
			if c.do() {
				c.ret = func() interface{} {
					return strings.Join(a, sep)
				}()
			}
		},
	}, {
		Name:   "HasPrefix",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			s, prefix := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.HasPrefix(s, prefix)
			}()
		},
	}, {
		Name:   "HasSuffix",
		Params: []kind{stringKind, stringKind},
		Result: boolKind,
		Func: func(c *callCtxt) {
			s, suffix := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.HasSuffix(s, suffix)
			}()
		},
	}, {
		Name:   "Repeat",
		Params: []kind{stringKind, intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, count := c.string(0), c.int(1)
			c.ret = func() interface{} {
				return strings.Repeat(s, count)
			}()
		},
	}, {
		Name:   "ToUpper",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strings.ToUpper(s)
			}()
		},
	}, {
		Name:   "ToLower",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strings.ToLower(s)
			}()
		},
	}, {
		Name:   "Trim",
		Params: []kind{stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, cutset := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Trim(s, cutset)
			}()
		},
	}, {
		Name:   "TrimLeft",
		Params: []kind{stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, cutset := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.TrimLeft(s, cutset)
			}()
		},
	}, {
		Name:   "TrimRight",
		Params: []kind{stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, cutset := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.TrimRight(s, cutset)
			}()
		},
	}, {
		Name:   "TrimSpace",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return strings.TrimSpace(s)
			}()
		},
	}, {
		Name:   "TrimPrefix",
		Params: []kind{stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, prefix := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.TrimPrefix(s, prefix)
			}()
		},
	}, {
		Name:   "TrimSuffix",
		Params: []kind{stringKind, stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, suffix := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.TrimSuffix(s, suffix)
			}()
		},
	}, {
		Name:   "Replace",
		Params: []kind{stringKind, stringKind, stringKind, intKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s, old, new, n := c.string(0), c.string(1), c.string(2), c.int(3)
			c.ret = func() interface{} {
				return strings.Replace(s, old, new, n)
			}()
		},
	}, {
		Name:   "Index",
		Params: []kind{stringKind, stringKind},
		Result: intKind,
		Func: func(c *callCtxt) {
			s, substr := c.string(0), c.string(1)
			c.ret = func() interface{} {
				return strings.Index(s, substr)
			}()
		},
	}},
	"text/tabwriter": []*builtin{{
		Name:   "Write",
		Params: []kind{topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			data := c.value(0)
			c.ret, c.err = func() (interface{}, error) {
				buf := &bytes.Buffer{}
				tw := tabwriter.NewWriter(buf, 0, 4, 1, ' ', 0)
				b, err := data.Bytes()
				if err != nil {
					return "", err
				}
				_, err = tw.Write(b)
				if err != nil {
					return "", err
				}
				return buf.String(), err
			}()
		},
	}},
	"text/template": []*builtin{{
		Name:   "Execute",
		Params: []kind{stringKind, topKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			templ, data := c.string(0), c.value(1)
			c.ret, c.err = func() (interface{}, error) {
				t, err := template.New("").Parse(templ)
				if err != nil {
					return "", err
				}
				buf := &bytes.Buffer{}
				if err := t.Execute(buf, data); err != nil {
					return "", err
				}
				return buf.String(), nil
			}()
		},
	}, {
		Name:   "HTMLEscape",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return template.HTMLEscapeString(s)
			}()
		},
	}, {
		Name:   "JSEscape",
		Params: []kind{stringKind},
		Result: stringKind,
		Func: func(c *callCtxt) {
			s := c.string(0)
			c.ret = func() interface{} {
				return template.JSEscapeString(s)
			}()
		},
	}},
}