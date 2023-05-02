package wget

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func CreateFolder(folderName string) error {
	_, ok := os.Stat(folderName)
	if os.IsExist(ok) {
		return nil
	} else if os.IsNotExist(ok) {
		if ok = os.MkdirAll(folderName, os.ModePerm); ok != nil {
			return ok
		}
	}
	return nil
}

func Wget(link string, wg *sync.WaitGroup) error {
	defer wg.Done()
	link = strings.TrimRight(link, "/")
	mainlink, err := url.ParseRequestURI(link)
	if err != nil {
		return err
	}
	err = CreateFolder(mainlink.Host)
	if err != nil {
		return err
	}
	col := colly.NewCollector(
		colly.AllowedDomains(mainlink.Host, "www."+mainlink.Host),
	)
	col.OnHTML("a[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		_ = col.Visit(ul)

	})

	col.OnHTML("link[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		_ = col.Visit(ul)

	})

	col.OnHTML("script[src]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("src"))
		_ = col.Visit(ul)

	})

	col.OnResponse(func(r *colly.Response) {
		entryName := r.Request.URL.Path
		fullName := mainlink.Hostname() + entryName

		// Если нет расширения, то значит это папка
		if path.Ext(fullName) == "" {
			err = CreateFolder(fullName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			last := strings.LastIndexByte(fullName, '/')
			if last > 0 {
				CreateFolder(fullName[:last])
			}
		}

		// Если имя страницы не имеет расширения, то пишем содержимое в index.html
		if path.Ext(entryName) == "" {
			if fullName[len(fullName)-1] != '/' {
				fullName += "/"
			}
			fullName += "index.html"
		}
		if err = r.Save(fullName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("saved:", mainlink.Hostname()+entryName)
	})

	if err = col.Visit(mainlink.String()); err != nil {
		return err
	}
	return nil
}
