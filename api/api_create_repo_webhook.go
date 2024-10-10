package api

import (
	"autodock-be/dto"
	"autodock-be/git"

	"github.com/gofiber/fiber/v2"
)

func CreateRepoWebhook(c *fiber.Ctx)error{
	inputObj := dto.WebHookRepo{}
	if err := c.BodyParser(&inputObj); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}
	

	err := git.CreateWebHookRepository(inputObj.Owner,inputObj.Repo,inputObj.Token,inputObj.WebHookUrl,inputObj.Secret)
	if err != nil{
		return err
	}
	return c.Status(200).JSON("Success")
}