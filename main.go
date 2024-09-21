package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "strings"

    "github.com/gocolly/colly/v2"
)

func main() {
    // Define short flags for user inputs
    pages := flag.Int("p", 1, "Number of pages to scrape")
    downloadPDFs := flag.Bool("d", false, "Download PDFs")
    includeKabProv := flag.Bool("k", false, "Include kabupaten and provincial domains")

    // Parse the flags
    flag.Parse()

    // Display ASCII art
    fmt.Println(`

    CPNS Scraper - 2024
    `)

    // Default search term for national government domains
    searchTerm := "CPNS pembukaan filetype:pdf site:.go.id 2024"

    // Modify searchTerm based on user's choice
    if *includeKabProv {
        searchTerm = "CPNS pembukaan filetype:pdf site:.go.id OR site:.kab.go.id OR site:.prov.go.id 2024"
    }

    baseURL := "https://www.google.com/search?q=%s&start=%d"

    downloadDir := "downloads" // Directory where PDFs will be saved
    os.MkdirAll(downloadDir, os.ModePerm) // Create the download directory if it doesn't exist

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
        if link != "" && strings.Contains(link, ".pdf") {
            domain := getDomain(link)
            if domain != "" {
                err := writer.Write([]string{domain, link})
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Printf("Ditemukan PDF: %s\n%s\n\n", e.Text, link)

                // Download the PDF if the user has enabled download option
                if *downloadPDFs {
                    downloadPDF(link, downloadDir)
                }
            }
        }
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
    })

    for i := 0; i < *pages; i++ {
        searchURL := fmt.Sprintf(baseURL, url.QueryEscape(searchTerm), i*10)
        fmt.Println("Visiting page:", searchURL)
        err := c.Visit(searchURL)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func getDomain(link string) string {
    parsedURL, err := url.Parse(link)
    if err != nil {
        return ""
    }
    return parsedURL.Host
}

// downloadPDF downloads the PDF file from the given URL and saves it to the specified directory
func downloadPDF(fileURL, dir string) {
    // Parse the file name from the URL
    parsedURL, err := url.Parse(fileURL)
    if err != nil {
        fmt.Println("Failed to parse URL:", fileURL)
        return
    }

    // Create a valid file name from the URL
    fileName := filepath.Base(parsedURL.Path)
    if !strings.HasSuffix(fileName, ".pdf") {
        fileName += ".pdf"
    }
    filePath := filepath.Join(dir, fileName)

    // Create the file
    out, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Failed to create file:", filePath)
        return
    }
    defer out.Close()

    // Download the PDF file
    resp, err := http.Get(fileURL)
    if err != nil {
        fmt.Println("Failed to download file:", fileURL)
        return
    }
    defer resp.Body.Close()

    // Copy the content to the file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        fmt.Println("Failed to save file:", filePath)
    } else {
        fmt.Println("File saved:", filePath)
    }
}
