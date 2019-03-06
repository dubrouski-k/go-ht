package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ParseArgs()Arguments{
	var operation = flag.String("operation", "", "operation")
	var item = flag.String("item", "", "item")
	var fileName = flag.String("fileName", "", "fileName")
	var id = flag.String("id", "", "id")
	args := Arguments{
		"id":        *id,
		"operation": *operation,
		"item":      *item,
		"fileName":  *fileName,
	}
	return args
}

func StoreToFile(filename string, persons []Person) error {
	bytesS, err := json.Marshal(persons)
	if err != nil {
		return err
	}

	os.Remove(filename)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	//file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(string(bytesS))
	if err != nil {
		return err
	}
	//err = file.Sync()
	//if err != nil {
	//	return err
	//}
	return nil
}


func Perform2(_ Arguments, _ io.Writer) error {
	return fmt.Errorf("Operation %s not allowed!", "abcd")
}

func Perform3(_ Arguments, _ io.Writer) error {
	return nil
}

func Perform(args Arguments, writer io.Writer) error {
	var err error
	opName := args["operation"]
	if opName == "" {
		return errors.New("-operation flag has to be specified")
	}
	fileName := args["fileName"]
	if fileName == "" {
		return errors.New("-fileName flag has to be specified")
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil{
		return err
	}
	bytesR, err := ioutil.ReadAll(f)
	err = f.Close()
	if err != nil{
		return err
	}
	persons := make([]Person,0)
	if len(bytesR) != 0 {
		err = json.Unmarshal(bytesR, &persons)
		if err != nil {
			return err
		}
	}

	//var item string
	item := args["item"]
	operatedId := args["id"]

	switch opName {
	case "add":
		{
			if item == "" {
				return errors.New("-item flag has to be specified")
			}
			var person Person
			err = json.Unmarshal([]byte(item), &person)
			if err != nil {
				return err
			}
			foundSame := false
			for _, value := range persons {
				if value.Id == person.Id {
					errString := `Item with id ` + person.Id + ` already exists`
					//var someBuffer bytes.Buffer
					//someBuffer.WriteString("Item with id ")
					//someBuffer.WriteString(person.Id)
					//someBuffer.WriteString(" already exists")
					writer.Write([]byte(errString))
					foundSame = true
				}
			}
			if !foundSame {
				persons = append(persons, person)
				if err = StoreToFile(fileName, persons); err != nil {
					return err
				}
			}
		}
	case "list":
		{
			_, err = writer.Write(bytesR)
			if err != nil {
				return err
			}
		}
	case "findById":
		{
			if operatedId == "" {
				return errors.New("-id flag has to be specified")
				//return errors.New("-id flag has to be specified")
			}
			for _, value := range persons {
				if value.Id == operatedId {
					bytesW, err := json.Marshal(value)
					if err != nil {
						return err
					}
					_, err = writer.Write(bytesW)
					if err != nil {
						return err
					}
				}
			}
		}
	case "remove":
		{
			if operatedId == "" {
				return errors.New("-id flag has to be specified")
			}
			newPersons := make([]Person, 0)
			passed := false
			for _, value := range persons {
				if value.Id != operatedId {
					newPersons = append(newPersons, value)
				}
				if value.Id == operatedId {
					passed = true
				}
			}
			if !passed {
				errString := "Item with id " + operatedId + " not found"
				_, err = writer.Write([]byte(errString))
				if err != nil {
					return err
				}
			}
			if passed {
				if err = StoreToFile(fileName, newPersons); err != nil {
					return err
				}
			}
		}
	default:
		//errString := "Operation " + opName + " not allowed!"
		return fmt.Errorf("Operation %s not allowed!", opName)
		//panic(errors.New(errString))
		//panic(errors.New("Operation " + opName + " not allowed!"))
		//return errors.New(`Operation abcd not allowed!`)
		//return errors.New(errString)
	}
	return nil
}


type Arguments map[string]string

type Person struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Age int `json:"age"`
}


func main() {
	err := Perform(ParseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
