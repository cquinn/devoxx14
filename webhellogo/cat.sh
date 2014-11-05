set +v
docker cp webhellogo:/data/counter .
cat counter
echo
rm counter
