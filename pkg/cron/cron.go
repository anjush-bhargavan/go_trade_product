package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/anjush-bhargavan/go_trade_product/config"
	"github.com/anjush-bhargavan/go_trade_product/pkg/clients/rabbitmq"
	userpb "github.com/anjush-bhargavan/go_trade_product/pkg/clients/user/pb"
	"github.com/anjush-bhargavan/go_trade_product/pkg/model"
	inter "github.com/anjush-bhargavan/go_trade_product/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/go_trade_product/utils"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type CronInstance struct {
	cron       *cron.Cron
	Repo       inter.ProductRepoInter
	UserClient userpb.UserServiceClient
	redis      *config.RedisService
}

func NewCron(repo inter.ProductRepoInter, userClient userpb.UserServiceClient, redis *config.RedisService) *CronInstance {
	return &CronInstance{
		cron:       cron.New(),
		Repo:       repo,
		UserClient: userClient,
		redis:      redis,
	}
}

func (c *CronInstance) InitCron() {
	c.cron.Start()
}

// ScheduleJob schedules a job to run at a specific time
func (c *CronInstance) ScheduleJob(scheduleTime time.Time, id uint) error {
	return c.cron.AddFunc(scheduleTime.Format("05 04 15 02 01 *"), func() {
		c.FindWinner(id)
	})
}

func (c *CronInstance) FindWinner(productId uint) {
	result, err := c.Repo.FindBidsOfProduct(productId)
	if err != nil {
		log.Println("error finding the bid table of product", err)
	}

	product, err := c.Repo.FindProductByID(productId)
	if err != nil {
		log.Println("error finding the product", err)
	}

	//Unmarshal the data
	var bidTable []model.Bids
	err = json.Unmarshal([]byte(result.Bids), &bidTable)
	if err != nil {
		log.Println("error unmarshaling the bid table of product", err)
	}

	subject := fmt.Sprintf("Congratulations! You've Won the Bid for %s", product.Name)

	for {
		ctx := context.Background()
		err = nil
		winner := utils.Pop(&bidTable)

		user, err := c.UserClient.ViewProfile(ctx, &userpb.ID{ID: uint32(winner.UserID)})
		if err != nil {
			log.Println("error finding the user profile by connecting user client", err)
		}

		msg := utils.NotifyWinner(*product, user.User_Name, winner.Amount)

		rabbitmq.SendMail(user.Email, msg, subject)

		winnerData, err := json.Marshal(&winner)
		if err != nil {
			log.Println("error in marshaling winner data", err)
		}

		key := fmt.Sprintf("winner of %v is %v",productId,winner.UserID)
		err = c.redis.SetDataInRedis(key, winnerData, time.Minute*30)
		if err != nil {
			log.Println("error in setiing data in redis", err)
		}

		time.Sleep(time.Minute * 30)

		_, err = c.Repo.FindOrder(winner.UserID, productId)
		if err == nil {
			return
		} else if err != gorm.ErrRecordNotFound {
			log.Println("the user did not completed order", err)
		}
	}
}
