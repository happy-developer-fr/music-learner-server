DOCKER_NAME=postgis_music_learner
docker stop $DOCKER_NAME ; docker rm $DOCKER_NAME ; docker run --name $DOCKER_NAME -e POSTGRES_PASSWORD=admin -p 5432:5432 -v /tmp:/tmp -d postgis/postgis:12-3.0-alpine
cp ./create.sql /tmp/create.sql
echo waiting for db starting 5s...
sleep 5s
echo creating
docker exec -it $DOCKER_NAME psql -U postgres postgres -f /tmp/create.sql
