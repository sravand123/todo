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
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskText string
var Priority int
var dueOpt bool
var Deadline string

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "change a task",
	Long:  `Edit a task`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := task.ReadItems(viper.GetString("datafile"))

		index, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(args[0], " is not a valid label\n", err)
		}
		if index > 0 && index <= len(items) {
			if taskText == "" && Priority == -1 && !dueOpt && Deadline == "" {
				fmt.Println("Invalid usage of edit command")
			} else {

				if taskText != "" {

					items[index-1].Text = taskText
				}
				if Priority != -1 {
					items[index-1].SetPriority(Priority)
				}
				if Deadline != "" {
					items[index-1].SetDeadline(Deadline)
				}
				if dueOpt {
					items[index-1].Done = false
				}
				item := items[index-1]
				w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
				fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.GetDeadline()+"\t"+item.PrettyDone())
				w.Flush()
				task.SaveItems(viper.GetString("datafile"), items)
			}
		} else {
			log.Println(args[0], "doesn't match any items")
		}

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&taskText, "task", "t", "", "change task name")
	editCmd.Flags().IntVarP(&Priority, "priority", "p", -1, "change priority of tasks ")
	editCmd.Flags().BoolVar(&dueOpt, "due", false, "change status of task")
	editCmd.Flags().StringVar(&Deadline, "deadline", "", "change deadline of task ")

}
