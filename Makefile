build:
	go build
	./Pix --path text/mecha.txt --env frame/postTaggerFrame.xml --save req0.xml -0
	./Pix --path text/mecha.txt --env frame/npChunkerFrame.xml --save req1.xml -1
	./Pix --path text/mecha.txt --env frame/fdgParserFrame.xml --save req2.xml -2
	./Pix --path text/mecha.txt --env frame/nameEntityRecogn.xml --save req3.xml -3
	./Pix --path text/mecha.txt --env frame/anaphoraResolution.xml --save req4.xml -4
	./Pix --path text/mecha.txt --env frame/clauseSplit.xml --save req5.xml -5
	./Pix --path text/mecha.txt --env frame/discourseParser.xml --save req6.xml -6
clean:
	rm -rf newRequest.xml
