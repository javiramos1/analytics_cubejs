run:
	go run .

db:
	docker-compose up -d
	sleep 3s
	docker cp ddl.sql postgres:/usr/local
	docker exec postgres psql -h localhost -d postgres -U postgres -p 5432 -a -q -f /usr/local/ddl.sql

stop:
	docker-compose down



