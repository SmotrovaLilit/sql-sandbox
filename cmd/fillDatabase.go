/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// fillDatabaseCmd represents the fillDatabase command
var fillDatabaseCmd = &cobra.Command{
	Use:   "fillDatabase",
	Short: "Fill database with test data",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := db.Save(&Book{Name: "Harry Potter", Author: "J.K. Rowling", ID: 1}).Error
		if err != nil {
			return err
		}
		err = db.Save(&Book{Name: "The Lord of the Rings", Author: "J.R.R. Tolkien", ID: 2}).Error
		if err != nil {
			return err
		}
		err = db.Save(&User{Name: "John Doe", Username: "johndoe", Password: "1234", ID: 1}).Error
		if err != nil {
			return err
		}
		err = db.Save(&User{Name: "Cap", Username: "admin", Password: "123", ID: 2}).Error
		if err != nil {
			return err
		}
		cmd.Println("Database filled successfully")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fillDatabaseCmd)
}
