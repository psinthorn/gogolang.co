# gogolang.co
Development 1st golang web project for Beginner to Deployment

# GoGolang.com Web Application development by Golang
>
> GoGolang.co I used it as a learning webapplication development from beginner to deployment online by golang  following MVC microserivces structure. And i will try to write every details on the way of learning maybe it can benefit to someone that start to learn golang web development like me. Go for golang and enjoys with Go :) 

# Structure MVC Microservices

# Getting Start
## Start Server

## We're use Getenv to check and get current server port running by using function below.

```golang
    func (s *server) PortRunning() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	return port
}
```

## Go Module Init
 ```
 go mod init github.com/your_github_usernams/your_repos

 ```
# Folders Structure
    -----
        |- app
        |- 
  

# Templates 
## How to set glbal templates

## Template Nested Template

# Database Section
## Connect to databse 

*** What is dufferrent about PUT and PATCH ***
- PUT needs all fields data input to update otherwise will update no data input fileds to blank.
- PATCH will update only a field that have data input and remain other fiedls data as the same.

# Golang Online Course Suggestion
## Beginner
## Web Deployment
## Algorythm

# Golang Book Suggestion 
## Learning Go by testing 

