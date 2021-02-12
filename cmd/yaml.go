/*
Copyright 2019 Google LLC
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
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type config struct {
	openshift map[string][]string
}

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yaml called")
		fmt.Println(args[0])
		fmt.Println(args[1])
		yamlConvert(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
	yamlCmd.Flags().BoolP("input", "i", false, "Path to the input file to covert, must be in Openshift format")
	yamlCmd.Flags().BoolP("output", "o", false, "Path to the output file for the results on the conversion")
}

func yamlConvert(input string, output string) {
	filename, _ := filepath.Abs(input)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", config.openshift)

}