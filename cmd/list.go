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
	"sort"
	"text/tabwriter"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneOpt bool
var allOpt bool
var orderBy string
var reverse bool
var head int
var tail int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Long:  `To list all the taks run ./tri list`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := task.ReadItems(viper.GetString("datafile"))
		if orderBy == "deadline" {
			sort.Sort(task.ByDeadline(items))
		} else if orderBy == "priority" {
			sort.Sort(task.ByPri(items))
		}
		if reverse {
			items = task.Reverse(items)
		}

		if head != -1 {
			items = items[:head]
		} else if tail != -1 {
			items = items[len(items)-tail:]
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
		for _, item := range items {
			if allOpt || item.Done == doneOpt {
				fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.GetDeadline()+"\t"+item.PrettyDone())
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "show all todos")
	listCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse ")
	listCmd.Flags().StringVarP(&orderBy, "order", "o", "priority", "order by a specific field")
	listCmd.Flags().IntVar(&head, "head", -1, "list head of the tasks")
	listCmd.Flags().IntVar(&tail, "tail", -1, "list tail of the tasks")

}
