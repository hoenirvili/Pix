build:
	go build
	./Pix --path text/mecha.txt --env frame/discourseParser.xml --save req6.xml -6
clean:
	rm -rf newRequest.xml
