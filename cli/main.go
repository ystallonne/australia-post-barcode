package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/b-b3rn4rd/4-state-barcode/src/australiapost"
	"github.com/spf13/cobra"
)

var (
	// Version application version, set using ldflags
	Version = "undefined"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	var barcode string
	var text string
	var filename string

	var rootCmd = &cobra.Command{
		Use:     "4-state-barcode",
		Short:   "Australia Post 4 state barcode generator",
		Version: Version,
		Long: `Australia Post 4 state barcode generator
Generates a SVG image containing barcode with an optional additional text
Example: 
4-state-barcode -b "6256439111HELLO" -f barcode.svg`,
		Run: func(cmd *cobra.Command, args []string) {
			file, _ := os.Create(filename)
			generator := australiapost.NewFourStateBarcode(barcode, file, text, australiapost.OptionLogger(logger))

			err := generator.Generate()
			if err != nil {
				logger.Fatal(err)
			}

		},
	}

	rootCmd.Flags().StringVarP(&barcode, "barcode", "b", "", "Barcode value")
	rootCmd.Flags().StringVarP(&text, "text", "t", "", "Optional barcode text")
	rootCmd.Flags().StringVarP(&filename, "filename", "f", "", "Output filename")
	rootCmd.MarkFlagRequired("barcode")
	rootCmd.MarkFlagRequired("filename")
	rootCmd.Execute()
}
