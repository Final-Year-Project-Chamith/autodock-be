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
	inputObj.WebHookUrl = "https://71d2-2402-d000-812c-1597-64f4-f09d-210e-2670.ngrok-free.app/autodock-be/api/webhook"
	inputObj.Secret = "Asd@1234"

	err := git.CreateWebHookRepository(inputObj.Owner,inputObj.Repo,inputObj.Token,inputObj.WebHookUrl,inputObj.Secret)
	if err != nil{
		return err
	}
	return c.Status(200).JSON("Success")
}