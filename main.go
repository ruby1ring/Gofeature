package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func regSentence(sentence string) {
	splitSentence:=regexp.MustCompile("[,|，]{1}").Split(sentence,-1)
	sentenceFlag:=make([]bool,len(splitSentence))
	file,_:=os.Open("./reg.txt")
	sc:=bufio.NewScanner(file)
	for sc.Scan(){
		sentence := strings.Split(sc.Text(),":")
		for _,v:=range sentence{
			feature:=strings.Split(v," ")
			featureSorce,_:=strconv.ParseFloat(feature[1],64)
			for index,sent:=range splitSentence{
				if len(regexp.MustCompile(strings.Replace(feature[0],".",".*",-1)).Find([]byte(sent)))!=0 && sentenceFlag[index]==false{
				fmt.Println(string(regexp.MustCompile(strings.Replace(feature[0],".",".*",-1)).Find([]byte(sent))),featureSorce)
				sentenceFlag[index]=true
			} 
			}
		}
	}
}







