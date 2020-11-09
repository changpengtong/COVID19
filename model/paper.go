package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Paper struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	Cord_uid     string `bson:"cord_uid" json:"cord_uid"`
	Sha          string `bson:"sha" json:"sha"`
	Source_x     string `bson:"source_x" json:"source_x"`
	Title        string `bson:"title" json:"title"`
	Doi          string `bson:"doi" json:"doi"`
	Pmcid        string `bson:"pmcid" json:"pmcid"`
	Pubmed_id    string `bson:"pubmed_id" json:"pubmed_id"`
	License      string `bson:"license" json:"license"`
	Abstract     string `bson:"abstract" json:"abstract"`
	Publish_time string `bson:"publish_time" json:"publish_time"`
	Authors      string `bson:"authors" json:"authors"`
	Journal      string `bson:"journal" json:"journal"`
	Mag_id       string `bson:"mag_id" json:"mag_id"`
	Who_covidence_id 	string `bson:"who_covidence_id" json:"who_covidence_id"`
	Arxiv_id 		string `bson:"arxiv_id" json:"arxiv_id"`
	Pdf_json_files string `bson:"pdf_json_files" json:"pdf_json_files"`
	Pmc_json_files string `bson:"pmc_json_files" json:"pmc_json_files"`
	Url string `bson:"url" json:"url"`
	S2_id string `bson:"s2_id" json:"s2_id"`
}


