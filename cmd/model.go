package cmd

import (
	"errors"

	"github.com/lizhennet/gorm-gen/pkg/core"
	"github.com/lizhennet/gorm-gen/pkg/generator"
	"github.com/lizhennet/gorm-gen/pkg/log"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "generator model file",
	Long: `generator model file,
For example:
	generator model file:	gorm-gen model table1
	generator model file with dal file:	gorm-gen model table1 --dal
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires a table name argument")
		}
		if len(args) > 1 {
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
		withDal := cmd.Flags().Changed("dal")
		table := args[0]
		ctx, err := generator.NewGenCtx(config, table)
		if err != nil {
			log.Error("generate context error,err:%s", err.Error())
			return nil
		}
		err = generator.GenModelFile(ctx, withDal)
		if err != nil {
			log.Error("generate file error,err:%s", err.Error())
			return nil
		}
		return err
	},
}

func init() {
	modelCmd.PersistentFlags().Bool("dal", false, "generator dal file")
	modelCmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./config/gorm-gen.yml)")
	rootCmd.AddCommand(modelCmd)
}
