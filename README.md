# Cinema ticket booking
Display all the seats in the theater and allow users to book it by clicking it. Only one user should be allowed to reserved a specific seat.
If another user clicks a seat that was booked, he should get an error. You must handle the concurrency scenarios and avoid data inconsistency.

Seats can be unbooked by clicking the booked seat again.

The solution should have:

Basic feature of the application should include:

    A single web page with the seats displayed in a grid (you could start by a smaller number of seats, maybe 20-30)

    If a seat is available, the user should be asked for their details like name, email ID, etc and send an email to them with a confirmation

    You do not need to gather the payment details

    Test coverage is ideal

    Full installation instruction must be available via readme.md.

## Prerequisites
    1. 'Go' should be installed on your machine (Min Version : 1.11)
    2. 'Mysql' should be installed on your machine

## Installation
	Please follow below instructions :

	1 To get the directory on local machine just run the command : git clone https://github.com/SharmaMahe/cinematicketbooking
	2.Then go to project directory : cd cinematicketbooking/src/app
	3 Set the local environment variable GOROOT,GOPATH and PATH variables 
		i.e 
		Project is installed/cloned in /var/www directory and 'go' installed in /usr/local
		then following command will work :
		export GOROOT=/usr/local/go
		export GOPATH=/var/www/cinematicketbooking/
		export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
	note : Above commands for linux environment 
    4 Run "go get github.com/astaxie/beego"
	5 Run "go get github.com/go-sql-driver/mysql"
	6 Run "go get github.com/beego/bee"
    7 Configure the database credential in src/app/conf
    
## Usage

	Just run "bee run"


## Test Case
	To execute test cases run 
	"go test -v tests/default_test.go"