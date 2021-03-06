/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"text/tabwriter"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches a task by a string",
	Long:  `To search a taks run ./todo search 'task' `,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := task.ReadItems(viper.GetString("datafile"))
		w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
		for _, item := range items {

			if strings.Contains(strings.ToLower(item.Text), strings.ToLower(args[0])) {
				fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.GetDeadline()+"\t"+item.PrettyDone())
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
