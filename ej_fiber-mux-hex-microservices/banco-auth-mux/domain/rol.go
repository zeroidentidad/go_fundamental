package domain

import (
	"strings"
)

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (p RolePermissions) IsAuthorizedFor(rol string, routeName string) bool {
	perms := p.rolePermissions[rol]
	for _, r := range perms {
		if r == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"getAllClientes", "getCliente", "postNewCuenta", "postNewTransaccion"},
		"user":  {"getCliente", "postNewTransaccion"},
	}}
}
