package main

import (
	"github.com/recative/recative-backend/domain"
	"github.com/recative/recative-backend/pkg"
)

func main() {
	dep := pkg.AutoInit(nil)

	domain.Init(dep)
}
