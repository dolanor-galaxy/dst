package decorator

import (
	"fmt"
	"github.com/dave/dst"
	"go/ast"
	"go/token"
)

func (r *fileRestorer) restoreNode(n dst.Node) ast.Node {
	if an, ok := r.Nodes[n]; ok {
		return an
	}
	switch n := n.(type) {
	case *dst.ArrayType:
		out := &ast.ArrayType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Lbrack
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))

		// Decoration: AfterLbrack
		r.applyDecorations(n.Decs.AfterLbrack)

		// Node: Len
		if n.Len != nil {
			out.Len = r.restoreNode(n.Len).(ast.Expr)
		}

		// Token: Rbrack
		r.cursor += token.Pos(len(token.RBRACK.String()))

		// Decoration: AfterLen
		r.applyDecorations(n.Decs.AfterLen)

		// Node: Elt
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.AssignStmt:
		out := &ast.AssignStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// List: Lhs
		for _, v := range n.Lhs {
			out.Lhs = append(out.Lhs, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: AfterLhs
		r.applyDecorations(n.Decs.AfterLhs)

		// Token: Tok
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))

		// Decoration: AfterTok
		r.applyDecorations(n.Decs.AfterTok)

		// List: Rhs
		for _, v := range n.Rhs {
			out.Rhs = append(out.Rhs, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.BadDecl:
		out := &ast.BadDecl{}
		r.Nodes[n] = out

		return out
	case *dst.BadExpr:
		out := &ast.BadExpr{}
		r.Nodes[n] = out

		return out
	case *dst.BadStmt:
		out := &ast.BadStmt{}
		r.Nodes[n] = out

		return out
	case *dst.BasicLit:
		out := &ast.BasicLit{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// String: Value
		out.ValuePos = r.cursor
		out.Value = n.Value
		r.cursor += token.Pos(len(n.Value))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Kind
		out.Kind = n.Kind

		return out
	case *dst.BinaryExpr:
		out := &ast.BinaryExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Op
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))

		// Decoration: AfterOp
		r.applyDecorations(n.Decs.AfterOp)

		// Node: Y
		if n.Y != nil {
			out.Y = r.restoreNode(n.Y).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.BlockStmt:
		out := &ast.BlockStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Lbrace
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))

		// Decoration: AfterLbrace
		r.applyDecorations(n.Decs.AfterLbrace)

		// List: List
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v).(ast.Stmt))
		}

		// Token: Rbrace
		out.Rbrace = r.cursor
		r.cursor += token.Pos(len(token.RBRACE.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.BranchStmt:
		out := &ast.BranchStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Tok
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))

		// Decoration: AfterTok
		r.applyDecorations(n.Decs.AfterTok)

		// Node: Label
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label).(*ast.Ident)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.CallExpr:
		out := &ast.CallExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Fun
		if n.Fun != nil {
			out.Fun = r.restoreNode(n.Fun).(ast.Expr)
		}

		// Decoration: AfterFun
		r.applyDecorations(n.Decs.AfterFun)

		// Token: Lparen
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))

		// Decoration: AfterLparen
		r.applyDecorations(n.Decs.AfterLparen)

		// List: Args
		for _, v := range n.Args {
			out.Args = append(out.Args, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: AfterArgs
		r.applyDecorations(n.Decs.AfterArgs)

		// Token: Ellipsis
		if n.Ellipsis {
			out.Ellipsis = r.cursor
			r.cursor += token.Pos(len(token.ELLIPSIS.String()))
		}

		// Decoration: AfterEllipsis
		r.applyDecorations(n.Decs.AfterEllipsis)

		// Token: Rparen
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.CaseClause:
		out := &ast.CaseClause{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Case
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.List == nil {
				return token.DEFAULT
			} else {
				return token.CASE
			}
		}().String()))

		// Decoration: AfterCase
		r.applyDecorations(n.Decs.AfterCase)

		// List: List
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: AfterList
		r.applyDecorations(n.Decs.AfterList)

		// Token: Colon
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))

		// Decoration: AfterColon
		r.applyDecorations(n.Decs.AfterColon)

		// List: Body
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v).(ast.Stmt))
		}

		return out
	case *dst.ChanType:
		out := &ast.ChanType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Begin
		out.Begin = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Dir == dst.RECV {
				return token.ARROW
			} else {
				return token.CHAN
			}
		}().String()))

		// Token: Chan
		if n.Dir == dst.RECV {
			r.cursor += token.Pos(len(token.CHAN.String()))
		}

		// Decoration: AfterBegin
		r.applyDecorations(n.Decs.AfterBegin)

		// Token: Arrow
		if n.Dir == dst.SEND {
			out.Arrow = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}

		// Decoration: AfterArrow
		r.applyDecorations(n.Decs.AfterArrow)

		// Node: Value
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Dir
		out.Dir = ast.ChanDir(n.Dir)

		return out
	case *dst.CommClause:
		out := &ast.CommClause{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Case
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Comm == nil {
				return token.DEFAULT
			} else {
				return token.CASE
			}
		}().String()))

		// Decoration: AfterCase
		r.applyDecorations(n.Decs.AfterCase)

		// Node: Comm
		if n.Comm != nil {
			out.Comm = r.restoreNode(n.Comm).(ast.Stmt)
		}

		// Decoration: AfterComm
		r.applyDecorations(n.Decs.AfterComm)

		// Token: Colon
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))

		// Decoration: AfterColon
		r.applyDecorations(n.Decs.AfterColon)

		// List: Body
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v).(ast.Stmt))
		}

		return out
	case *dst.CompositeLit:
		out := &ast.CompositeLit{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(ast.Expr)
		}

		// Decoration: AfterType
		r.applyDecorations(n.Decs.AfterType)

		// Token: Lbrace
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))

		// Decoration: AfterLbrace
		r.applyDecorations(n.Decs.AfterLbrace)

		// List: Elts
		for _, v := range n.Elts {
			out.Elts = append(out.Elts, r.restoreNode(v).(ast.Expr))
		}

		// Token: Rbrace
		out.Rbrace = r.cursor
		r.cursor += token.Pos(len(token.RBRACE.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Incomplete
		out.Incomplete = n.Incomplete

		return out
	case *dst.DeclStmt:
		out := &ast.DeclStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Decl
		if n.Decl != nil {
			out.Decl = r.restoreNode(n.Decl).(ast.Decl)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.DeferStmt:
		out := &ast.DeferStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Defer
		out.Defer = r.cursor
		r.cursor += token.Pos(len(token.DEFER.String()))

		// Decoration: AfterDefer
		r.applyDecorations(n.Decs.AfterDefer)

		// Node: Call
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call).(*ast.CallExpr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.Ellipsis:
		out := &ast.Ellipsis{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Ellipsis
		out.Ellipsis = r.cursor
		r.cursor += token.Pos(len(token.ELLIPSIS.String()))

		// Decoration: AfterEllipsis
		r.applyDecorations(n.Decs.AfterEllipsis)

		// Node: Elt
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.EmptyStmt:
		out := &ast.EmptyStmt{}
		r.Nodes[n] = out

		// Token: Semicolon
		if !n.Implicit {
			out.Semicolon = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}

		// Value: Implicit
		out.Implicit = n.Implicit

		return out
	case *dst.ExprStmt:
		out := &ast.ExprStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.Field:
		out := &ast.Field{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// List: Names
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v).(*ast.Ident))
		}

		// Decoration: AfterNames
		r.applyDecorations(n.Decs.AfterNames)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(ast.Expr)
		}

		// Decoration: AfterType
		r.applyDecorations(n.Decs.AfterType)

		// Node: Tag
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag).(*ast.BasicLit)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.FieldList:
		out := &ast.FieldList{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Opening
		if n.Opening {
			out.Opening = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}

		// Decoration: AfterOpening
		r.applyDecorations(n.Decs.AfterOpening)

		// List: List
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v).(*ast.Field))
		}

		// Token: Closing
		if n.Closing {
			out.Closing = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.File:
		out := &ast.File{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Package
		out.Package = r.cursor
		r.cursor += token.Pos(len(token.PACKAGE.String()))

		// Decoration: AfterPackage
		r.applyDecorations(n.Decs.AfterPackage)

		// Node: Name
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name).(*ast.Ident)
		}

		// Decoration: AfterName
		r.applyDecorations(n.Decs.AfterName)

		// List: Decls
		for _, v := range n.Decls {
			out.Decls = append(out.Decls, r.restoreNode(v).(ast.Decl))
		}

		// Scope: Scope
		out.Scope = r.restoreScope(n.Scope)

		return out
	case *dst.ForStmt:
		out := &ast.ForStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: For
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))

		// Decoration: AfterFor
		r.applyDecorations(n.Decs.AfterFor)

		// Node: Init
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init).(ast.Stmt)
		}

		// Token: InitSemicolon
		if n.Init != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}

		// Decoration: AfterInit
		r.applyDecorations(n.Decs.AfterInit)

		// Node: Cond
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond).(ast.Expr)
		}

		// Token: CondSemicolon
		if n.Post != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}

		// Decoration: AfterCond
		r.applyDecorations(n.Decs.AfterCond)

		// Node: Post
		if n.Post != nil {
			out.Post = r.restoreNode(n.Post).(ast.Stmt)
		}

		// Decoration: AfterPost
		r.applyDecorations(n.Decs.AfterPost)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.FuncDecl:
		out := &ast.FuncDecl{}
		r.Nodes[n] = out

		// Init: Type
		out.Type = &ast.FuncType{}

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Func
		if true {
			out.Type.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}

		// Decoration: AfterFunc
		r.applyDecorations(n.Decs.AfterFunc)

		// Node: Recv
		if n.Recv != nil {
			out.Recv = r.restoreNode(n.Recv).(*ast.FieldList)
		}

		// Decoration: AfterRecv
		r.applyDecorations(n.Decs.AfterRecv)

		// Node: Name
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name).(*ast.Ident)
		}

		// Decoration: AfterName
		r.applyDecorations(n.Decs.AfterName)

		// Node: Params
		if n.Type.Params != nil {
			out.Type.Params = r.restoreNode(n.Type.Params).(*ast.FieldList)
		}

		// Decoration: AfterParams
		r.applyDecorations(n.Decs.AfterParams)

		// Node: Results
		if n.Type.Results != nil {
			out.Type.Results = r.restoreNode(n.Type.Results).(*ast.FieldList)
		}

		// Decoration: AfterResults
		r.applyDecorations(n.Decs.AfterResults)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.FuncLit:
		out := &ast.FuncLit{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(*ast.FuncType)
		}

		// Decoration: AfterType
		r.applyDecorations(n.Decs.AfterType)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.FuncType:
		out := &ast.FuncType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Func
		if n.Func {
			out.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}

		// Decoration: AfterFunc
		r.applyDecorations(n.Decs.AfterFunc)

		// Node: Params
		if n.Params != nil {
			out.Params = r.restoreNode(n.Params).(*ast.FieldList)
		}

		// Decoration: AfterParams
		r.applyDecorations(n.Decs.AfterParams)

		// Node: Results
		if n.Results != nil {
			out.Results = r.restoreNode(n.Results).(*ast.FieldList)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.GenDecl:
		out := &ast.GenDecl{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Tok
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))

		// Decoration: AfterTok
		r.applyDecorations(n.Decs.AfterTok)

		// Token: Lparen
		if n.Lparen {
			out.Lparen = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}

		// Decoration: AfterLparen
		r.applyDecorations(n.Decs.AfterLparen)

		// List: Specs
		for _, v := range n.Specs {
			out.Specs = append(out.Specs, r.restoreNode(v).(ast.Spec))
		}

		// Token: Rparen
		if n.Rparen {
			out.Rparen = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.GoStmt:
		out := &ast.GoStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Go
		out.Go = r.cursor
		r.cursor += token.Pos(len(token.GO.String()))

		// Decoration: AfterGo
		r.applyDecorations(n.Decs.AfterGo)

		// Node: Call
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call).(*ast.CallExpr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.Ident:
		out := &ast.Ident{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// String: Name
		out.NamePos = r.cursor
		out.Name = n.Name
		r.cursor += token.Pos(len(n.Name))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Object: Obj
		out.Obj = r.restoreObject(n.Obj)

		return out
	case *dst.IfStmt:
		out := &ast.IfStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: If
		out.If = r.cursor
		r.cursor += token.Pos(len(token.IF.String()))

		// Decoration: AfterIf
		r.applyDecorations(n.Decs.AfterIf)

		// Node: Init
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init).(ast.Stmt)
		}

		// Decoration: AfterInit
		r.applyDecorations(n.Decs.AfterInit)

		// Node: Cond
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond).(ast.Expr)
		}

		// Decoration: AfterCond
		r.applyDecorations(n.Decs.AfterCond)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Token: ElseTok
		if n.Else != nil {
			r.cursor += token.Pos(len(token.ELSE.String()))
		}

		// Decoration: AfterElse
		r.applyDecorations(n.Decs.AfterElse)

		// Node: Else
		if n.Else != nil {
			out.Else = r.restoreNode(n.Else).(ast.Stmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.ImportSpec:
		out := &ast.ImportSpec{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Name
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name).(*ast.Ident)
		}

		// Decoration: AfterName
		r.applyDecorations(n.Decs.AfterName)

		// Node: Path
		if n.Path != nil {
			out.Path = r.restoreNode(n.Path).(*ast.BasicLit)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.IncDecStmt:
		out := &ast.IncDecStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Tok
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.IndexExpr:
		out := &ast.IndexExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Lbrack
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))

		// Decoration: AfterLbrack
		r.applyDecorations(n.Decs.AfterLbrack)

		// Node: Index
		if n.Index != nil {
			out.Index = r.restoreNode(n.Index).(ast.Expr)
		}

		// Decoration: AfterIndex
		r.applyDecorations(n.Decs.AfterIndex)

		// Token: Rbrack
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.InterfaceType:
		out := &ast.InterfaceType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Interface
		out.Interface = r.cursor
		r.cursor += token.Pos(len(token.INTERFACE.String()))

		// Decoration: AfterInterface
		r.applyDecorations(n.Decs.AfterInterface)

		// Node: Methods
		if n.Methods != nil {
			out.Methods = r.restoreNode(n.Methods).(*ast.FieldList)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Incomplete
		out.Incomplete = n.Incomplete

		return out
	case *dst.KeyValueExpr:
		out := &ast.KeyValueExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Key
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key).(ast.Expr)
		}

		// Decoration: AfterKey
		r.applyDecorations(n.Decs.AfterKey)

		// Token: Colon
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))

		// Decoration: AfterColon
		r.applyDecorations(n.Decs.AfterColon)

		// Node: Value
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.LabeledStmt:
		out := &ast.LabeledStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Label
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label).(*ast.Ident)
		}

		// Decoration: AfterLabel
		r.applyDecorations(n.Decs.AfterLabel)

		// Token: Colon
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))

		// Decoration: AfterColon
		r.applyDecorations(n.Decs.AfterColon)

		// Node: Stmt
		if n.Stmt != nil {
			out.Stmt = r.restoreNode(n.Stmt).(ast.Stmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.MapType:
		out := &ast.MapType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Map
		out.Map = r.cursor
		r.cursor += token.Pos(len(token.MAP.String()))

		// Token: Lbrack
		r.cursor += token.Pos(len(token.LBRACK.String()))

		// Decoration: AfterMap
		r.applyDecorations(n.Decs.AfterMap)

		// Node: Key
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key).(ast.Expr)
		}

		// Token: Rbrack
		r.cursor += token.Pos(len(token.RBRACK.String()))

		// Decoration: AfterKey
		r.applyDecorations(n.Decs.AfterKey)

		// Node: Value
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.Package:
		out := &ast.Package{}
		r.Nodes[n] = out

		// Value: Name
		out.Name = n.Name

		// Scope: Scope
		out.Scope = r.restoreScope(n.Scope)

		// Map: Imports
		out.Imports = map[string]*ast.Object{}
		for k, v := range n.Imports {
			out.Imports[k] = r.restoreObject(v)
		}

		// Map: Files
		out.Files = map[string]*ast.File{}
		for k, v := range n.Files {
			out.Files[k] = r.restoreNode(v).(*ast.File)
		}

		return out
	case *dst.ParenExpr:
		out := &ast.ParenExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Lparen
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))

		// Decoration: AfterLparen
		r.applyDecorations(n.Decs.AfterLparen)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Rparen
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.RangeStmt:
		out := &ast.RangeStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: For
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))

		// Decoration: AfterFor
		r.applyDecorations(n.Decs.AfterFor)

		// Node: Key
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key).(ast.Expr)
		}

		// Token: Comma
		if n.Value != nil {
			r.cursor += token.Pos(len(token.COMMA.String()))
		}

		// Decoration: AfterKey
		r.applyDecorations(n.Decs.AfterKey)

		// Node: Value
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value).(ast.Expr)
		}

		// Decoration: AfterValue
		r.applyDecorations(n.Decs.AfterValue)

		// Token: Tok
		if n.Tok != token.ILLEGAL {
			out.Tok = n.Tok
			out.TokPos = r.cursor
			r.cursor += token.Pos(len(n.Tok.String()))
		}

		// Token: Range
		r.cursor += token.Pos(len(token.RANGE.String()))

		// Decoration: AfterRange
		r.applyDecorations(n.Decs.AfterRange)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.ReturnStmt:
		out := &ast.ReturnStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Return
		out.Return = r.cursor
		r.cursor += token.Pos(len(token.RETURN.String()))

		// Decoration: AfterReturn
		r.applyDecorations(n.Decs.AfterReturn)

		// List: Results
		for _, v := range n.Results {
			out.Results = append(out.Results, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.SelectStmt:
		out := &ast.SelectStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Select
		out.Select = r.cursor
		r.cursor += token.Pos(len(token.SELECT.String()))

		// Decoration: AfterSelect
		r.applyDecorations(n.Decs.AfterSelect)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.SelectorExpr:
		out := &ast.SelectorExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Token: Period
		r.cursor += token.Pos(len(token.PERIOD.String()))

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Node: Sel
		if n.Sel != nil {
			out.Sel = r.restoreNode(n.Sel).(*ast.Ident)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.SendStmt:
		out := &ast.SendStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Chan
		if n.Chan != nil {
			out.Chan = r.restoreNode(n.Chan).(ast.Expr)
		}

		// Decoration: AfterChan
		r.applyDecorations(n.Decs.AfterChan)

		// Token: Arrow
		out.Arrow = r.cursor
		r.cursor += token.Pos(len(token.ARROW.String()))

		// Decoration: AfterArrow
		r.applyDecorations(n.Decs.AfterArrow)

		// Node: Value
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.SliceExpr:
		out := &ast.SliceExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Lbrack
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))

		// Decoration: AfterLbrack
		r.applyDecorations(n.Decs.AfterLbrack)

		// Node: Low
		if n.Low != nil {
			out.Low = r.restoreNode(n.Low).(ast.Expr)
		}

		// Token: Colon1
		r.cursor += token.Pos(len(token.COLON.String()))

		// Decoration: AfterLow
		r.applyDecorations(n.Decs.AfterLow)

		// Node: High
		if n.High != nil {
			out.High = r.restoreNode(n.High).(ast.Expr)
		}

		// Token: Colon2
		if n.Slice3 {
			r.cursor += token.Pos(len(token.COLON.String()))
		}

		// Decoration: AfterHigh
		r.applyDecorations(n.Decs.AfterHigh)

		// Node: Max
		if n.Max != nil {
			out.Max = r.restoreNode(n.Max).(ast.Expr)
		}

		// Decoration: AfterMax
		r.applyDecorations(n.Decs.AfterMax)

		// Token: Rbrack
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Slice3
		out.Slice3 = n.Slice3

		return out
	case *dst.StarExpr:
		out := &ast.StarExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Star
		out.Star = r.cursor
		r.cursor += token.Pos(len(token.MUL.String()))

		// Decoration: AfterStar
		r.applyDecorations(n.Decs.AfterStar)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.StructType:
		out := &ast.StructType{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Struct
		out.Struct = r.cursor
		r.cursor += token.Pos(len(token.STRUCT.String()))

		// Decoration: AfterStruct
		r.applyDecorations(n.Decs.AfterStruct)

		// Node: Fields
		if n.Fields != nil {
			out.Fields = r.restoreNode(n.Fields).(*ast.FieldList)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		// Value: Incomplete
		out.Incomplete = n.Incomplete

		return out
	case *dst.SwitchStmt:
		out := &ast.SwitchStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Switch
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))

		// Decoration: AfterSwitch
		r.applyDecorations(n.Decs.AfterSwitch)

		// Node: Init
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init).(ast.Stmt)
		}

		// Decoration: AfterInit
		r.applyDecorations(n.Decs.AfterInit)

		// Node: Tag
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag).(ast.Expr)
		}

		// Decoration: AfterTag
		r.applyDecorations(n.Decs.AfterTag)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.TypeAssertExpr:
		out := &ast.TypeAssertExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Token: Period
		r.cursor += token.Pos(len(token.PERIOD.String()))

		// Decoration: AfterX
		r.applyDecorations(n.Decs.AfterX)

		// Token: Lparen
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))

		// Decoration: AfterLparen
		r.applyDecorations(n.Decs.AfterLparen)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(ast.Expr)
		}

		// Token: TypeToken
		if n.Type == nil {
			r.cursor += token.Pos(len(token.TYPE.String()))
		}

		// Decoration: AfterType
		r.applyDecorations(n.Decs.AfterType)

		// Token: Rparen
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.TypeSpec:
		out := &ast.TypeSpec{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Node: Name
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name).(*ast.Ident)
		}

		// Token: Assign
		if n.Assign {
			out.Assign = r.cursor
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}

		// Decoration: AfterName
		r.applyDecorations(n.Decs.AfterName)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.TypeSwitchStmt:
		out := &ast.TypeSwitchStmt{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Switch
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))

		// Decoration: AfterSwitch
		r.applyDecorations(n.Decs.AfterSwitch)

		// Node: Init
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init).(ast.Stmt)
		}

		// Decoration: AfterInit
		r.applyDecorations(n.Decs.AfterInit)

		// Node: Assign
		if n.Assign != nil {
			out.Assign = r.restoreNode(n.Assign).(ast.Stmt)
		}

		// Decoration: AfterAssign
		r.applyDecorations(n.Decs.AfterAssign)

		// Node: Body
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body).(*ast.BlockStmt)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.UnaryExpr:
		out := &ast.UnaryExpr{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// Token: Op
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))

		// Decoration: AfterOp
		r.applyDecorations(n.Decs.AfterOp)

		// Node: X
		if n.X != nil {
			out.X = r.restoreNode(n.X).(ast.Expr)
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	case *dst.ValueSpec:
		out := &ast.ValueSpec{}
		r.Nodes[n] = out

		// Decoration: Start
		r.applyDecorations(n.Decs.Start)

		// List: Names
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v).(*ast.Ident))
		}

		// Decoration: AfterNames
		r.applyDecorations(n.Decs.AfterNames)

		// Node: Type
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type).(ast.Expr)
		}

		// Token: Assign
		if n.Values != nil {
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}

		// Decoration: AfterAssign
		r.applyDecorations(n.Decs.AfterAssign)

		// List: Values
		for _, v := range n.Values {
			out.Values = append(out.Values, r.restoreNode(v).(ast.Expr))
		}

		// Decoration: End
		r.applyDecorations(n.Decs.End)

		return out
	default:
		panic(fmt.Sprintf("%T", n))
	}
	return nil
}
