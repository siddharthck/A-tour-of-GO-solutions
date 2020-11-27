package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(buffer []byte) (int,error){
	for i := range buffer{
		buffer[i] = 'A'
	}
	return len(buffer),nil
}

func main() {
	reader.Validate(MyReader{})
}
