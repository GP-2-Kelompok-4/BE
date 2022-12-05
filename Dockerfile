# FROM golang:1.19

# # membuat direktori app
# RUN mkdir /app

# # set working directory /app
# WORKDIR /app

# COPY ./ /app

# RUN go mod tidy

# RUN go build -o immersive-dashboard

# CMD ["./immersive-dashboard"]

FROM golang:1.19-alpine3.16

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . .

##buat executeable
RUN go build -o immersive-dashboard .

##jalankan executeable
CMD ["./immersive-dashboard"]
