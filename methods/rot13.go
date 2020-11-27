package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (r13 rot13Reader) Read(buffer []byte) (int,error){
	n,err := r13.r.Read(buffer)
	for i:=0;i<n;i++{
		if (buffer[i]>='A' && buffer[i]<='Z'){
			if (buffer[i]+13 > 'Z'){
			buffer[i] = 'A'-1 + buffer[i]+13-'Z'
		}else{buffer[i]=buffer[i]+13}
		
		
	
		}else if(buffer[i]>='a' && buffer[i]<='z'){ 
			if (buffer[i]+13 > 'z'){
			buffer[i] = 'a'-1 + buffer[i]+13-'z'
		}else{buffer[i]=buffer[i]+13}
		
	}else{}
	}
	
	return len(buffer),err
}

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
