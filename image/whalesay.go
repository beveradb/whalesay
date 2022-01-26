package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var name, _ = os.Hostname()

  fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>nikovirtala/whalesay</title>
</head>
<body>
<pre>
    .___                           
  __| _/____   ____   ____   ____  
 / __ |/  _ \ / ___\ / ___\ /  _ \ 
/ /_/ (  <_> ) /_/  > /_/  >  <_> )
\____ |\____/\___  /\___  / \____/ 
     \/     /_____//_____/         

This request was processed by: %s
</pre>
</body>
</html>
`, name)
log.Print("Served request ",r,"\n")
}

func main() {
  log.SetOutput(os.Stderr)
  log.Println("Starting server ...")
  http.HandleFunc("/", handler)
  err := http.ListenAndServe(":8080",nil)
  if err != nil {
    log.Fatal("ListenAndServer: ", err)
  }
}
