package main

import "log"
import "github.com/parnurzeal/gorequest"
import "github.com/fatih/color"

type testCase struct {
    desc string
    method string
    url string
    params string
    expected string
}

func main() {
    var body string
    var errs []error
    request := gorequest.New()
    
    testCases := create_testcases()

    for _, tCase :=  range testCases {
        if tCase.method == `GET` {
            _, body, errs = request.Get(tCase.url).End()
        } else {
            _, body, errs = request.Post(tCase.url).Send(tCase.params).End()
        }

        if errs != nil {
            color.Set(color.FgRed)
            log.Panicln(errs)
            color.Unset()
        } else if body != tCase.expected {
            log.Println(color.RedString("Fail"), tCase.desc+`:`, tCase.method, tCase.url, tCase.params +
                        "\n" + color.RedString("expected: " + tCase.expected) + 
                        "\n" + color.RedString("return: " + body))
        } else {
            log.Println(color.HiGreenString("Pass"), tCase.desc+`:`, tCase.method, tCase.url, tCase.params)
        }
    }
}

func create_testcases() []testCase {
    var testCases []testCase
    return append(testCases,
        testCase{desc: `basic_test`,
                method: `GET`,
                url: `http://localhost:8080/test`,
                expected: `{"token":"allumetesttoken","command":"allumetestcommand","results":[{"foo":"bar"}]}` },
        testCase{desc: `basic_testcypher`,
                method: `GET`,
                url: `http://localhost:8080/testcypher`,
                expected: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6ImFsbHVtZXRlc3R0b2tlbiIsImNvbW1hbmQiOiJhbGx1bWV0ZXN0Y29tbWFuZCIsInJlc3VsdHMiOlt7ImZvbyI6ImJhciJ9XX0.sN1U3_A4gLUYbrxYnalgcW-AuTO7gmJSX8AcBldrDco` },
}