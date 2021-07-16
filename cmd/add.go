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
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int
var deadline string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your TODO list",
	Long:  `To add a task run ./todo add your-task `,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, _ := task.ReadItems(viper.GetString("datafile"))
	for _, arg := range args {
		item := task.Item{Text: arg}
		item.SetPriority(priority)
		item.SetDate()
		item.SetDeadline(deadline)
		items = append(items, item)
	}
	task.SaveItems(viper.GetString("datafile"), items)

}
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	addCmd.Flags().StringVar(&deadline, "deadline", "none", "DD/MM/YYYY HH:MM")

}
