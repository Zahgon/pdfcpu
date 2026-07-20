package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

var (
	conf             string
	force            bool
	kpw              string
	needStackTrace   bool
	offline          bool
	offlineSet       bool
	opw              string
	perm             string
	quiet            bool
	removeEncryption bool
	removeSignatures bool
	selectedPages    string
	unit             string
	upw              string
	verbose          int
)

var rootCmd = &cobra.Command{
	Use:   "pdfcpu",
	Short: "PDF tooling for Go and the command line",
	Long: `pdfcpu provides command-line tools for working with PDF files.
It is built on a Go API for direct PDF control.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error { _ = "STUB: not implemented"; return nil }

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&conf, "conf", "c", "", "set or disable config dir: $path | disable")
	rootCmd.PersistentFlags().BoolVar(&force, "force", false, "overwrite existing output files")
	rootCmd.PersistentFlags().BoolVarP(&offline, "offline", "o", false, "disable http traffic")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "disable output")
	rootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "Increase verbosity. Use -v or -vv.")
	rootCmd.AddCommand(commands()...)
}

func commands() []*cobra.Command { _ = "STUB: not implemented"; return nil }

func initConfig() { _ = "STUB: not implemented"; return }

func validateConfigDirFlag() error { _ = "STUB: not implemented"; return nil }

func ensureDefaultConfig() (*model.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getConfig() (*model.Configuration, error) { _ = "STUB: not implemented"; return nil, nil }
