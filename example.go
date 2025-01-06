package main

import (
	lib "das/lib"
	"fmt"
)

func main() {
	con := lib.Sql{}
	con.New(nil)

	defer con.Close()

	con.CommitMultiple("insert into category(id, name) values(?, ?)", [][]any{
		{1, "Jane"},
		{2, "Shang"},
	})

	ses2 := storage.GetSession(key2)
	con.Go()

	rows, err := con.Get("select id, name from category where id = ?", 1)

	if err == nil {
		for rows.Next() {
			var id int
			var name string

			rows.Scan(&id, &name)
			fmt.Println(id, name)
		}
		con.Commit("delete from category where id = ?", 1)
		con.Go()

	} else {
		fmt.Println(err)
	}
	for rows.Next() {

		//rows.Scan(&id, &name)

		fmt.Print(1)
	}
	fmt.Print(ses2.Get("name"))
}
