package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// nama struct adalah Blog
// yang membangung dari object
type Blog struct {
	Id          int
	Title       string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	JavaScript  bool
	Html        bool
	Php         bool
	React       bool
}

// data - data yang ditampung
var dataBlog = []Blog{
	{
		Title:       "Dragoon Nest",
		StartDate:   "07/06/2023",
		EndDate:     "10/06/2023",
		Duration:    "3 Hari",
		Description: "Petani Gold di Dragon Nest Return",
		JavaScript:  true,
		Html:        true,
		Php:         false,
		React:       false,
	},
	{
		Title:       "Kimetsu No Yaiba",
		StartDate:   "07/06/2023",
		EndDate:     "10/06/2023",
		Duration:    "3 Hari",
		Description: "Macam - macam pilar di Dragon Slayer",
		JavaScript:  true,
		Html:        true,
		Php:         true,
		React:       true,
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	// mengambil data
	e.GET("/", home)
	e.GET("/addproject", addProject)
	e.GET("/contact", contact)
	e.GET("detailproject/:id", detailProject)
	e.GET("updateproject/:id", updateProject)

	// mengirim data
	e.POST("/addblog", addBlog)
	e.POST("/deleteblog/:id", deleteBlog)
	e.POST("/updateproject/:id", updateProjectDone)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": dataBlog,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/addproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func detailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var detailProject = Blog{}

	// for melakukan perulangan
	// i = penampung index dari range
	// data = penampung data dari range
	// range = jarak data/banyaknya data
	// dataBlog = sumber data yang ingin dilakukan perulangan
	for i, data := range dataBlog {
		if id == i {
			detailProject = Blog{
				Title:       data.Title,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				JavaScript:  data.JavaScript,
				Html:        data.Html,
				Php:         data.Php,
				React:       data.React,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": detailProject,
	}

	var tmpl, err = template.ParseFiles("views/detailproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("project-name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	duration := hitungDuration(startDate, endDate)
	description := c.FormValue("description")
	javascript := c.FormValue("javascript")
	html := c.FormValue("html")
	php := c.FormValue("php")
	react := c.FormValue("react")

	println("Title : " + title)
	println("Star Date : " + startDate)
	println("End Date : " + endDate)
	println("Duration : " + duration)
	println("Description : " + description)
	println("Technologies : " + javascript)
	println("Technologies : " + html)
	println("Technologies : " + php)
	println("Technologies : " + react)

	var newBlog = Blog{
		Title:       title,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    duration,
		Description: description,
		JavaScript:  (javascript == "javascript"),
		Html:        (html == "html"),
		Php:         (php == "php"),
		React:       (react == "react"),
	}

	// append bertugas untuk menambahakan data newBlog ke dalam slice dataBlog
	// mirip fungsi push() pada JS
	// param 1 = dimana data ditampung
	// param 2 = data yang akan ditampug
	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var detailProject = Blog{}

	for i, data := range dataBlog {
		if id == i {
			detailProject = Blog{
				Id:          id,
				Title:       data.Title,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				JavaScript:  data.JavaScript,
				Html:        data.Html,
				Php:         data.Php,
				React:       data.React,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": detailProject,
	}

	var tmpl, err = template.ParseFiles("views/updateproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProjectDone(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : ", id)

	title := c.FormValue("project-name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	duration := hitungDuration(startDate, endDate)
	description := c.FormValue("description")
	javascript := c.FormValue("javascript")
	html := c.FormValue("html")
	php := c.FormValue("php")
	react := c.FormValue("react")

	println("Title : " + title)
	println("Star Date : " + startDate)
	println("End Date : " + endDate)
	println("Duration : " + duration)
	println("Description : " + description)
	println("Technologies : " + javascript)
	println("Technologies : " + html)
	println("Technologies : " + php)
	println("Technologies : " + react)

	var updateProject = Blog{
		Title:       title,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    duration,
		Description: description,
		JavaScript:  (javascript == "javascript"),
		Html:        (html == "html"),
		Php:         (php == "php"),
		React:       (react == "react"),
	}

	dataBlog[id] = updateProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func hitungDuration(StartDate, EndDate string) string {
	startTime, _ := time.Parse("2006-01-02", StartDate)
	endTime, _ := time.Parse("2006-01-02", EndDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " Minggu"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " Minggu"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " Hari"
				} else {
					duration = strconv.Itoa(durationDays) + " Hari"
				}
			}
		}
	}

	return duration
}
