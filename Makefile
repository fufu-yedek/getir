build-docker:
	@echo 'Building project'
	@docker build --no-cache -t getir-challange .

	@echo 'Running docker project'
	@docker run -p 6060:8080 --name getir-challange --rm getir-challange

run-tests:
	@echo 'Building project'
	@docker build -f Dockerfile.test --no-cache -t getir-challange-tests .

	@echo 'Running tests'
	@docker run --name getir-challange-tests --rm getir-challange-tests


heroku-deploy:
	@heroku login
	@heroku container:login
	@heroku container:push web -a fufu-getir-challange
	@heroku container:release web -a fufu-getir-challange

heroku-logs:
	@heroku login
	@heroku logs --tail -a fufu-getir-challange