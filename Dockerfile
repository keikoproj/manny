# Build Stage
FROM golang:1.13-alpine AS builder

LABEL REPO="https://github.com/keikoproj/manny"

WORKDIR /go/src/github.com/keikoproj/manny
COPY . .

RUN make build

# Final Stage
FROM scratch

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/keikoproj/manny"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

COPY --from=builder /go/src/github.com/keikoproj/manny/bin/manny /bin/manny

CMD ["/bin/manny", "--help"]
