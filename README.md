# URL Shortener

This project is a URL shortener service written in Go. It allows users to shorten long URLs into more manageable, shorter links.

## Features

- Shorten long URLs
- Redirect to the original URL using the shortened link
- Track click statistics

## Installation

To install the URL Shortener, follow these steps:

1. Clone the repository:
    ```
    git clone https://github.com/jamwitk/url-shortener.git
    ```
2. Navigate to the project directory:
    ```
    cd url-shortener
    ```
3. Build the project:
    ```
    go build
    ```

## Usage

To run the URL shortener service, use the following command:
```
./url-shortener
```

By default, the service will run on `http://localhost:8080`.

## API Endpoints

### Shorten URL

- **Endpoint:** `/shorten`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "url": "http://example.com"
    }
    ```
- **Response:**
    ```json
    {
        "shortened_url": "http://localhost:8080/abc123"
    }
    ```

### Redirect to Original URL

- **Endpoint:** `/{shortened_url}`
- **Method:** `GET`
- **Response:** Redirects to the original URL.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
