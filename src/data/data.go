package data
import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Info struct{
	MaCH string
	link string
}

func GetData() map[string]string{
	db,err := sql.Open("mysql","root:0304@/info")
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	var info []Info
	rows,err := db.Query("select MaCH,link from link")
	if err!=nil{
		 panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var temp Info
		err = rows.Scan(&temp.MaCH,&temp.link)
		if err!=nil{
			panic(err)
		}
		info = append(info,temp)
	}

	data := map[string]string{
		info[0].MaCH : info[0].link,
		info[1].MaCH : info[1].link,
	}
	return data

}