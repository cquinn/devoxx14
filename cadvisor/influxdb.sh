#https://github.com/cquinn/tutum-docker-influxdb
#https://registry.hub.docker.com/u/tutum/influxdb/

docker run -d -p 8083:8083 -p 8086:8086 --expose 8090 --expose 8099 --name="influxdb" tutum/influxdb
