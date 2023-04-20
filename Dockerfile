FROM golang:1.19-alpine as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o languagle_api_app .

FROM alpine
WORKDIR /app/
COPY --from=builder /app/languagle_api_app .
EXPOSE 3000
ENTRYPOINT ["./languagle_api_app"]