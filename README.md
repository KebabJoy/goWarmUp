# Job Application

* Create DB named `zip_dev`
* Run `cd migrations && goose mysql "<username>:<pswrd>@/zip_dev" up`
* Run `go run main.go`
* Test cURL to zip a file `curl --location --request POST 'localhost:3000/api/v1/zip' \
  --header 'Content-Type: multipart/form-data; boundary="myFile"' \
  --form 'myFile=@"<FULL_PATH_TO_FILE>"'`
