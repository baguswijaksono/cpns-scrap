
# CPNS Scraper

This is a Go-based scraper for finding and downloading CPNS-related PDF files from government websites. The program uses Google search queries to locate the PDFs and allows the user to choose whether or not to download them.

## Features

- Scrapes PDF files from `.go.id`, `.kab.go.id`, and `.prov.go.id` domains.
- Optionally downloads the PDF files.
- Saves the scraped data (domain and PDF link) in a `results.csv` file.
- Customizable number of

## Prerequisites

- Go installed on your machine
- Internet access to perform the web scraping

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/baguswijaksono/cpns-scrap.git
   ```
2. Navigate to the project directory:
   ```bash
   cd cpns-scrap
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

You can run the program using various flags to customize its behavior.

### Flags

| Flag      | Description                                   | Example          |
| --------- | --------------------------------------------- | ---------------- |
| `-p`      | The number of Google search result pages to scrape. Default is 1. | `-p 5`           |
| `-d`      | Enable downloading of PDFs. This is a boolean flag. | `-d`             |
| `-k`      | Include scraping from kabupaten (`.kab.go.id`) and provincial (`.prov.go.id`) domains. This is a boolean flag. | `-k`             |

### Examples

1. **Scrape 5 pages without downloading PDFs:**
   ```bash
   go run main.go -p 5
   ```

2. **Scrape 5 pages and download the PDFs:**
   ```bash
   go run main.go -p 5 -d
   ```

3. **Scrape 5 pages, download PDFs, and include kabupaten/provinsi domains:**
   ```bash
   go run main.go -p 5 -d -k
   ```

### CSV Output

The scraper will generate a `results.csv` file in the root of the project. This file contains two columns:
- `Domain`: The domain from which the PDF file was found.
- `File PDF`: The direct URL to the PDF file.

### Download Directory

If the `-d` flag is used, the PDFs will be downloaded to a directory called `downloads`. This directory is created automatically by the program.

## Notes

- **User-Agent**: The scraper uses a custom `User-Agent` header to mimic a browser request.
- **Rate Limiting**: Be mindful of Google’s rate limits when running the scraper with high page counts.

## License

This project is licensed under the MIT License. search result pages to scrape.
 
