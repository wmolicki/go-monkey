package evaluator

import (
	"fmt"
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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported: %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got: %d, want: %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError(
					"argument to `first` not supported, must be %s, got %s", object.ARRAY_OBJ,
					args[0].Type())

			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got: %d, want: %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError(
					"argument to `last` not supported, must be %s, got %s", object.ARRAY_OBJ,
					args[0].Type())

			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[len(arr.Elements) - 1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got: %d, want: %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError(
					"argument to `rest` not supported, must be %s, got %s", object.ARRAY_OBJ,
					args[0].Type())

			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElems := make([]object.Object, length-1)
				copy(newElems, arr.Elements[1:length])
				return &object.Array{Elements: newElems}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments, got: %d, want: %d", len(args), 2)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError(
					"argument to `push` not supported, must be %s, got %s", object.ARRAY_OBJ,
					args[0].Type())

			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElems := make([]object.Object, length+1)
			copy(newElems, arr.Elements)
			newElems[length] = args[1]

			return &object.Array{Elements: newElems}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, a := range args {
				fmt.Print(a.Inspect())
			}
			fmt.Print("\n")
			return NULL
		},
	},
}
