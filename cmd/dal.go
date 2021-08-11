package cmd

import (
	"errors"

	"github.com/lizhennet/gorm-gen/pkg/core"
	"github.com/lizhennet/gorm-gen/pkg/generator"
	"github.com/lizhennet/gorm-gen/pkg/log"
	"github.com/spf13/cobra"
)

var dalCmd = &cobra.Command{
	Use:   "dal",
	Short: "generator dal file",
	Long: `generator  dal file,
For example:	
			gorm-gen dal table1
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires a table name argument")
		}
		if len(args) > 2 {
			return errors.New("too many arguments")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return errors.New("get config flag error,err:" + err.Error())
		}
		config = core.GetStringWithDefault(config, "./config/gorm-gen.yml")
		table := args[0]
		ctx, err := generator.NewGenCtx(config, table)
		if err != nil {
			log.Error("generate context error,err:%s", err.Error())
			return nil
		}
		err = generator.GenDalFile(ctx)
		if err != nil {
			log.Error("generate file error,err:%s", err.Error())
			return nil
		}
		return err
	},
}

func init() {
	dalCmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./config/gorm-gen.yml)")
	rootCmd.AddCommand(dalCmd)
}
