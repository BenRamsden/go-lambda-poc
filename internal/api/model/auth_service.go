package model

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

const PolicyStr string = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`

type authService struct {
	enforcer *casbin.Enforcer
}

// Can implements AuthService.
func (svc *authService) Can(actor string, object string, permission Permission) bool {
	ok, err := svc.enforcer.Enforce(actor, object, string(permission))
	if err != nil {
		return false
	}
	return ok
}

// Enforce implements AuthService.
func (svc *authService) Enforce(actor string, object string, permission Permission) error {
	ok, err := svc.enforcer.AddPolicy(actor, object, string(permission))
	if err != nil {
		return err
	}
	if ok {
		// Policy was added, save it
		return svc.enforcer.SavePolicy()
	}
	return nil
}

// Revoke implements AuthService.
func (svc *authService) Revoke(actor string, object string, permission Permission) error {
	ok, err := svc.enforcer.RemovePolicy(actor, object, string(permission))
	if err != nil {
		return err
	}
	if ok {
		return svc.enforcer.SavePolicy()
	}
	return nil
}

func NewAuthService() AuthService {
	a := fileadapter.NewAdapter("policy.csv")

	m, err := casbinmodel.NewModelFromString(PolicyStr)
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}

	return &authService{
		enforcer: e,
	}
}
