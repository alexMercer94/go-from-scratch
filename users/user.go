package users

import (
	"fmt"
	"time"

	"github.com/go-from-scratch/models"
)

func AddUser() {
	us := new(models.User)
	us.AddUser(10, "Alex", time.Now(), true)
	fmt.Println(us)
}
