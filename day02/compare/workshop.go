package compare

import (
	"bytes"
	"encoding/json"
)

type Post struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`    
	Title  string `json:"title"` 
	Body   string `json:"body"`  
}

func V1() []byte {
	p := Post{}
	b, _ := json.Marshal(p)
	return b
}

func V2() {
	p := Post{}
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	enc.Encode(p)
}