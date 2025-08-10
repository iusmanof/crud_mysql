go mod init github.com/iusmnaof/go-project/
go get "github.com/jinzu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"

sudo apt update
sudo apt install mysql-server
sudo systemctl status mysql
mysql --version
sudo mysql_secure_installation


CREATE DATABASE testdb;
USE testdb;
CREATE TABLE testdb (id INT AUTO_INCREMENT PRIMARY KEY, ......);


d, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
