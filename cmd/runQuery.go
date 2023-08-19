package cmd

import (
	"errors"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// runQueryCmd represents the runSQL command
var runQueryCmd = &cobra.Command{
	Use:   "runQuery",
	Short: "Run query",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("query is required")
		}
		query := args[0]
		rows, err := stdDB.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for {
			cols, err := rows.Columns()
			if err != nil {
				return err
			}
			var header []interface{}
			for _, col := range cols {
				header = append(header, col)
			}

			t := table.NewWriter()
			t.SetOutputMirror(cmd.OutOrStdout())
			t.AppendHeader(header)
			for rows.Next() {
				columns := make([]*string, len(cols))
				columnPointers := make([]interface{}, len(cols))
				for i, _ := range columns {
					columnPointers[i] = &columns[i]
				}
				err := rows.Scan(columnPointers...)
				if err != nil {
					return err
				}
				var tableRow []interface{}
				for _, column := range columns {
					if column == nil {
						tableRow = append(tableRow, "<null>")
					} else {
						tableRow = append(tableRow, *column)
					}
				}
				t.AppendRow(tableRow)
			}
			t.Render()
			if !rows.NextResultSet() {
				break
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runQueryCmd)
}
