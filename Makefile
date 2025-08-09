## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run/backend: run the backend of the application
run/backend:
	cd ./backend && go run ./cmd/moxer

## run/frontend: run the frontend of the application
run/frontend:
	cd ./frontend && npm run dev