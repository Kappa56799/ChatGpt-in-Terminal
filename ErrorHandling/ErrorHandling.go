package ErrorHandling

import (
  "fmt"
  "os"
)

// If an Error Occurs in the OCR function, it will print the error message and exit the program
func OCRError(err error) {
  if err != nil {
    fmt.Println("An Error Occured in OCR: ", err)
    os.Exit(2)
  }
}

// A Prompt incase the user enters the wrong input
func CorrectInput() {
  fmt.Fprintln(os.Stderr, "usage: ./<filename> <flag> <Your Prompt> or ./<filename> \nflags: -c (Code) -m (Full Message) -o (OCR from Picture)")
}

// Checks if the APIKey is set, if not, prints the error message and exits the program
func CheckAPIKey(apiKey string) {
    if apiKey == "" {
      fmt.Println("Please set the OPENAI_API_KEY environment variable or Set APIKey")
      os.Exit(2)
    }
}

// If an Error Occurs when it sends a request, it will print the error message and exit the program
func RequestError(err error) {
  if err != nil {
    fmt.Println("An Error Occured in Request: ", err)
    os.Exit(2)
  }
}

// If an Error Occurs in while unmarshaling JSON, it will print the error message and exit the program
func JSONError(err error) {
  if err != nil {
    fmt.Println("An Error Occured in while unmarshaling JSON: ", err)
    os.Exit(2)
  }
}

// If an Error Occurs when the GPT query is sent, it will print the error message and exit the program
func GPTError(data map[string]interface{}) {
  if errDetails, ok := data["error"].(map[string]interface{}); ok {

    fmt.Println("Error detected:")
    fmt.Println("Message:", errDetails["message"])
    fmt.Println("Type:", errDetails["type"])
    fmt.Println("Param:", errDetails["param"])
    fmt.Println("Code:", errDetails["code"])
    os.Exit(2)
  }
}
