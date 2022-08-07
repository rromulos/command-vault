cv/a:
	@./main -a ='$(args)'
cv/l:
	@./main -l
cv/d:
	@./main -d=$(id)
cv/scmd:
	@./main -scmd ='$(args)'
cv/scat:
	@./main -scat ='$(args)'
cv/sdes:
	@./main -sdes ='$(args)'		