package tile

// tile :
type tile struct {
	Id       string `bson:"_id" json:"id" uri:"id"`
	Name     string `bson:"name" json:"name"`
	Title    string `bson:"title" json:"title"`
	SubTitle string `bson:"sub_title" json:"sub_title"`
	Type     string `bson:"type" json:"type"`
	Path     string `bson:"path" json:"path"`
}
