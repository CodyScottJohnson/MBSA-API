package config

import(
  "fmt"
  "bytes"

  "github.com/spf13/viper"
  "github.com/ttacon/chalk"
)


func Get()(*viper.Viper){
  viper.SetConfigName("production")     // no need to include file extension
  viper.AddConfigPath("config")  // set the path of your config file
  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println(chalk.Magenta.Color("Config file not found"))
    fmt.Println("  -loading default")
    viper.SetConfigType("toml")
    viper.ReadConfig(bytes.NewBuffer(default_config()))
  } else{
      viper.WatchConfig()
      fmt.Println(chalk.Blue.Color("(*) Loaded Config"))
  }
  //fmt.Println(viper.GetString("sql_api.url"))
  return viper.GetViper()

}
