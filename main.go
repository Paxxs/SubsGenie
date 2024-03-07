package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	// cspell:words joho godotenv
	MyLog "github.com/Paxxs/SubsGenie/myLog"
	"github.com/joho/godotenv"
)

var mLog *MyLog.MyLogger

func main() {
	// cspell:words subconvert
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		os.Exit(1)
	}

	mLog = MyLog.NewMyLogger()

	mLog.Info(":: 读取配置")
	allSubscriptions, onlyCoreAndCf, requestUrlTemplate, gistID, githubToken := readConfig()

	var configText string
	// var err error
	// try first with allSub
	mLog.Info(":: 开始获取订阅 - ALL")
	if configText, err = fetchConfig(fmt.Sprintf(requestUrlTemplate, allSubscriptions)); err != nil {
		// if the first attempt fails, try with onlyCoreAndCf
		mLog.Info("订阅ALL获取失败：", err)
		mLog.Info(":: 开始获取订阅 - CoreAndCF")
		if configText, err = fetchConfig(fmt.Sprintf(requestUrlTemplate, onlyCoreAndCf)); err != nil {
			fmt.Println("订阅CoreAndCF获取失败：", err)
		}
	}
	// fmt.Println(configText)
	mLog.Debug("看看前几行：\n", ReadFirstNLines(configText, 5))
	if gistID != "" && githubToken != "" {
		// 获取当前时间，并按照指定格式格式化
		now := time.Now().In(loc)
		formattedTime := now.Format("06/01/02 15:04:05")
		// 将格式化后的时间字符串插入到给定字符串中
		descriptionContent := fmt.Sprintf("clash 订阅文件 - %s [SubsGenie]", formattedTime)
		err = UpdateGist(githubToken, gistID, descriptionContent, "coreAndCF.txt", configText)
		if err != nil {
			fmt.Println("failed upload to gist:", err)
			os.Exit(1)
		}
	}
}

// readConfig returns the configuration settings for the application.
//
// It retrieves the values of the environment variables CORE_SUBSCRIPTION_URL,
// CF_SUBSCRIPTION_URL, OTHER_SUBSCRIPTION_URLS, CONFIG_URL, EXTRA_PARAMS, and
// SUBCONVERT_SERVICE_URL. If CORE_SUBSCRIPTION_URL is empty, it prints an error
// message and exits the program. If CONFIG_URL is empty, it sets a default value.
// The function then constructs the request URL template and returns the
// subscriptions and the request URL template.
//
// Returns:
// - allSubscriptions: a string containing all the subscription URLs.
// - onlyCoreAndCf: a string containing the core and CF subscription URLs.
// - requestUrlTemplate: a string containing the request URL template.
func readConfig() (allSubscriptions, onlyCoreAndCf, requestUrlTemplate, gistID, githubToken string) {
	coreSubScription := os.Getenv("CORE_SUBSCRIPTION_URL")
	cfSubScription := os.Getenv("CF_SUBSCRIPTION_URL")
	otherSubScription := os.Getenv("OTHER_SUBSCRIPTION_URLS")
	gistID = os.Getenv("GIST_ID")
	githubToken = os.Getenv("GITHUB_TOKEN")

	mLog.Debug(fmt.Sprintf("读取环境变量\n【CORE_SUBSCRIPTION_URL】:%s\n【CF_SUBSCRIPTION_URL】:%s\n【OTHER_SUBSCRIPTION_URLS】:%s\n【GIST_ID】:%s\n【GITHUB_TOKEN】:%s\n", coreSubScription, cfSubScription, otherSubScription, gistID, githubToken))

	configUrl := os.Getenv("CONFIG_URL")
	extraParams := os.Getenv("EXTRA_PARAMS")
	subconvertServiceUrl := os.Getenv("SUBCONVERT_SERVICE_URL")
	if coreSubScription == "" {
		fmt.Println("CORE_SUBSCRIPTION is required")
		os.Exit(1)
	}
	if configUrl == "" {
		configUrl = "https://raw.githubusercontent.com/Paxxs/subconverter-clash-rule/main/ACL4SSR_Online_Full_AdblockPlus.ini"
	}
	mLog.Debug(fmt.Sprintf("读取环境变量\n【CONFIG_URL】:%s\n【EXTRA_PARAMS】:%s\n【SUBCONVERT_SERVICE_URL】:%s", configUrl, extraParams, subconvertServiceUrl))

	configUrl = url.QueryEscape(configUrl)

	if subconvertServiceUrl == "" {
		subconvertServiceUrl = "https://api.dler.io/sub"
	}

	allUrls := []string{coreSubScription}
	if cfSubScription != "" {
		allUrls = append(allUrls, cfSubScription)
	}
	if otherSubScription != "" {
		allUrls = append(allUrls, strings.Split(otherSubScription, ",")...)
	}
	mLog.Debug("合并链接：", allUrls)

	// 全部的订阅
	allSubscriptions = url.QueryEscape(strings.Join(allUrls, "|"))

	// 核心和大善人
	onlyCoreAndCf = coreSubScription
	if cfSubScription != "" {
		onlyCoreAndCf += "|" + cfSubScription
	}
	onlyCoreAndCf = url.QueryEscape(onlyCoreAndCf)

	mLog.Debug(fmt.Sprintf("编码后的订阅链接：\n【allSubscriptions】:%s\n【onlyCoreAndCf】:%s", allSubscriptions, onlyCoreAndCf))

	// 生成url模板
	requestUrlTemplate = fmt.Sprintf("%s?target=clash&url=%%s&config=%s%s", subconvertServiceUrl, configUrl, extraParams)

	mLog.Debug(fmt.Sprintf("生成的url模板\n【requestUrlTemplate】:%s", requestUrlTemplate))
	return allSubscriptions, onlyCoreAndCf, requestUrlTemplate, gistID, githubToken
}

