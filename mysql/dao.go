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

// bioentity

func BioentityTotalData(keyword string) map[string]interface{} {
	bioentity := make(map[string]interface{})
	c := make(chan interface{})
	go BarGraphPapersByYear(c, keyword)
	bioentity["bar"] = <-c
	go BioentityArticles(c, keyword)
	bioentity["articles"] = <-c
	go BioEntityAuthors(c, keyword)
	bioentity["coauthor"] = <-c
	go BioentityInstitutions(c, keyword)
	bioentity["institution"] = <-c
	close(c)
	return bioentity
}
func BioentityType(keyword string) interface{} {
	return GenerateSQL("SELECT DISTINCT Type as type FROM KaggleAllBioentitiesCombined WHERE entity_list = '%" + keyword + "%'")
}

func BioentityArticles(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT title  as ArticleTitle, SUBSTRING(publish_time,1,4) as PubYear, authors as Authors, journal  as Journal_Title, abstract, url FROM KaggleAllPaperDetails WHERE pubmed_id IN (SELECT DISTINCT pmid FROM KaggleAllBioentitiesCombined WHERE entity_list LIKE '%" + keyword + "%') AND (title != '' OR title is not null) ORDER BY PubYear DESC")

}

func BioentityInstitutions(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT concat(authorAffiliation, '', authorAffiliationLocation) as Affiliation, count(DISTINCT paperTitle) as NumberOfPapers FROM KaggleAllAuthors WHERE paperTitle IN (SELECT title FROM KaggleAllPaperDetails WHERE pubmed_id IN (SELECT DISTINCT PMID FROM KaggleAllBioentitiesCombined WHERE entity_list LIKE '%" + keyword + "%')) GROUP BY Affiliation ORDER BY NumberOfPapers DESC")
}
func BioEntityAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT  DISTINCT SUBSTRING_INDEX(SUBSTRING_INDEX(authorName, ' ', 1), ' ', -1) AS ForeName,   TRIM( SUBSTR(authorName, LOCATE(' ', authorName)) ) AS LastName,     authorAffiliation as Affiliation,   authorAffiliationLocation as Location,       authorName as FullName FROM KaggleAllAuthors WHERE paperTitle IN (SELECT title FROM KaggleAllPaperDetails WHERE pubmed_id IN (SELECT DISTINCT PMID FROM KaggleAllBioentitiesCombined WHERE entity_list LIKE '%" + keyword + "%')) ORDER BY LastName")
}

func BarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT count(id) as NumberOfPapers, SUBSTRING(publish_time,1,4) as PubYear FROM KaggleAllPaperDetails WHERE pubmed_id IN (SELECT DISTINCT pmid FROM KaggleAllBioentitiesCombined WHERE entity_list LIKE '%" + keyword + "%' ) GROUP BY PubYear DESC LIMIT 10")
	//return GenerateSQL("SELECT count(pubmed_id) as NumberOfPapers, SUBSTRING(publish_time,1,4) as PubYear FROM KaggleAllPaperDetails WHERE pubmed_id IN\n    (SELECT DISTINCT PMID FROM entities_revised_finalTSV WHERE entity LIKE '%" + keyword + "%' )\n    AND (title != '' OR title is not null) GROUP BY PubYear DESC LIMIT 10")
}

func BioentityWordCloud() interface{} {
	return GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.Ace2MeshWords ORDER BY occurences DESC;")
}

//instition
func InstitutionTotalData(keyword string) map[string]interface{} {
	institution := make(map[string]interface{})
	c := make(chan interface{})
	go InstitutionBarGraphPapersByYear(c, keyword)
	institution["bar"] = <-c
	go InstitutionPapers(c, keyword)
	institution["articles"] = <-c
	go InstitutionAuthors(c, keyword)
	institution["coauthor"] = <-c
	go InstitutionWordCloud(c, keyword)
	institution["wordcloud"] = <-c
	close(c)
	return institution
}
func AllInstitutionNames() interface{} {
	return GenerateSQL("SELECT DISTINCT authorAffiliation, authorAffiliationLocation\nFROM KaggleAllAuthors\nWHERE authorAffiliation is not null AND authorAffiliation != ''")
}
func InstitutionBarGraphPapersByYear(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT count(DISTINCT pubmed_id) as NumberOfPapers, SUBSTRING(publish_time,1,4) as PubYear\nFROM KaggleAllPaperDetails\nWHERE title IN (SELECT DISTINCT paperTitle FROM KaggleAllAuthors WHERE authorAffiliation LIKE '%" + keyword + "%' AND (paperTitle != '' OR paperTitle is not null))\nAND title is not null AND title != ' ' AND pubmed_id is not null AND pubmed_id!=' '\nGROUP BY PubYear ORDER BY PubYear DESC LIMIT 10")
}

