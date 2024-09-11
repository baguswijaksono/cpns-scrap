package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "net/url"
    "os"
    "github.com/gocolly/colly/v2"
)

func main() {
    searchTerm := "CPNS pembukaan filetype:pdf site:.go.id 2024"
    searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(searchTerm))

    c := colly.NewCollector()

    file, err := os.Create("results.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    err = writer.Write([]string{"Domain", "File PDF"})
    if err != nil {
        log.Fatal(err)
    }

    c.OnHTML("a", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        if link != "" && len(e.Text) > 15 {
            err := writer.Write([]string{getDomain(link), link})
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("Ditemukan: %s\n%s\n\n", e.Text, link)
        }
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
    })

    err = c.Visit(searchURL)
    if err != nil {
        log.Fatal(err)
    }
}

func getDomain(link string) string {
    parsedURL, err := url.Parse(link)
    if err != nil {
        return ""
    }
    return parsedURL.Host
}
