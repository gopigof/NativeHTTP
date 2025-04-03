# NativeHTTP - HTTP Server from Scratch

This project is a custom HTTP/1.1 server written in Go that handles concurrent connections, parses HTTP requests, and routes them to appropriate handlers. It features a modular design with clean separation between routing logic and request handling.

## Usage
- Starting the Server: ```./your_program.sh --directory <path_to_dir>```
- directory is the file directory for file operations

## Implemented Routes
- GET /: Returns a 200 OK response
- GET /echo/{message}: Returns the specified message in the response body
- GET /user-agent: Returns the User-Agent header from the request
- GET /files/{filename}: Returns the contents of the specified file
- POST /files/{filename}: Uploads a file to the server
