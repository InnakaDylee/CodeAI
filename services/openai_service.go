package services

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type CodeService interface {
	CodeGenerator(userInput string) (string, error)
}

type codeService struct{
	Token string
}

func NewCodeService(token string) CodeService {
	return &codeService{
		Token: token,
	}
}

func (cs *codeService) CodeGenerator(userInput string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(cs.Token)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Anda adalah code generator, anda hanya membantu user menjawab menggunakan output code dari bahasa apapun, jika user tidak meminta spesifik bahasa programming apapun maka anda akan menggunakan bahasa golang sebagai default",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := cs.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (cs *codeService) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}