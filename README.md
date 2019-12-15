# monit
Docker Compose stack for Grafana - Prometheus

## How to Run

- clone this repo
```bash
$ git clone https://github.com/faruqisan/monit
$ cd monit
```
- run docker-compose
```bash
$ docker-compose up
```
Just add `-d` on option to run in background

## Container Names

Usually docker will named the container with prefix `monit_` and suffix `_1`, or just check your terminal for container name

| Image                     | Name                       |
|---------------------------|----------------------------|
| Grafana                   | monit_grafana_1            |
| Prometheus                | monit_prometheus_1         |
| Statsd-Exporter           | monit_statsd_exp_1         |
| Example Go-Prome App      | monit_example_go_1         |
| Example Go-Statsd-exp App | monit_example_statsd_app_1 |

## Prometheus Configuration

Prometheus configuration located at `/prometheus/prometheus.yml`, edit the configuration to add metrics to scape.

## Statsd-Exporter Configuration

Statsd-Exporter configuration located at `/statsd_exporter/statsd_mapping.yml`, edit the configuration to add metrics mapping

## Grafana

Grafana use persistence storage, located at `/grafana-persist`, so your dashboard will stay. for the first run, add data `prometheus` data source by using `http://monit_prometheus_1` address

## Troubleshooting 

- Run docker-compose without -d option, and check the log there.

### Known Issue

- when run on ubuntu server as a root, you need to manually create grafana-persist folder `/grafana-persist` and add the permission using `chmod 777 ./grafana-persist`
