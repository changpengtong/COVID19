package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:Covid19arrivingsoon!@tcp(subaru.ischool.utexas.edu:3306)/Pubmed20_C04")
	if err != nil {
		fmt.Println("failed")
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	db.SetConnMaxLifetime(time.Second * 14400)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(0)

	//Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Println("failed")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

//auto-complete api
func AllInstitutionNames() interface{} {
	return GenerateSQL("SELECT institution as name, id as id FROM COVID19_Institution")
}

func AllAuthorNames() interface{} {
	return GenerateSQL("SELECT name AS name, id AS id FROM COVID19_AuthorAutoComplete")
}

func AllBiontityNames() interface{} {
	return GenerateSQL("SELECT entity AS name, id AS id FROM COVID19_Bioentity")
}

// bioentity

func BioentityTotalData(keyword string) map[string]interface{} {
	bioentity := make(map[string]interface{})
	c0 := make(chan interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan interface{})
	c5 := make(chan interface{})
	go BioentityName(c0, keyword)
	go BarGraphPapersByYear(c1, keyword)
	go BioentityArticles(c2, keyword)
	go BioEntityAuthors(c3, keyword)
	go BioentityInstitutions(c4, keyword)
	go BioEntityClinicals(c5, keyword)
	bioentity["name"] = <-c0
	bioentity["bar"] = <-c1
	bioentity["articles"] = <-c2
	bioentity["coauthor"] = <-c3
	bioentity["institution"] = <-c4
	bioentity["clinicals"] = <-c5
	close(c0)
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	close(c5)
	return bioentity
}

func BioentityName(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT entity FROM COVID19_Bioentity WHERE id = '" + keyword + "'")
}

func BioentityArticles(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as ArticleTitle, a.abstract, a.author as Authors, SUBSTRING(a.publish_time, 1, 4) AS PubYear, a.journal as Journal_Title, a.url, a.tweet FROM COVID19_Paper a JOIN COVID19_Paper2Bioentity b ON a.cord_uid = b.paper_id WHERE b.entity_id = '" + keyword + "' ORDER BY PubYear DESC LIMIT 1000")
}

func BioentityInstitutions(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT d.id AS id, d.institution AS Affiliation, count(c.paper_id) AS NumberOfPapers FROM COVID19_Paper a JOIN COVID19_Paper2Bioentity b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Institution c ON a.cord_uid = c.paper_id JOIN COVID19_Institution d ON c.institution_id = d.id WHERE b.entity_id = '" + keyword + "' GROUP BY id, Affiliation ORDER BY NumberOfPapers DESC")
}

func BioEntityAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT SUBSTRING_INDEX(d.author_name, ' ', 2) AS ForeName, SUBSTRING_INDEX(d.author_name, ' ', -2) AS LastName, e.institution AS Affiliation, d.author_name as FullName, e.location AS Location, d.id AS aid FROM COVID19_Paper a JOIN COVID19_Paper2Bioentity b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Author c ON a.cord_uid = c.paper_id JOIN COVID19_Author d ON c.author_id = d.id JOIN COVID19_Institution e ON d.institution_id = e.id WHERE b.entity_id = '" + keyword + "' LIMIT 3000;")
}

func BarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT COUNT(a.cord_uid) AS NumberOfPapers, SUBSTRING(a.publish_time, 1, 4) as PubYear FROM COVID19_Paper a JOIN COVID19_Paper2Bioentity b ON a.cord_uid = b.paper_id WHERE b.entity_id = '" + keyword + "' GROUP BY PubYear DESC LIMIT 10")
}

func BioEntityClinicals(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as title, a.researchers, a.year, a.sponsors, a.url FROM COVID19_Clinical a JOIN COVID19_Clinical2Entity b ON a.id = b.clinical_id WHERE b.entity_id = '" + keyword + "' ORDER BY year DESC LIMIT 1000")
}

///institution

