package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-file-go/azfile"

	// "path/filepath"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	Sqldb = sqlDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/title", getTitle).Methods("GET")
	r.HandleFunc("/api/sign-up", signUp).Methods("POST")
	r.HandleFunc("/api/file-upload", fileUpload).Methods("POST")

	r.HandleFunc("/api/login", login).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(r)
	srv := &http.Server{Handler: handler, Addr: ":3000"}
	log.Fatal(srv.ListenAndServe())
}

type Data struct {
	Title string `json: "title"`
}
type UserData struct {
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
	Id       int    `json:_id`
}
type fileDetails struct {
	File string `json: file`
	Type string `json: type`
	Name string `json: name`
}
type FileMetaData struct {
	UserId       int    `json:userId`
	UserName     string `json :userName`
	Email        string `json :email`
	UploadedTime string `json :uploadedTime`
	FileName     string `json:fileName`
	FileFormat   string `json :fileFormat`
}
type StatusData struct {
	StatusData string `json: "statusData"`
	Id         int    `json :id`
	Name       string `json:name`
	Email      string `json:email`
}

type Sizer interface {
	Size() int64
}

func getTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	title := Data{"Data Aggregation"}
	//user := UserData{"d", "d", "d"}
	json.NewEncoder(w).Encode(&title)
}

func fileUpload(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, _ := r.FormFile("file")

	size := file.(Sizer).Size()
	byteContainer := make([]byte, size)

	file.Read(byteContainer)

	f, err := os.Create(fileHeader.Filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	val := string(byteContainer)
	// fmt.Println("val: ", val)
	data := []byte(val)
	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	url := main1(f)

	fmt.Println(url)

	var statusData StatusData

	statusData.StatusData = "success"

	statusData.Email = url

	json.NewEncoder(w).Encode(statusData)

}
func signUp(w http.ResponseWriter, r *http.Request) {
	//main1()
	Sqldb = sqlDB()
	fmt.Println("POST METHOD")
	w.Header().Set("Content-Type", "application/json")
	var userData UserData
	//var book Book
	_ = json.NewDecoder(r.Body).Decode(&userData)
	fmt.Println(userData)
	// sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	// res, err := sqlStatement.Exec("banana", 150)
	stmt, err1 := Sqldb.Prepare("insert into userdetails(email,password,name) values(?, ?, ?);")
	if err1 != nil {
		fmt.Println(err1)
	}
	res, err1 := stmt.Exec(userData.Email, userData.Password, userData.Username)
	fmt.Println("res=  -====-=-=-=-=-=-=", res)

	var statusData StatusData

	statusData.StatusData = "success"

	json.NewEncoder(w).Encode(statusData)

}

func login(w http.ResponseWriter, r *http.Request) {
	Sqldb = sqlDB()
	var userData UserData
	_ = json.NewDecoder(r.Body).Decode(&userData)
	var userName string = userData.Username
	var pass string = userData.Password
	// fmt.Println("usernaem: ", userName)
	// ctx := context.Background()
	fmt.Println("login method")
	w.Header().Set("Content-Type", "application/json")

	//var book Book
	_ = json.NewDecoder(r.Body).Decode(&userData)

	// Read employees
	var stat string = " "
	var name, email, password string
	// rows,err := Sqldb.Query("select * from userdetails where name=?")
	rows, err := Sqldb.Query("select * from userdetails where name= ? and password= ?", userName, pass)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email, &password, &name)
		if err != nil {
			log.Fatal(err)
		}
		if userData.Username == userName {
			stat = "success"
			break
		}
		// fmt.Println(name)
		fmt.Println("Name : ", name)

	}
	if stat == " " {
		stat = "failure"
	}
	// stat="failure"

	// rows,err:=Sqldb.Query("Select * from userdetails")
	// if err!=nil{
	// 	log.Println(err)
	// }
	// var username,password string
	// for rows.Next(){
	// 	err=rows.Scan(&username,&password)
	// 	if err!=nil{
	// 		log.Println(err)
	// 	}
	// }
	defer Sqldb.Close()
	fmt.Println("status : ", stat)

	var statusData StatusData
	if stat == "failure" {
		fmt.Println("stat : ", stat)
		statusData.StatusData = "failure"
		json.NewEncoder(w).Encode(statusData)
	}

	if stat == "success" {
		fmt.Print("status : ", stat)
		statusData.StatusData = "success"
		json.NewEncoder(w).Encode(statusData)
	}

}

