# MailJet

A Simple API to send mail using MailJet API.
To setup on local, need to install go environment, then clone this repository at local machine in `/go/src/` folder.
After then install `gorilla mux` and `mail jet` package then change `API key` and `secret key`  in main function on line 21, 22.

Then execute command `go build` then after run the `exe` file.
Go to the postman tool and write `localhost:9090/sendmail` with attach header as `key = token` and `value = x-api`.


