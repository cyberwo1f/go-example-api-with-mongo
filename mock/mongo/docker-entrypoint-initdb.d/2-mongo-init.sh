mongoimport --db example --collection user --drop --file /docker-entrypoint-initdb.d/user.json --jsonArray
mongoimport --db example --collection message --drop --file /docker-entrypoint-initdb.d/message.json --jsonArray
