package controllers

import "github.com/revel/revel"

type Objects struct {
    *revel.Controller
}

type Person struct {
    Id int32
    Name string
}

func (c Objects) List() revel.Result {
    var objects = Person{20, "Bob"}
    return c.RenderJson(objects)
}

func (c Objects) Get(id int32) revel.Result {
    var objects = Person{id, "Bob"}
    return c.RenderJson(objects)
}