func InstitutionTotalData(keyword string) map[string]interface{} {
	institution := make(map[string]interface{})
	c0 := make(chan interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan map[string]interface{})
	c5 := make(chan interface{})
	go InstitutionName(c0, keyword)
	go InstitutionBarGraphPapersByYear(c1, keyword)
	go InstitutionPapers(c2, keyword)
	go InstitutionAuthors(c3, keyword)
	go InstitutionWordCloud(c4, keyword)
	go InstitutionClinicals(c5, keyword)
	institution["name"] = <-c0
	institution["bar"] = <-c1
	institution["articles"] = <-c2
	institution["coauthor"] = <-c3
	institution["wordcloud"] = <-c4
	institution["clinicals"] = <-c5
	close(c0)
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	close(c5)
	return institution
}

func InstitutionName(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT institution FROM COVID19_Institution WHERE id = '" + keyword + "'")
}

func InstitutionBarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT COUNT(a.cord_uid) AS NumberOfPapers, SUBSTRING(a.publish_time, 1, 4) as PubYear FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id WHERE b.institution_id = '" + keyword + "' GROUP BY PubYear DESC LIMIT 10")
}

func InstitutionPapers(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as ArticleTitle, a.abstract, a.author as Authors, SUBSTRING(a.publish_time, 1, 4) AS PubYear, a.journal as Journal_Title, a.url, a.tweet FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id WHERE b.institution_id = '" + keyword + "' ORDER BY PubYear DESC LIMIT 1000")
}

func InstitutionAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT SUBSTRING_INDEX(d.author_name, ' ', 2) AS ForeName, SUBSTRING_INDEX(d.author_name, ' ', -2) AS LastName, e.institution AS Affiliation, d.author_name as FullName, e.location AS Location, d.id AS aid FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Author c ON a.cord_uid = c.paper_id JOIN COVID19_Author d ON c.author_id = d.id JOIN COVID19_Institution e ON d.institution_id = e.id WHERE b.institution_id = '" + keyword + "' LIMIT 3000")
}

func InstitutionDisease(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.institution_id = '" + keyword + "' AND c.type = 'Disease' GROUP BY id, Mention ORDER BY occurences DESC")
}

func InstitutionDrug(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.institution_id = '" + keyword + "' AND c.type = 'Chemical' GROUP BY id, Mention ORDER BY occurences DESC")
}

func InstitutionGene(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Institution b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.institution_id = '" + keyword + "' AND c.type = 'Gene' GROUP BY id, Mention ORDER BY occurences DESC")
}

func InstitutionClinicals(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as title, a.researchers, a.year, a.sponsors, a.url FROM COVID19_Clinical a JOIN COVID19_Clinical2Institution b ON a.id = b.clinical_id WHERE b.institution_id = '" + keyword + "' ORDER BY year DESC LIMIT 1000")
}

func InstitutionWordCloud(c chan map[string]interface{}, keyword string) {
	res := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	go InstitutionDisease(c1, keyword)
	go InstitutionDrug(c2, keyword)
	go InstitutionGene(c3, keyword)
	res["disease"] = <-c1
	res["drug"] = <-c2
	res["gene"] = <-c3
	close(c1)
	close(c2)
	close(c3)
	c <- res
}

//authorlist
func AuthorList(keyword string) interface{} {
	return GenerateSQL("SELECT a.id as aid, a.email as Email, a.author_name AS Name, b.institution AS Affiliation, b.id AS institution_id, b.location AS Location FROM COVID19_Author a LEFT JOIN COVID19_Institution b ON a.institution_id = b.id WHERE author_name LIKE '%" + keyword + "%'")
}

//author
func AuthorTotalData(aid string) map[string]interface{} {
	author := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan map[string]interface{})
	c5 := make(chan interface{})
	c6 := make(chan interface{})
	go AuthorArticles(c1, aid)
	go AuthorcoAuthors(c2, aid)
	go AuthorBarGraphPapersByYear(c3, aid)
	go AuthorWordCloudType(c4, aid)
	go AuthorCard(c5, aid)
	go AuthorClinicals(c6, aid)
	author["articles"] = <-c1
	author["coauthor"] = <-c2
	author["bar"] = <-c3
	author["wordcloud"] = <-c4
	author["card"] = <-c5
	author["clinicals"] = <-c6
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	close(c5)
	close(c6)
	return author
}

