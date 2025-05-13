package main

import "github.com/huugiii/hello-world-api/internal/app/application"

func main() {
  app := application.New()
  app.Start()
}