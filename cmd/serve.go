// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	"Hush_API/sqldb"
    "Hush_API/server"
    "Hush_API/config"
	"Hush_API/monitoring"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))
		fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("--JFS API--")))
		fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))

		config := config.Get();
		config.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
		if(!config.IsSet("sql.api")){
			panic(chalk.Magenta.Color("DB Settings Missing from Config"))
		}

		if(!config.IsSet("server")){
			panic(chalk.Magenta.Color("Server Settings Missing from Config"))
		}

		if(config.IsSet("monitoring.enabled") && config.GetBool("monitoring.enabled")){
			monitoring.Init(config.Sub("monitoring"))
		}
		db, err := sqldb.Connect(config.Sub("sql.api"))

		defer db.Close()
		err = server.Start(db, config.GetString("server.port"))
		if err != nil{
			fmt.Println(chalk.Red.Color( "Unable To Connect to Database:"), err)
			panic(chalk.Red.Color(err.Error()))
		}

	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
