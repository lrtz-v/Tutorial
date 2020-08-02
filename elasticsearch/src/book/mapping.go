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
