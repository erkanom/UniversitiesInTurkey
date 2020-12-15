package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type uni struct {
	name  string
	id    int
	depts []int
}

// type dept struct {
// 	id   int
// 	name string
// }

func main() {
	url := "https://yokatlas.yok.gov.tr/lisans-anasayfa.php"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Page not found")
		panic(err)

	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	strContent := string(content[:])
	scanner := bufio.NewScanner(strings.NewReader(strContent))
	index := 0
	mapOfUni := make(map[int]uni)
	for scanner.Scan() {
		var temp string
		temp = scanner.Text()
		if strings.Contains(temp, "option value") && strings.Contains(temp, "ÜNİVERSİTESİ") {
			temp = strings.Trim(temp, " ")
			temp = strings.TrimLeft(temp, " ")
			temp = strings.TrimLeft(temp, "															")
			temp = strings.ReplaceAll(temp, "<option value=\"", "")
			temp = strings.ReplaceAll(temp, "\">", " ")
			temp = strings.ReplaceAll(temp, "</option>", "")
			temp = strings.ReplaceAll(temp, "<", "")
			num, err := strconv.Atoi(temp[0:4])
			if err != nil {
				fmt.Print("atoi not working")
			}
			var tempUni = new(uni)
			tempUni.id = num
			tempUni.name = temp[5:]
			mapOfUni[index] = *tempUni
		}
	}
	for _, value := range mapOfUni {
		getDep(&value)
	}
}
func getDep(uni *uni) {

	newUrl := "https://yokatlas.yok.gov.tr/lisans-univ.php?u="

	urlId := strconv.Itoa(uni.id)

	newUrl += urlId
	resp, err := http.Get(newUrl)
	if err != nil {
		log.Fatal("Page not found")
		panic(err)

	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	strContent := string(content[:])
	scanner := bufio.NewScanner(strings.NewReader(strContent))
	//index := 0
	for scanner.Scan() {
		var temp string
		temp = scanner.Text()
		if strings.Contains(temp, "lisans.php?y=") {
			temp = strings.ReplaceAll(temp, "<a data-parent=\"#\" href=\"lisans.php?y=", " ")
			temp = strings.ReplaceAll(temp, "\">", "")
			temp = strings.Trim(temp, " ")
			tempInt, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Print("convert error string to int ")
			}
			uni.depts = append(uni.depts, tempInt)
		}
	}
}
