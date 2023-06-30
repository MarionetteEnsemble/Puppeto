package engine

import (
	"math"
	"path"
	"strconv"
	"strings"
)

type ReadFileFunc = func(stack Stack, specifier string) (string, any)

func StrToF(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)

	if e != nil {
		return math.NaN()
	}

	return f
}

func StrToInt(s string) int64 {
	i, e := strconv.ParseInt(s, 10, 64)

	if e != nil {
		//@TODO: NaN in int??

		return int64(math.NaN())
	}

	return i
}

func IntToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func TypeOf(v any) string {
	switch v.(type) {
	case int64:
		return "integer"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "boolean"
	case nil:
		return "nil"
	case PArray:
		return "array"
	case PObject:
		return "object"
	}

	return "unknown"
}

func ToStr(v any) string {
	switch v := v.(type) {
	case int64:
		return IntToStr(v)
	case float64:
		return FToStr(v)
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	case nil:
		return "nil"
	case PArray:
		ss := []string{}

		for _, v := range v {
			ss = append(ss, ToStr(v))
		}

		return "[" + strings.Join(ss, ", ") + "]"
	case PObject:
		ss := []string{}

		for k, v := range v {
			ss = append(ss, "\n\t"+k+": "+ToStr(v))
		}

		return "{" + strings.Join(ss, ", ") + "}"
	}

	return "unknown"
}

func ToBool(v any) bool {
	switch v := v.(type) {
	case int64:
		return v != 0
	case float64:
		return v != 0
	case string:
		return len(v) > 0
	case bool:
		return v
	case nil:
		return false
	case PArray:
		return true
	case PObject:
		return true
	}

	return false
}

func ToInt(v any) int64 {
	switch v := v.(type) {
	case int64:
		return v
	case float64:
		return int64(v)
	case string:
		return int64(len(v))
	case bool:
		if v {
			return 1
		}

		return 0
	//case nil:
	case PArray:
		return int64(len(v))
	case PObject:
		return int64(len(v))
	}

	return int64(math.NaN())
}

func And(a, b any) bool {
	return ToBool(a) && ToBool(b)
}

func Or(a, b any) bool {
	return ToBool(a) || ToBool(b)
}

func Eq(a, b any) bool {
	return a == b
}

func Neq(a, b any) bool {
	return a != b
}

func Mtoe(a, b any) bool {
	if a == b {
		return true
	}

	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a >= float64(b)
		case float64:
			return a >= b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a >= b
		case float64:
			return float64(a) >= b
		}
	}

	return ToInt(a) >= ToInt(b)
}

func Ltoe(a, b any) bool {
	if a == b {
		return true
	}

	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a <= float64(b)
		case float64:
			return a <= b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a <= b
		case float64:
			return float64(a) <= b
		}
	}

	return ToInt(a) <= ToInt(b)
}

func Mt(a, b any) bool {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a > float64(b)
		case float64:
			return a > b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a > b
		case float64:
			return float64(a) > b
		}
	}

	return ToInt(a) > ToInt(b)
}

func Lt(a, b any) bool {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a < float64(b)
		case float64:
			return a < b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a < b
		case float64:
			return float64(a) < b
		}
	}

	return ToInt(a) < ToInt(b)
}

func Sum(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a + float64(b)
		case float64:
			return a + b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a + b
		case float64:
			return float64(a) + b
		}
	}

	return ToInt(a) + ToInt(b)
}

func Sub(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a - float64(b)
		case float64:
			return a - b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a - b
		case float64:
			return float64(a) - b
		}
	}

	return ToInt(a) - ToInt(b)
}

func Mul(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a * float64(b)
		case float64:
			return a * b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a * b
		case float64:
			return float64(a) * b
		}
	}

	return ToInt(a) * ToInt(b)
}

func Div(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return a / float64(b)
		case float64:
			return a / b
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a / b
		case float64:
			return float64(a) / b
		}
	}

	return ToInt(a) / ToInt(b)
}

