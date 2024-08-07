package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

/*
type displayer interface {
	Display()
}
*/

/*
Option 1:
	type outputtable interface {
		Save() error
		Display()
	}

Option 2:
	type outputtable interface {
		saver
		Display()
	}
*/

type outputtable interface {
	saver
	Display()
}

/* Option 1: Nhuoc diem khong xac dinh duoc kieu du lieu dau ra
func Add(a, b any) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}
	return nil
}
*/

func Add[T int | float64 | string](a, b T) T {
	return a + b
}
func OutputData(data outputtable) error {
	data.Display()
	return SaveData(data)
}

func SaveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Save failed")
		return err
	}
	fmt.Println("Save successed")
	return nil
}
func GetInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func GetTodoData() string {
	text := GetInput("Todo text: ")
	return text
}

func GetNoteData() (string, string) {
	title := GetInput("Note title: ")

	content := GetInput("Note content: ")

	return title, content
}
