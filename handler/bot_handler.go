package handler

import (
	"code-ai/models/domain"
	"code-ai/services"
	"fmt"
	"net/http"
	"runtime"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	BotService services.BotService
    OpenAIService services.CodeService
	User services.UserServiceImp
}

func NewBotHandler(botService services.BotService, openAiService services.CodeService, user services.UserServiceImp) *BotHandler {
	return &BotHandler{
		BotService: botService,
        OpenAIService: openAiService,
		User: user,
	}
}

func (wh *BotHandler) HandleBot(c echo.Context)error{
	update := new(tgbotapi.Update)
    if err := c.Bind(update); err != nil {
        return err
    }

	runtime.GOMAXPROCS(2)
	
	var userFound = make(chan *domain.User)

    // Handle the update here using Bot Service
    if update.Message != nil {
		go func () {
			user, err := wh.User.FindUserByID(update.Message.From.ID)
			if err != nil {
				createUser := domain.User{
					ID: update.Message.From.ID,
					Name: update.Message.From.UserName,
					Status: "UNPAID",
					TextCount: 10,
				}
				user, _ = wh.User.CreateUser(&createUser)
			}
			userFound <- user
		}()

		user := <- userFound
		if update.Message.Photo != nil { 
			wh.BotService.SendMessage(update.Message.Chat.ID, "Tunggu Sebentar...")
			return c.String(http.StatusOK, "OK")
		}

		if update.Message.Photo == nil{
			if user.TextCount > 0 || user.Status == "PAID" {
				if update.Message.Text == "/start" {
					wh.BotService.SendMessage(update.Message.Chat.ID, "Halo, "+update.Message.From.UserName+"! Silahkan masukkan kode yang anda inginkan. Untuk melihat daftar perintah, silahkan ketik /help")
					return c.String(http.StatusOK, "OK")
				}
				if update.Message.Text == "/upgrade" {
					wh.BotService.SendMessage(update.Message.Chat.ID, "Berikut Harga Credit yang tersedia:\n\nRp. 30.000,00 per 250 Credit\nRp. 50.000,00 per 500 Credit\nRp. 100.000,00 per 1200 Credit\n\ntransfer ke rekening 1234567890 a/n Innaka Dylee. Setelah itu, \nkirim bukti transfer ke @InnakaDylee.")
					return c.String(http.StatusOK, "OK")
				}
				if update.Message.Text == "/credit" {
					wh.BotService.SendMessage(update.Message.Chat.ID, "Bot ini dibuat oleh @InnakaDylee.")
					return c.String(http.StatusOK, "OK")
				}
				if update.Message.Text == "/limit" {
					var text string
					if user.Status == "PAID" {
						text = fmt.Sprintf("Limit anda tidak terbatas.")
					} else {
						text = fmt.Sprintf("Limit anda tersisa %d kali.", user.TextCount)
					}
					wh.BotService.SendMessage(update.Message.Chat.ID, text)
					return c.String(http.StatusOK, "OK")
				}
				if update.Message.Text == "/help" {
					wh.BotService.SendMessage(update.Message.Chat.ID, "Anda dapat melakukan /upgrade untuk mengupgrade akun anda menjadi premium. , dan /limit untuk melihat limit anda saat ini. /credit untuk melihat credit pembuat bot ini.")
					return c.String(http.StatusOK, "OK")
				}

				wh.BotService.SaveMessage(update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

				if user.TextCount > 0 && user.Status == "UNPAID"{
					wh.User.ReduceLimitText(update.Message.From.ID, user.TextCount)
				}

				// answer, _ := wh.OpenAIService.CodeGenerator(update.Message.Text)
				msg, err := wh.BotService.SendMessage(update.Message.Chat.ID, "test answer")
				if err != nil {
					return err
				}
				// Handle the response message returned from the service
				return c.JSON(http.StatusOK, msg)
			}

			wh.BotService.SendMessage(update.Message.Chat.ID, "Mohon maaf, anda sudah melebihi batas penggunaan. Silahkan upgrade ke premium untuk menggunakan layanan ini.")
			return c.String(http.StatusOK, "OK")
		}
    }

    return c.String(http.StatusOK, "OK")
}