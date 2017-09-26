# compiler settings
CC=go build

# Binaries
BIN=starwar

# the default target
all: 
	${CC} -o $(BIN)

# tests
test:	test0.out test1.out test2.out

# construct a .out file from a .in file and check it
%.out:	${BIN} %.in %.exp
#		sed -e 's/[ 	]*#.*$$//' <$*.in  >$*.in~  2>/dev/tty   
#		sed -e 's/[ 	]*#.*$$//' <$*.exp >$*.out~ 2>/dev/tty  
		./$(BIN) <$*.in >$*.out 2>/dev/tty
#		./$(BIN) <$*.in~ 2>/dev/tty | tee $*.out
		diff -q $*.out $*.exp

# cleaning
clean:
	-$(RM) -rf *.out $(BIN) *~

