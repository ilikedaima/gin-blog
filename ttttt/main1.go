package ttttt

// import (
// 	"fmt"
// 	"gin-blog/pkg/logging"
// 	"gin-blog/pkg/setting"
// 	"gin-blog/routers"
// 	"net/http"
// 	"syscall"

// 	"github.com/fvbock/endless"
// )

// func main() {
// 	endless.DefaultReadTimeOut = setting.ReadTimeout
//     endless.DefaultWriteTimeOut = setting.WriteTimeout
//     endless.DefaultMaxHeaderBytes = 1 << 20
//     endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

//     router := routers.InitRouter()
// 	server := endless.NewServer(endPoint, router)
//     server.BeforeBegin = func(add string) {
//         logging.Info("Actual pid is %d", syscall.Getpid())
//     }
// 	s := &http.Server{
// 		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
// 		Handler:        router,
// 		ReadTimeout:    setting.ReadTimeout,
// 		WriteTimeout:   setting.WriteTimeout,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	server.Server = *s

// 	err := server.ListenAndServe()
// 	if err != nil {
//         logging.Info("Server err: %v", err)
//     }
// }