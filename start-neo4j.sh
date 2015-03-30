#!/bin/sh

DIR="neo4j-community-2.2.0"
FILE="$DIR-unix.tar.gz"

wget "http://dist.neo4j.org/$FILE"
tar zxf $FILE
sed -i "s/auth_enabled\=true/auth_enabled\=false/g" $DIR/conf/neo4j-server.properties
$DIR/bin/neo4j start
sleep 3
