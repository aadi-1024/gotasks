package cmd

import (
	"fmt"
	"github.com/aadi-1024/gotasks/database"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

)

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get all tasks",
	Long: "Get all active tasks",
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		database.Db.View(func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte("Tasks")).ForEach(func(k, v []byte) error {
				fmt.Printf("%v - %v\n", string(k), string(v))
				return nil
			})
			return err
		})
	},
}