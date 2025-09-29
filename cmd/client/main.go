package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/httputils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("cmd/client/main: failed to load env")
	}

	fmt.Println("Введите длинный URL")
	reader := bufio.NewReader(os.Stdin)
	long, err := reader.ReadString('\n')
	if err != nil {
		slog.Error("cmd/client/main: " + err.Error())

		return
	}

	long = strings.TrimSuffix(long, "\n")
	urls := url.Values{
		"url": strings.Split(long, " "),
	}

	client := req.C()
	res := client.Post(env.URL()).
		SetHeader(httputils.ContentType, httputils.URLEncoded).
		SetFormDataFromValues(urls).
		Do()

	if res.Err != nil {
		slog.Error("cmd/client/main: " + res.Err.Error())

		return
	}

	fmt.Println("Статус-код " + res.Status)
	fmt.Println(res.String())
}
