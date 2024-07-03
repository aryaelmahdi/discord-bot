package handler

import (
	"fmt"
	"log"
	"time"
)

func (handler *DiscordHandlerImpl) Login6BTuesday() {
	_, err := handler.Cron.AddFunc("30 7 * * 2", func() {
		fmt.Println("tuesday job 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		fmt.Println("res B : ", res)
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				log.Printf("error: %v", err)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BTuesday2() {
	_, err := handler.Cron.AddFunc("50 10 * * 2", func() {
		fmt.Println("tuesday job 2 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				log.Printf("error: %v", err)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BWednesday() {
	_, err := handler.Cron.AddFunc("30 7 * * 3", func() {
		fmt.Println("wednesday job 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BWednesday2() {
	_, err := handler.Cron.AddFunc("50 10 * * 3", func() {
		fmt.Println("wednesday job 2 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BWednesday3() {
	_, err := handler.Cron.AddFunc("15 15 * * 3", func() {
		fmt.Println("wednesday job 3 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BThursday() {
	_, err := handler.Cron.AddFunc("30 7 * * 4", func() {
		fmt.Println("thursday job 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BThursday2() {
	_, err := handler.Cron.AddFunc("50 10 * * 4", func() {
		fmt.Println("thursday job 2 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (handler *DiscordHandlerImpl) Login6BThursday3() {
	_, err := handler.Cron.AddFunc("15 15 * * 5", func() {
		fmt.Println("thursday job 3 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6B(user.Username, user.Password); err != nil {
				log.Printf("error: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}
