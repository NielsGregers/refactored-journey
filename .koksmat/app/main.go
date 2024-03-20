package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/koksmat-test1/magicapp"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: koksmat-test1
description: Describe the main purpose of this kitchen
---

# koksmat-test1
`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("koksmat-test1", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.Execute(name, "koksmat-test1", "")
}
