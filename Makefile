build:
	go build
	./Pix --path mecha.txt --env frame/postTaggerFrame.xml  --save  req1.xml -0
clean:
	rm -rf newRequest.xml
