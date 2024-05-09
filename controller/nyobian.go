package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	gile"github.com/jul003/BackEnd_Jul"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetGadgetGetAll(c *fiber.Ctx) error {
    gadgets, err := gile.GadgetGetAll()
    if err != nil {
    
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }
    
    return c.JSON(gadgets)
}
