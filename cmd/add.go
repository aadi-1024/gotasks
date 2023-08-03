package cmd

import (
	// "os"
	"fmt"
	"hash/fnv"

	"github.com/aadi-1024/gotasks/database"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [task]",
	Short: "Add a new task",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := ""
		for _, i := range args {
			task += i + " "
		}
		database.Db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			f := fnv.New32a()
			f.Write([]byte(task))
			err := b.Put([]byte(fmt.Sprint(f.Sum32())), []byte(task))
			if err == nil {
				fmt.Printf("Added task - %v", task)
			}
			return err
		})
	},
}
