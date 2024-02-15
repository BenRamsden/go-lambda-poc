package main

import "github.com/jugo-io/go-poc/internal/api/model"

func main() {
	auth := model.NewAuthService()

	if err := auth.Enforce("tom", "data1", model.PermissionRead); err != nil {
		panic(err)
	}

	if !auth.Can("tom", "data1", model.PermissionRead) {
		panic("tom should have permission")
	}

	if err := auth.Revoke("tom", "data1", model.PermissionRead); err != nil {
		panic(err)
	}

	if auth.Can("tom", "data1", model.PermissionRead) {
		panic("tom should not have permission")
	}
}
