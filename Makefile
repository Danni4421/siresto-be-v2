auth-secret:
	@echo "" >> .env
	@echo "AUTH_SECRET=\"$(shell openssl rand -base64 32 | tr -d '\n')\"" >> .env
	@echo "AUTH_SECRET generated and saved to .env"