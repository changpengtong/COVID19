package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(subaru.ischool.utexas.edu:3306)/Pubmed20_C04")
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
	return GenerateSQL("SELECT DISTINCT authorAffiliation as name, authorAffiliationLocation\nFROM KaggleAllAuthors\nWHERE authorAffiliation is not null AND authorAffiliation != ''")
}
func AllTitleNames() interface{} {
	return GenerateSQL("SELECT DISTINCT title\n as name FROM KaggleAllPaperDetails\nWHERE title is not null  AND title!=''")
}
func AllAuthorNames() interface{} {
	return GenerateSQL("SELECT DISTINCT authorName\n as name FROM KaggleAllAuthors\nWHERE authorName is not null AND authorName!=''\nAND authorName REGEXP '^[A-Za-z0-9]'")
}
func AllBiontityNames() interface{} {
	return GenerateSQL("SELECT DISTINCT entity\nas name FROM KaggleAllBioentitiesCleaned\nWHERE entity is not null AND entity!=''")
}

// bioentity

func BioentityTotalData(keyword string) map[string]interface{} {
	bioentity := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan interface{})
	go BarGraphPapersByYear(c1, keyword)
	go BioentityArticles(c2, keyword)
	go BioEntityAuthors(c3, keyword)
	go BioentityInstitutions(c4, keyword)
	bioentity["bar"] = <-c1
	bioentity["articles"] = <-c2
	bioentity["coauthor"] = <-c3
	bioentity["institution"] = <-c4
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	return bioentity
}
func BioentityType(keyword string) interface{} {
	return GenerateSQL("SELECT DISTINCT Type as type FROM KaggleAllBioentitiesCombined WHERE entity_list = '%" + keyword + "%'")
}

func BioentityArticles(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT pubmed_id, TRIM(BOTH '['  FROM title) as ArticleTitle, SUBSTRING(publish_time,1,4) as PubYear, authors as Authors, journal as Journal_Title, abstract, Tweets as tweet,(case when (url LIKE '%; %') THEN concat('https://www.ncbi.nlm.nih.gov/pubmed/',pubmed_id) ELSE url END)as url FROM KaggleAllPaperDetails FORCE INDEX (`title_pubmedid_index`) WHERE EXISTS (SELECT DISTINCT pmid FROM KaggleAllBioentitiesCleaned FORCE INDEX (`entity_pmid_index`) WHERE pmid=KaggleAllPaperDetails.pubmed_id AND  entity LIKE '%" + keyword + "%') ORDER BY PubYear DESC")

}

func BioentityInstitutions(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT authorAffiliation as Affiliation, authorAffiliationLocation as Location, count(DISTINCT paperTitle) as NumberOfPapers FROM KaggleAllAuthors\n    WHERE concat(authorAffiliation, '', authorAffiliationLocation) is not null AND paperTitle IN\n     (SELECT title FROM KaggleAllPaperDetails WHERE pubmed_id IN\n        (SELECT DISTINCT PMID FROM KaggleAllBioentitiesCleaned WHERE entity LIKE '%" + keyword + "%'))\n    GROUP BY authorAffiliation, authorAffiliationLocation ORDER BY NumberOfPapers DESC")
}

func BioEntityAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT KaggleAllAuthors.paperTitle, KaggleAllAuthors.paper_id,SUBSTRING_INDEX(KaggleAllAuthors.authorName, ' ', 2) AS ForeName, SUBSTRING_INDEX(KaggleAllAuthors.authorName, ' ', -2) AS LastName, KaggleAllAuthors.authorAffiliation as Affiliation, KaggleAllAuthors.authorName as FullName, KaggleAllAuthors.authorAffiliationLocation as Location FROM KaggleAllAuthors LEFT JOIN KaggleAllPaperDetails ON paperTitle=title INNER JOIN KaggleAllBioentitiesCleaned ON KaggleAllPaperDetails.pubmed_id=KaggleAllBioentitiesCleaned.pmid WHERE entity LIKE '%" + keyword + "%' ORDER BY LastName")
}

func BarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT count(DISTINCT title) as NumberOfPapers, SUBSTRING(publish_time,1,4) as PubYear FROM KaggleAllPaperDetails WHERE pubmed_id IN (SELECT DISTINCT pmid FROM KaggleAllBioentitiesCleaned WHERE entity LIKE '%" + keyword + "%' ) GROUP BY PubYear DESC LIMIT 10")
}

