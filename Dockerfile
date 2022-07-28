# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang as base

ADD . /usr/src/llegend

FROM base as build
WORKDIR /usr/src/llegend
RUN go build

FROM build as exec
WORKDIR /app
COPY --from=build /usr/src/llegend/llegend /app
COPY --from=build /usr/src/llegend/*.html /app
CMD ["/app/llegend"]