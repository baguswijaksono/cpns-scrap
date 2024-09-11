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
    // Membuat query Google Dorking
    searchTerm := "CPNS pembukaan"
    fileType := "pdf"
    site := ".go.id"

    // Menghasilkan query URL
    dorkQuery := fmt.Sprintf("%s filetype:%s site:%s", url.QueryEscape(searchTerm), fileType, site)
    searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", dorkQuery)

    fmt.Printf("Mencari: %s\n", searchURL)

    // Membuat collector untuk scraping
    c := colly.NewCollector()

    // Membuka file CSV untuk menulis
    file, err := os.Create("results.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Menulis header ke file CSV
    err = writer.Write([]string{"Domain", "File PDF"})
    if err != nil {
        log.Fatal(err)
    }

    // Event handler untuk scraping setiap hasil pencarian
    c.OnHTML("a", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        if link != "" && len(e.Text) > 15 {
            // Menulis baris ke file CSV
            err := writer.Write([]string{e.Request.URL.Host, link})
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("Ditemukan: %s\n%s\n\n", e.Text, link)
        }
    })

    // Event handler untuk mengatasi masalah saat scraping
    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Error:", err)
    })

    // Memulai scraping
    err = c.Visit(searchURL)
    if err != nil {
        log.Fatal(err)
    }
}
