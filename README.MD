# ASCII Art Web

## Description

ASCII Art Web is a small Go web application that converts ordinary text into banner-style ASCII art. It provides a focused browser interface and supports the required `standard`, `shadow`, and `thinkertoy` banners, along with the additional banner files included in the project.

## Authors

ASCII Art Web team.

## Usage: how to run

From the project root:

```sh
go run .
```

Open [http://localhost:8080](http://localhost:8080) in a browser. Enter text, select a banner, and choose **Generate art**. The server can also be built as a standalone binary:

```sh
go build -o ascii-art-web .
./ascii-art-web
```

## Implementation details: algorithm

1. The HTTP server serves the home page at `GET /`, static CSS under `/static/`, and accepts form submissions at `POST /ascii-art`.
2. The selected banner is checked against an allowlist before any file is opened. This prevents path traversal and gives invalid selections a `400 Bad Request` response.
3. A banner file contains 95 glyph records, one for each printable ASCII character. The loader determines the record height from the file, validates its dimensions, and stores each glyph in a rune map.
4. Input is normalized to Unix line endings and the escaped sequence `\\n` is treated as a line break. Each input character is replaced with the matching glyph row; unsupported characters are rendered as spaces.
5. Successful renders return `200 OK`. Malformed form data returns `400`, missing/unreadable banner or template files return `404`/`500` as appropriate, and unsupported methods return `405 Method Not Allowed`.
