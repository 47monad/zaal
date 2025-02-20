package interface

import (
  "github.com/47monad/zaal/cue/common@v0"
)

#GRPC: {
  port: *4748 | common.#Port
  features: {
    reflection: *false | bool
    healthCheck: *false | bool
    logging: *false | bool
  }
}

grpc: #GRPC
