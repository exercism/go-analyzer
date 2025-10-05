FROM golang:1.25.1-alpine3.22 as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser && mkdir /go-analyzer

# source code
WORKDIR /go-analyzer
COPY ./go.mod /go-analyzer/go.mod

# download dependencies
RUN go mod download

# Create analyze.sh
RUN printf '%s\n' '#!/bin/sh' '/opt/analyzer/bin/analyzer "$@"' > /go/bin/analyze.sh
RUN chmod +x /go/bin/analyze.sh

# get the rest of the source code
COPY . /go-analyzer

# build
RUN go generate .
RUN GOOS=linux GOARCH=amd64 go build --tags=build -o /go/bin/analyzer .

# Build a minimal and secured container
# The ast parser needs Go installed for import statements.
# Therefore, unfortunately we cannot build from scratch as we would normally do with Go.
FROM golang:1.25.1-alpine3.22
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin /opt/analyzer/bin
USER appuser
WORKDIR /opt/analyzer
ENTRYPOINT ["/opt/analyzer/bin/analyze.sh"]
