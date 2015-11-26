/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     desktop.go
 * +===============================================
 */
package gosimac

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
	"os/user"
)

func ChangeDesktopBackground(path string) error {
	usr, err := user.Current()
	if err != nil {
		glog.Errorf("OS.User: %v\n", err)
		return err
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/Library/Application Support/Dock/desktoppicture.db", usr.HomeDir))
	if err != nil {
		glog.Errorf("Database.SQL: %v\n", err)
		return err
	}
	glog.V(2).Infof("Database was opened successfully")
	defer db.Close()

	sqlSmt := fmt.Sprintf("update data set value = '%s';", path)
	glog.V(2).Infof("SQL Query: %s", sqlSmt)

	_, err = db.Exec(sqlSmt)
	if err != nil {
		glog.Errorf("Database.SQL: %v\n", err)
		return err
	}
	glog.V(2).Infof("%s was executed successfully", sqlSmt)

	return nil
}
