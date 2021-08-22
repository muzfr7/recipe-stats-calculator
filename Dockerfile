FROM alpine:3.7

# set working dir
WORKDIR /app

# copy binary to /app/bin/
COPY ./bin ./bin

# set binary as an entrypoint
ENTRYPOINT ["/app/bin/main"]
