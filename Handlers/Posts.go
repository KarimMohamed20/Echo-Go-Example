package Handlers

import (
	"echo-framework/Models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func (h * Handler) CreatePost(c echo.Context) (err error){
	u := &Models.User{
		ID: bson.ObjectIdHex(getUID(c)),
	}
	p := &Models.Post{
		ID: bson.NewObjectId(),
		UID: getUID(c),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	if p.Title == "" && p.Content == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Fields"}
	}
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("FlutterEgypt").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("FlutterEgypt").C("posts").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	userID := getUID(c)


	// Retrieve posts from database
	var posts []*Models.Post
	db := h.DB.Clone()
	if err = db.DB("FlutterEgypt").C("posts").
		Find(bson.M{"uid": userID}).
		All(&posts); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, posts)
}
