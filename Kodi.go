package main

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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
	Idseason   int     `gorm:"column:idSeason"`
	Season     seasons `gorm:"foreignkey:idSeason;association_foreignkey:idSeason"`
	File       files   `gorm:"foreignkey:idFile;association_foreignkey:idFile"`
}

type seasons struct {
	Idseason   int `gorm:"column:idSeason;primary_key"`
	IdShow     int `gorm:"column:idShow"`
	Season     int
	Name       string
	Userrating int
}

type files struct {
	Idfile      int    `gorm:"column:idFile;primary_key"`
	Idpath      int    `gorm:"column:idPath"`
	Strfilename string `gorm:"column:strFilename"`
	Playcount   int    `gorm:"column:playCount"`
	Lastplayed  string `gorm:"column:lastPlayed"`
	Dateadded   string `gorm:"column:dateAdded"`
	Path        path   `gorm:"foreignkey:idPath;association_foreignkey:idPath"`
}

type path struct {
	Idpath         int    `gorm:"column:idPath;primary_key"`
	Strpath        string `gorm:"column:strPath"`
	Strcontent     string `gorm:"column:strContent"`
	Strscraper     string `gorm:"column:strScraper"`
	Strhash        string `gorm:"column:strHash"`
	Scanrecursive  int    `gorm:"column:scanRecursive"`
	Usefoldernames int    `gorm:"column:useFolderNames"`
	Strsettings    string `gorm:"column:strSettings"`
	Noupdate       int    `gorm:"column:noUpdate"`
	Exclude        int    `gorm:"column:exclude"`
	Dateadded      string `gorm:"column:dateAdded"`
	Idparentpath   int    `gorm:"column:idParentPath"`
}

func (Episode) TableName() string {
	return "episode"
}

func (e Episode) deleteChildren() error {
	var err error
	return err
}

func (e Episode) print() {
	fmt.Printf("Episode: %v\n", e)
	fmt.Printf("Season: %v\n", e.Season)
	fmt.Printf("File: %v\n", e.File)
	fmt.Printf("Path: %v\n", e.File.Path)
}

var dbMySql *gorm.DB

func main() {

	dbMySql, err := gorm.Open("mysql", "dan:gismo@tcp(localhost:3306)/MyVideos107?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Printf("failed to connect database on MySql %s", err)
		panic("failed to connect database on MySql")
	} else {
		fmt.Println("Connection established to MySql")
	}
	defer dbMySql.Close()

	dbMySql.SingularTable(true)
	dbMySql.LogMode(true)

	var episodes []Episode
	if err := dbMySql.Debug().Set("gorm:auto_preload", true).Where("idShow = ? And c18 like '%Better Call Saul%'", -1).Find(&episodes).Error; err != nil {
		// if err := dbMySql.Debug().Set("gorm:auto_preload", true).Where("idShow = ?", -1).Find(&episodes).Error; err != nil {
		panic(err)
	} else {
		fmt.Printf("\v\n", episodes)
		// episodes.print()
	}

	for _, episode := range episodes {
		// episode.print()
		dbMySql.Debug().Delete(seasons{}, "idSeason = ?", &episode.Idseason)
		dbMySql.Debug().Delete(path{}, "idPath = ?", &episode.File.Idpath)
		dbMySql.Debug().Delete(files{}, "idFile = ?", &episode.File.Idfile)
		dbMySql.Debug().Delete(Episode{}, "idEpisode = ?", &episode.Idepisode)
	}

	fmt.Println("Done, thank you")
}
