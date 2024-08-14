# ASCII Art Web Dockerize

This project is a web-based application that allows users to generate ASCII art using different banners. The application is built using Go and Dockerized for easy deployment.

## Features

- Web-based GUI for generating ASCII art.
- Supports multiple banners for text rendering.
- Dockerized for ease of deployment.

## Project Structure

- **`main.go`**: The main entry point of the application.
- **`ascii-art/`**: Contains the ASCII art generation logic.
- **`handlers/`**: Contains the HTTP handlers for serving web requests.
- **`static/`**: Static files like CSS and JavaScript.
- **`templates/`**: HTML templates for the web interface.
- **`Dockerfile`**: Docker configuration file for building the Docker image.
- **`script.sh`**: Shell script for setting up or running the application.

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed for containerization.

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://learn.zone01kisumu.ke/git/quochieng/ascii-art-web-dockerize.git
   cd ascii-art-web-dockerize
   ```
2. **Build a Docker image and run it in a  container:**

   To build the application image and run it in a container, just run script.sh:
   ```bash
   ./script.sh
   ```

3. **Access the application:**

   Open your web browser and go to http://localhost:8080 to start using the application.

## Docker Implementation
- The Dockerfile defines a multi-stage build for an ASCII art web application. It uses the golang:1.22.6 image for building the Go application, then copies the compiled binary into a minimal alpine:3.18 image. 
- Metadata labels provide details about the project. 
- The final image includes the executable, static files, templates, and banner files. It exposes port 8080 and runs the application with CMD ["./main"].

## Contributors

This project has been authored by:
  - Quinter Ochieng ([GitHub](https://github.com/apondi-art))
  - John Opiyo ([GitHub](https://github.com/SidneyOps75))
  - Hilary Okello ([GitHub](https://github.com/HilaryOkello))
  
## License

This project is licensed under the [MIT](LICENSE) License
