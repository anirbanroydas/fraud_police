FROM golang:1.10 AS builder

# Add golang dep tool
RUN apt-get update && \
	curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \
	chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/anirbanroydas/fraud_police
WORKDIR /go/src/github.com/anirbanroydas/fraud_police

COPY Gopkg.toml Gopkg.lock ./

# instal depnedecies from Gopkg.lock without considering the source code
RUN dep ensure -vendor-only

# copy every file and folder selectively to avoid sending unwanted or confidential data 
COPY pkg ./pkg
COPY cmd ./cmd

# build/compile the project
RUN cd cmd/fraud_police_server && \
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o fraud_police_server


# Stage 2
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/anirbanroydas/fraud_police/cmd/fraud_police_server  .

EXPOSE 8081

ENTRYPOINT ["./fraud_police_server"]

# CMD ["docker-entrypoint.sh"]