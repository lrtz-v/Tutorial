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
			"index": true,
			"analyzer": "smartcn",
			"fields": {
				"raw": {
					"type":  "keyword",
					"index": true
				}
			}
		},
		"size" : {
			"type" : "long"
		},
		"type" : {
			"type" : "text",
			"index": true,
			"analyzer": "english",
			"fields": {
				"raw": {
					"type":  "keyword",
					"index": true
				}
			}
		}
	}
}
`
