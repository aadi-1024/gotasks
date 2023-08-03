package cmd

import (
	"fmt"
	"strings"
	"github.com/aadi-1024/gotasks/database"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use: "rm [id]",
	Short: "Remove en element",
	Long: "Remove any task by matching with the id",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id_match := args[0]
		m := make(map[string] string)
		database.Db.View(func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte("Tasks")).ForEach(func(k, v []byte) error {
				m[string(k)] = string(v)
				return nil
			})
			return err
		})
		for k := range m {
			if strings.HasPrefix(k, id_match) {
				database.Db.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("Tasks"))
					err := b.Delete([]byte(k))
					return err
				})	
				break
			}
			fmt.Println("Couldnt find anything that starts with", id_match)
		}
	},
}