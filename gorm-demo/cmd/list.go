package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gorm-demo/internal/config"
	gormdb "gorm-demo/internal/db"
	"gorm-demo/internal/models"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all records from the database",
}

var listUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all records from the users table",
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "User Name", "Email"})
		for _, user := range users {
			table.Append([]string{
				fmt.Sprintf("%d", user.ID),
				user.Username,
				user.Email,
			})
		}
		table.Render()

		return nil
	},
}

var listPostsCmd = &cobra.Command{
	Use:   "posts",
	Short: "List all records from the posts table",
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Body", "User"})
		for _, post := range posts {

			table.Append([]string{
				fmt.Sprintf("%d", post.ID),
				post.Title,
				post.Body[:20],
				post.User.Username,
			})
		}
		table.Render()

		return nil
	},
}

var listCommentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "List all records from the comments table",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		comments := make([]models.Comment, 0)
		db.Preload("Post").Find(&comments)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Body", "Post"})
		for _, comment := range comments {
			table.Append([]string{
				fmt.Sprintf("%d", comment.ID),
				comment.Body[:20],
				comment.Post.Title,
			})
		}
		table.Render()
		return nil
	},
}

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all records from the tags table",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		tags := make([]models.Tag, 0)
		db.Find(&tags)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Created At"})
		for _, tag := range tags {
			table.Append([]string{
				fmt.Sprintf("%d", tag.ID),
				tag.Name,
				tag.CreatedAt.String(),
			})
		}
		table.Render()
		return nil
	},
}

func init() {
	listCmd.AddCommand(listUsersCmd)
	listCmd.AddCommand(listPostsCmd)
	listCmd.AddCommand(listCommentsCmd)
	listCmd.AddCommand(listTagsCmd)
	rootCmd.AddCommand(listCmd)
}
