package api

import (
	"autodock-be/dto"
	"autodock-be/git"

	"github.com/gofiber/fiber/v2"
)

func CreateRepoWebhook(c *fiber.Ctx)error{
	inputObj := dto.WebHookRepo{}
	inputObj.Owner = "generator-check"
	inputObj.Repo = "HRManagement-ws131-cgaas"
	inputObj.Token = "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"
	inputObj.WebHookUrl = "https://wwwwgoogle.com"
	inputObj.Secret = "Asd@1234"

	err := git.CreateWebHookRepository(inputObj.Owner,inputObj.Repo,inputObj.Token,inputObj.WebHookUrl,inputObj.Secret)
	if err != nil{
		return err
	}
	return c.Status(200).JSON("Success")
}