func InstitutionPapers(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT DISTINCT K01.title as ArticleTitle, K01.abstract, K01.authors as Authors, K01.pmcid, K01.pubmed_id, SUBSTRING(K01.publish_time,1,4) as PubYear, K01.url, K01.journal as Journal_Title, T01.tweet_count as tweet \nFROM KaggleAllPaperDetails K01\nLEFT JOIN tweets_final T01\nON K01.pubmed_id = T01.pmid\nWHERE title is not null AND title != ' ' AND pubmed_id is not null AND pubmed_id !=' ' AND\n      title IN(\nSELECT DISTINCT paperTitle FROM KaggleAllAuthors WHERE authorAffiliation LIKE '%" + keyword + "%')\nORDER BY PubYear DESC")
}
func InstitutionAuthors(c chan interface{}, keyword string) {
	c <- GenerateSQL("SELECT\n    DISTINCT SUBSTRING_INDEX(authorName, ' ', 2) AS ForeName,\n     SUBSTRING_INDEX(authorName, ' ', -2) AS LastName,\n       authorAffiliation as Affiliation, authorName  as FullName,\n       authorAffiliationLocation as Location\n   FROM KaggleAllAuthors\n    WHERE authorName is not null AND authorAffiliation LIKE '%" + keyword + "%' \n    AND authorName REGEXP '^[A-Za-z0-9]'\n    ORDER BY LastName")
}
func InstitutionWordCloud(c chan interface{}, keyword string) {
	res := make(map[string]interface{})
	disease := GenerateSQL("SELECT entity_list as Mention, count(DISTINCT pmid) as occurences FROM disease1TSV WHERE pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN (SELECT paperTitle FROM KaggleAllAuthors WHERE authorAffiliation LIKE '%" + keyword + "%') AND title != '' AND title is not null) GROUP BY entity_list ORDER BY occurences DESC;")
	drug := GenerateSQL("SELECT entity_list as Mention, count(DISTINCT pmid) as occurences FROM chemical_v2TSV WHERE pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN (SELECT paperTitle FROM KaggleAllAuthors WHERE authorAffiliation LIKE '%" + keyword + "%') AND title != '' AND title is not null) GROUP BY entity_list ORDER BY occurences DESC;")
	gene := GenerateSQL("SELECT entity_list as Mention, count(DISTINCT pmid) as occurences FROM gene1TSV    WHERE pmid IN (SELECT pubmed_id FROM KaggleAllPaperDetails WHERE title IN (SELECT paperTitle FROM KaggleAllAuthors WHERE authorAffiliation LIKE '%" + keyword + "%' ) AND title != '' AND title is not null) GROUP BY entity_list ORDER BY occurences DESC;")
	res["disease"] = disease
	res["drug"] = drug
	res["gene"] = gene
	c <- res
}

//author
func AuthorTotalData(aid string) map[string]interface{} {
	author := make(map[string]interface{})
	c := make(chan interface{})
	go AuthorArticles(c, aid)
	author["articles"] = <-c
	go AuthorcoAuthors(c, aid)
	author["coauthor"] = <-c
	go AuthorWordCloudType(c)
	author["wordcloud"] = <-c
	go AuthorBarGraphPapersByYear(c, aid)
	author["bar"] = <-c
	go AuthorCard(c, aid)
	author["card"] = <-c
	close(c)
	//	author["articles"] = AuthorArticles()
	//	author["coauthor"] = AuthorcoAuthors()
	//	author["wordcloud"] = AuthorWordCloudType()
	////	author["institution"] = AuthorInstitutions()
	//	author["bar"] = AuthorBarGraphPapersByYear()
	//	fmt.Println(author)
	return author
}