func AuthorArticles(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as ArticleTitle, a.abstract, a.author as Authors, SUBSTRING(a.publish_time, 1, 4) AS PubYear, a.journal as Journal_Title, a.url, a.tweet FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id WHERE b.author_id = '" + aid + "' ORDER BY PubYear DESC LIMIT 1000")
}

func AuthorcoAuthors(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT SUBSTRING_INDEX(e.author_name, ' ', 2) AS ForeName, SUBSTRING_INDEX(e.author_name, ' ', -2) AS LastName, f.institution AS Affiliation, e.author_name as FullName, f.location AS Location, e.id AS aid FROM (SELECT a.cord_uid AS cord_uid FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id WHERE b.author_id = '" + aid + "') c JOIN COVID19_Paper2Author d ON c.cord_uid = d.paper_id JOIN COVID19_Author e ON d.author_id = e.id JOIN COVID19_Institution f ON e.institution_id = f.id WHERE d.author_id != '" + aid + "' LIMIT 3000")
}

func AuthorBarGraphPapersByYear(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT COUNT(a.cord_uid) AS NumberOfPapers, SUBSTRING(a.publish_time, 1, 4) as PubYear FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id WHERE b.author_id = '" + aid + "' GROUP BY PubYear DESC LIMIT 10")
}

func AuthorCard(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT SUBSTRING_INDEX(a.author_name, ' ', 2) AS ForeName, SUBSTRING_INDEX(a.author_name, ' ', -2) AS LastName, b.institution AS Affiliation, b.id AS institution_id, a.author_name AS FullName, b.location AS Location, a.id AS aid, a.email AS Email FROM COVID19_Author a LEFT JOIN COVID19_Institution b ON a.institution_id = b.id WHERE a.id = '" + aid + "'")
}

func AuthorDisease(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.author_id = '" + aid + "' AND c.type = 'Disease' GROUP BY id, Mention ORDER BY occurences DESC")
}

func AuthorDrug(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.author_id = '" + aid + "' AND c.type = 'Chemical' GROUP BY id, Mention ORDER BY occurences DESC")
}

func AuthorGene(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT d.id AS id, d.entity AS Mention, COUNT(a.cord_uid) AS occurences FROM COVID19_Paper a JOIN COVID19_Paper2Author b ON a.cord_uid = b.paper_id JOIN COVID19_Paper2Bioentity c ON a.cord_uid = c.paper_id JOIN COVID19_Bioentity d ON c.entity_id = d.id WHERE b.author_id = '" + aid + "' AND c.type = 'Gene' GROUP BY id, Mention ORDER BY occurences DESC")
}

func AuthorClinicals(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT TRIM(BOTH '['  FROM a.title) as title, a.researchers, a.year, a.sponsors, a.url FROM COVID19_Clinical a JOIN COVID19_Clinical2Author b ON a.id = b.clinical_id WHERE b.author_id = '" + aid + "' ORDER BY year DESC LIMIT 1000")
}

func AuthorWordCloudType(c chan map[string]interface{}, aid string) {
	res := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	go AuthorDisease(c1, aid)
	go AuthorDrug(c2, aid)
	go AuthorGene(c3, aid)
	res["disease"] = <-c1
	res["drug"] = <-c2
	res["gene"] = <-c3
	close(c1)
	close(c2)
	close(c3)
	c <- res
}

//sql
func GenerateSQL(sql string) interface{} {
	db := ConnectMySQL()
	rows, _ := db.Query(sql)
	columns, _ := rows.Columns()
	fmt.Println(columns)
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	var result []map[string]string
	for rows.Next() {
		_ = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		result = append(result, record)
	}
	rows.Close()
	db.Close()
	return result
}

func main() {
	//ConnectMySQL()
}
