package monitoring

import (
  "time"

  "github.com/rcrowley/go-metrics"
  "github.com/vrischmann/go-metrics-influxdb"
)

func Init(config config) {

    r := metrics.NewRegistry()
    metrics.RegisterDebugGCStats(r)
    metrics.RegisterRuntimeMemStats(r)

    go metrics.CaptureDebugGCStats(r, time.Second*5)
    go metrics.CaptureRuntimeMemStats(r, time.Second*5)
    if(config.IsSet("influx") && config.GetBool("influx.enabled")){
      go influxdb.InfluxDB(
          r,                       // metrics registry
          time.Second*5,           // interval
          "http://localhost:8086", // the InfluxDB url
          "Hush_API",                 // your InfluxDB database
          "",                  // your InfluxDB user
          "",            // your InfluxDB password
      )
    }
  }


  type config interface{
    GetBool(key string)(bool)
    GetInt(key string)(int)
    GetFloat64(key string)(float64)
    GetString(key string)(string)
    IsSet(key string)(bool)
  }
