#https://registry.hub.docker.com/u/tutum/grafana/
#https://github.com/cquinn/tutum-docker-grafana

docker run -d -p 80:80 \
   -e INFLUXDB_HOST=0.0.0.0 \
   -e INFLUXDB_PORT=8086 \
   -e INFLUXDB_NAME=testdb \
   -e INFLUXDB_USER=root \
   -e INFLUXDB_PASS=root \
   -e INFLUXDB_IS_GRAFANADB=true \
   --name="grafana" \
   tutum/grafana
