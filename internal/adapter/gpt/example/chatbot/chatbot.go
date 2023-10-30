package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	_ "github.com/sashabaranov/go-openai"
	"os"
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5TurboInstruct,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
