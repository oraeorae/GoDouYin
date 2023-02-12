package test

import (
	"bufio"
	"fmt"
	"github.com/syyongx/go-wordsfilter"
	"go_douyin/utils/sensitive_word_filter"
	"os"
	"testing"
)

//测试敏感词过滤
func TestSensitive(t *testing.T) {
	// 创建前缀树
	trie := sensitive_word_filter.NewTrie()

	// 从文件中读取敏感词
	file, _ := os.Open("../config/sensitive_words.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trie.Insert(scanner.Text())
	}

	// 过滤敏感词
	text := "这是一段评论，里面有敏感词,快手"
	fmt.Println("原始评论:", text)
	filteredText := trie.Filter(text)
	fmt.Println("过滤后评论:", filteredText)
}

func TestSensitive2(t *testing.T) {
	wf := wordsfilter.New()
	root, _ := wf.GenerateWithFile("../config/sensitive_words.txt")
	//替换
	newStr := wf.Replace("不可描快", root)
	//打印替换好的文本
	fmt.Println(newStr)
}
