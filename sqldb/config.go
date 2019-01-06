package sqldb

type config interface{
  GetBool(key string)(bool)
  GetInt(key string)(int)
  GetFloat64(key string)(float64)
  GetString(key string)(string)
  IsSet(key string)(bool)
}
