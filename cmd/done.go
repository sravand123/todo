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
	"strconv"
	"todo/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Set a task completed",
	Long:  `To set  taks complete , run ./tri done taskNo`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {

			i, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal(arg, " is not a valid label\n", err)
			}
			items, _ := task.ReadItems(viper.GetString("datafile"))
			if i > 0 && i <= len(items) {
				items[i-1].Done = true
				fmt.Printf("%q %v\n", items[i-1].Text, "marked done")
				task.SaveItems(viper.GetString("datafile"), items)
			} else {
				log.Println(i, " doesn't match any items")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
