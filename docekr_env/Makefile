.PHONY: env_up
server_env_up:
	docker compose -f docker-compose.yml up -d
	mysql -h127.0.0.1 -uroot -p -P3326 <./carbon.sql
	mysql -h127.0.0.1 -uroot -p -P3326 <./pay.sql
	mysql -h127.0.0.1 -uroot -p -P3326 <./trade.sql

.PHONY: env_down
server_env_down:
	docker compose -f docker-compose.yml down