package main

import (
	"fmt"

	"github.com/tiancai110a/gin-blog/pkg"
)

func main() {
	fmt.Println(pkg.RunMode, pkg.HTTPPort, pkg.ReadTimeout, pkg.WriteTimeout, pkg.JwtSecret, pkg.PageSize)
}
