package main

import (
	_ "hellogf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hellogf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
