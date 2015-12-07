#!bin/bash
curl	-H "Content-Type: text/xml;charset=UTF-8" -H "parseText_TXT" --data-binary @../frame/postTaggerFrame.xml -X POST http://nlptools.infoiasi.ro:80/WebPosRo/PosTaggerRoWS