func Mod(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) % b
		case float64:
			return int64(a) % int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a % b
		case float64:
			return a % int64(b)
		}
	}

	return ToInt(a) % ToInt(b)
}

func Pow(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return math.Pow(a, float64(b))
		case float64:
			return math.Pow(a, b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return math.Pow(float64(a), float64(b))
		case float64:
			return float64(math.Pow(float64(a), b))
		}
	}

	return math.Pow(float64(ToInt(a)), float64(ToInt(b)))
}

func Band(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) & b
		case float64:
			return int64(a) & int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a & b
		case float64:
			return a & int64(b)
		}
	}

	return ToInt(a) & ToInt(b)
}

func Bor(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) | b
		case float64:
			return int64(a) | int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a | b
		case float64:
			return a | int64(b)
		}
	}

	return ToInt(a) | ToInt(b)
}

func Bxor(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) | b
		case float64:
			return int64(a) | int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a | b
		case float64:
			return a | int64(b)
		}
	}

	return ToInt(a) | ToInt(b)
}

func Bls(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) << b
		case float64:
			return int64(a) << int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a << b
		case float64:
			return a << int64(b)
		}
	}

	return ToInt(a) << ToInt(b)
}

func Brs(a, b any) any {
	switch a := a.(type) {
	case float64:
		switch b := b.(type) {
		case int64:
			return int64(a) >> b
		case float64:
			return int64(a) >> int64(b)
		}
	case int64:
		switch b := b.(type) {
		case int64:
			return a >> b
		case float64:
			return a >> int64(b)
		}
	}

	return ToInt(a) >> ToInt(b)
}

func Neg(a any) any {
	switch a := a.(type) {
	case float64:
		return -a
	case int64:
		return -a
	}

	return -ToInt(a)
}

func Not(a any) bool {
	switch a := a.(type) {
	case bool:
		return !a
	}

	return !ToBool(a)
}

func Get(a, k any) (any, any) {
	o, ok := a.(PObject)

	if !ok {
		a, ok := a.(PArray)

		if !ok {
			return nil, "Expected object or array"
		}

		return a[ToInt(k)], nil
	}

	return o[ToStr(k)], nil
}

func Set(a, k, v any) (any, any) {
	o, ok := a.(PObject)

	if !ok {
		a, ok := a.(PArray)

		if !ok {
			return nil, "Expected object or array"
		}

		a[ToInt(k)] = v

		return v, nil
	}

	o[ToStr(k)] = v

	return v, nil
}

func StringifyError(stack Stack, e PError) string {
	if o, ok := e.(PObject); ok {
		f, ok := o["toString"].(PFunction)

		if !ok {
			return ToStr(e)
		}

		stack := append(Stack{NewPosition(1, 1, "[StringifyError]")}, stack...)
		v, e := f(stack, PArray{})

		if e != nil {
			return StringifyError(stack, e)
		}

		return ToStr(v)
	}

	return ToStr(e)
}

func EvalStr(scope *Scope, s string) (any, any) {
	return RunString(scope, "[eval]", s)
}

func EvalHTMLStr(scope *Scope, s string) (string, any) {
	return RunHTMLString(scope, "[eval]", s)
}

func ImportStr(scope *Scope, file string, s string) (PObject, any) {
	_, e := RunString(scope, file, s)

	if e != nil {
		return nil, e
	}

	return scope.List, nil
}

func ImportFileStr(scope *Scope, file string, readFile ReadFileFunc, stack Stack) (PObject, any) {
	s, e := readFile(stack, file)

	if e != nil {
		return nil, e
	}

	_, e = RunString(scope, file, s)

	if e != nil {
		return nil, e
	}

	return scope.List, nil
}

func JoinPath(rootDir string, file string) string {
	if strings.HasPrefix(file, "/") || strings.HasPrefix(file, "~/") || strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") || strings.HasPrefix(file, "file://") {
		return file
	}

	return path.Join(rootDir, file)
}
