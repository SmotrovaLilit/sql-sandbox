package cmd

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dbDriver string
var db *gorm.DB
var stdDB *sql.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db-sandbox",
	Short: "The application demonstrate difference between mysql, postgres, mssql in terms of sql injections",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)
	rootCmd.PersistentFlags().StringVar(&dbDriver, "dbDriver", "", "Supported values: mssql, mysql, postgres")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./db-sandbox.yaml)")
	err := rootCmd.MarkPersistentFlagRequired("dbDriver")
	cobra.CheckErr(err)
}

func initialize() {
	initConfig()
	cobra.CheckErr(initDB())
}

func initDB() error {
	dsn := viper.GetString("dsn." + dbDriver)
	if dsn == "" {
		return fmt.Errorf("dsn for %s is not set", dbDriver)
	}
	var err error

	switch dbDriver {
	case "mysql":
		stdDB, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect database mysql: %w", err)
		}
	case "postgres":
		stdDB, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect database postgres: %w", err)
		}
	case "mssql":
		stdDB, err = sql.Open("sqlserver", dsn)
		if err != nil {
			return err
		}
		db, err = gorm.Open(sqlserver.New(sqlserver.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect database mssql: %w", err)
		}
	default:
		return fmt.Errorf("unsupported db driver: %s", dbDriver)
	}
	//db = db.Debug()
	err = db.AutoMigrate(Book{}, User{})
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName("db-sandbox")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
