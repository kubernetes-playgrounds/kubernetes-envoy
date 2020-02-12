# FROM envoyproxy/envoy-alpine:latest
# RUN rm /etc/envoy/envoy.yaml
# COPY frontend.yaml /etc/envoy/frontend.yaml

FROM envoyproxy/envoy-dev:latest
COPY frontend.json /etc/envoy/envoy.json
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.json --service-cluster frontend --mode serve
