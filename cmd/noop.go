/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	ezImage "github.com/bangau1/ezimage/pkg/image"
)

type NoopFlags struct{
	Source string
	Destination string
	JpegQuality int
}

var noopFlags NoopFlags

// noopCmd represents the noop command
var noopCmd = &cobra.Command{
	Use:   "noop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		img, err := ezImage.NewImageFromLocation(noopFlags.Source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var t ezImage.Transformation
		t = ezImage.NewNoOpTransformation()
		result := t.Apply(img)
		if result.Error != nil {
			fmt.Println(result.Error)
			os.Exit(1)
		}
		mimeType := ezImage.MIME_TYPE_JPEG
		if strings.HasSuffix(strings.ToLower(noopFlags.Destination), ".png"){
			mimeType = ezImage.MIME_TYPE_PNG
		}
		if err := result.Data.Save(noopFlags.Destination, mimeType, noopFlags.JpegQuality); err != nil{
			fmt.Println(err)
			os.Exit(1)
		}else{
			fmt.Printf("Succesfully saved to %s\n", noopFlags.Destination)
		}
	},
}

func init() {
	rootCmd.AddCommand(noopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noopCmd.PersistentFlags().String("foo", "", "A help for foo")
	noopCmd.PersistentFlags().StringVarP(&noopFlags.Source, "src", "s", "", "Image source")
	noopCmd.PersistentFlags().StringVarP(&noopFlags.Destination, "dest", "d", "", "Image destination")
	noopCmd.PersistentFlags().IntVarP(&noopFlags.JpegQuality, "quality", "q", 75, "JPEG Image Quality. This will be ignored if the destination used is non-jpeg format")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
