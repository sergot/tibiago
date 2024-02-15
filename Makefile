.PHONY: backup load resetdb

backup:
	docker exec -it tibiago-db-1 mariadb-dump --no-create-info -uroot -pmariadb tibiago > db/data.sql

load:
	cat db/data.sql | docker exec -i tibiago-db-1 mariadb -uroot -pmariadb tibiago

resetdb:
	docker exec -it tibiago-db-1 mariadb -uroot -pmariadb -e "DROP DATABASE IF EXISTS tibiago; CREATE DATABASE tibiago;"
