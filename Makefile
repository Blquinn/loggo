all: loggo gcpstream

.PHONY: loggo
loggo:
	(cd cli && go build -ldflags="-s -w" -o ../loggo .)

.PHONY: gcpstream
gcpstream:
	(cd gcpstream && go build -ldflags="-s -w" -o ../loggo-gcp-stream .)

clean:
	rm loggo
	rm loggo-gcp-stream
