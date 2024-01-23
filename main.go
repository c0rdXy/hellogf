package main

import (
	_ "hellogf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"hellogf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
