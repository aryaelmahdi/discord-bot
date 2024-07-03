package handler

import (
	"fmt"
	"time"
)

func (handler *DiscordHandlerImpl) Login6AMonday() {
	fmt.Println("masuk automated login")
	_, err := handler.Cron.AddFunc("50 10 * * 1", func() {
		fmt.Println("monday job 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("res A : ", res)
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AMonday2() {
	_, err := handler.Cron.AddFunc("15 15 * * 1", func() {
		fmt.Println("monday job 2 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6ATuesday() {
	_, err := handler.Cron.AddFunc("30 7 * * 2", func() {
		fmt.Println("tuesday job 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AWednesday() {
	_, err := handler.Cron.AddFunc("30 7 * * 3", func() {
		fmt.Println("wednesday job 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AWednesday2() {
	_, err := handler.Cron.AddFunc("50 10 * * 3", func() {
		fmt.Println("wednesday job 2 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AThursday() {
	_, err := handler.Cron.AddFunc("30 7 * * 4", func() {
		fmt.Println("thursday job 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AThursday2() {
	_, err := handler.Cron.AddFunc("15 15 * * 4", func() {
		fmt.Println("thursday job 2 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}

func (handler *DiscordHandlerImpl) Login6AFriday() {
	_, err := handler.Cron.AddFunc("20 13 * * 5", func() {
		fmt.Println("friday job 6A")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot6A(user.Username, user.Password); err != nil {
				fmt.Println(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
	})
	if err != nil {
		fmt.Println("cron job error: ", err)
	}
}
