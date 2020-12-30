package handler

import (
	"race_condition/database"
	"race_condition/model"

	"github.com/gofiber/fiber/v2"
)

type TransferRequest struct {
	From   uint `json:"fromID"`
	To     uint `json:"toID"`
	Amount int  `json:"amount"`
}

func TransferMoney(c *fiber.Ctx) error {
	db := database.DBConn
	var transferReq TransferRequest
	var fromAccount model.Account
	var toAccount model.Account
	var err error

	if err = c.BodyParser(&transferReq); err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": err})
	}

	if !CheckBalance(transferReq.From, transferReq.Amount) {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "Yetersiz bakiye!"})
	}

	if err = db.Take(&fromAccount, transferReq.From).Error; err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "Bir hata oluştu"})
	}

	if err = db.Take(&toAccount, transferReq.To).Error; err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "Bir hata oluştu"})
	}

	fromAccount.Balance = fromAccount.Balance - transferReq.Amount
	toAccount.Balance = toAccount.Balance + transferReq.Amount

	db.Save(&fromAccount)
	db.Save(&toAccount)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Transfer başarılı."})

}

func CheckBalance(userID uint, amount int) bool {
	db := database.DBConn
	var account model.Account

	db.Take(&account, userID)

	if account.Balance < amount {
		return false
	}

	return true
}
