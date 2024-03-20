// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Ping
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-test1/execution"
	"github.com/365admin/koksmat-test1/utils"
)

func HealthPingPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-test1", "00-health", "10-ping.ps1", "", "-pong", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
