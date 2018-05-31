package main

import (
    "encoding/json"
    "log"

    "github.com/fatih/color"
    "github.com/parnurzeal/gorequest"
)

type testCase struct {
    I    uint64     `json:"i"`
}

func main() {
    var body string
    var errs []error
    url := `http://localhost:8080/`
    request := gorequest.New()
    
    for i:=(^uint64(0)); i>=uint64(0); i-=1 {
        var in testCase
        in.I = i

         _, body, errs = request.Post(url).Send(in).End()

        if errs != nil {
            color.Set(color.FgRed)
            log.Panicln(errs)
            color.Unset()
        } else {
            var out testCase
            json.Unmarshal([]byte(body), &out)
            if out.I != in.I {
                log.Panicln(color.RedString("Fail"), `:`, in.I,
                            "\n" + color.RedString("expected: "), in.I, 
                            "\n" + color.RedString("return: "), out.I)
            } else {
                log.Println(color.HiGreenString("Pass"), `:`, in.I)
            }
        }
    }
}