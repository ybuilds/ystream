package models

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value"`
	RankingName  string `bson:"ranking_name" json:"ranking_name"`
}
