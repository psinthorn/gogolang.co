package global

const (
	mongouri    = "mongodb+srv://f2dev:F2deV@f2code.ducv6.mongodb.net/blogapps?retryWrites=true&w=majority"
	dbname      = "blogapps"
	performance = 100
)

var (
	jwtSecret = []byte("blogappsSecret")
)
