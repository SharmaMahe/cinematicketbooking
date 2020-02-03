package controllers

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_"fmt"
	"math"
	"strconv"
	"app/mail"
	"unicode/utf8"
)

// init To load the models
func init() {
    orm.RegisterModel(new(models.Seat),new(models.User),new(models.BookedSeat))
}

// MainController operations for cinema booking
type MainController struct {
	beego.Controller
}

// GET ...
// @Title GET
// @Description Show ticket booking form
// @Param	body	body 	MainController	true		"body for MainController content"
// @Success 201 {object} MainController
// @Failure 403 body is empty
func (c *MainController) Get() {

	// Initialize orm
	o := orm.NewOrm()

	// Declare Seat model
	var records []models.Seat

	// Declare BookedSeat model
	var allBookedSeats []models.BookedSeat

	// fetch all packets records
	o.QueryTable("seats").All(&records)

	o.QueryTable("booked_seats").Filter("status",1).RelatedSel().All(&allBookedSeats)

	c.Data["seats"] = customizeSeats(records,allBookedSeats)

	flash := beego.ReadFromRequest(&c.Controller)

	// Flash messages
	var flashMessage,flashClass = "","";

	if _, ok := flash.Data["success"]; ok {
		flashMessage = flash.Data["success"]
		flashClass = "alert alert-success"
	} else if _, ok = flash.Data["error"]; ok {
		flashMessage = flash.Data["error"]
		flashClass = "alert alert-danger"
	}

	// Pass variables to view.
	c.Data["message"] = flashMessage
	c.Data["flashClass"] = flashClass
	c.Data["allBookedSeats"] = allBookedSeats

	// Render Template
	c.TplName = "index.html"
}

// customizeSeats To break the seats in proper format for display.
func customizeSeats(seats []models.Seat,allBookedSeats []models.BookedSeat) map[int][]models.Seat {
	var breakPointLength,UpdatesValues = 7,make(map[int][]models.Seat);
	seatLoop := int(math.Ceil((float64(len(seats)))/7))

	// Make booked seat
	for j := 0; j < len(seats); j++ {
		k := contains(allBookedSeats,seats[j].Id)
		if k {
			seats[j].Booked = true	
		}else{
			seats[j].Booked = false
		}
	}
	
	// Break down into rows
	for i := 0; i < seatLoop; i++ {
		UpdatesValues[i] = seats[(i*breakPointLength):(i+1)*breakPointLength]		
	}

	return UpdatesValues
}

func contains(s []models.BookedSeat, id int) bool {
    for _,v := range s {
        if v.Seat.Id == id {
            return true
        }
    }
    return false
}

// Post ...
// @Title Create
// @Description Book seat
// @Param	body	body 	models.MainController	true		"body for MainController content"
// @Success 201 {object} models.MainController
// @Failure 403 body is empty
func (c *MainController) BookSeat() {

	// Booking status 
	// status = 0,not booked yet
	// status = 1,booked
	// status = 2 ,cancelled

	o := orm.NewOrm()

	flash := beego.NewFlash()

	// initialize user model
	user := models.User{}

	// initialize the seat model
	seat := models.Seat{}

	//Get form value.
	email := c.GetString("email")
	name := c.GetString("name")
	seats := c.GetStrings("seats")

	if len(seats) == 0{
		flash.Error("Couldn't book tickets")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	valid := validation.Validation{}
    valid.Required(email, "Email")
    valid.Required(name, "Name")
    valid.Required(seats, "Seats")

	// validate form
    b := valid.HasErrors()
 
    if !b {

		// Begin transaction
		o.Begin()

		// Insert User Record
		user.Email = email
		user.Name = name

		o.ReadOrCreate(&user, "Email");

		// Loop through the seats
		for _,val := range seats{

			// initialize the bookedSeat model
			bookedSeat := models.BookedSeat{}

			// Query for seat object
			seat.SeatNumber = val
			o.Read(&seat,"SeatNumber");

			// Assgn field values
			bookedSeat.User = &user
			bookedSeat.Seat = &seat
			bookedSeat.Status = 1

			// Finally insert the record and check if not exist.
			if created, _, err := o.ReadOrCreate(&bookedSeat, "Seat","Status"); err == nil {
			    if !created {
			        o.Rollback();
				    flash.Error("Seat already booked!")
	    			flash.Store(&c.Controller)
	    			c.Redirect("/", 302)
	    			return
			    }
			}
		}

		// Commit the changes.
		o.Commit();	

		// Send ticket confirmation
		go mail.SendBookingMail(name,email,seats)

		msg := "Tickets"

		switch len(seats) {
			case 1:
				msg = "Ticket"
		}

		flash.Success("Your "+msg+" booked successfully!")
	    flash.Store(&c.Controller)
	    c.Redirect("/", 302)
		return 
    }else{
    	// validation does not pass
        var errs string = "";
        for _, err := range valid.Errors {
            errs = errs +","+ err.Key+" "+err.Message
        }
        flash.Error(trimFirstRune(errs))
        flash.Store(&c.Controller)
        c.Redirect("/", 302)
        return
    }
}

// Post ...
// @Title Cancel Ticket
// @Description To canel the movie ticket 
// @Param	body	body 	models.MainController	true		"body for MainController content"
// @Success 201 {object} models.MainController
// @Failure 403 body is empty
func (c *MainController) CancelSeat() {
	o := orm.NewOrm()

	flash := beego.NewFlash()

	//Get form value.
	email := c.GetString("email")
	seats := c.GetStrings("cancelseats")

	if len(seats) == 0{
		flash.Error("Couldn't cancel booking")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	cancelledSeats := make([]string,0)

	// Begin transaction
	o.Begin()

	// Insert User Record
	user := models.User{Email: email}

	err := o.Read(&user,"Email");

	if err == nil {
		// Loop through the seats
		for _,val := range seats{

			// initialize the bookedSeat model
			id,err := strconv.Atoi(val)
			bookedSeat := models.BookedSeat{Id:id,User:&user}

			if err = o.Read(&bookedSeat,"Id","User"); err == nil {
				// Change the status
				bookedSeat.Status = 2

				// Update record 
			    if _, err = o.Update(&bookedSeat,"Status"); err != nil {
			    	o.Rollback();
				    flash.Error("Something went wrong!")
	    			flash.Store(&c.Controller)
	    			c.Redirect("/", 302)
	    			return
			    }
			    seat := models.Seat{Id:bookedSeat.Seat.Id}
			    o.Read(&seat)

			    cancelledSeats = append(cancelledSeats,seat.SeatNumber)
			}else{
				o.Rollback();
				flash.Error("You are not authorized to cancel this booking, please enter correct email id")
				flash.Store(&c.Controller)
				c.Redirect("/", 302)
				return
			}
		}
	}else{
		flash.Error("You are not authorized to cancel this booking, please enter correct email id")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	// Commit the changes.
	o.Commit();	

	// Send ticket confirmation
	go mail.SendBookingCancelMail(user.Name,user.Email,cancelledSeats)

	msg := "Tickets"

	switch len(seats) {
		case 1:
			msg = "Ticket"
	}

	flash.Success(msg+ " cancelled successfully!")
    flash.Store(&c.Controller)
    c.Redirect("/", 302)
	return 
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
