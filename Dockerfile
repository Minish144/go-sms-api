FROM golang:alpine AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o api .

FROM alpine AS bin
COPY --from=build /go/src/app/api /api
CMD ["/api"]
