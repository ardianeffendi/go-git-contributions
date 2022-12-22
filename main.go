package main

import (
	"flag"
	"fmt"
)

// scan traverses a folder given by path and its subfolders
func scan(path string) {
    fmt.Printf("Found folders:\n\n")
    repositories := recursiveScanFolder(path)
    filePath := getDotFilePath()
    addNewSliceElementsToFile(filePath, repositories)
    fmt.Printf("\n\nAdded Successfully!\n\n")
}


// recursiveScanFolder starts the recursive search of git repositories
// in the `path` subtree
func recursiveScanFolder(path string) []string {
    return scanGitFolders(make([]string, 0), path)
}


// stats generates a graph of given email Git contributions
func stats(email string) {
    print("stats")
}


func main() {
    var folder string
    var email string
    flag.StringVar(&folder, "add", "", 
        "add a new folder to scan for Git repositories")
    flag.StringVar(&email, "email", "your@email.com", "the email to scan")
    flag.Parse()

    if folder != "" {
        scan(folder)
        return
    }

    stats(email)
}
