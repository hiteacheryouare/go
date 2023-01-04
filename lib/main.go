package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

const defaultPort = "3000"
const dmv1HTML = `
			<!--
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWWWWWMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMNKOxdolcc::,,,,;;cxNMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMWKxc'..    .  ......   lNMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMNOo;.  .................  .OMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMN0o,. .........  .........  ..;0WMMMMMMW0ONMMMMMM
MMMMMMMMMNk;. ............   ........ 'xo..lO000Oxc.'OMMMMMM
MMMMMMMMKc. .....  .................  l0xoc,,,,,.   :XMMMMMM
MMMMMMM0, ......   ................  ;0O,,ONXKKK0d' cNMMMMMM
MMMMMMNc .............   ........  .l0NX0OXWWNWWW0,.xWMMMMMM
MMMMMMk. ...  ...............   .'o0NXo;o0XNNNXNO;.lXMMMMMMM
MMMMMNc ....  ........... ..,:cokXNNWNk;'',,;,,,..dNMMMMMMMM
MMMMM0' ............  .':ok0XNWWWWWWWWNNKOko:. .:0WMMMMMMMMM
MMMMM0' .......  ..,lx0XNWWWWWWWWWWWWWNKkoc,. ;OWMMMMMMMMMMM
MMMMMNc    ...,:okKNWWWWWNWWWNNNNKOkdl;..,::..dWMMMMMMMMMMMM
MMMMMMXl..ckOKNNWWWNWWWNNX0kdl::,'',;:;..,;'.cKMMMMMMMMMMMMM
MMMMMMMNc.:0NNNNNNXKOxl:;,',;:loxOXNWMW0xodx0WMMMMMMMMMMMMMM
MMMMMMMM0, .,::;;,,'',:ldk0XWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMK; .,'. :xO0KNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMNx;,,,,lKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMWNXXNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM

Powerful, cross-platform, and open-source.
porkyproductions.github.io


-->
<!DOCTYPE html>
<html>
    <br>
    <image src = "data:image/jpeg;base64,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"  onclick = "GenerateInsult()"/ class = "center"> 
    <br/>
    <div  class = "center">
        <p id = "insult"></p>
        
    </div>
    <script>
        const insults = 
              [
                  "you are dog water",
                  "you bad",
                  "you have the same chance of reproducing as a computer mouse",
                  "you are the human equivalent of a participation award",
                  "There are two things I hate, you and poop. They smell really bad. The poop too.",
                  "You are about as useful as a broken clock telling the time.",
                  "At least criminals can go to jail.",
                  "Get gud at the game.",
                  "I see your future... people are celebrating ... your funeral :(",
                  "Never gonna GIVE YOU UP. Never gonna LET YOU DOWN. <br>Never gonna TURN AROUND. and DESERT YOU!",
                  "You’re the reason God created the middle finger.",
                  "I’ll never forget the first time we met. But I’ll keep trying.",
                  "Your kind of people is the reason shampoo has instructions",
                  "You are about as tasteless as an unsalted pretzel. NO insult to the pretzel",
                  "Hold still. I’m trying to imagine you with a personality.",
                  "Your face makes onions cry.",
                  "You bring a lot of happiness when you leave",
                  "aslidhfnclih, sorry I sneezed. I'm allergic to stupidity",
                  "Just opening your mouth brings down the IQ of the humanity",
                  "Trees work hard to produce oxygen for you. Please apologize to them",
                  "Punching bags must be jealous of your face",
                  "Don’t be ashamed of who you are. That’s your parents’ job.",
                  "Light travels faster than sound, which is why you seemed bright until you spoke.",
                  "If I had a face like yours, I'd sue my parents.",
                  "You're so ugly, when your mom dropped you off at school she got a fine for littering.",
                  "They say opposites attract. <br>I hope you meet someone who is good-looking, intelligent, and cultured.",
                  "If being good looking was a crime, you'd be in paradise",
                  "Your house is so dirty you have to wipe your feet before you go outside.",
                  "How did you get here? Did someone leave your cage open?",
                  "We can always tell when you are lying. Your lips move.",
                  "I have seen",
                  "Doctor Strange has seen Fourteen million, six hundred and five parallel universes <br> only one of them are you likable",
                  "As an outsider, what do you think of the human race?",
                  "If you really spoke your mind, you'd be speechless.",
                  "RING RING, Oh is that the phone. No wait. It's just you making my brain hurt.",
                  "Sorry, I only speak to people with positive IQs"
                  
                  

              ]

        function GenerateInsult(){
          insult = insults[Math.floor(Math.random() * (insults.length-0.01))];
          console.log(insult);
          document.getElementById("insult").innerHTML = insult;
        }

    </script>
    <style>
        .center {
            text-align: center;
            display: block;
            margin-left: auto;
            margin-right: auto 
        }
    </style>
<html/>
		
		`

type responseData struct {
	Data    string `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func reverseString(s string) string {
	// Convert the string to a slice of runes (characters)
	r := []rune(s)

	// Reverse the slice of runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// Return the reversed string
	return string(r)
}

func simpleJsonResponse(e, d string, c int) {
	http.HandleFunc(e, func(w http.ResponseWriter, r *http.Request) {
		var code int = c
		w.Header().Set("server", runtime.Version())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		data := responseData{
			Data: d,
			Code: c,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})
	return
}
func simpleQueryParamsJsonResponse(e, d, p string, c int) {
	http.HandleFunc(e, func(w http.ResponseWriter, r *http.Request) {
		// Get the user input from the query string
		input := r.URL.Query().Get(p)
		// Reverse the input string
		reversedInput := reverseString(input)
		var code int = c
		w.Header().Set("server", runtime.Version())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		data := responseData{
			Data:    reversedInput,
			Message: d,
			Code:    c,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})
	return
}
func simpleHTMLResponse(e, h string, c int) {
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		var code int = c
		w.Header().Set("server", runtime.Version())
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(code)
		html := h
		w.Write([]byte(html))
	})
	return
}
func redirectNotFound(e string) {
	http.HandleFunc(e, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://porkyproductions.github.io/404.html", http.StatusNotFound)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	simpleJsonResponse("/", "It works!", 200)
	simpleQueryParamsJsonResponse("/reverse", "Use data", "input", 203)
	simpleHTMLResponse("/dmv1", dmv1HTML, 200)
	redirectNotFound("/goober")
	fmt.Println("Server is listening at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
