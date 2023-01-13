# Job Application

* Create DB named `zip_dev`
* Run `cd migrations && goose mysql "<username>:<pswrd>@/zip_dev" up`
* Run `go run main.go`
* Test cURL to zip a file `curl --location --request POST 'localhost:3000/api/v1/zip' \
  --header 'Content-Type: multipart/form-data; boundary="myFile"' \
  --form 'myFile=@"<FULL_PATH_TO_FILE>"'`
* Test cURL to get a list of files `curl --location --request GET 'localhost:3000/api/v1/files'`
* Grep `filename` from previous request and use it to download the file. cURL: `curl --location --request GET 'localhost:3000/api/v1/files/ee21ed2d-e226-4467-bcad-c11ef2d7e0df'` 


* Не реализовал множественную загрузку файлов, т.к. итак много времени потратил, часов 5 на проект