docker/run:
	docker run --name=proxy --rm \
	-p 9901:9901 \
  	-p 10000:10000 \
  	-v $(CURDIR)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
  	envoyproxy/envoy:v1.14.1
