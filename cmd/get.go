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
	"strconv"
	"text/tabwriter"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display a task given id",
	Long: `To display a taks with an  id run ./todo get ID
To display multiple tasks run ./todo get id1 id2 id3 ...`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := task.ReadItems(viper.GetString("datafile"))
		w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)

		for _, arg := range args {
			index, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(arg, " is not a valid label")
			}
			if index > 0 && index <= len(items) {
				item := items[index-1]
				fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.GetDeadline()+"\t"+item.PrettyDone())
			} else {
				fmt.Println(arg, " is not a valid label")
			}

		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
