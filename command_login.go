package main

import (
  "fmt"
  "net/http"
  "io"
)

func commandLogin() error {
	url := "https://wiki.wizard101central.com/wiki/api.php?action=query&meta=tokens&type=login&format=json"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Cookie", "wikidb3_mw__session=eo3v1utgpnpfn5ltnnet9v8r3tc8jlpl")
	req.Header.Add("User-Agent", "WizWord101/0.0 (https://github.com/samassembly/WizardingWordle101)")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}