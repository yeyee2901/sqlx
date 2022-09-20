FROM golang:1.19.1-alpine3.15

WORKDIR /aplikasi
RUN apk update && apk add libc-dev && apk add gcc && apk add make

# copy go mod & download required pkg
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download && go mod verify

# Copy entire project
COPY . .

# compile the project
RUN go build

CMD ["./sqlx"]
