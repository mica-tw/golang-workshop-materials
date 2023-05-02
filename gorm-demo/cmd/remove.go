package cmd

import (
	"fmt"
	"gorm-demo/internal/config"
	gormdb "gorm-demo/internal/db"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete all records from all the tables",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := gormdb.NewDB(cfg)
		if err != nil {
			return err
		}

		// Delete all records from the tables
		if err := db.Exec("DELETE FROM post_tags").Error; err != nil {
			return err
		}

		if err := db.Exec("DELETE FROM comments").Error; err != nil {
			return fmt.Errorf("failed to delete comments: %v", err)
		}
		if err := db.Exec("DELETE FROM tags").Error; err != nil {
			return fmt.Errorf("failed to delete tags: %v", err)
		}
		if err := db.Exec("DELETE FROM posts").Error; err != nil {
			return fmt.Errorf("failed to delete posts: %v", err)
		}
		if err := db.Exec("DELETE FROM users").Error; err != nil {
			return fmt.Errorf("failed to delete users: %v", err)
		}

		fmt.Println("All records deleted successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
