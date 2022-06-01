FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM node:0.12.4-onbuild
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
