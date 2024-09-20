FROM golang as builder

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o  ./
FROM golang as production

COPY --from=builder /app/app ./

EXPOSE 8090
ENTRYPOINT ["./app","serve","--http=0.0.0.0:8090"]