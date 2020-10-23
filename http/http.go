package http

import (
  "net"
  "net/http"
  "time"
)

var Client = http.Client{
  Timeout: 15 * time.Second,
  Transport: &http.Transport{
    Dial: (&net.Dialer{
      Timeout:   30 * time.Second,
      KeepAlive: 30 * time.Second,
    }).Dial,
    TLSHandshakeTimeout:   10 * time.Second,
    ResponseHeaderTimeout: 10 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,
  },
}
