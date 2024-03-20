package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/365admin/koksmat-test1/cmds"
	"github.com/365admin/koksmat-test1/utils"
)

func RegisterCmds() {
	RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  `Describe the main purpose of this kitchen`,
	}
	HealthPingPostCmd := &cobra.Command{
		Use:   "ping  pong",
		Short: "Ping",
		Long:  `Simple ping endpoint`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthPingPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthPingPostCmd)
	HealthCoreversionPostCmd := &cobra.Command{
		Use:   "coreversion ",
		Short: "Core Version",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthCoreversionPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthCoreversionPostCmd)

	RootCmd.AddCommand(healthCmd)
	testCmd := &cobra.Command{
		Use:   "test",
		Short: "Test",
		Long:  `Describe the main purpose of this kitchen`,
	}
	TestBasicPostCmd := &cobra.Command{
		Use:   "basic ",
		Short: "Basic Test",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TestBasicPost(ctx, args)
		},
	}
	testCmd.AddCommand(TestBasicPostCmd)
	TestSharepointPostCmd := &cobra.Command{
		Use:   "sharepoint ",
		Short: "SharePoint Test",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TestSharepointPost(ctx, args)
		},
	}
	testCmd.AddCommand(TestSharepointPostCmd)
	TestExchangePostCmd := &cobra.Command{
		Use:   "exchange ",
		Short: "Exchange Test",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TestExchangePost(ctx, args)
		},
	}
	testCmd.AddCommand(TestExchangePostCmd)

	RootCmd.AddCommand(testCmd)
}
