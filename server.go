package main


import (
  // "io"
  // "fmt"
  "net/http"
  "github.com/bmizerany/pat"
  "log"
  "strings"
  "html/template"
)

// func main() {
//   resp, err := http.Get("http://google.com")
//   if err != nil {
//     fmt.Println(err)
//   }
//   defer resp.Body.Close()
//   fmt.Println(resp.Header["Expires"])
// }

type ResponseHeaderInfo struct {
  Public bool
}

// hello world, the web server
func HeaderServer(w http.ResponseWriter, req *http.Request) {
  req.ParseForm()
  url := req.Form.Get("url")
  header := ResponseHeader(url)
  public := false
  for _, value := range header["Cache-Control"] {
    if strings.HasPrefix(value, "public") {
      public = true
    }
  }
  info := ResponseHeaderInfo{public}
  tmpl, err := template.ParseFiles("./views/headers.html.tmpl")
  if err != nil { }
  tmpl.Execute(w, info)
}

func main() {
  m := pat.New()
  m.Get("/headers", http.HandlerFunc(HeaderServer))
  http.Handle("/", m)
  err := http.ListenAndServe(":12345", nil)
  log.Println("Serving at :12345")
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func ResponseHeader(url string) http.Header {
  resp, err := http.Get(url)
  if err != nil {
    log.Println("ResponseHeader: ", err)
  }
  defer resp.Body.Close()
  return resp.Header
}