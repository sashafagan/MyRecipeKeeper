package data 

type recipes struct{
	Id string ``````````````````````json:"Id"
	RecipeName string`````````````` json:"RecipeName"
	Source string`````````````````` json:"Source"
	PrepTime string ````````````````json:"Preparation Time"
	CookTime string ````````````````json:"Cooking Time"
	ServingSize int ````````````````json:"Serving Size"
	Ingredients map[string]string   json:"Ingredients"
	Direction map[int]string````````json:"Direction"
	Tag []string````````````````````json:"Tags"
}
var ThisVariable Recipe