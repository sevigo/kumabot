# docker build --rm -f docker/Dockerfile -t sevigo/kumabot .
# docker rm kumabot --force 
# docker run -e OWM_API_KEY=$OWM_API_KEY -e LINE_CHANNEL_SECRET=$LINE_CHANNEL_SECRET -e LINE_CHANNEL_TOKEN=$LINE_CHANNEL_TOKEN \
#   --publish 8081:8081 --rm --name kumabot kumabot:latest
FROM golang:1.14.2

COPY . /go/src/github.com/sevigo/kumabot
WORKDIR /go/src/github.com/sevigo/kumabot
RUN make install

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=0 /go/bin/kumabot /bin/kumabot
ENTRYPOINT ["/bin/kumabot"]
EXPOSE 8081
