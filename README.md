# GoGolang.co "เรียนรู้การสร้างเว็บแอพพลิเคชั่นด้วยภาษาโก" 
## Golang web appication case study project
Development 1st golang web project for Beginner to Deployment

# GoGolang.com Web Application development by Golang
> ยินดีต้อนรับสู่ GoGolang.co ซึ่งเป็นโครงการ "การเรียนรู้การสร้างเว็บแอพพลิเคชั่นด้วยภาษาโก" เว็บไซร์อยู่ในช่วงการจัดทำและพัฒนาด้วยภาษาโก (golang) โดยมีเป้าหมายเพื่อเรียนรู้และให้ความรู้ในการพัฒนา เว็บแอพพลิเคชั่นด้วยภาษาโก ตั้งแต่เบื้องต้นจนถึงการนำขึ้นโฮสในเซิร์ฟเวอร์แบบต่างๆ คุณสามารถเข้ามาเยี่ยมชมเว็บไซร์ของเราได้ทุกวันเพื่อติดตามความเคลื่อนไหวเปลี่ยแปลงของเว็บไซร์ คุณจะได้เห็นทิศทางการพัฒนาและลำดับขั้นตอนในการพัฒนาที่ต่อเนื่อง โดยเราจะพยายามเขียนบทความ ในทุกช่วงทุกขั้นตอนในระหว่างการพัฒนาเว็บไซร์แห่งนี้ โดยร่วมถึง การพบเจอปัญหา และการแก้ปัญหา โดยหวังว่าจะเป็นประโยชน์ต่อผู้เยียมชม ที่เริ่มเรียนรู้ภาษาโกเพื่อพัฒนาต่อยอดต่อไป. :) 

>
> GoGolang.co I used it as a learning webapplication development from beginner to deployment by golang  following MVC microserivces structure. And i will try to write every details on the way of my learning. May it can benefit to someone that start to learn golang web development like me. Go for golang and enjoys with Go :) 


# โครงสร้างโฟเดอร์และไฟล์ระบบ
## Structure MVC Microservices

# เริ่มต้นสร้าง Http Server โดย Gin-Gonic
## Getting Start

## สั่ง start http server 
## Start Server
```
    $ go run main.go

```


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

