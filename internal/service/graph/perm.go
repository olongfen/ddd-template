package graph

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/olongfen/toolkit/scontext"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Perm struct {
	Description string
	Tag         string
	Domain      string
	Action      string
}

type Perms []Perm

func (p Perms) Get(domain string, action string) (Perm, error) {
	for _, v := range p {
		if v.Domain == domain && v.Action == action {
			return v, nil
		}
	}
	return Perm{}, errors.New("no data")
}

var P Perms = []Perm{
	{Description: "用户管理", Tag: "添加用户", Domain: "user", Action: "add"},
	{Description: "用户管理", Tag: "编辑用户", Domain: "user", Action: "edit"},
	{Description: "用户管理", Tag: "删除用户", Domain: "user", Action: "delete"},
	{Description: "用户管理", Tag: "查询用户", Domain: "user", Action: "list"},
}

func CheckPerm() func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		const userContextKey = "__local_user_context__"
		acc := graphql.GetOperationContext(ctx)
		for _, v := range acc.Operation.SelectionSet {

			s := v.(*ast.Field)
			for _, _v := range s.SelectionSet {
				_s := _v.(*ast.Field)
				// todo 通过rabc验证
				if _, err = P.Get(s.Name, _s.Name); err != nil {
					return nil, err
				}
			}
		}
		userUuid := scontext.GetUserUuid(ctx.Value(userContextKey).(context.Context))
		if userUuid == "" {
			return nil, &gqlerror.Error{
				Message: "Access Denied",
			}
		}
		ctx = scontext.SetUserUuid(ctx, userUuid)
		return next(ctx)
	}
}
