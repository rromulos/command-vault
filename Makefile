cv/a:
	@go run ./cmdvault/cmd -a =$(args)
cv/l:
	@go run ./cmdvault/cmd -l
cv/d:
	@go run ./cmdvault/cmd -d=$(id)