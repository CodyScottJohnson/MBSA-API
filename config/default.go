package config

func default_config()([]byte) {

  return []byte(`
    [sql]
      [sql.api]#This is the connection info for the database where basic app data will be stored
      url = "postgresql://postgres@localhost:32791/Hush?sslmode=disable"
      port = 32
      type = "postgres"
      [oauth]#This is the connection info for the database where Oauth info will be stored
    [server]
    port = 1234
`)
}
