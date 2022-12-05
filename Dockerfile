# FROM golang:1.19

# # membuat direktori app
# RUN mkdir /app

# # set working directory /app
# WORKDIR /app

# COPY ./ /app

# RUN go mod tidy

# RUN go build -o immersive-dashboard

# CMD ["./immersive-dashboard"]
