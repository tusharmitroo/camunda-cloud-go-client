/*
Copyright © 2021 Matheus Cruz <matheuscruz.dev@gmail.com>

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
)

var (
	zeebeClientDeleteExample = ""
	zeebeClientGetExample    = ""
	zeebeClientCreateExample = ""
)

// zbClientCmd represents the zb-client command
var zbClientCmd = &cobra.Command{
	Use:   "zb-client [options]",
	Short: "Manage your zeebe clients resources on Camunda Cloud",
	Long: `Used together [OPTIONS] like get, create, delete for manage your zeebe clients resources on Camunda Cloud. For example:` +
		zeebeClientGetExample + zeebeClientCreateExample + zeebeClientDeleteExample,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zb-client called")
	},
}

func init() {
	rootCmd.AddCommand(zbClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zbClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zbClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
