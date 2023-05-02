package cmd

import (
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/spf13/cobra"
	"gorm-demo/internal/config"
	gormdb "gorm-demo/internal/db"
	"gorm-demo/internal/models"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed database with fake data",
}

var seedUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "seed users table with fake data",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		// Insert fake data into the users table
		fk := faker.New()
		for i := 0; i < 10; i++ {
			user := models.User{
				Username: fk.Person().Name(),
				Email:    fk.Internet().Email(),
				Password: fk.Internet().Password(),
			}
			db.Create(&user)
		}

		fmt.Println("Fake users inserted successfully")
		return nil
	},
}

var seedPostsCmd = &cobra.Command{
	Use:   "posts",
	Short: "seed posts table with fake data",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		users := make([]models.User, 0)
		db.Find(&users)
		fk := faker.New()
		for _, u := range users {
			for i := 0; i < 10; i++ {
				post := models.Post{
					Title:    fk.Company().CatchPhrase(),
					Body:     fk.Lorem().Paragraph(10),
					UserID:   u.ID,
					Comments: []models.Comment{},
					Tags:     []models.Tag{},
				}
				db.Create(&post)
			}
		}
		fmt.Println("Fake posts inserted successfully")
		return nil
	},
}

var seedTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "seed tags table with fake data",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		posts := make([]models.Post, 0)
		db.Preload("User").Find(&posts)
		fk := faker.New()
		for _, p := range posts {
			for i := 0; i < 5; i++ {
				tag := models.Tag{
					Name:  fk.Lorem().Word(),
					Posts: []models.Post{p},
				}
				db.Create(&tag)
			}
		}

		fmt.Println("Fake tags inserted successfully")
		return nil
	},
}

var seedCommentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "seed comments table with fake data",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		posts := make([]models.Post, 0)
		db.Preload("User").Find(&posts)
		fk := faker.New()
		for _, p := range posts {
			for i := 0; i < 5; i++ {
				comment := models.Comment{
					Body:   fk.Lorem().Sentence(5),
					PostID: p.ID,
					UserID: p.UserID,
				}
				db.Create(&comment)
			}
		}

		fmt.Println("Fake comments inserted successfully")
		return nil
	},
}

func init() {
	seedCmd.AddCommand(seedUsersCmd)
	seedCmd.AddCommand(seedPostsCmd)
	seedCmd.AddCommand(seedTagsCmd)
	seedCmd.AddCommand(seedCommentsCmd)
	rootCmd.AddCommand(seedCmd)
}
