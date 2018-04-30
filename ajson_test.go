package ajson

import (
	"encoding/json"
	"errors"

	"fmt"
	"reflect"
	"testing"

	"github.com/fatih/structs"
)

type post struct {
	ID      int64
	Title   string
	Authors []author
}

type author struct {
	ID   int64
	Name string
}

type wrapPost struct {
	ID      int64           `json:"id"`
	Title   string          `json:"title"`
	Authors json.RawMessage `json:"authors"`
}

type wrapAuthor struct {
	ID   json.Number `json:"id"`
	Name string      `json:"name"`
}

func (obj *wrapPost) DecodeRawMessage(data []byte, scfield *structs.Field, swfield *structs.Field) error {
	switch scfield.Value().(type) {
	case []author:
		var raw []json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return err
		}
		scval := make([]author, len(raw))
		for k, one := range raw {
			if err := DecodeObject(one, &wrapAuthor{}, &scval[k]); err != nil {
				return err
			}
		}
		scfield.Set(scval)
	}
	return nil
}

func (obj *wrapAuthor) DecodeNumber(data json.Number, scfield *structs.Field, swfield *structs.Field) error {
	switch scfield.Value().(type) {
	case int64:
		var val int64
		var err error
		if val, err = data.Int64(); err != nil {
			return errors.New("author.ID not int64")
		}
		scfield.Set(val)
	}
	return nil
}

func TestAst(t *testing.T) {
	demo := make(map[string]string)
	// demo.posts = `[{"id":1,"title":"heyha","authors":[{"id":2,"name":"abos"},{"id":3,"name":"freeman"}]}]`
	demo["post"] = `{"id":1,"title":"heyha","authors":[{"id":2,"name":"abos"},{"id":3,"name":"freeman"}]}`

	post := post{}
	err := DecodeObject([]byte(demo["post"]), &wrapPost{}, &post)
	if err != nil {
		panic(err)
	}

	fmt.Println("-----")
	la := len(post.Authors)
	if la > 0 {
		fmt.Print("Ques ID: ", reflect.TypeOf(post.Authors[0].ID), " ")
	}
	// fmt.Println(post)
	fmt.Print(reflect.TypeOf(post), "")
	fmt.Print(post, "")
	fmt.Print(reflect.TypeOf(post.Authors), la, "")

	t.Fail()
}
