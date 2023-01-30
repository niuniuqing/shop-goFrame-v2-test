package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"shop-goFrame-v2-test/internal/cmd"
	_ "shop-goFrame-v2-test/internal/logic"
	_ "shop-goFrame-v2-test/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
