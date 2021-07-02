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
	"log"
	"strconv"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Task",
	Long:  `Delete a Task`,
	Run: func(cmd *cobra.Command, args []string) {
		items, _ := task.ReadItems(viper.GetString("datafile"))
		if allOpt {
			items = []task.Item{}
		} else {

			for _, arg := range args {

				index, err := strconv.Atoi(arg)
				if err != nil {
					log.Fatal(args[0], " is not a valid label\n", err)
				}
				if index > 0 && index <= len(items) {
					items = append(items[:index-1], items[index:]...)
				} else {
					log.Fatal("Cannot find the task", arg)
				}
			}
		}
		_ = task.SaveItems(viper.GetString("datafile"), items)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVar(&allOpt, "all", false, "show all todos")
}
