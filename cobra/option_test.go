package cobra_test

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestExtendCommand(t *testing.T) {
	t.Run("after run", func(t *testing.T) {
		var command cobra.Command

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			After(&command.Run, func(cmd *cobra.Command, args []string) {
				cmd.Use = "first"
			})
		})
		assert.Empty(t, command.Use)
		command.Run(&command, nil)
		assert.Equal(t, command.Use, "first")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			After(&command.Run, func(cmd *cobra.Command, args []string) {
				cmd.Use = "last"
			})
		})
		assert.Equal(t, command.Use, "first")
		command.Run(&command, nil)
		assert.Equal(t, command.Use, "last")
	})

	t.Run("after run with error", func(t *testing.T) {
		var command cobra.Command

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			AfterE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "first"
				return nil
			})
		})
		assert.Empty(t, command.Use)
		assert.NoError(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "first")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			AfterE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "last"
				return errors.New("test")
			})
		})
		assert.Equal(t, command.Use, "first")
		assert.Error(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "last")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			AfterE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "first"
				return errors.New("test")
			})
		})
		assert.Equal(t, command.Use, "last")
		assert.Error(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "last")
	})

	t.Run("before run", func(t *testing.T) {
		var command cobra.Command

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			Before(&command.Run, func(cmd *cobra.Command, args []string) {
				cmd.Use = "first"
			})
		})
		assert.Empty(t, command.Use)
		command.Run(&command, nil)
		assert.Equal(t, command.Use, "first")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			Before(&command.Run, func(cmd *cobra.Command, args []string) {
				cmd.Use = "last"
			})
		})
		assert.Equal(t, command.Use, "first")
		command.Run(&command, nil)
		assert.Equal(t, command.Use, "first")
	})

	t.Run("before run with error", func(t *testing.T) {
		var command cobra.Command

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			BeforeE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "first"
				return nil
			})
		})
		assert.Empty(t, command.Use)
		assert.NoError(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "first")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			BeforeE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "last"
				return errors.New("test")
			})
		})
		assert.Equal(t, command.Use, "first")
		assert.Error(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "last")

		Apply(&command, viper.New(), func(command *cobra.Command, container *viper.Viper) {
			BeforeE(&command.RunE, func(cmd *cobra.Command, args []string) error {
				cmd.Use = "first"
				return errors.New("test")
			})
		})
		assert.Equal(t, command.Use, "last")
		assert.Error(t, command.RunE(&command, nil))
		assert.Equal(t, command.Use, "first")
	})
}
