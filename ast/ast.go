package ast

import (
	"regexp"
	"strings"
)

type Ast struct {
	imports []string

	// for rendering go src
	functionName string
	indent       int
	returnType   string
}

func NewAst() *Ast {
	return &Ast{
		imports: []string{"fmt"},
	}
}

func (a *Ast) Imports() []string {
	return a.imports
}

func (a *Ast) addImport(name string) {
	for _, i := range a.imports {
		if i == name {
			// already imported
			return
		}
	}

	a.imports = append(a.imports, name)
}

func (a *Ast) importType(name string) string {
	if strings.Index(name, ".") != -1 {
		parts := strings.Split(name, ".")
		a.addImport(strings.Join(parts[:len(parts)-1], "."))

		parts2 := strings.Split(name, "/")
		return parts2[len(parts2)-1]
	}

	return name
}

type Node interface {
	render(ast *Ast) (string, string)
	AddChild(node Node)
}

func Parse(line string) Node {

	nodeName := strings.SplitN(line, " ", 2)[0]

	switch nodeName {
	case "AlwaysInlineAttr":
		return parseAlwaysInlineAttr(line)
	case "ArraySubscriptExpr":
		return parseArraySubscriptExpr(line)
	case "AsmLabelAttr":
		return parseAsmLabelAttr(line)
	case "AvailabilityAttr":
		return parseAvailabilityAttr(line)
	case "BinaryOperator":
		return parseBinaryOperator(line)
	case "BreakStmt":
		return parseBreakStmt(line)
	case "BuiltinType":
		return parseBuiltinType(line)
	case "CallExpr":
		return parseCallExpr(line)
	case "CharacterLiteral":
		return parseCharacterLiteral(line)
	case "CompoundStmt":
		return parseCompoundStmt(line)
	case "ConditionalOperator":
		return parseConditionalOperator(line)
	case "ConstAttr":
		return parseConstAttr(line)
	case "ConstantArrayType":
		return parseConstantArrayType(line)
	case "CStyleCastExpr":
		return parseCStyleCastExpr(line)
	case "DeclRefExpr":
		return parseDeclRefExpr(line)
	case "DeclStmt":
		return parseDeclStmt(line)
	case "DeprecatedAttr":
		return parseDeprecatedAttr(line)
	case "ElaboratedType":
		return parseElaboratedType(line)
	case "Enum":
		return parseEnum(line)
	case "EnumConstantDecl":
		return parseEnumConstantDecl(line)
	case "EnumDecl":
		return parseEnumDecl(line)
	case "EnumType":
		return parseEnumType(line)
	case "FieldDecl":
		return parseFieldDecl(line)
	case "FloatingLiteral":
		return parseFloatingLiteral(line)
	case "FormatAttr":
		return parseFormatAttr(line)
	case "FunctionDecl":
		return parseFunctionDecl(line)
	case "FunctionProtoType":
		return parseFunctionProtoType(line)
	case "ForStmt":
		return parseForStmt(line)
	case "IfStmt":
		return parseIfStmt(line)
	case "ImplicitCastExpr":
		return parseImplicitCastExpr(line)
	case "IntegerLiteral":
		return parseIntegerLiteral(line)
	case "MallocAttr":
		return parseMallocAttr(line)
	case "MemberExpr":
		return parseMemberExpr(line)
	case "ModeAttr":
		return parseModeAttr(line)
	case "NoThrowAttr":
		return parseNoThrowAttr(line)
	case "NonNullAttr":
		return parseNonNullAttr(line)
	case "ParenExpr":
		return parseParenExpr(line)
	case "ParmVarDecl":
		return parseParmVarDecl(line)
	case "PointerType":
		return parsePointerType(line)
	case "PredefinedExpr":
		return parsePredefinedExpr(line)
	case "QualType":
		return parseQualType(line)
	case "Record":
		return parseRecord(line)
	case "RecordDecl":
		return parseRecordDecl(line)
	case "RecordType":
		return parseRecordType(line)
	case "RestrictAttr":
		return parseRestrictAttr(line)
	case "ReturnStmt":
		return parseReturnStmt(line)
	case "StringLiteral":
		return parseStringLiteral(line)
	case "TranslationUnitDecl":
		return parseTranslationUnitDecl(line)
	case "Typedef":
		return parseTypedef(line)
	case "TypedefDecl":
		return parseTypedefDecl(line)
	case "TypedefType":
		return parseTypedefType(line)
	case "UnaryOperator":
		return parseUnaryOperator(line)
	case "VarDecl":
		return parseVarDecl(line)
	case "WhileStmt":
		return parseWhileStmt(line)
	case "NullStmt":
		return nil
	default:
		panic("Unknown node type: '" + line + "'")
	}
}

func groupsFromRegex(rx, line string) map[string]string {
	// We remove tabs and newlines from the regex. This is purely cosmetic,
	// as the regex input can be quite long and it's nice for the caller to
	// be able to format it in a more readable way.
	fullRegexp := "(?P<address>[0-9a-fx]+) " +
		strings.Replace(strings.Replace(rx, "\n", "", -1), "\t", "", -1)
	re := regexp.MustCompile(fullRegexp)

	match := re.FindStringSubmatch(line)
	if len(match) == 0 {
		panic("could not match regexp '" + fullRegexp +
			"' with string '" + line + "'")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}