// func Readusers(username string, pass string) string {

// 	fmt.Println("userName : ", username)
// 	fmt.Println("password: ", pass)
// 	ctx := context.Background()

// 	// Check if database is alive.
// 	err := Sqldb.PingContext(ctx)
// 	if err != nil {
// 		return "failure"
// 	}

// 	tsql := fmt.Sprintf("SELECT email, password, name FROM usertable.userdetails;")

// 	// Execute query
// 	rows, err := Sqldb.QueryContext(ctx, tsql)
// 	if err != nil {
// 		return "failure"
// 	}

// 	defer rows.Close()

// 	// Iterate through the result set.
// 	for rows.Next() {
// 		var name, email, password string

// 		// Get values from row.
// 		err := rows.Scan(&email, &password, &name)
// 		if err != nil {
// 			return "failure"
// 		}
// 		if name == username {
// 			fmt.Printf("email: %s, password: %s, name: %s\n", email, password, name)
// 			return "success"
// 		} else {
// 			return "failure"
// 		}
// 	}
// 	return "failure"
// }

var Sqldb *sql.DB

func sqlDB() *sql.DB {
	var server = "akswebappdb21.mysql.database.azure.com"
	// var port = 3306

	var database = "usertable"
	var user = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, server, database)

	var err error

	// Create connection pool
	Sqldb, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	if Sqldb != nil {
		fmt.Printf("Connected!\n")
	}

	return Sqldb

}

// Please set environment variable ACCOUNT_NAME and ACCOUNT_KEY to your storage accout name and account key,
// before run the examples.
func accountInfo() (string, string) {
	//return os.Getenv("ACCOUNT_NAME"), os.Getenv("ACCOUNT_KEY")
	return "csg10032001b930ea2c", "wjS5LrRRAIKtGIT6f5aWiK/7S/CyhHgwAuQr4zxSAo4xbSiciSrt5P514W2SejQGHAYEEcgLUKQORPKzgjEAHg=="
}

func main1(file *os.File) string {
	//file, err := os.Open("BigFile.bin") // Open the file we want to upload (we assume the file already exists).
	// file, err := os.Create("Sample.txt")
	// fmt.Println(file.Name())

	defer file.Close()
	fileSize, err := file.Stat() // Get the size of the file (stream)
	if err != nil {
		log.Fatal(err)
	}

	// From the Azure portal, get your Storage account file service URL endpoint.
	accountName, accountKey := accountInfo()
	credential, err := azfile.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create a FileURL object to a file in the share (we assume the share already exists).
	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.windows.net/myshare/%s", accountName, file.Name()))
	vars := fmt.Sprintf("https://%s.file.core.windows.net/myshare/%s", accountName, file.Name())
	fileURL := azfile.NewFileURL(*u, azfile.NewPipeline(credential, azfile.PipelineOptions{}))

	ctx := context.Background() // This example uses a never-expiring context

	// Trigger parallel upload with Parallelism set to 3. Note if there is an Azure file
	// with same name exists, UploadFileToAzureFile will overwrite the existing Azure file with new content,
	// and set specified azfile.FileHTTPHeaders and Metadata.
	err = azfile.UploadFileToAzureFile(ctx, file, fileURL,
		azfile.UploadToAzureFileOptions{
			Parallelism: 3,
			FileHTTPHeaders: azfile.FileHTTPHeaders{
				CacheControl: "no-transform",
			},
			Metadata: azfile.Metadata{
				"createdby": "ananth&gokul",
			},
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {
				fmt.Printf("Uploaded %d of %d bytes.\n", bytesTransferred, fileSize.Size())
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	return vars
}

/*

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "akswebappdb21.mysql.database.azure.com"
	database = "usertable"
	user     = "dbadminuser@akswebappdb21"
	password = "Maxmin@321"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")

	// Insert some data into table.
	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("banana", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("orange", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("apple", 100)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
}
*/
