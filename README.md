# greeter
>Send automated birthday greeting!

Remembering people's birthday and greeting them have never been easier.
This is a golang package to automate birthday greeting.

## prerequisite

Make sure you have [Golang](https://golang.org/) installed and API KEY from [SENDGRID](https://sendgrid.com/).

## install
`go get github.com/vikashvverma/greeter`

Create a file named `config.json` in `$GOPATH/bin` as follows:
```
{
  "time": "20:02",
  "apiKey": "API_KEY",
  "subject": "Happy B'day!"
}
```
Replace the `API_KEY` with your sendgrid `API_KEY`, change the `time` or `subject` if you want.
 The greeting will be sent at specified `time` as per the server time.








