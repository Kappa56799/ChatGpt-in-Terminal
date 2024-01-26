/*
Author: Kacper Palka palka@kacper.boo
Description: This is a simple CLI tool that uses OpenAI's GPT-3 API to generate code or text based on the prompt you provide.
*/

package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "strings"
  "bufio"

  //Libraries used to achieve loading animation and rest client
  "github.com/briandowns/spinner"
  "github.com/go-resty/resty/v2"
)

// API endpoint
const ( 
  apiEndpoint = "https://api.openai.com/v1/chat/completions" 
  flagCode    = "-c"
  flagMessage = "-m"
)

func main() {
    // Check if the user provided the correct input
    checkInput()
    
    // Prints the generated code or text based on the input provided by the user
    SelectMode()
}

// Check if the user provided the correct input and if not, print the error message
func checkInput() {
    // Checks if number of arguments is 3 and if the third argument is not empty and if there is only 2 arguments and the second one is not empty
    if len(os.Args) == 3 && os.Args[2] == "" || len(os.Args) == 2 && os.Args[1] != ""{
      fmt.Fprintln(os.Stderr, "usage: ./<filename> <flag> <Your Prompt> or ./terGPT \nflags: -c (Code) -m (Full Message) ")
      os.Exit(2)
    }
}

// Selects the mode based on the input provided by the user
func SelectMode() {

  // If there is no arguments provided by the user, the program will ask for the input and print the generated text
  if len(os.Args) == 1 {
    // Allows the user to keep inputting the prompt until the user exits the program
    fmt.Println("--------------------------Please enter your prompt (CTRL + C to Quit)--------------------------")
    for true {
      inputText := ScanInput()
      content := GPTQuery(inputText)
      fmt.Println(content + "\n")
    }
  } else if os.Args[1] == flagCode { // If the user provided the flag -c, the program will print the generated code
      input := os.Args[2]
      content := GPTQuery(input)
      CodeContent := strings.Join(extractContent(content), "") // Changes the array of strings into one string
      fmt.Println(CodeContent)
  } else if os.Args[1] == flagMessage { // If the user provided the flag -m, the program will print the generated text and code (Full Message)
      input := os.Args[2]
      content := GPTQuery(input)
      fmt.Println(content)
  } else { // If the user provided the wrong flag, the program will print the error message
      fmt.Fprintln(os.Stderr, "usage: ./<filename> <flag> <Your Prompt> or ./terGPT \nflags: -c (Code) -m (Full Message) ")
  }

}

// Scans for the input provided by the user from the terminal and returns it
func ScanInput() (string){
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()

  inputText := scanner.Text()

  return inputText
}

// Extracts only the code partt from the generated text by splitting it and returning it
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

	return formattedParagraphs
}

// Sends the request to the API and returns the generated text
func GPTQuery(input string) (string){
    // Gets the API key from the environment variable (You can change name if needed)
    apiKey := os.Getenv("OPENAI_API_KEY")

    if apiKey == "" {
      fmt.Println("Please set the OPENAI_API_KEY environment variable")
      os.Exit(2)
    }

  // Data variable is used to store the JSON response from the API
    var data map[string]interface{}
  // Loading animation
    loading := spinner.New(spinner.CharSets[33], 100*time.Millisecond)

    loading.Start() // Start the loading animation
    // Create a new rest client and send the request to the API
    client := resty.New()
    response, err := client.R().
                    SetAuthToken(apiKey).
                    SetHeader("Content-Type", "application/json").
                    SetBody(map[string]interface{}{"model": "gpt-3.5-turbo",
                                                  "messages":   []interface{}{map[string]interface{}{"role": "assistant", "content": input}},
                                                  "max_tokens": 500,
    }).
    Post(apiEndpoint)

    loading.Stop() // Stop the loading animation

    // If there is an error with the request, print the error message and exit the program
    if err != nil {
      fmt.Printf("Error in API request. Status Code: %d\n", response.StatusCode())
      fmt.Println("Response:", response.String())
      os.Exit(2)
    } 
    
    // Stores the response body into the body variable
    body := response.Body()
    
    // Decodes the JSON response into the data variable using marshal
    err = json.Unmarshal(body, &data)

    // If there is an error with decoding the JSON response, print the error message and exit the program
    if err != nil {
        fmt.Println("Error while decoding JSON response:", err)
        os.Exit(2)
    }

    // If there is an error in the request, print the error message and exit the program
    if errDetails, ok := data["error"].(map[string]interface{}); ok {

      fmt.Println("Error detected:")
      fmt.Println("Message:", errDetails["message"])
      fmt.Println("Type:", errDetails["type"])
      fmt.Println("Param:", errDetails["param"])
      fmt.Println("Code:", errDetails["code"])
      os.Exit(2)
    } 

    // Stores the generated text into the content variable
    content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

    return content
}


