package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lilleyz7/fileTerminal/types"
	"github.com/rivo/tview"
)

var application *tview.Application
var currDir *types.DirectoryStorage
var list *tview.List
var alphabet = []rune{'a', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func UpdateCurrentDirectory(directoryName string) {
	var path string
	if currDir == nil {
		path = directoryName
	} else {

		path = filepath.Join(currDir.Path, directoryName)
		fmt.Println(path)
	}

	interiorComponents, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	if len(interiorComponents) > 24 {
		interiorComponents = interiorComponents[:23]
	}
	temp := currDir
	currDir = types.NewDirectoryStorage(path, interiorComponents, temp.PreviousDir)
}

func GoToPreviousDirectory() {
	if currDir.PreviousDir == nil {
		return
	} else {
		currDir = currDir.PreviousDir

	}
	ShowScreen()
}

func UpdateList() *tview.List {

	if list.GetItemCount() != 0 {
		list.Clear()
	}

	list.AddItem("Search", "Search for a filename within the current directory", 's',
		func() {
			//Search for file
		})

	for idx, val := range currDir.InnerComponents {
		if val.IsDir() {
			list.AddItem(val.Name(), strconv.Itoa(idx), alphabet[idx], func() {
			})
		} else {
			valInfo, err := val.Info()
			if err != nil {
				list.AddItem("Invalid", "Invalid", 'q', func() {
					application.Stop()
				})
			}
			text := `
			SIZE: %d KB -----
			LAST MOD: %s -----
			`
			list.AddItem(val.Name(), fmt.Sprintf(text, valInfo.Size(), valInfo.ModTime().String()), 'f',
				func() {
				})
		}
	}

	if currDir.PreviousDir != nil {
		list.AddItem("Back", "Return to previous dir", 'b', func() {
			GoToPreviousDirectory()
		})

	}

	list.SetSelectedFunc(HandleResponse)

	return list

}

func ShowScreen() {
	newList := UpdateList()
	if err := application.SetRoot(newList, true).SetFocus(newList).Run(); err != nil {
		panic("This is fucked")
	}

}

func HandleResponse(index int, mainText string, secondaryText string, shortcut rune) {
	if currDir.InnerComponents[index].IsDir() {
		UpdateCurrentDirectory(mainText)
		ShowScreen()
	} else if mainText == "Back" && currDir.PreviousDir != nil {
		GoToPreviousDirectory()
	} else if mainText == "Back" {
		fmt.Println("IM SLOW")
	} else if mainText == "search" {
		//
	}
}

func root() string {
	return os.Getenv("SystemDrive") + string(os.PathSeparator)
}

func InitializeGlobals() {
	currDir = &types.DirectoryStorage{}
	UpdateCurrentDirectory(root())
	list = tview.NewList()
	application = tview.NewApplication()

}

func main() {
	InitializeGlobals()
	ShowScreen()

}
