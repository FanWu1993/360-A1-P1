# compiler settings
CC=go build

# Binaries
BIN=starwar

# the default target
all: 
	${CC} -o $(BIN)

# tests
test:	test0.out test1.out test2.out test3.out test4.out test5.out

# construct a .out file from a .in file and check it
SED=/usr/bin/sed
DIFF=/usr/bin/diff
%.out:	$(BIN)
		$(SED) -e 's/[ 	]*#.*$$//' <$*.in  >$*.in~  2>/dev/tty   
		$(SED) -e 's/[ 	]*#.*$$//' <$*.exp >$*.out~ 2>/dev/tty  
		./$(BIN)<$*.in~ >$*.out 2>/dev/tty
		./$(BIN) <$*.in~ 2>/dev/tty | tee $*.out
		$(DIFF) -q $*.out $*.out~

# cleaning
clean:
	-$(RM) -rf *.out $(BIN) *~

