package interface

import (
  "github.com/47monad/zaal/cue/common@v0"
)

#HTTP: {
  port: *4747 | common.#Port
}

http: #HTTP
