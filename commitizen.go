package main

import (
    "fmt"
    "bytes"
    "log"
)

func main() {
    var args Arguments
    parseArgs(&args)

    // open debug switch
    if args.debug {
        log.SetFlags(log.Lshortfile | log.LstdFlags)
    } else {
        log.SetFlags(0)
    }

    if args.install {
        if path, err := Install(); err != nil {
            log.Fatal(err)
        } else {
            fmt.Printf("Install commitizen to %s\n", path)
        }
    } else {
        // ask the question
        var answers Answers
        if err := AskForCommitMessage(&answers); err != nil {
            log.Fatal(err)
        }

        // assemble the answers to commit message 
        var buf bytes.Buffer
        answers.AssembleIntoMessage(&buf)

        // do git commit
        result, err := CommitMessage(buf.Bytes(), args.all)
        if err != nil {
            fmt.Printf("run git commit failed, commit message is: \n\n\t%s\n\n", buf.String())
            log.Fatal(err)
        }

        fmt.Print(result)
    }
}

