package main

import (
	"fmt"
	"regexp"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	new, _ = gse.New("zh,testdata/test_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func main() {
	// 加载默认词典
	seg.LoadDict()
	// 加载默认 embed 词典
	// seg.LoadDictEmbed()
	//
	// 加载简体中文词典
	// seg.LoadDict("zh_s")
	// seg.LoadDictEmbed("zh_s")
	//
	// 加载繁体中文词典
	// seg.LoadDict("zh_t")
	//
	// 加载日文词典
	// seg.LoadDict("jp")
	//
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	cut()

	segCut()
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", new.Analyze(hmm, text))

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	reg := regexp.MustCompile(`(\d+年|\d+月|\d+日|[\p{Latin}]+|[\p{Hangul}]+|\d+\.\d+|[a-zA-Z0-9]+)`)
	text1 := `헬로월드 헬로 서울, 2021年09月10日, 3.14`
	hmm = seg.CutDAG(text1, reg)
	fmt.Println("Cut with hmm and regexp: ", hmm, hmm[0], hmm[6])
}

func analyzeAndTrim(cut []string) {
	a := seg.Analyze(cut, "")
	fmt.Println("analyze the segment: ", a)

	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))
}

func cutPos() {
	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)

	posSeg.WithGse(seg)
	po = posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func segCut() {
	// 分词文本
	tb := []byte("山达尔星联邦共和国联邦政府")

	// 处理分词结果
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(string(tb), true))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(string(tb)))

	segments := seg.Segment(tb)
	// 处理分词结果, 普通模式
	fmt.Println(gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	// 搜索模式
	fmt.Println(gse.ToString(segments1, true))
}
