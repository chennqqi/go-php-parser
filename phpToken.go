package php

import (
	"fmt"
	"sort"
	"strconv"
)

type Item struct {
	typ ItemType
	pos Location
	val string
}

type Location struct {
	Pos  int
	Line int
	File string
}

type ItemType int

const (
	itemEOF ItemType = iota
	itemHTML
	itemPHP
	itemPHPBegin
	itemPHPEnd
	itemPHPToken
	itemError
	itemSpace
	itemFunction
	itemStatic
	itemSelf
	itemParent
	itemFinal
	itemFunctionName
	itemTypeHint
	itemVariableOperator
	itemBlockBegin
	itemBlockEnd
	itemGlobal

	itemNamespace
	itemUse

	itemComment

	itemIgnoreErrorOperator

	itemReturn
	itemFunctionCallBegin
	itemFunctionCallEnd
	itemArgumentType
	itemArgumentName
	itemComma
	itemStatementEnd
	itemEcho

	itemIf
	itemElse
	itemElseIf
	itemFor
	itemForeach
	itemEndIf
	itemEndFor
	itemEndForeach
	itemEndWhile
	itemEndSwitch
	itemAsOperator
	itemWhile
	itemContinue
	itemBreak
	itemDo
	itemOpenParen
	itemCloseParen
	itemSwitch
	itemCase
	itemDefault

	itemTry
	itemCatch
	itemFinally
	itemThrow

	itemClass
	itemAbstract
	itemPrivate
	itemPublic
	itemProtected
	itemInterface
	itemImplements
	itemExtends
	itemNewOperator
	itemConst

	itemNull
	itemStringLiteral
	itemNumberLiteral
	itemBooleanLiteral

	itemShellCommand

	itemIdentifier

	itemAssignmentOperator
	itemNegationOperator
	itemAdditionOperator
	itemSubtractionOperator
	itemMultOperator
	itemConcatenationOperator
	itemUnaryOperator
	itemComparisonOperator
	itemInstanceofOperator

	itemAndOperator
	itemOrOperator
	itemWrittenAndOperator
	itemWrittenXorOperator
	itemWrittenOrOperator

	itemObjectOperator
	itemScopeResolutionOperator

	itemCastOperator

	itemArray
	itemArrayKeyOperator
	itemArrayLookupOperatorLeft
	itemArrayLookupOperatorRight
	itemList
	itemBitwiseShiftOperator
	itemEqualityOperator
	itemAmpersandOperator
	itemBitwiseXorOperator
	itemBitwiseOrOperator
	itemTernaryOperator1
	itemTernaryOperator2

	itemInclude
	itemExit
)

