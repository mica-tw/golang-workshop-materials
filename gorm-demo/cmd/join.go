package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gorm-demo/internal/config"
	gormdb "gorm-demo/internal/db"
	"os"

	"github.com/spf13/cobra"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join users, posts and tags tables",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		var result []struct {
			UserName  string `gorm:"column:username"`
			PostID    uint   `gorm:"column:post_id"`
			PostTitle string `gorm:"column:post_title"`
			TagName   string `gorm:"column:tag_name"`
		}

		if err := db.Table("users").
			Select("users.username, posts.id as post_id, posts.title as post_title, tags.name as tag_name").
			Joins("INNER JOIN posts ON posts.user_id = users.id").
			Joins("INNER JOIN post_tags ON post_tags.post_id = posts.id").
			Joins("INNER JOIN tags ON tags.id = post_tags.tag_id").
			Scan(&result).Error; err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"User", "Post ID", "Post", "Tag"})
		for _, item := range result {
			table.Append([]string{
				item.UserName,
				fmt.Sprintf("%d", item.PostID),
				item.PostTitle,
				item.TagName,
			})
		}
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
}
