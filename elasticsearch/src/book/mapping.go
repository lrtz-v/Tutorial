package book

// Mapping of books
const Mapping = `
{
	"properties" : {
	"created" : {
		"type" : "date"
	},
	"id" : {
		"type" : "long"
	},
	"name" : {
		"type" : "text",
		"analyzer": "smartcn",
		"fields" : {
		"keyword" : {
			"type" : "keyword",
			"ignore_above" : 256
		}
		}
	},
	"size" : {
		"type" : "long"
	},
	"type" : {
		"type" : "text",
		"analyzer": "english",
		"fields" : {
		"keyword" : {
			"type" : "keyword",
			"ignore_above" : 256
		}
		}
	}
	}
}
`
