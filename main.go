package main

import (
	"fmt"
	note "practice_section5/Note"
	todo "practice_section5/Todo"
	user "practice_section5/User"
)

func main() {
	title, content := user.GetNoteData()
	todoText := user.GetTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*	userNote.Display()
		err = todo.Save()
		if err != nil {
			fmt.Println("Saving the Todo failed")
			return
		}
		fmt.Println("Saving the Todo successed")
	*/
	//user.SaveData(userNote)
	//user.SaveData(todo)
	result1 := user.Add(2, 5)
	result2 := user.Add(1.5, 2.6)
	result3 := user.Add("Hello", "Halo")
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	err = user.OutputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.OutputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}
