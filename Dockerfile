FROM golang:1.23.2
LABEL authors="mishukoffs"

# Move to working directory /build
WORKDIR /app

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o telegram_weather_bot

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/app/telegram_weather_bot"]