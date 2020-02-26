package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/tools"
)

func main() {
	tools.RunWithMacOSApp()
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()
}
