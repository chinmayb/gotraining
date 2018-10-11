import ("fmt"
  "net/http"
  ""
  "os")


func fetch(url string, chanb chan string){
  resp, err := http.Get(url)
  if err != nil {
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    defer resp.Body.Close()
    if err != nil {
      chanb <- fmt.Sprintf()
      }
  }
}

func main(){
  urls := os.Args[1:]
  chanb := make(chan string)
  for _,url := range urls {
    go fetch(url, chanb)
  }
  fmt.Println("")
}
