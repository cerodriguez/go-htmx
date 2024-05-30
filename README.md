# GO HTMX CRUD Demo

This is a simple CRUD (Create, Read, Update, Delete) demo application using HTMX for the frontend and Go for the backend. The application demonstrates how to dynamically update parts of a web page without reloading, using HTMX with a Go backend.

## Prerequisites

- Go 1.22.3 or later
- Internet connection to load HTMX from CDN

## Project Structure
htmx-demo/
├── main.go
├── templates/
│ ├── index.html
│ ├── item_list.html
│ └── form.html


## Setup and Running

1. **Clone the repository or download the zip file**:
    ```sh
    git clone git@github.com:cerodriguez/go-htmx.git
    cd go-htmx
    ```

2. **Run the Go server**:
    ```sh
    go run main.go
    ```

3. **Open your browser and navigate to**:
    ```
    http://localhost:8080
    ```

## Description

- **main.go**: The main Go file that sets up the web server and handles the CRUD operations.
- **templates/index.html**: The main HTML file that includes the HTMX script and the initial structure for the CRUD interface.
- **templates/item_list.html**: The HTML template for displaying the list of items.
- **templates/form.html**: The HTML template for the form used to create and update items.

## Endpoints

- `GET /`: Displays the main page and the list of items.
- `GET /create`: Displays the form to create a new item.
- `POST /create`: Handles the creation of a new item.
- `GET /update`: Displays the form to update an existing item.
- `POST /update`: Handles the update of an existing item.
- `GET /delete`: Handles the deletion of an item.

## Using HTMX

HTMX is used to enhance the interactivity of the web application by making AJAX requests and updating parts of the page without a full reload. Key HTMX attributes used in this project:
- `hx-get`: Used to make a GET request.
- `hx-post`: Used to make a POST request.
- `hx-target`: Specifies the target element to be updated.
- `hx-swap`: Specifies how the response should be swapped into the target element.

