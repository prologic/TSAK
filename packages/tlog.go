package packages

import (
  "github.com/vulogov/TSAK/internal/log"
  "reflect"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["tlog"] = map[string]reflect.Value{
    "Trace":    reflect.ValueOf(log.Trace),
    "Info":     reflect.ValueOf(log.Info),
    "Error":    reflect.ValueOf(log.Error),
    "Warning":  reflect.ValueOf(log.Warning),
  }
}
