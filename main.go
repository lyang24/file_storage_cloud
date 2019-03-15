package main
import(
	"net/http"
	"file_storage_cloud/handler"
)

func main(){
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err:=http.ListenAndServe(":8080:, nil)
	if err!=nil{
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
