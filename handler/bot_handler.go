package handler

import (
	"code-ai/models/domain"
	"code-ai/services"
	"code-ai/utils/helper"
	"fmt"
	"net/http"
	"runtime"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	BotService services.BotService
    OpenAIService services.CodeServiceInterface
	User services.UserServiceInterface
}

func NewBotHandler(botService services.BotService, openAiService services.CodeServiceInterface, user services.UserServiceInterface) *BotHandler {
	return &BotHandler{
		BotService: botService,
        OpenAIService: openAiService,
		User: user,
	}
}

func (wh *BotHandler) HandleBot(ctx echo.Context) error {
	var updateChan = make(chan *tgbotapi.Update)
	go func() error {
		fmt.Println(time.Now())
		update := new(tgbotapi.Update)
		if err := ctx.Bind(update); err != nil {
			return err
		}
		updateChan <- update
		return nil
	}()
	update := <-updateChan
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
					Credit: 10,
				}
				user, _ = wh.User.CreateUser(&createUser)
			}
			userFound <- user
		}()

		if update.Message.Photo != nil { 
			msg, err := wh.BotService.SendMessage(update.Message.Chat.ID, "Tidak dapat menerima gambar.")
			if err != nil {
				return err
			}
			return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
		}

		if update.Message.Photo == nil{
			user := <- userFound

			if user.Credit > 0 || user.Status == "PAID" {
				if update.Message.Text == "/start" {
					msg, _ :=wh.BotService.SendMessage(update.Message.Chat.ID, "Halo, "+update.Message.From.UserName+"! Silahkan masukkan kode yang anda inginkan. Untuk melihat daftar perintah, silahkan ketik /help")
					return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
				}
				if update.Message.Text == "/upgrade" {
					msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, "Berikut Harga Credit yang tersedia:\n\nRp. 30.000,00 per 250 Credit\nRp. 50.000,00 per 500 Credit\nRp. 100.000,00 per 1200 Credit\n\ntransfer ke rekening 1234567890 a/n Innaka Dylee. Setelah itu, \nkirim bukti transfer ke @InnakaDylee.")
					return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
				}
				if update.Message.Text == "/credit" {
					msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, "Bot ini dibuat oleh @InnakaDylee.")
					return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
				}
				if update.Message.Text == "/limit" {
					var text string
					if user.Status == "PAID" {
						text = fmt.Sprintf("Limit anda tidak terbatas.")
					} else {
						text = fmt.Sprintf("Limit anda tersisa %d kali.", user.Credit)
					}
					msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, text)
					return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
				}
				if update.Message.Text == "/help" {
					msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, "Anda dapat melakukan /upgrade untuk mengupgrade akun anda menjadi premium. , dan /limit untuk melihat limit anda saat ini. /credit untuk melihat credit pembuat bot ini.")
					return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
				}

				wh.BotService.SaveMessage(update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

				if user.Credit > 0 && user.Status == "UNPAID"{
					wh.User.ReduceLimitText(update.Message.From.ID, user.Credit)
				}

				var message = make(chan *domain.Message)

				go func() {
					answer, _ := wh.OpenAIService.CodeGenerator(update.Message.Text)
					msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, answer)
					message <- msg
				}()
				msg := <- message

				// Handle the response message returned from the service
				return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
			}

			msg, _ := wh.BotService.SendMessage(update.Message.Chat.ID, "Mohon maaf, anda sudah melebihi batas penggunaan. Silahkan upgrade ke premium untuk menggunakan layanan ini.")
			return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", msg))
		}
    }

    return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Send Message", "OK"))
}