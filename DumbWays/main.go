package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// nama dari struct nya adalah blog
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
	// Author      string
	// PostDate    string
}

// data - data yang ditampung
var dataBlog = []Blog{
	{
		Title:       "Halo Guys",
		StartDate:   "07/06/2023",
		EndDate:     "10/06/2023",
		Duration:    "3 Bulan",
		Description: "Mari Makan Guys",
		JavaScript:  true,
		Html:        true,
		Php:         true,
		React:       true,
		// Author:      "Power Ranger",
		// PostDate:    "07/06/2023",
	},
	{
		Title:       "Halo Ultarmen",
		StartDate:   "07/06/2023",
		EndDate:     "10/06/2023",
		Duration:    "3 Bulan",
		Description: "Mari Makan Guys",
		JavaScript:  true,
		Html:        true,
		Php:         true,
		React:       true,
		// Author:      "Ultaramen",
		// PostDate:    "07/06/2023",
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	// e.GET("/hello", hellowordl)
	e.GET("/", home)
	e.GET("/addproject", addProject)
	e.GET("/contact", contact)
	e.GET("detailproject/:id", detailProject)
	e.GET("updateproject/:id", updateProject)

	e.POST("/addblog", addBlog)
	e.POST("/deleteblog/:id", deleteBlog)
	e.POST("/updateproject/:id", updateProjectDone)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// func hellowordl(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello Wordl")
// }

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

	// data := map[string]interface{}{
	// 	"id":      id,
	// 	"Title":   "DumbWays Web App",
	// 	"Content": "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Temporibus dignissimos eum sit voluptatum suscipit. Sequi dolorem, ipsa dolores, optio quod officiis quisquam, atque nihil omnis magnam quia tempora minus doloribus tempore et laboriosam voluptates. Mollitia beatae at officia quisquam placeat delectus, cumque ipsum facilis pariatur praesentium aliquid quo eaque fugiat nisi atque corporis sequi. Alias unde nihil maiores earum sint soluta repellat quasi nisi numquam quia, illum et molestias eum, hic voluptates delectus possimus accusantium temporibus dolorem qui tempora animi odio quo! Quaerat accusantium numquam quibusdam, ratione quod cumque culpa totam, eaque alias obcaecati mollitia earum! Porro aspernatur similique itaque. Lorem ipsum dolor sit amet consectetur, adipisicing elit. Exercitationem, delectus quisquam. Atque, iusto. Error impedit nemo voluptate quis ea enim esse, commodi aliquam ex ipsam debitis veritatis libero quasi autem dolor natus dolorem hic magni magnam, asperiores molestias incidunt odit eaque doloremque! Est alias repellendus ex libero consectetur adipisci similique, ipsa earum, fugit laboriosam et dolorum sunt facere ad accusamus nemo tempore distinctio odio! Accusamus rem veniam, nobis eveniet ipsam molestias dicta ut. Pariatur accusamus accusantium vel sunt illum alias cum, est incidunt perspiciatis animi? Ipsum esse accusantium soluta, incidunt voluptate, optio ratione expedita consequuntur natus quos velit enim itaque!",
	// }

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
				// Author:      data.Author,
				// PostDate:    data.PostDate,
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
		// Author:      "Arsandi Wira P",
		// PostDate:    time.Now().String(),
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
				// Author:      data.Author,
				// PostDate:    data.PostDate,
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
		// Author:      "Arsandi Wira P",
		// PostDate:    time.Now().String(),
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