func BioentityWordCloud() interface{} {
	return GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.Ace2MeshWords ORDER BY occurences DESC;")
}

///institution

func InstitutionTotalData(keyword string) map[string]interface{} {
	institution := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan map[string]interface{})
	go InstitutionBarGraphPapersByYear(c1, keyword)
	go InstitutionPapers(c2, keyword)
	go InstitutionAuthors(c3, keyword)
	go InstitutionWordCloud(c4, keyword)
	institution["bar"] = <-c1
	institution["articles"] = <-c2
	institution["coauthor"] = <-c3
	institution["wordcloud"] = <-c4
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	return institution
}

func InstitutionBarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT count(DISTINCT pubmed_id) as NumberOfPapers, SUBSTRING(publish_time,1,4) as PubYear\nFROM KaggleAllPaperDetails\nWHERE EXISTS\n      (SELECT DISTINCT paperTitle FROM KaggleAllAuthors\n      WHERE title=KaggleAllAuthors.paperTitle AND\n        authorAffiliation LIKE '%" + keyword + "%'\n  AND (paperTitle != '' OR paperTitle is not null))\n  AND pubmed_id is not null AND pubmed_id!=' '\n    GROUP BY PubYear ORDER BY PubYear DESC LIMIT 10")
}

func InstitutionPapers(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT title as ArticleTitle, abstract, authors as Authors, pmcid, pubmed_id,\n    SUBSTRING(publish_time,1,4) as PubYear, journal as Journal_Title, Tweets as tweet,\n    (case when (url LIKE '%; %') THEN concat('https://www.ncbi.nlm.nih.gov/pubmed/',pubmed_id) ELSE url END) as url\n    FROM KaggleAllPaperDetails\n    WHERE EXISTS (SELECT DISTINCT paperTitle FROM KaggleAllAuthors\n        WHERE paperTitle=KaggleAllPaperDetails.title AND\n              authorAffiliation LIKE '%" + keyword + "%')\n  AND pubmed_id is not null AND pubmed_id !=' '\n   ORDER BY PubYear DESC;")
	//c<-GenerateSQL("SELECT SLEEP(1);")

}
func InstitutionAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT\n    DISTINCT SUBSTRING_INDEX(authorName, ' ', 2) AS ForeName,\n     SUBSTRING_INDEX(authorName, ' ', -2) AS LastName,\n       authorAffiliation as Affiliation, authorName  as FullName,\n       authorAffiliationLocation as Location\n   FROM KaggleAllAuthors\n    WHERE authorName is not null AND authorAffiliation LIKE '%" + keyword + "%' \n    AND authorName REGEXP '^[A-Za-z0-9]'\n    ORDER BY LastName")
	//c<-GenerateSQL("SELECT SLEEP(2);")

}
func InstitutionDisease(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned WHERE type='disease' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE EXISTS (SELECT paperTitle FROM KaggleAllAuthors WHERE paperTitle=KaggleAllPaperDetails.title AND authorAffiliation LIKE '%" + keyword + "%' ) ) GROUP BY Mention ORDER BY occurences DESC")
}
func InstitutionDrug(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned WHERE type='chemical' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE EXISTS (SELECT paperTitle FROM KaggleAllAuthors WHERE paperTitle=KaggleAllPaperDetails.title AND authorAffiliation LIKE '%" + keyword + "%')) GROUP BY Mention ORDER BY occurences DESC")
}
func InstitutionGene(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned WHERE type='gene' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE EXISTS (SELECT paperTitle FROM KaggleAllAuthors WHERE paperTitle=KaggleAllPaperDetails.title AND authorAffiliation LIKE '%" + keyword + "%')) GROUP BY Mention ORDER BY occurences DESC")
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
	return GenerateSQL("SELECT DISTINCT authorEmail as Email, authorName as Name, authorAffiliationLocation as Location, authorAffiliation as Affiliation FROM KaggleAllAuthors WHERE authorName LIKE '%" + keyword + "%' AND authorAffiliation is not null AND authorAffiliationLocation is not null AND authorEmail is not null AND authorEmail!=' '")
}

