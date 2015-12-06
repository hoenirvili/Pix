build:
	go build
	./Pix --path mecha.txt --env envelope.xml --save newRequest.xml -0 -1 -2
clean:
	rm -rf newRequest.xml
