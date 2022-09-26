FROM golang:1.19-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o book-api .

FROM scratch

COPY --from=builder /build/book-api /

COPY --from=builder /build/.env /

EXPOSE 3000

ENTRYPOINT /book-api