//author
func AuthorTotalData(key1 string, key2 string) map[string]interface{} {
	author := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	c4 := make(chan map[string]interface{})
	go AuthorArticles(c1, key1, key2)
	go AuthorcoAuthors(c2, key1, key2)
	go AuthorBarGraphPapersByYear(c3, key1, key2)
	go AuthorWordCloudType(c4, key1, key2)
	author["articles"] = <-c1
	author["coauthor"] = <-c2
	author["bar"] = <-c3
	author["wordcloud"] = <-c4
	close(c1)
	close(c2)
	close(c3)
	close(c4)
	return author
}

func AuthorArticles(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT DISTINCT title as ArticleTitle, abstract, authors as Authors, pmcid, pubmed_id,\n SUBSTRING(publish_time,1,4) as PubYear, journal as Journal_Title, Tweets as tweet,\n                (case when (url LIKE '%; %')\n THEN concat('https://www.ncbi.nlm.nih.gov/pubmed/',pubmed_id)\n                    ELSE url END)as url\nFROM KaggleAllPaperDetails WHERE title IN(\nSELECT DISTINCT paperTitle FROM KaggleAllAuthors WHERE\n    authorName LIKE'%" + name + "%'\n    AND authorAffiliation LIKE '%" + affi + "%'\n    )\nAND pubmed_id is not null AND pubmed_id !=' '\nORDER BY PubYear DESC;")
}

func AuthorcoAuthors(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT DISTINCT SUBSTRING_INDEX(authorName, ' ', 2) AS ForeName, SUBSTRING_INDEX(authorName, ' ', -2) AS LastName,\n    authorAffiliation as Affiliation, authorName as FullName, authorAffiliationLocation as Location\nFROM KaggleAllAuthors\nWHERE paperTitle IN (\nSELECT DISTINCT paperTitle FROM KaggleAllAuthors WHERE authorName LIKE'%" + name + "%'\n    AND authorAffiliation LIKE '%" + affi + "%')\nAND authorName is not null  and authorName!='%" + name + "%' AND authorName REGEXP '^[A-Za-z0-9]'\nORDER BY LastName")

}
func AuthorBarGraphPapersByYear(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT count(title) as NumberOfPapers,\n                SUBSTRING(publish_time,1,4) as PubYear\nFROM KaggleAllPaperDetails WHERE title IN(\nSELECT DISTINCT paperTitle FROM KaggleAllAuthors WHERE authorName LIKE'%" + name + "%'\n    AND authorAffiliation LIKE '%" + affi + "%')\nAND pubmed_id is not null AND pubmed_id !=' '\nGROUP BY PubYear ORDER BY PubYear DESC LIMIT 10;")
}
func AuthorCard(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT a2.LastName, a2.ForeName, a13.Affiliation " +
		"FROM Pubmed20_C04.A02_AuthorList a2 " +
		"JOIN Pubmed20_C04.A13_AffiliationList a13 ON a13.PMID = a2.PMID " +
		"WHERE a2.AID = " + aid + " LIMIT 1")
}

func AuthorDisease(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned\n    WHERE type='disease' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN\n    (SELECT paperTitle FROM KaggleAllAuthors WHERE authorName LIKE'%" + name + "%'\n    AND authorAffiliation LIKE '%" + affi + "%'))\n     GROUP BY Mention ORDER BY occurences DESC")
}
func AuthorDrug(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned WHERE type='chemical' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN (SELECT paperTitle FROM KaggleAllAuthors WHERE authorName LIKE'%" + name + "%' AND authorAffiliation LIKE '%" + affi + "%')) GROUP BY Mention")
}
func AuthorGene(c chan interface{}, name string, affi string) {
	c <- GenerateSQL("SELECT entity as Mention, count(DISTINCT pmid) as occurences FROM KaggleAllBioentitiesCleaned WHERE type='gene' AND pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN (SELECT paperTitle FROM KaggleAllAuthors WHERE authorName LIKE'%" + name + "%' AND authorAffiliation LIKE '%" + affi + "%')) GROUP BY Mention")
}
func AuthorWordCloudType(c chan map[string]interface{}, name string, affi string) {
	res := make(map[string]interface{})
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	go AuthorDisease(c1, name, affi)
	go AuthorDrug(c2, name, affi)
	go AuthorGene(c3, name, affi)
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
		//	fmt.Println(record)
	}
	rows.Close()
	db.Close()
	return result
}

func main() {
	//ConnectMySQL()
}
