package main

import (
	"encoding/json"
	"github.com/Azunyan1111/block-site/s"
	"honnef.co/go/js/dom"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

func main() {
	// URLの正規表現とGoogleのURLの正規表現
	u,_ := regexp.Compile(`^http(s?)://(.*)(/?)`)
	uGoogle,_ := regexp.Compile(`^http(s?)://(.*)google(.*)(/?)`)
	uGoogle2,_ := regexp.Compile(`^http(s?)://(.*)googleusercontent(.*)(/?)`)

	// DOMを操作する ドキュメントを取得
	d := dom.GetWindow().Document()

	// Gクラスが存在する場合
	gClass := d.GetElementsByClassName("g")
	// 要素を格納する変数
	//var elementsG []dom.Element

	//
	for _, g := range gClass{
		// aタグを探す
		for i, a := range g.GetElementsByTagName("a"){
			href := a.GetAttribute("href")
			if u.MatchString(href){
				if !uGoogle.MatchString(href){
					if !uGoogle2.MatchString(href){
						// ここに来るのは対象となるエレメンツのみとなる
						siteReport := GetSiteReport(href)
						if siteReport.Block{
							g.SetAttribute("style","display:none")
							g.SetAttribute("my",strconv.Itoa(i))
							text := `
ここに表示されたサイトはクソサイトなので非表示にされました。<br>
再度表示する場合はこのエリアをクリック<br>
投票結果  Good:` + strconv.Itoa(siteReport.Good) + `  Bad:` + strconv.Itoa(siteReport.Bad)

							// Div 要素を作成
							div := d.CreateElement("div")
							// テキストを注入
							div.SetInnerHTML(text)
							// 自身にタグをつける
							div.SetAttribute("my","h-" + strconv.Itoa(i))
							div.SetAttribute("style","border: 1px solid;")
							// onclickで表示するスクリプトを注入
							div.SetAttribute("onclick",`document.querySelector("div[my='` + strconv.Itoa(i) + `']").style.display = "block";document.querySelector("div[my='h-0']").style.display="none"`)
							// 自身の子要素に非表示した要素を注入
							//div.SetOuterHTML(div.OuterHTML()+ g.OuterHTML())
							g.SetOuterHTML(div.OuterHTML() + g.OuterHTML())


							//g.SetOuterHTML(divStart + g.OuterHTML() + "</div>")
							//g.SetAttribute()
							//js.Global.Call("alert", g)
							//g.SetAttribute("my","h-" + strconv.Itoa(i))
							//js.Global.Call("alert", g)
							//g.SetOuterHTML(e.OuterHTML())
							//js.Global.Call("alert", g)
						}

					}
				}
			}
		}
	}


	/*/




	// Gクラスがなかった場合はすべてのAタグ

	// すべてのAタグを取得
	AtagElements := d.GetElementsByTagName("a")

	// ブロッカーに送信するURLだけを抽出する
	for _, element := range AtagElements {
		href := element.GetAttribute("href")
		if u.MatchString(href){
			if !uGoogle.MatchString(href){
				if !uGoogle2.MatchString(href){
					elementsG = append(elementsG, element)
				}
			}
		}
	}

	for _,e := range elementsG{
		ur := url.Values{}
		ur.Set("",e.GetAttribute("href"))
		resp,err := http.Get("http://150.95.213.47:8080/report/" + ur.Encode())
		if err != nil{
			println(err.Error())
			continue
		}
		var rs s.ResponseSite
		if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil{
			println(err.Error())
			continue
		}
		println(rs)
		e.SetInnerHTML("※ここだよ！※" +e.InnerHTML())
	}
	//*/
}

func GetSiteReport(u string)s.ResponseSite{
	ur := url.Values{}
	ur.Set("",u)
	resp,err := http.Get("http://150.95.213.47:8080/report/" + ur.Encode()[1:])
	if err != nil{
		println(err.Error())
		return s.ResponseSite{}
	}
	var rs s.ResponseSite
	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil{
		println(err.Error())
		return s.ResponseSite{}
	}
	return rs
}

