# ChatGpt-in-Terminal

## Description
This is a simple chatbot that uses GPT 3.5-turbo to generate responses and code straight from the terminal only. It is written in Go and uses the OpenAI API to generate responses.
I made this because I wanted to make a chatbot that I could use in the terminal allowing me to easily output code to a file using redirects on Linux as well as learning to 
program in Go :). I tested this on Arch Linux using kitty terminal and everything worked fine including running the program using 'go run' and compiling it using 'go build'.

## Requirements
- Go 1.16 or higher 
- OpenAI API key
- Linux (Haven't tested on Windows)

# How to find your API key
Go to this website [Your API Key!](https://platform.openai.com/api-keys), log in and create a key!


## How to Install
Either download the compiled code or clone the repository using the following command:
```git clone https://github.com/Kappa56799/ChatGpt-in-Terminal```

From here follow the instructions in the usage section to run the program.

## Usage
To run the program you firstly need to set the environment variable OPENAI_API_KEY to your API key. You can do this by running the following command in the terminal or by putting this command into your shell config (fish,zsh,bash):
```bash
export OPENAI_API_KEY="<your key here>"
```
Then you can run the program using the following command:
``` go run main.go``` or ```go build -o  main.go <your output name here>```

Compiling it using go build is recommended and allows you to run the program as an executable.

Once compiled you can run the program using the following command (on Linux):
```./main``` or ```./<your output name here>```
Running the program without flags allows you to input and receive multiple responses from and into the chatbot.

Other way of running the program is using command line arguments which requires a flag and your prompt.
```./main <flag> "Hello, how are you?"``` or ```./<your output name here> <flag> "Hello, how are you?"```

There is only 2 flags that can be used with the program:
-m: This flag shows the entire message (including code) that is received from the chatbot. This includes the prompt and the response.
-c: This flag shows only the code that is received from the chatbot. This does not include explanations about the code.

## Examples

## Helpful links
To complete this project these two libraries came in handy:
- [Spinner](github.com/briandowns/spinner)
- [Resty/Rest API for Go](github.com/go-resty/resty/v2)

Spinner allows for a cool loading animation while Resty allows for easy communication with the OpenAI API and formatting of the JSON responses.




