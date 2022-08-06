cv/a:
	@./main -a ='$(args)'
	# @go run ./cmdvault/cmd -a ='$(args)'
cv/l:
	@./main -l
	# @go run ./cmdvault/cmd -l
cv/d:
	@./main -d=$(id)
	# @go run ./cmdvault/cmd -d=$(id)