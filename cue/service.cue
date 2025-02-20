package cue

import (
  "github.com/47monad/zaal/cue/db"
  "github.com/47monad/zaal/cue/log"
  "github.com/47monad/zaal/cue/interface"
  "github.com/47monad/zaal/cue/monitoring"
)

#Env: "production" | "staging" | "dev"
#Mode: "normal" | "debug"

#Schema: {
  name: *"go-app" | string
  title: *"Go App" | string
  version: *"1.0.0" | string
  host: *"127.0.0.1" | string
  env: *"dev" | #Env
  mode: *"normal" | #Mode
  logging: log.config
  mongodb?: db.mongodb
  prometheus?: monitoring.prometheus
  grpc?: interface.grpc
  http?: interface.http
}

service: #Schema

