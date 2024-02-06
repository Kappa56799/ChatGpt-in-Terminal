/*
Author: Kacper Palka palka@kacper.boo
Description: This is a simple CLI tool that uses OpenAI's API to generate code or text based on the prompt you provide.
*/

package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "strings"
  "bufio"
  "os/exec"
  "ChatGpt-in-Terminal/ErrorHandling"

  //Libraries used to achieve loading animation, rest client and OCR
  "github.com/briandowns/spinner"
  "github.com/go-resty/resty/v2"
  "github.com/otiai10/gosseract/v2"

)

// API endpoint
const ( 
  apiEndpoint = "https://api.openai.com/v1/chat/completions" 
  flagCode    = "-c"
  flagMessage = "-m"
  flagOCR     = "-o"
  Clipboard   = "wl-paste"
)

var (
  flag string
  message string
  data map[string]interface{}

  // Gets the API key from the environment variable (You can change name if needed)
  apiKey = os.Getenv("OPENAI_API_KEY")
)

func main() {

  // If there are arguments provided by the user, the program will use them as the input
	if len(os.Args) > 1 {
		flag = os.Args[1]
	}
	if len(os.Args) > 2 {
		message = os.Args[2]
	}

  // If there is no arguments provided by the user, the program will ask for the input and print the generated text
  if len(os.Args) == 1 {
    // Allows the user to keep inputting the prompt until the user exits the program
    fmt.Println("--------------------------Please enter your prompt (CTRL + C to Quit)--------------------------")

    for true {
      inputText := ScanInput()
      content := GPTQuery(inputText)
      fmt.Println(content + "\n")
    }

  } else if flag == flagCode && message != "" { 
      // If the user provided the flag -c, the program will print only the generated code
      content := GPTQuery(message)
      CodeContent := strings.Join(extractContent(content), "") // Changes the array of strings into one string
      fmt.Println(CodeContent)

  } else if flag == flagMessage && message != "" { 
      // If the user provided the flag -m, the program will print the generated text and code (Full Message)
      content := GPTQuery(message)
      fmt.Println(content)

  } else if flag == flagOCR {
      // If the user provided the flag -o, the program will use OCR to extract the text from the picture and use that in the prompt
      client := gosseract.NewClient()
      defer client.Close() // waits until everything is done and then closes the client
      // Uses the clipboard to get the picture and then uses OCR to extract the text from the picture
      cmd := exec.Command(Clipboard)
      clipboardImage, err := cmd.Output()

      // If there is an error with the OCR, print the error message and exit the program
      ErrorHandling.OCRError(err)

      // Sets the image from the bytes and then gets the text from the picture
      client.SetImageFromBytes(clipboardImage)
      text, _ := client.Text()

      // Queries the API with the text extracted from the picture and prints the generated text
      content := GPTQuery(text + " " + message)
      fmt.Println(content)

  } else { // If the user provided the wrong flag, the program will print the error message
      ErrorHandling.CorrectInput()
  }

}

// Scans for the input provided by the user from the terminal and returns it
func ScanInput() string {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

// Extracts only the code part from the generated text by splitting it and returning it
func extractContent(input string) []string {

  // Splits the input into paragraphs
  paragraphs := strings.Split(input, "```")
  formattedParagraphs := []string{}

  // If there code generated without the ``` tags, return the input
  if strings.Contains(input, "```") == false {
    return paragraphs
  }

  // Joins the paragraphs with the code together by appending them to an array
	for i := 1; i < len(paragraphs); i += 2 {
    formattedParagraphs = append(formattedParagraphs, paragraphs[i])
	}

  // Returns the array of strings
	return formattedParagraphs
}

// Sends the request to the API and returns the generated text
func GPTQuery(input string) string {
    
    //Checks if apiKey is set, if not, prints the error message and exits the program
    ErrorHandling.CheckAPIKey(apiKey)

    // Loading animation
    loading := spinner.New(spinner.CharSets[33], 100*time.Millisecond)

    loading.Start() // Start the loading animation
    // Create a new rest client and send the request to the API
    client := resty.New()
    response, err := client.R().
                    SetAuthToken(apiKey).
                    SetHeader("Content-Type", "application/json").
                    SetBody(map[string]interface{}{"model": "gpt-4",
                                                  "messages":   []interface{}{map[string]interface{}{"role": "assistant", "content": input}},
                                                  "max_tokens": 1000,
    }).
    Post(apiEndpoint)

    loading.Stop() // Stop the loading animation

    // If there is an error with the request, print the error message and exit the program
    ErrorHandling.RequestError(err)
    
    // Stores the response body into the body variable
    body := response.Body()
    
    // Decodes the JSON response into the data variable using marshal
    err = json.Unmarshal(body, &data)

    // If there is an error with decoding the JSON response, print the error message and exit the program
    ErrorHandling.JSONError(err)

    // If there is an error in the request, print the error message and exit the program
    ErrorHandling.GPTError(data)

    // Stores the generated text into the content variable
    content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

    // Returns the generated text
    return content
}


