set -v
docker ps -q | xargs docker kill
docker ps -aq | xargs docker rm
