# GoGolang.co "เรียนรู้การสร้างเว็บแอพพลิเคชั่นด้วยภาษาโก"

## Golang web appication case study project

Development 1st golang web project for Beginner to Deployment

# GoGolang.com Web Application development by Golang

> ยินดีต้อนรับสู่ GoGolang.co ซึ่งเป็นโครงการ "การเรียนรู้การสร้างเว็บแอพพลิเคชั่นด้วยภาษาโก" เว็บไซร์อยู่ในช่วงการจัดทำและพัฒนาด้วยภาษาโก (golang) โดยมีเป้าหมายเพื่อเรียนรู้และให้ความรู้ในการพัฒนา เว็บแอพพลิเคชั่นด้วยภาษาโก ตั้งแต่เบื้องต้นจนถึงการนำขึ้นโฮสในเซิร์ฟเวอร์แบบต่างๆ คุณสามารถเข้ามาเยี่ยมชมเว็บไซร์ของเราได้ทุกวันเพื่อติดตามความเคลื่อนไหวเปลี่ยแปลงของเว็บไซร์ คุณจะได้เห็นทิศทางการพัฒนาและลำดับขั้นตอนในการพัฒนาที่ต่อเนื่อง โดยเราจะพยายามเขียนบทความ ในทุกช่วงทุกขั้นตอนในระหว่างการพัฒนาเว็บไซร์แห่งนี้ โดยร่วมถึง การพบเจอปัญหา และการแก้ปัญหา โดยหวังว่าจะเป็นประโยชน์ต่อผู้เยียมชม ที่เริ่มเรียนรู้ภาษาโกเพื่อพัฒนาต่อยอดต่อไป. :)

> GoGolang.co I used it as a learning webapplication development from beginner to deployment by golang following MVC microserivces structure. And i will try to write every details on the way of my learning. May it can benefit to someone that start to learn golang web development like me. Go for golang and enjoys with Go :)

# มาเริ่มกันเลย
## Getting Start 

# ขั้นที่ 1 สร้าง Folder และ Start Server กันก่อน
## Section 1 Create folder structure and Start Server

### ภาพรวมในขั้นตอนที่ 1 นี้มีอะไรกันบ้าง
> - 1. สร้าง Folder Structure
>   - 1.1 Git เตรียมและจัดเก็บงานที่ github.com 
>   - 1.2 จัดการ package ด้วย go module
> - 2. สร้าง function สำหรับ Start Server 
> - 3. กำหนด Folder สำหรับ HTML Template (ฺBasic)
> - 4. กำหนด Folder สำหรับ Static Assets ต่างๆ เช่น css, images, js เป็นต้น
> - 5. ติดตั้ง Bootstrap, fonts-awesome
> - 5. ติดตั้ง node-sass สำหรับ sass compileer เพื่อเพิ่มความยืดหยุ่นให้การจัดการ css 
> - 6. สร้าง Function สำหรับ เลือก Server Port อัตโนมัติ ระหว่าง Dev และ Production Enverontment
> - 7. ทดสอบ Run Server
> - 8. Push to github

## โครงสร้างโฟเดอร์และไฟล์ระบบ MVC Microservices
### Structure MVC Microservices

> ขั้นแรกเรามาเริ่มสร้าง Folder Structure สำหรับโปรเจคกันเลย โดยแนะนำให้สร้างเฉพาะ Folder ทั้งหมดก่อน ส่วน files ต่างๆ ค่อยเริ่มสร้างตามขั้นตอนในแต่ละตอน
### 1. Folders Structure

    -----
        |- app
            |- app.go
            |- urls_mapping.go
        |- assets
            |-css
                |-style.css
            |-images
            |-sass
                |- abstracts
                    |- _functions.scss
                    |- _mixins.scss
                    |- _variables.scss
                |- base
                    |- _animations.scss
                    |- _base.scss
                    |- _typography.scss
                    |- _utilities.scss
                |- components
                    |- component_name.scss
                    |- ...
                |- layout
                    |- layout_name.scss
                    |- ...
                |- pages
                    |- page_name.scss
                    |- ...
                main.scss
        |- configs
            |- server_port_selector.go
            |- ...
        |- controllers
            |- controller-folder-name
                |- controller_name.go
                |- ...
            |- ...
        |- datasources
        | |- database-name-folder
            |- database_name.go
            |- ...
        | |- ...
        |- domain
            |- models-folder-name
                |- model_name.dao
                |- model_anme.dto
                |- ...
        |- routes
            |- route_name.go
            |- ...
        |- services
        | |- services-name-folder
            |- service_name.go
            |- ...
        |- utils
            |- utils_folder_name
                |- utils_name.go
                |- ...
            |- ...
        |- views
            |- views-folder-name
                |- views_file_name.html
                |- ...
            |- ...
        - main.go
        - .gitignore
        - go.mod
        - go.sum
        - package.json
        - package-lock.json
        - README.md
    -----
