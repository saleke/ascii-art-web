# ASCII Art Web

ASCII Art Web is a Go-based web application for generating ASCII art from user-provided text using configurable banner styles.

## Overview

The application provides a focused browser interface for entering text, selecting a banner, generating formatted ASCII output, and copying the result. It uses Go's standard HTTP library and server-rendered HTML without authentication, tracking, or an external database.

## Features

- Converts text into banner-style ASCII art.
- Includes 13 selectable banner styles.
- Supports line breaks and printable ASCII characters.
- Preserves generated whitespace in a horizontally scrollable output area.
- Offers live result refinement and copy-to-clipboard support.
- Provides responsive, keyboard-accessible forms and project links.

## Technologies

- Go 1.22
- `net/http`
- Go HTML templates
- HTML, CSS, and small progressive-enhancement JavaScript

## How It Works

1. `GET /` renders the generator form.
2. `POST /ascii-art` validates the submitted text and banner against an allowlist.
3. The banner loader reads the selected font file and maps its 95 glyph records to printable ASCII characters.
4. The renderer replaces each input character with the corresponding glyph rows while preserving line breaks.
5. The server returns a rendered result page or JSON when the request asks for `application/json`.

Banner and template paths are resolved from the project root, so the application does not depend on the shell's current directory. Set `APP_ROOT` when a deployment stores these assets somewhere other than the working directory.

## Running Locally

From the project root:

```sh
go run .
```

Open [http://localhost:8080](http://localhost:8080). To use another port, run `PORT=3000 go run .`.

To build and run a binary:

```sh
go build -o ascii-art-web .
./ascii-art-web
```

Keep the `templates`, `static`, and `banners` directories with the application when moving the binary outside the source tree, and set `APP_ROOT` to their parent directory.

## Usage

Enter up to 500 characters, choose a banner, and select **Generate art**. On the result page, edit the message or banner to refresh the preview, or use **Copy output** to copy the generated art.

## Configuration

| Variable | Default | Purpose |
| --- | --- | --- |
| `PORT` | `8080` | HTTP server port supplied by local or production environments. |
| `APP_ROOT` | Source project root | Parent directory for templates, static files, and banners. |
| `PORTFOLIO_URL` | `https://saleke.vercel.app` | Portfolio link shown in the interface. |
| `GITHUB_REPOSITORY_URL` | `https://github.com/saleke/ascii-art-web` | Repository link shown in the interface. |

## Testing

Run the full test suite with `go test ./...`. Tests cover handlers, JSON previews, validation, all included banner files, production port configuration, routes, and static asset serving.

## Deployment

The application is a standard Go HTTP service. A compatible host should run `go build -o bin/ascii-art-web .` and `./bin/ascii-art-web`, provide `PORT`, and retain the `templates`, `static`, and `banners` directories. No production URL is hardcoded in the application. `render.yaml` provides a minimal Render Blueprint for this service.

## Project Structure

```text
ascii-art-web/
|-- ascii-art/   # Banner loading and ASCII rendering
|-- banners/     # Font/banner data files
|-- handlers/    # HTTP handlers and project configuration
|-- static/      # Stylesheet
|-- templates/   # Server-rendered pages
`-- main.go      # Server and routing setup
```

## Developer

Built by **Aleke Emmanuel Solomon**, Software Developer in Training.

- [GitHub](https://github.com/saleke)
- [Portfolio](https://saleke.vercel.app)
- [Live Site](https://asciiartweb.vercel.app/)

## Repository

[github.com/saleke/ascii-art-web](https://github.com/saleke/ascii-art-web)
