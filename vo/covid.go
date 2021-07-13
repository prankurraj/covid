package vo

type Response struct {
	CasesTimeSeries []struct {
		Dailyconfirmed string `json:"dailyconfirmed"`
		Dailydeceased  string `json:"dailydeceased"`
		Dailyrecovered string `json:"dailyrecovered"`
		Date           string `json:"date"`
		Dateymd        string `json:"dateymd"`
		Totalconfirmed string `json:"totalconfirmed"`
		Totaldeceased  string `json:"totaldeceased"`
		Totalrecovered string `json:"totalrecovered"`
	} `json:"cases_time_series"`
	Statewise []struct {
		Active          string `json:"active"`
		Confirmed       string `json:"confirmed"`
		Deaths          string `json:"deaths"`
		Deltaconfirmed  string `json:"deltaconfirmed"`
		Deltadeaths     string `json:"deltadeaths"`
		Deltarecovered  string `json:"deltarecovered"`
		Lastupdatedtime string `json:"lastupdatedtime"`
		Migratedother   string `json:"migratedother"`
		Recovered       string `json:"recovered"`
		State           string `json:"state"`
		Statecode       string `json:"statecode"`
		Statenotes      string `json:"statenotes"`
	} `json:"statewise"`
	Tested []struct {
		Dailyrtpcrsamplescollectedicmrapplication     string `json:"dailyrtpcrsamplescollectedicmrapplication"`
		Firstdoseadministered                         string `json:"firstdoseadministered"`
		Frontlineworkersvaccinated1Stdose             string `json:"frontlineworkersvaccinated1stdose"`
		Frontlineworkersvaccinated2Nddose             string `json:"frontlineworkersvaccinated2nddose"`
		Healthcareworkersvaccinated1Stdose            string `json:"healthcareworkersvaccinated1stdose"`
		Healthcareworkersvaccinated2Nddose            string `json:"healthcareworkersvaccinated2nddose"`
		Over45Years1Stdose                            string `json:"over45years1stdose"`
		Over45Years2Nddose                            string `json:"over45years2nddose"`
		Over60Years1Stdose                            string `json:"over60years1stdose"`
		Over60Years2Nddose                            string `json:"over60years2nddose"`
		Positivecasesfromsamplesreported              string `json:"positivecasesfromsamplesreported"`
		Registration1845Years                         string `json:"registration18-45years"`
		Registrationabove45Years                      string `json:"registrationabove45years"`
		Registrationflwhcw                            string `json:"registrationflwhcw"`
		Registrationonline                            string `json:"registrationonline"`
		Registrationonspot                            string `json:"registrationonspot"`
		Samplereportedtoday                           string `json:"samplereportedtoday"`
		Seconddoseadministered                        string `json:"seconddoseadministered"`
		Source                                        string `json:"source"`
		Source2                                       string `json:"source2"`
		Source3                                       string `json:"source3"`
		Source4                                       string `json:"source4"`
		Testedasof                                    string `json:"testedasof"`
		Testsconductedbyprivatelabs                   string `json:"testsconductedbyprivatelabs"`
		To60YearswithcoMorbidities1Stdose             string `json:"to60yearswithco-morbidities1stdose"`
		To60YearswithcoMorbidities2Nddose             string `json:"to60yearswithco-morbidities2nddose"`
		Totaldosesadministered                        string `json:"totaldosesadministered"`
		Totaldosesavailablewithstates                 string `json:"totaldosesavailablewithstates"`
		Totaldosesavailablewithstatesprivatehospitals string `json:"totaldosesavailablewithstatesprivatehospitals"`
		Totaldosesinpipeline                          string `json:"totaldosesinpipeline"`
		Totaldosesprovidedtostatesuts                 string `json:"totaldosesprovidedtostatesuts"`
		Totalindividualsregistered                    string `json:"totalindividualsregistered"`
		Totalindividualstested                        string `json:"totalindividualstested"`
		Totalindividualsvaccinated                    string `json:"totalindividualsvaccinated"`
		Totalpositivecases                            string `json:"totalpositivecases"`
		Totalrtpcrsamplescollectedicmrapplication     string `json:"totalrtpcrsamplescollectedicmrapplication"`
		Totalsamplestested                            string `json:"totalsamplestested"`
		Totalsessionsconducted                        string `json:"totalsessionsconducted"`
		Totalvaccineconsumptionincludingwastage       string `json:"totalvaccineconsumptionincludingwastage"`
		Updatetimestamp                               string `json:"updatetimestamp"`
		Years1Stdose                                  string `json:"years1stdose"`
		Years2Nddose                                  string `json:"years2nddose"`
	} `json:"tested"`
}