func AuthorArticles(c chan interface{}, aid string) {
	//return GenerateSQL("SELECT DISTINCT b.LastName,a.PMID,a.ArticleTitle,a.PubYear,b.PMID " +
	//	"FROM Pubmed20_C04.GeneEntityArticleDetails a,Pubmed20_C04.GeneEntityAuthors b WHERE a.PMID=b.PMID LIMIT 20")
	//c <- GenerateSQL("SELECT ArticleTitle, PubYear,Authors,Journal_Title FROM Pubmed20_C04.BinChenListOfPapers ORDER BY PubYear DESC")
	c <- GenerateSQL("SELECT a1.ArticleTitle,a1.Journal_Title,d3.Authors,d2.PubYear FROM Pubmed20_C04.A02_AuthorList a2 " +
		"JOIN Pubmed20_C04.A01_Articles a1 ON a2.PMID = a1.PMID " +
		"JOIN Pubmed20_C04.D02_PublishingYear d2 ON d2.PMID = a2.PMID " +
		"JOIN Pubmed20_C04.D03_PaperAuthors d3 ON d3.PMID = a2.PMID " +
		"WHERE a2.AID = " + aid)
}

func AuthorcoAuthors(c chan interface{}, aid string) {
	//c <- GenerateSQL("SELECT FullName,ForeName, LastName, Affiliation, NumberofPapers FROM Pubmed20_C04.BinChenCoauthorsFinal ORDER BY NumberofPapers DESC")
	c <- GenerateSQL("SELECT DISTINCT c4.AID, c4.LastName, c4.ForeName, c4.Affiliation, count(a22.PMID) as rank " +
		"FROM (SELECT a2.PMID FROM Pubmed20_C04.A02_AuthorList a2 " +
		"JOIN Pubmed20_C04.D03_PaperAuthors d3 ON d3.PMID = a2.PMID WHERE a2.AID = " + aid + ") a22 " +
		"JOIN Pubmed20_C04.A02_AuthorList c4 ON c4.PMID = a22.PMID " +
		"WHERE c4.AID is not null and c4.AID != " + aid + " " +
		"GROUP BY c4.AID ORDER BY rank DESC")

}
func AuthorBarGraphPapersByYear(c chan interface{}, aid string) {
	//c <- GenerateSQL("SELECT PubYear, Number_of_Papers FROM Pubmed20_C04.BinChenBarGraphPapersByYear")
	c <- GenerateSQL("SELECT d2.PubYear,count(a2.PMID) as NumberOfPapers FROM Pubmed20_C04.A02_AuthorList a2 " +
		"JOIN Pubmed20_C04.A01_Articles a1 ON a2.PMID = a1.PMID " +
		"JOIN Pubmed20_C04.D02_PublishingYear d2 ON d2.PMID = a2.PMID " +
		"JOIN Pubmed20_C04.D03_PaperAuthors d3 ON d3.PMID = a2.PMID " +
		"WHERE a2.AID = " + aid + " GROUP BY d2.PubYear DESC LIMIT 10")
}
func AuthorCard(c chan interface{}, aid string) {
	c <- GenerateSQL("SELECT a2.LastName, a2.ForeName, a13.Affiliation " +
		"FROM Pubmed20_C04.A02_AuthorList a2 " +
		"JOIN Pubmed20_C04.A13_AffiliationList a13 ON a13.PMID = a2.PMID " +
		"WHERE a2.AID = " + aid + " LIMIT 1")
}

func AuthorWordCloud() interface{} {
	return GenerateSQL("SELECT Short_Keywords, frequency FROM Pubmed20_C04.BinChenAbstractKeywordsShort ORDER BY frequency DESC;")
}

func AuthorWordCloudType(c chan interface{}) { //map[string]interface{} {
	res := make(map[string]interface{})
	disease := GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.BinChenDiseaseCount ORDER BY occurences DESC;")
	drug := GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.BinChenDrugCount ORDER BY occurences DESC;")
	gene := GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.BinChenGeneCount ORDER BY occurences DESC;")
	//mutation := GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.BinChenMutationcount ORDER BY occurences DESC;")
	//species := GenerateSQL("SELECT Mention, occurences FROM Pubmed20_C04.BinChenSpeciesCount ORDER BY occurences DESC;")
	res["disease"] = disease
	res["drug"] = drug
	res["gene"] = gene
	//res["mutation"] = mutation
	//res["species"] = species
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
