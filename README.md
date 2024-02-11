## SE2 Assignment 1

This assignment requires you to work with docker-compose, prometheus, and grafana.

### Overall idea

Three different services collect statistics about the host machine, the docker containers, and kernel (host) metrics.
Prometheus, a time-series database collects (scraps) the metrics from the aforementioned services.
Grafana, a visualization web-service, will continuously present the data.

### Technical setup

A start script handles preliminary steps and calls docker-compose.
The start script builds local files and create docker images.
Docker-Compose starts all five services.
Docker-Compose doesn't work well while your have an active VPN connection!

#### Requirements for local testing:

- libelf-dev
- clang
- docker.io
- docker-compose
- make
- stress
- maybe more

#### Task list

- Read the readme at https://github.com/cloudflare/ebpf_exporter
- In a seperate, temporary folder, clone & compile the ebpf_exporter without using docker.
- Run the syscalls example as described in the readme of ebpf_exporter
- Execute `stress -d 6 --hdd-bytes 10GB` to see if the amout of write syscalls increase substatially.
- Hint: The examples need be compiled locally and later mounted into the ebpf-container

- Read the readme at https://github.com/google/cadvisor
- Read the readme at https://github.com/prometheus/node_exporter

- Fork this repository.
- Make sure your repo is private.
- Add the user 'se-gitlab machine' as reporter to the members of your repository.

- In general, "ADD CONTENT HERE" are to be replaced with your solution.
- Start with `start.sh`, then `docker-compose.yml`
- Use `docker-compose run SERVICENAME` to test individual services
- Adapt `prometheus.yml` according to your definitions in docker-compose.yml
- Ensure that all services start
- Use `docker ps` to check
- `start.sh` should be a working one-click-be-happy script
- Open the urls listed at the end of `start.sh` for verification.

- Open your Prometheus instance
- Ensure that Prometheus is fetching metrics from all three sources (ebpf-exporter, cAdvisor, node-exporter)
  - Status -> Targets

- Familiarize yourself a bit with Promql, and Grafana
- Create a toy-query in Prometheus to check if you have data, e.g. 'node_memory_Cached_bytes'
- Open your Grafana instance and login: admin (PW defined in `config.monitoring`)
- Add your prometheus instance as source (use the container name as url)
- Create an empty dashboard
- Create multiple panels in your dashboard using (but not limited to) these metrics:
  - `ebpf_exporter_syscalls_total`: Show the rate of all read and write syscalls (host).
  - `container_cpu_usage_seconds_total`: Show the CPU utilization of only the Prometheus container and only the Grafana container.
  - `ebpf_exporter_llc_misses_total`: Show the current rate of cache misses on the host.
  - `container_network_receive_bytes_total`, `container_network_transmit_bytes_total`: Show the network traffic of only the Prometheus container.
  - `node_memory_MemTotal_bytes`, `node_memory_MemAvailable_bytes`: Show the node's (host's) memory utilization.
  - `node_cpu_seconds_total`: Show the node's CPU utilization (needs a longer formula).
  - `ebpf_exporter_bio_latency_seconds_bucket`: show a historgram of your disk's write latency (hints: heatmap, https://grafana.com/blog/2020/06/23/how-to-visualize-prometheus-histograms-in-grafana/, ebpf-exporter readme)
  - `container_???`: Create a panel that shows the memory utilization of the Grafana container after you added the other panels.

- How much memory does Grafana actually use?
- Adapt the `docker-compose.yml` (at line ~101) file to limit the available memory of the grafana container to it's average usage plus ~20% extra.

- Stress your machine with, e.g., `stress`, `dd`, or youtube, to see if your metrics change.
- Try `stress -d 6 --hdd-bytes 10GB`
- Add other panels as you see fit.

- On your dashboard, find the sharing button and share (save) your dashboad for external-use as `dashboard.yml`.
- Save and upload your `dashboard.yml` to your repository.
- Upload a screenshot of your dashboard as `dashboard.png` into your repository.


### Evaluation

- Ensure that you have enrolled properly at your examination office (jexam/...)
- Ensure that you filled this form: https://docs.google.com/forms/d/e/1FAIpQLSf6w_kFM6DpEKLLpDbXipOTEj4TY64o2cFfY8KhJPKbrjaxbg/viewform
- Your repository will be evaluated by our CI job.
- You have to fork this repository.
- You need to add the user `se-gitlab machine` as reporter to your repository.
- You need to make your repository private (default).
- You'll receive e-mails about evaluations.
- We will, additionally, look into your repository.
- The CI-checker will run tests and execute `start.sh`, ensure that it works.

### Deadline

- 15.12.2023 at 18:00 (Dresden local time)