### 1.1 ติดตั้ง git และสร้าง github.com
> เนื่องจากปัจจุบันนั้นการใช้ git นั้นน่าจะเป็นพื้นฐานทั่วไปสำหรับ Developers ทุกคน จึงไม่ขอกล่าวถึงวิธีการติดตั้งในขั้นตอนนี้ ร่วมถึงการสร้าง git account ในผู้ให้บริการต่างๆ ไม่ว่าจะเป็น github.com gitlab.com หรืออื่นๆ จะขอกล่าวถึงเฉพาะคำสั่งที่ใช้ในการทำงานเท่านั้น

### 1.2 จัดการ package go module

## เริ่มต้นสร้าง Http Server โดย Gin-Gonic
### กำหนด Folder สำหรับ HTML Template
### กำหนด Folder สำหรับ Static Assets

## 2. สร้าง Function สำหรับ เลือก server port อัตโนมัติ
### We're use Getenv to check and get current server port running by using function below.

```golang
    func (s *server) PortRunning() string {
     port := os.Getenv("PORT")
    if port == "" {
		port = "8089"
	}
	return port
}
```

### Start Server

```
    $ go run main.go
```
## ใช้ Go Module ในการจัดการ Package ใน Project
## How to init Go Module

```
go mod init github.com/your_github_username/your_repo_name

```

## 3. Templates
### How to set glbal templates
### Template Nested Template

## Static Folder for Assets
### Link Bootstrap, Fonts-Awesome and Google Fonts
### install node-sass and jquery  
```
    $ npm init
    $ npm install node-sass --save-dev
    $ npm install jquery --save
```
#### Add script command to compile sass by add line below package.json and auto watching and auto compile when make any change to files.

#### เพิ่ม node-sass เพื่อช่วยเพิ่มประสิทธิภาพและยืดหยุ่นในการเขียน css รวมถึงการคอมพลายแบบอัตโนมัติเมื่อมีการเขียนเพิ่มหรือลดโค้ด โดยใช้คำสั่ง -w 

```
    "scripts": {
        "compile:sass": "node-sass ./assets/sass/main.scss ./assets/css/style.css -w"
    }
```
#### Run compile command during development 

#### คำสั่งในการคอมพลาย sass ไปเป็น css 

```
    $ npm run compile:sass

```

## 8. Push to github account 
> จัดเก็บไฟล์ ที่ github.com account


## Database Section

### เชื่อมต่อฐานข้อมูล - Connect to databse

folder structure for datasource
    -----
        |- datasource
            |- mysql
                |- users_db.go
            |- postgres
            |- ...

### Connect to MySql
```
    func init() {
        
        dataSourceName := fmt.Printf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1:3306", "users_db")
    }
```

### สร้างตัวแปรแวดล้อม (env variables)

```
// รูปแบบการ set ค่าแปร environment variable
$ export variable_name='value'

// set ค่า user name ที่ใช้เชื่อมต่อ database
$ export mysql_users_usersname='root'

// set ค่า password สำหรับเชื่อมต่อ database
$ export mysql_users_password='p11#$12ss'

// ชื่อ Host และ port ที่ database server ติดตั้งอยู่ 
// หากเป็นลงบนเครื่องที่ใช้เขียนโปรแกรมจะใช้ตามพื้นฐานทั่วไปตามด้านล่าง 
$ export mysql_users_host='127.0.0.1:3306'

// ชื่อ database ที่ต้องการติดต่อ
$ export mysql_users_schema='users'

```

### การตรวจสอบค่าตัวแปร

```
// คำสั่งสำหรับแสดงค่าตัวแปรที่ set แล้วเพื่อตรวจสอบว่าถูกต้องหรือไม่ โดยใช้คำสั่ง echo ตามด้วยชื่อ ตัวแปร
$ echo $variable_name

// การใช้งาน
$ echo $mysql_users_username 


```

**_ What is dufferrent about PUT and PATCH _**

- PUT needs all fields data input to update otherwise will update no data input fileds to blank.
- PATCH will update only a field that have data input and remain other fiedls data as the same.

## Golang Online Course Suggestion

### Beginner

### Web Deployment

### Algorythm

## Golang Book Suggestion

### Learning Go by testing
