package main

import (
    "flag"
)


// scan traverses a folder given by path and its subfolders
func scan(path string) {
    print("scan")
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