// itemTypeMap maps itemType to strings that may be used for debugging and error messages
var itemTypeMap = map[ItemType]string{
	itemHTML:             "HTML",
	itemPHP:              "PHP",
	itemPHPBegin:         "PHP Begin",
	itemPHPEnd:           "PHP End",
	itemPHPToken:         "PHP Token",
	itemEOF:              "EOF",
	itemError:            "Error",
	itemSpace:            "Space",
	itemFunction:         "Function",
	itemStatic:           "static",
	itemSelf:             "self",
	itemParent:           "parent",
	itemFinal:            "final",
	itemFunctionName:     "Function Name",
	itemTypeHint:         "Function Type Hint",
	itemVariableOperator: "$",
	itemBlockBegin:       "Block Begin",
	itemBlockEnd:         "Block End",

	itemGlobal:            "global",
	itemReturn:            "Return",
	itemFunctionCallBegin: "Function Call Begin",
	itemFunctionCallEnd:   "Function Call End",
	itemArgumentType:      "Function Argument Type",
	itemArgumentName:      "Function Argument Name",
	itemComma:             "Function Argument Separator",
	itemStatementEnd:      "Statement End",
	itemEcho:              "Echo",

	itemNamespace: "namespace",
	itemUse:       "use",

	itemIgnoreErrorOperator: "@",

	itemIf:         "If",
	itemElse:       "Else",
	itemElseIf:     "ElseIf",
	itemFor:        "for",
	itemForeach:    "foreach",
	itemSwitch:     "switch",
	itemCase:       "case",
	itemDefault:    "default",
	itemAsOperator: "as",
	itemWhile:      "while",
	itemDo:         "do",
	itemOpenParen:  "open-paren",
	itemCloseParen: "close-paren",
	itemContinue:   "continue",
	itemBreak:      "break",
	itemNull:       "null",

	itemComment: "/* */",

	itemTry:     "try",
	itemCatch:   "catch",
	itemFinally: "finally",
	itemThrow:   "throw",

	itemClass:       "Class",
	itemConst:       "Const",
	itemAbstract:    "abstract",
	itemPrivate:     "Private",
	itemProtected:   "Protected",
	itemPublic:      "Public",
	itemInterface:   "Interface",
	itemImplements:  "implements",
	itemExtends:     "extends",
	itemNewOperator: "new",

	itemShellCommand:   "`",
	itemStringLiteral:  "sting-literal",
	itemNumberLiteral:  "number-literal",
	itemBooleanLiteral: "bool-literal",

	itemIdentifier: "identifier",

	itemAssignmentOperator:      "=",
	itemNegationOperator:        "!",
	itemAdditionOperator:        "+",
	itemSubtractionOperator:     "-",
	itemMultOperator:            "*/%",
	itemConcatenationOperator:   ".",
	itemUnaryOperator:           "++|--",
	itemComparisonOperator:      "==<>",
	itemObjectOperator:          "->",
	itemScopeResolutionOperator: "::",
	itemInstanceofOperator:      "instanceof",

	itemAndOperator:        "&&",
	itemOrOperator:         "||",
	itemWrittenAndOperator: "logical-and",
	itemWrittenXorOperator: "logical-xor",
	itemWrittenOrOperator:  "logical-or",
	itemCastOperator:       "(type)",

	itemList:                     "list",
	itemArray:                    "array",
	itemArrayKeyOperator:         "=>",
	itemArrayLookupOperatorLeft:  "[",
	itemArrayLookupOperatorRight: "]",
	itemBitwiseShiftOperator:     "<<>>",
	itemEqualityOperator:         "!===",
	itemAmpersandOperator:        "&",
	itemBitwiseXorOperator:       "^",
	itemBitwiseOrOperator:        "|",
	itemTernaryOperator1:         "?",
	itemTernaryOperator2:         ":",

	itemInclude: "include",
	itemExit:    "exit",
}

var tokenList []string

func init() {
	tokenList = make([]string, len(tokenMap))
	i := 0
	for token := range tokenMap {
		tokenList[i] = token
		i += 1
	}
	sort.Sort(sort.Reverse(sort.StringSlice(tokenList)))
}

