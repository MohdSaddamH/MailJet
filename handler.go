package main

import(
	"net/http"
	"log"
	"os"
	"encoding/json"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	mux "github.com/gorilla/mux"

)

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("token")
		if token != "x-api" {
			msg := map[string]string{
						"message" : "Un Authorized",
						"token" : "token is missing",
					}			
			json.NewEncoder(w).Encode(msg)
			return
		}
		
        next.ServeHTTP(w, r)
    })
}


func handler(){
	log.Println(" Handler function called")

	r := mux.NewRouter()
	r.HandleFunc("/sendmail", sendMail)
	r.Use(authMiddleware)

	log.Fatal(http.ListenAndServe(":9090", r))
}


func sendMail(w http.ResponseWriter, r *http.Request){

	// Get your environment Mailjet keys and connect
	publicKey := os.Getenv("APIKEY_PUBLIC")
	secretKey := os.Getenv("APIKEY_PRIVATE")	
	
	mailjetClient := mailjet.NewMailjetClient(publicKey, secretKey)
	messagesInfo := []mailjet.InfoMessagesV31 {
      mailjet.InfoMessagesV31{
        From: &mailjet.RecipientV31{
          Email: "pilot@mailjet.com",
          Name: "Mailjet Pilot",
        },
        To: &mailjet.RecipientsV31{
          mailjet.RecipientV31 {
            Email: "mailjet@gmail.com",
            Name: "MailJet1",
          },
        },
        Cc: &mailjet.RecipientsV31{
          mailjet.RecipientV31 {
            Email: "example2@gmail.com",
            Name: "MailJet2",
          },
        },
        Bcc: &mailjet.RecipientsV31{
          mailjet.RecipientV31 {
            Email: "example3@gmail.com",
            Name: "MailJet3",
          },
        },
        Subject: "Email testing",
        TextPart: "Dear receiver, welcome to Mailjet!",        
      },
    }
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	
	log.Println("Successfullt transffered", res)
	
	json.NewEncoder(w).Encode("Message send successfully")
}
