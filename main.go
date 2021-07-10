package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
    e, err := casbin.NewEnforcer("./examples/rbac_model.conf","./examples/rbac_policy.csv")
    if err != nil {
        panic(err)
    }

    fmt.Println(e.Enforce("bob","data2","read"))

    fmt.Println(e.DeleteRolesForUser("bob"))

    fmt.Println(e.Enforce("bob","data2","read"))
}
