package main

func main() {
	url := "http://localhost:8080/videos"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic xxxx")
	req.Header.Add("cache-control",  "no-cache")

	res, _ : http.DefaultClient.Do(req)

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	
}