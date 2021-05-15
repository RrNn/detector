package jobs

import (
  "fmt"
  "log"

  "sync"
  "time"

  "github.com/RrNn/detector/controller"
  "github.com/RrNn/detector/http"
  "github.com/RrNn/detector/models"
)

func StartPinging(c *controller.Controller) {
  // urlsToPing := []models.Url{}
  // c.DB.Find(&urlsToPing)
  // URLPinger(urlsToPing, c)
  URLPinger(c)
}

var wg = sync.WaitGroup{}

func fetchUrls(cont *controller.Controller) (urls []models.Url) {
  urls = []models.Url{}
  cont.DB.Find(&urls)
  return urls
}

func URLPinger(cont *controller.Controller) {
  var count int = 0
  ticker := time.NewTicker(30 * time.Second)
  go func() {
    urlsToPing := fetchUrls(cont)
    for {
      select {
      case t := <-ticker.C:
        log.Println("-------------- Pinging starting starting now", t, " --------------")
        for k := range urlsToPing {
          ping := models.Ping{
            Time:  time.Now(),
            UrlID: urlsToPing[k].ID,
          }
          if resp, err := http.Client.Get(urlsToPing[k].Link); err != nil {
            log.Print("An error occured: ", time.Now(), err)
            ping.Error = err.Error()
            CreatePing(ping, urlsToPing[k], cont)
          } else {
            ping.Status = resp.StatusCode
            CreatePing(ping, urlsToPing[k], cont)
            count++
            fmt.Printf("URL : %s -- STATUS_CODE : %d\n", urlsToPing[k].Link, resp.StatusCode)
          }
        }
        if count == len(urlsToPing) {
          log.Println("-------------- Pinging ended at", time.Now(), ": All URL's were pinged --------------")
          urlsToPing = fetchUrls(cont)
        } else {
          log.Println("-------------- Pinging ended at", time.Now(), ": Some URL's were not pinged --------------")
          urlsToPing = fetchUrls(cont)
        }
      }
    }
  }()
}

func CreatePing(ping models.Ping, url models.Url, cont *controller.Controller) {
  if err := cont.DB.Create(&ping).Error; err != nil {
    log.Print("Error creating ping record for", url, err)
  }
}
