run-app:
	cd backend && go run cmd/main/main.go
	cd frontend && npm run dev

run-db:
	docker run --name name-of-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -p 5987:5432 -d postgres:14