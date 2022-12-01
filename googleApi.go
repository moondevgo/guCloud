package guCloud

// [Updating Google Sheets using Golang.](https://dev.to/mediocredevops/playing-with-google-sheets-api-using-golang-14en)
// [Sheets API Guides Go quickstart](https://developers.google.com/sheets/api/quickstart/go)

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/moondevgo/guBasic"
)

var Scopes = map[string][]string{
	"sheets": []string{"https://www.googleapis.com/auth/spreadsheets"},
	"keep":   []string{"https://www.googleapis.com/auth/keep", "https://www.googleapis.com/auth/keep.readonly"},
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getGoogleJsonPath(folder, nick, authType string) string {
	// fmt.Println(folder + "google_" + authType + "_" + nick + ".json")
	return folder + "google_" + authType + "_" + nick + ".json"
}

// api string *http.Client
func ApiClient(api_name, bot_nick, user_nick string) *http.Client {
	// folder := guBasic.GetConfigFolder()
	// folder := guBasic.GetConfigFolder("CLOUD", "GOOGLE", "CONFIG")

	folder := `C:\MoonDev\withLang\inGo\goUtils\_config\`
	nick := bot_nick
	authType := "bot"
	if user_nick != "" {
		nick = user_nick
		authType = "user"
	}
	path := getGoogleJsonPath(folder, nick, authType)
	// path := folder + "google_" + authType + "_" + nick + ".json"

	// // b, err := ioutil.ReadFile("google_bot_moonsats.json")
	// b, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	b := guBasic.GetConfigBuf(path)

	var client *http.Client
	if authType == "bot" {
		config, err := google.JWTConfigFromJSON(b, Scopes[api_name]...)
		fmt.Printf("config: %T", config)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		client = config.Client(oauth2.NoContext)
	} else if authType == "bot" {
		config, err := google.ConfigFromJSON(b, Scopes[api_name]...)
		fmt.Printf("config: %T", config)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		// client = config.Client(oauth2.NoContext)
		client = getClient(config)
	}

	return client
}