// tokenMap maps source code string tokens to item types when strings can
// be represented directly. Not all item types will be represented here.
var tokenMap = map[string]ItemType{
	"class":          itemClass,
	"clone":          itemUnaryOperator,
	"const":          itemConst,
	"abstract":       itemAbstract,
	"interface":      itemInterface,
	"implements":     itemImplements,
	"extends":        itemExtends,
	"new":            itemNewOperator,
	"if":             itemIf,
	"else":           itemElse,
	"elseif":         itemElseIf,
	"while":          itemWhile,
	"do":             itemDo,
	"for":            itemFor,
	"foreach":        itemForeach,
	"switch":         itemSwitch,
	"endif;":         itemEndIf,
	"endfor;":        itemEndFor,
	"endforeach;":    itemEndForeach,
	"itemendwhile;":  itemEndWhile,
	"itemendswitch;": itemEndSwitch,
	"case":           itemCase,
	"break":          itemBreak,
	"continue":       itemContinue,
	"default":        itemDefault,
	"function":       itemFunction,
	"static":         itemStatic,
	"final":          itemFinal,
	"self":           itemSelf,
	"parent":         itemParent,
	"return":         itemReturn,
	"{":              itemBlockBegin,
	"}":              itemBlockEnd,
	";":              itemStatementEnd,
	"(":              itemOpenParen,
	")":              itemCloseParen,
	",":              itemComma,
	"echo":           itemEcho,
	"throw":          itemThrow,
	"try":            itemTry,
	"catch":          itemCatch,
	"finally":        itemFinally,
	"private":        itemPrivate,
	"public":         itemPublic,
	"protected":      itemProtected,
	"true":           itemBooleanLiteral,
	"false":          itemBooleanLiteral,
	"instanceof":     itemInstanceofOperator,
	"global":         itemGlobal,
	"list":           itemList,
	"array":          itemArray,
	"exit":           itemExit,
	"include":        itemInclude,
	"include_once":   itemInclude,
	"require":        itemInclude,
	"require_once":   itemInclude,
	"@":              itemIgnoreErrorOperator,
	"null":           itemNull,
	"NULL":           itemNull,

	"use":       itemUse,
	"namespace": itemNamespace,

	"(int)":     itemCastOperator,
	"(integer)": itemCastOperator,
	"(bool)":    itemCastOperator,
	"(boolean)": itemCastOperator,
	"(float)":   itemCastOperator,
	"(double)":  itemCastOperator,
	"(real)":    itemCastOperator,
	"(string)":  itemCastOperator,
	"(array)":   itemCastOperator,
	"(object)":  itemCastOperator,
	"(unset)":   itemCastOperator,

	"/*": itemComment,
	"*/": itemComment,
	"//": itemComment,
	"#":  itemComment,

	"->": itemObjectOperator,
	"::": itemScopeResolutionOperator,

	"+=":  itemAssignmentOperator,
	"-=":  itemAssignmentOperator,
	"*=":  itemAssignmentOperator,
	"/=":  itemAssignmentOperator,
	".=":  itemAssignmentOperator,
	"%=":  itemAssignmentOperator,
	"&=":  itemAssignmentOperator,
	"|=":  itemAssignmentOperator,
	"^=":  itemAssignmentOperator,
	"<<=": itemAssignmentOperator,
	">>=": itemAssignmentOperator,
	"=>":  itemArrayKeyOperator,

	"===": itemComparisonOperator,
	"==":  itemComparisonOperator,
	"=":   itemAssignmentOperator,
	"!==": itemComparisonOperator,
	"!=":  itemComparisonOperator,
	"<>":  itemComparisonOperator,
	"!":   itemNegationOperator,
	"++":  itemUnaryOperator,
	"--":  itemUnaryOperator,
	"+":   itemAdditionOperator,
	"-":   itemSubtractionOperator,
	"*":   itemMultOperator,
	"/":   itemMultOperator,
	">=":  itemComparisonOperator,
	">":   itemComparisonOperator,
	"<=":  itemComparisonOperator,
	"<":   itemComparisonOperator,
	"%":   itemMultOperator,
	".":   itemConcatenationOperator,

	"&&":  itemAndOperator,
	"||":  itemOrOperator,
	"&":   itemAmpersandOperator,
	"^":   itemBitwiseXorOperator,
	"|":   itemBitwiseOrOperator,
	"<<":  itemBitwiseShiftOperator,
	">>":  itemBitwiseShiftOperator,
	"?":   itemTernaryOperator1,
	":":   itemTernaryOperator2,
	"and": itemWrittenAndOperator,
	"xor": itemWrittenXorOperator,
	"or":  itemWrittenOrOperator,
	"as":  itemAsOperator,

	"[": itemArrayLookupOperatorLeft,
	"]": itemArrayLookupOperatorRight,

	"$": itemVariableOperator,
}

func (i ItemType) String() string {
	itemTypeName, ok := itemTypeMap[i]
	if !ok {
		return strconv.Itoa(int(i))
	}
	return itemTypeName
}

func (i Item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%v:%.10q...", i.typ, i.val)
	}
	return fmt.Sprintf("%v:%q", i.typ, i.val)
}