// fetchConfig fetches the config from the given requestUrl.
//
// requestUrl string
// (string, error)
func fetchConfig(requestUrl string) (string, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	stringBody := string(body)
	if strings.HasPrefix(stringBody, "The following link") {
		return "", fmt.Errorf(stringBody)
	}

	if !strings.HasPrefix(stringBody, "port: 7890") {
		fmt.Println("🫥 内容可能不对")
	}
	return stringBody, nil
}

// UpdateGist 更新一个GitHub Gist的描述和文件内容。
// 需要提供GitHub的访问令牌（YOUR_TOKEN），Gist的ID（GIST_ID），文件名（FileName）和内容（ContentText）。
func UpdateGist(token, gistID, description, fileName, contentText string) error {
	url := fmt.Sprintf("https://api.github.com/gists/%s", gistID)
	// payload := fmt.Sprintf(`{"description":"%s","files":{"%s":{"content":"%s"}}}`, description, fileName, contentText)

	type GistFileContent struct {
		Content string `json:"content"`
	}

	type GistUpdateRequest struct {
		Description string                     `json:"description"`
		Files       map[string]GistFileContent `json:"files"`
	}

	// 构造请求体
	requestData := GistUpdateRequest{
		Description: description,
		Files: map[string]GistFileContent{
			fileName: {
				Content: contentText,
			},
		},
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json") // 确保设置Content-Type为application/json

	client := &http.Client{}

	mLog.Info(":: 开始上传至gist")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体（DEBUG输出）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 判断是否请求成功
	stringBody := string(body)
	if resp.StatusCode == 404 {
		return fmt.Errorf("Github Gist %s not found", gistID)
	} else if resp.StatusCode == 422 {
		return fmt.Errorf("Validation failed, or the endpoint has been spammed.")
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Github API error:(%s) %s", resp.Status, stringBody)
	}
	mLog.Debug("更新 gist 成功：", stringBody)
	mLog.Info("上传至gist成功")
	return nil
}

// ReadFirstNLines reads and returns the first N lines of the given string.
//
// s: the string to read from
// n: the number of lines to read
// string: the first N lines of the string
func ReadFirstNLines(s string, n int) string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) >= n {
			break
		}
	}
	return strings.Join(lines, "\n")
}
