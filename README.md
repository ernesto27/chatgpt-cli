# Command line ChatGPT 


## Example 

![Example chatgpt-cli](demo.gif)

## Setup 
You must obtain a OpenAI auth_token from here 

https://platform.openai.com/


## Environment variables setup

You must configure a  env variable in your OS.

Example

```
export AUTH_TOKEN_OPEN_AI=myenv
```


Another option is create an .env file with the definition of the key, this will override the OS env var


Create a .env file 

```sh
cp .env-example .env
```

update key on .env file 
```
AUTH_TOKEN_OPEN_AI=mykey
```






## Develop init
 
```sh
go run .
```


