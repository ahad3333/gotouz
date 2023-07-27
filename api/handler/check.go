package handler

import (
	"add/config"
	"add/models"
)

func Check(x int) (models.UserLogin) {
	switch {
	case x == 1:
		return models.UserLogin{
			Username: config.LoginGmail,
			Password: config.PasswordGmail,
		}
	case x == 2:
		return models.UserLogin{
			Username: config.LoginGmail1,
			Password: config.PasswordGmail1,
		}
	case x == 3:
		return models.UserLogin{
			Username: config.LoginGmail2,
			Password: config.PasswordGmail2,
		}
	case x == 4:
		return models.UserLogin{
			Username: config.LoginGmail3,
			Password: config.PasswordGmail3,
		}
	default:
		return models.UserLogin{
			Username: "",
			Password: "",
		}
	}
}


// func Check(x string, num string) (models.UserLogin, string) {

// 	if x == "x" && num == "" || num == "4"{
// 		fmt.Println("qadam.....2")
// 		return models.UserLogin{
// 			Username: config.LoginGmail,
// 			Password: config.PasswordGmail,
// 		},"1"
// 	}else if x == "x" && num == "1" {
// 		fmt.Println("qadam.....3")
// 		return models.UserLogin{
// 			Username: config.LoginGmail1,
// 			Password: config.PasswordGmail1,
// 		},"2"
// 	}else if x == "x" && num == "2" {
// 		fmt.Println("qadam.....4")
// 		return models.UserLogin{
// 			Username: config.LoginGmail2,
// 			Password: config.PasswordGmail2,
// 		},"3"
// 	}else if x == "x" && num == "3" {
// 		fmt.Println("qadam.....5")
// 		return models.UserLogin{
// 			Username: config.LoginGmail3,
// 			Password: config.PasswordGmail3,
// 		},"4"
// 	}
// 	fmt.Println("qadam.....6")
// 	return models.UserLogin{
// 		Username: "",
// 		Password: "",
// 	},""
// }
