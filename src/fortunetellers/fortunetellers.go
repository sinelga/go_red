package fortunetellers

import (
	"domains"
	"encoding/csv"
	"encoding/json"
	"log/syslog"
	"os"
	"strconv"
)

func GetAll(golog syslog.Writer) []byte {

	csvFile, err := os.Open("/home/juno/git/go_red/fortunetellers.csv")

	if err != nil {
		golog.Crit(err.Error())
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		golog.Crit(err.Error())
	}

	var allfortunetellers []domains.Fortuneteller
	var fortuneteller domains.Fortuneteller

	for _, each := range csvData {

		fortuneteller.Id, _ = strconv.Atoi(each[0])
		fortuneteller.Name = each[1]
		fortuneteller.Phone = each[2]
		fortuneteller.Location = each[3]
		fortuneteller.Moto = each[4]
		fortuneteller.Desc = each[5]

		allfortunetellers = append(allfortunetellers, fortuneteller)
	}

	jsondata, err := json.Marshal(allfortunetellers)
	
	if err != nil {
		golog.Crit(err.Error())

	}

	return jsondata
}
