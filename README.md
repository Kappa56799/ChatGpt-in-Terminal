# ChatGpt-in-Terminal

## Description
This is a simple chatbot that uses OpenAI GPT to generate responses and code straight from the terminal only. It is written in Go and uses the OpenAI API to generate responses.
I made this because I wanted to make a chatbot that I could use in the terminal allowing me to easily output code to a file using redirects on Linux as well as learning to 
program in Go :). I tested this on Arch Linux using kitty terminal and everything worked fine including running the program using 'go run' and compiling it using 'go build'. 
Feel free to fork and change this repository to your linking.

## Requirements
- Go 1.16 or higher 
- OpenAI API key
- Linux (Haven't tested on Windows)

# How to find your API key
Go to this website [Your API Key!](https://platform.openai.com/api-keys), log in and create a key!
![image](https://github.com/Kappa56799/ChatGpt-in-Terminal/assets/114831362/6e1759ab-44b8-4a41-bc01-4ce63405576f)



## How to Install
Either download the compiled code or clone the repository using the following command:

```git clone https://github.com/Kappa56799/ChatGpt-in-Terminal```

From here follow the instructions in the usage section to run the program.

## Usage
To run the program you firstly need to set the environment variable OPENAI_API_KEY to your API key. You can do this by running the following command in the terminal or by putting this command into your shell config (fish,zsh,bash):
```bash
export OPENAI_API_KEY="<your key here>"
```
Or you can change the value of the variable in the main.go file to your key and compile the app yourself.

Then you can run the program using the following command:

```go run main.go``` 
or 
```go build -o  main.go <your output name here>```

Compiling it using go build is recommended and allows you to run the program as an executable.

Once compiled you can run the program using the following command (on Linux):
```./main``` or ```./<your output name here>```
Running the program without flags allows you to input and receive multiple responses from and into the chatbot.

Other way of running the program is using command line arguments which requires a flag and your prompt.

```./main <flag> "Hello, how are you?"``` 
or 
```./<your output name here> <flag> "Hello, how are you?"```

There is only 3 flags that can be used with the program:
-m: This flag shows the entire message (including code) that is received from the chatbot. This includes the prompt and the response.
-c: This flag shows only the code that is received from the chatbot. This does not include explanations about the code.
-o: This flag uses OCR to extract text from an image in clipboard and sends it to the chatbot.

## Examples
![image](https://github.com/Kappa56799/ChatGpt-in-Terminal/assets/114831362/f17fdd93-5662-44df-b82c-21171be4a6b6)

![image](https://github.com/Kappa56799/ChatGpt-in-Terminal/assets/114831362/2fbb4600-256a-41ae-be7e-654f0a6d00f2)

![image](https://github.com/Kappa56799/ChatGpt-in-Terminal/assets/114831362/ad859a3e-6f4e-4899-b858-b0075796793a)


## Helpful links
To complete this project these two libraries came in handy:
- [Spinner](https://github.com/briandowns/spinner)
- [Resty/Rest API for Go](https://github.com/go-resty/resty)
- [Tesseract for Golang](https://github.com/otiai10/gosseract)

Spinner allows for a cool loading animation while Resty allows for easy communication with the OpenAI API and formatting of the JSON responses.




