BINARY_NAME=MarkDown.exe
APP_NAME=MarkDown
VERSION=1

##build	bin	and	pakcage	app
build:
	rm	-rf	${BINARY_NAME}
	rm	-f	fyne-md
	fyne	package	-os	windows	icon	icon.png

##run:	builds	and	runs	the	application
run:
	go	run	.

##clean	runs	go	clean	and	deletes	binaries
clean:
	@echo	"Cleaning..."
	@go	clean
	@rm	-rf	${BINARY_NAME}
	@echo	"Cleaned!"
#test	runs	all	tests
test:
	go	test	-v	./...