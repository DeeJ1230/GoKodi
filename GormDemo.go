package main

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tvshow struct {
	Idshow     int `gorm:"column:idShow;primary_key"`
	C00        string
	C01        string
	C02        string
	C03        string
	C04        string
	C05        string
	C06        string
	C07        string
	C08        string
	C09        string
	C10        string
	C11        string
	C12        string
	C13        string
	C14        string
	C15        string
	C16        string
	C17        string
	C18        string
	C19        string
	C20        string
	C21        string
	C22        string
	C23        string
	Userrating int
	Duration   int
	Episodes   []Episode `gorm:"foreignkey:idShow;association_foreignkey:idShow"`
}

type Episode struct {
	Idepisode  int `gorm:"column:idEpisode;primary_key"`
	Idfile     int `gorm:"column:idFile"`
	C00        string
	C01        string
	C02        string
	C03        string
	C04        string
	C05        string
	C06        string
	C07        string
	C08        string
	C09        string
	C10        string
	C11        string
	C12        string
	C13        string
	C14        string
	C15        string
	C16        string
	C17        string
	C18        string
	C19        string
	C20        string
	C21        string
	C22        string
	C23        string
	Idshow     int `gorm:"column:idShow"`
	Userrating int
	Idseason   int `gorm:"column:idSeason"`
}

/*
func (Tvshow) TableName() string {
   return "tvshow"
}
*/

func main() {

	dbMySql, err := gorm.Open("mysql", "dan:gismo@tcp(localhost:3306)/MyVideos75?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Printf("failed to connect database on MySql %s", err)
		panic("failed to connect database on MySql")
	} else {
		fmt.Println("Connection established to MySql")
	}
	defer dbMySql.Close()

	dbMySql.SingularTable(true)
	dbMySql.LogMode(true)

	var serie Tvshow
	if err := dbMySql.Debug().Where("idShow = ?", 666).First(&serie).Error; err != nil {
		checkerror(err)
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", serie)
		// fmt.Printf("%s\n", serie.Episodes)
	}

	var episode Episode
	if err := dbMySql.Debug().Where("idShow = ?", 666).First(&episode).Error; err != nil {
		panic(err)
	} else {
		fmt.Println(episode)
	}

	//var ep []Episode
	if err = dbMySql.Model(&serie).Related(&serie.Episodes, "Episodes").Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", serie)
		fmt.Printf("%s\n", serie.Episodes)
		//fmt.Printf("%s\n", ep)
	}

	fmt.Println("Deleting")
	if err = dbMySql.Debug().Delete(&serie).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", serie)
	}

	/*
	   if err = dbMySql.Model(&serie).Related(&Episode{}).Error; err != nil {
	      fmt.Println(err)
	   } else {
	      fmt.Printf("%s\n", serie)
	      fmt.Printf("%s\n", serie.Episodes)
	   }
	*/

	/*
	   if err := dbMySql.Debug().Model(&serie).Association("Episodes").Find(&serie.Episodes).Error; err != nil {
	      fmt.Println(err)
	   } else {
	      fmt.Printf("%s\n", serie)
	      fmt.Printf("%s\n", serie.Episodes)
	   }
	*/
}

/*
func checkerror(err Error) {
	panic(err)
}
*/
