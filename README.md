# Command line ChatGPT 


## Example 

![Example chatgpt-cli](demo.gif)



## Setup 

You must obtain a OpenAI auth_token from here 

https://platform.openai.com/

You have to configure a env variable in your OS.

NOTE: if you use linux or mac you can add this on the .bashrc file.

Example

```
export AUTH_TOKEN_OPEN_AI=myapikey
```

## Download release

https://github.com/ernesto27/chatgpt-cli/releases


## Local develop setup

If you clone the repo another options is to create an .env file with the definition of the key, this will override the OS env var


Create a .env file 

```sh
cp .env-example .env
```

update key on .env file 
```
AUTH_TOKEN_OPEN_AI=myapikey
```


Run 
 
```sh
go run .
```


