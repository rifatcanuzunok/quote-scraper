# Goodreads Quotes Scraper

## Overview

The Goodreads Quotes Scraper is a Golang application that extracts quotes from Goodreads and stores them in a PostgreSQL database. The application fetches quotes along with their authors and quote categories, providing an easy way to build a collection of inspiring and thought-provoking quotes.
Features

   - **Web Scraping**: Utilizes Golang and the PuerkitoBio/goquery library to scrape quotes from Goodreads.
   - **Database Storage**: Saves quotes, authors, and categories in a PostgreSQL database.
   - **Configuration**: Supports configuration through environment variables or a `config.yaml` file.

## Prerequisites

Before running the application, ensure you have:

   - Golang installed on your machine.
   - A running PostgreSQL database with the necessary configuration details.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/rifatcanuzunok/quote-scraper
cd your-project-name
```


2. Install dependencies:

```bash
go mod download
```

3. Set up your PostgreSQL database and configure connection details in config.yaml or through environment variables.

Run the application:

```bash
go run main.go
```

## Configuration

Configure the application using environment variables or a config.yaml file. Example config.yaml:

```yaml

DB_USERNAME: your_db_username
DB_PASSWORD: your_db_password
DB_HOST: localhost
DB_PORT: 5432
DB_NAME: your_db_name
```
