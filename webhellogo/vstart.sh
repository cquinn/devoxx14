set -v
docker run -d -p 9090:8080 -v /home/docker:/data --name="webhellogo" cquinn/webhellogo
