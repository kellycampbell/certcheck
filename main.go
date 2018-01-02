package main

import (
  "os"
  "fmt"
  "github.com/cloudflare/cfssl/certinfo"
)



func main() {
  for _, url := range os.Args[1:] {
    cert, err := certinfo.ParseCertificateDomain(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }

    fmt.Printf("%v: ", cert.SANs)
    fmt.Printf("%v\n", cert.NotAfter)
  }
}
