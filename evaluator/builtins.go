package evaluator

import (
	"unicode/utf8"

	"github.com/wmolicki/go-monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got: %d, want: %d", len(args), 1)
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(utf8.RuneCountInString(arg.Value))}
			default:
				return newError("argument to `len` not supported: %s", args[0].Type())
			}
		},
	},
}
