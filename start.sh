#!/bin/bash

# Fail on fist error:
set -e

# Get ebpf_exporter:
# Use git to fetch https://github.com/cloudflare/ebpf_exporter.git into the folder exporters/ebpf_exporter.
# Skip the step if directory existsp

if [ ! -d "exporters/ebpf_exporter" ]; then
	git clone https://github.com/cloudflare/ebpf_exporter.git exporters/ebpf_exporter
fi

# ebpf_exporter contains examples that need to be build locally (on the host) beforehand
# Use make to compile the code in the examples subfolder

cd exporters/ebpf_exporter
make -C examples clean build

# Build the container for ebpf_exporter (Dockerfile exists already)
# Use docker build
docker build --tag ebpf_exporter --target ebpf_exporter .
docker cp $(docker create ebpf_exporter):/ebpf_exporter ./
#docker build -t ebpf_exporter .
#docker cp $(docker create ebpf_exporter):/ebpf_exporter ./
cd ../../



# Storage folder for Prometheus
cd prometheus
if [ -d "storage" ]; then
	echo "re-using old storage dir"
else
	mkdir storage
fi
cd ..


# Starting all services
# Do not change this
CURRENT_UID=$(id -u):$(id -g) docker-compose --compatibility up -d


# Check if you can open these:
echo "You may now open Ebpf_exporter at http://localhost:9440/metrics"
echo "You may now open Node-Exporter at http://localhost:9441/metrics"
echo "You may now open Cadvisor at http://localhost:9442/metrics"
echo "You may now open Prometheus at http://localhost:9090"
echo "You may now open Grafana at http://localhost:9091"




