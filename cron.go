// package main

// import (
// 	"log"
// 	"time"

// 	"gin-blog/models"

// 	"github.com/robfig/cron"
// )

// func main() {
// 	log.Println("Starting...")

// 	c := cron.New()
// 	c.AddFunc("* 1 * * * *", func() {
// 		log.Println("Run models.CleanAllTag...")
// 		models.CleanAllTag()
// 	})
// 	c.AddFunc("* 1 * * * *", func() {
// 		log.Println("Run models.CleanAllArticle...")
// 		models.CleanAllArticle()
// 	})

// 	c.Start()

// 	t1 := time.NewTimer(time.Second * 10)
// 	for {
// 		select {
// 		case <-t1.C:
// 			t1.Reset(time.Second * 10)
// 		}
// 	}
// }