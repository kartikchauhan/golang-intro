package main

import "fmt"

const (
	isAdmin = 1 << iota
	isManager
	isDeveloper

	canSeeCanada
	canSeeNoida
	canSeeHyderabad // camelCasing for internal constants
	CanSeeGorakhpur // PascalCasing for external constants
)

func roles() {
	var allRoles byte = isManager | canSeeCanada | canSeeHyderabad

	fmt.Printf("%b\n", allRoles)

	fmt.Printf("isManager? %v\n", allRoles&isAdmin == isManager)

	fmt.Printf("isManager? %v\n", allRoles&isDeveloper == isManager)

	fmt.Printf("isManager? %v\n", allRoles&isManager == isManager)

}
