package model

type pipeline string

const(

	Radar string = `[
  {
    "$addFields": {
      "words": {
        "$map": {
          "input": {
            "$split": [
              "$source_x",
              " "
            ]
          },
          "as": "str",
          "in": {
            "$trim": {
              "input": {
                "$toLower": [
                  "$$str"
                ]
              },
              "chars": " ,|(){}-<>.;"
            }
          }
        }
      }
    }
  },
  {
    "$unwind": "$words"
  },
  {
    "$match": {
      "words": {
        "$nin": [
          "",
          "also",
          "i",
          "me",
          "my",
          "myself",
          "we",
          "us",
          "our",
          "ours",
          "ourselves",
          "you",
          "your",
          "yours",
          "yourself",
          "yourselves",
          "he",
          "him",
          "his",
          "himself",
          "she",
          "her",
          "hers",
          "herself",
          "it",
          "its",
          "itself",
          "they",
          "them",
          "their",
          "theirs",
          "themselves",
          "what",
          "which",
          "who",
          "whom",
          "whose",
          "this",
          "that",
          "these",
          "those",
          "am",
          "is",
          "are",
          "was",
          "were",
          "be",
          "been",
          "being",
          "have",
          "has",
          "had",
          "having",
          "do",
          "does",
          "did",
          "doing",
          "will",
          "would",
          "should",
          "can",
          "could",
          "ought",
          "i'm",
          "you're",
          "he's",
          "she's",
          "it's",
          "we're",
          "they're",
          "i've",
          "you've",
          "we've",
          "they've",
          "i'd",
          "you'd",
          "he'd",
          "she'd",
          "we'd",
          "they'd",
          "i'll",
          "you'll",
          "he'll",
          "she'll",
          "we'll",
          "they'll",
          "isn't",
          "aren't",
          "wasn't",
          "weren't",
          "hasn't",
          "haven't",
          "hadn't",
          "doesn't",
          "don't",
          "didn't",
          "won't",
          "wouldn't",
          "shan't",
          "shouldn't",
          "can't",
          "cannot",
          "couldn't",
          "mustn't",
          "let's",
          "that's",
          "who's",
          "what's",
          "here's",
          "there's",
          "when's",
          "where's",
          "why's",
          "how's",
          "a",
          "an",
          "the",
          "and",
          "but",
          "if",
          "or",
          "because",
          "as",
          "until",
          "while",
          "of",
          "at",
          "by",
          "for",
          "with",
          "about",
          "against",
          "between",
          "into",
          "through",
          "during",
          "before",
          "after",
          "above",
          "below",
          "to",
          "from",
          "up",
          "upon",
          "down",
          "in",
          "out",
          "on",
          "off",
          "over",
          "under",
          "again",
          "further",
          "then",
          "once",
          "here",
          "there",
          "when",
          "where",
          "why",
          "how",
          "all",
          "any",
          "both",
          "each",
          "few",
          "more",
          "most",
          "other",
          "some",
          "such",
          "no",
          "nor",
          "not",
          "only",
          "own",
          "same",
          "so",
          "than",
          "too",
          "very",
          "say",
          "says",
          "said",
          "shall"
        ]
      }
    }
  },
  {
    "$group": {
      "_id": {
        "__alias_0": "$words"
      },
      "__alias_1": {
        "$sum": {
          "$cond": [
            {
              "$ne": [
                {
                  "$type": "$words"
                },
                "missing"
              ]
            },
            1,
            0
          ]
        }
      }
    }
  },
  {
    "$project": {
      "_id": 0,
      "__alias_0": "$_id.__alias_0",
      "__alias_1": 1
    }
  },
  {
    "$project": {
      "group": "$__alias_0",
      "value": "$__alias_1",
      "_id": 0
    }
  },
  {"$sort":{"value":-1}},
  {
    "$limit": 5
  }
]`
	Linefold string = `[
  {
    "$match": {
      "publish_time": {
        "$gt": "2020-01-00",
		"$lt":"2021-01-00"
      }
    }
  },
  {
    "$group": {
      "_id": {
        "__alias_0": "$publish_time"
      },
      "__alias_1": {
        "$sum": {
          "$cond": [
            {
              "$ne": [
                {
                  "$type": "$publish_time"
                },
                "missing"
              ]
            },
            1,
            0
          ]
        }
      }
    }
  },
  {
    "$project": {
      "_id": 0,
      "__alias_0": "$_id.__alias_0",
      "__alias_1": 1
    }
  },
  {
    "$project": {
      "x": "$__alias_0",
      "y": "$__alias_1",
      "_id": 0
    }
  },
  {"$sort":{"x":1}},
  {
    "$limit": 500000
  }
]`
	Pie string = `[
  {
    "$addFields": {
      "words": {
        "$map": {
          "input": {
            "$split": [
              "$title",
              " "
            ]
          },
          "as": "str",
          "in": {
            "$trim": {
              "input": {
                "$toLower": [
                  "$$str"
                ]
              },
              "chars": " ,|(){}-<>.;:?!_~#$%&*;/"
            }
          }
        }
      }
    }
  },
  {
    "$unwind": "$words"
  },
  {
    "$match": {
      "words": {
        "$nin": [
          "",
          "also",
          "i",
          "me",
          "my",
          "myself",
          "we",
          "us",
          "our",
          "ours",
          "ourselves",
          "you",
          "your",
          "yours",
          "yourself",
          "yourselves",
          "he",
          "him",
          "his",
          "himself",
          "she",
          "her",
          "hers",
          "herself",
          "it",
          "its",
          "itself",
          "they",
          "them",
          "their",
          "theirs",
          "themselves",
          "what",
          "which",
          "who",
          "whom",
          "whose",
          "this",
          "that",
          "these",
          "those",
          "am",
          "is",
          "are",
          "was",
          "were",
          "be",
          "been",
          "being",
          "have",
          "has",
          "had",
          "having",
          "do",
          "does",
          "did",
          "doing",
          "will",
          "would",
          "should",
          "can",
          "could",
          "ought",
          "i'm",
          "you're",
          "he's",
          "she's",
          "it's",
          "we're",
          "they're",
          "i've",
          "you've",
          "we've",
          "they've",
          "i'd",
          "you'd",
          "he'd",
          "she'd",
          "we'd",
          "they'd",
          "i'll",
          "you'll",
          "he'll",
          "she'll",
          "we'll",
          "they'll",
          "isn't",
          "aren't",
          "wasn't",
          "weren't",
          "hasn't",
          "haven't",
          "hadn't",
          "doesn't",
          "don't",
          "didn't",
          "won't",
          "wouldn't",
          "shan't",
          "shouldn't",
          "can't",
          "cannot",
          "couldn't",
          "mustn't",
          "let's",
          "that's",
          "who's",
          "what's",
          "here's",
          "there's",
          "when's",
          "where's",
          "why's",
          "how's",
          "a",
          "an",
          "the",
          "and",
          "but",
          "if",
          "or",
          "because",
          "as",
          "until",
          "while",
          "of",
          "at",
          "by",
          "for",
          "with",
          "about",
          "against",
          "between",
          "into",
          "through",
          "during",
          "before",
          "after",
          "above",
          "below",
          "to",
          "from",
          "up",
          "upon",
          "down",
          "in",
          "out",
          "on",
          "off",
          "over",
          "under",
          "again",
          "further",
          "then",
          "once",
          "here",
          "there",
          "when",
          "where",
          "why",
          "how",
          "all",
          "any",
          "both",
          "each",
          "few",
          "more",
          "most",
          "other",
          "some",
          "such",
          "no",
          "nor",
          "not",
          "only",
          "own",
          "same",
          "so",
          "than",
          "too",
          "very",
          "say",
          "says",
          "single",
          "using",
          "high",
          "host",
          "-",
          "versus",
          "1",
          "type",
          "data",
          "multiple",
          "era",
          "based",
          "two",
          "said",
          "among",
          "role",
          "la",
          "first",
          "controlled",
          "shall",
          "not",
          "use",
          "en",
          "2",
          "new",
          "by",
          "its",
          "or",
          "its",
          "do",
          "as",
          "or",
          "de"
        ]
      }
    }
  },
  {
    "$match": {
      "words": {
        "$nin": [
          null,
          "",
          "A",
          "AN",
          "AND",
          "FOR",
          "FROM",
          "IN",
          "OF",
          "ON",
          "THE",
          "TO",
          "WITH"
        ]
      }
    }
  },
  {
    "$group": {
      "_id": {
        "__alias_0": "$words",
        "__alias_1": "$words"
      },
      "__alias_2": {
        "$sum": {
          "$cond": [
            {
              "$ne": [
                {
                  "$type": "$words"
                },
                "missing"
              ]
            },
            1,
            0
          ]
        }
      }
    }
  },
  {
    "$project": {
      "_id": 0,
      "__alias_0": "$_id.__alias_0",
      "__alias_1": "$_id.__alias_1",
      "__alias_2": 1
    }
  },
  {
    "$project": {
      "text": "$__alias_0",
      "size": "$__alias_2",
      "color": "$__alias_1",
      "_id": 0
    }
  },
  {"$sort":{"size":-1}},
  {
    "$limit": 10
  }
]`

	WordCloud string = `[
  {
    "$addFields": {
      "words": {
        "$map": {
          "input": {
            "$split": [
              "$title",
              " "
            ]
          },
          "as": "str",
          "in": {
            "$trim": {
              "input": {
                "$toLower": [
                  "$$str"
                ]
              },
              "chars": " ,|(){}-<>.;:?!_~#$%&*;/"
            }
          }
        }
      }
    }
  },
  {
    "$unwind": "$words"
  },
  {
    "$match": {
      "words": {
        "$nin": [
          "",
          "also",
          "i",
          "me",
          "my",
          "myself",
          "we",
          "us",
          "our",
          "ours",
          "ourselves",
          "you",
          "your",
          "yours",
          "yourself",
          "yourselves",
          "he",
          "him",
          "his",
          "himself",
          "she",
          "her",
          "hers",
          "herself",
          "it",
          "its",
          "itself",
          "they",
          "them",
          "their",
          "theirs",
          "themselves",
          "what",
          "which",
          "who",
          "whom",
          "whose",
          "this",
          "that",
          "these",
          "those",
          "am",
          "is",
          "are",
          "was",
          "were",
          "be",
          "been",
          "being",
          "have",
          "has",
          "had",
          "having",
          "do",
          "does",
          "did",
          "doing",
          "will",
          "would",
          "should",
          "can",
          "could",
          "ought",
          "i'm",
          "you're",
          "he's",
          "she's",
          "it's",
          "we're",
          "they're",
          "i've",
          "you've",
          "we've",
          "they've",
          "i'd",
          "you'd",
          "he'd",
          "she'd",
          "we'd",
          "they'd",
          "i'll",
          "you'll",
          "he'll",
          "she'll",
          "we'll",
          "they'll",
          "isn't",
          "aren't",
          "wasn't",
          "weren't",
          "hasn't",
          "haven't",
          "hadn't",
          "doesn't",
          "don't",
          "didn't",
          "won't",
          "wouldn't",
          "shan't",
          "shouldn't",
          "can't",
          "cannot",
          "couldn't",
          "mustn't",
          "let's",
          "that's",
          "who's",
          "what's",
          "here's",
          "there's",
          "when's",
          "where's",
          "why's",
          "how's",
          "a",
          "an",
          "the",
          "and",
          "but",
          "if",
          "or",
          "because",
          "as",
          "until",
          "while",
          "of",
          "at",
          "by",
          "for",
          "with",
          "about",
          "against",
          "between",
          "into",
          "through",
          "during",
          "before",
          "after",
          "above",
          "below",
          "to",
          "from",
          "up",
          "upon",
          "down",
          "in",
          "out",
          "on",
          "off",
          "over",
          "under",
          "again",
          "further",
          "then",
          "once",
          "here",
          "there",
          "when",
          "where",
          "why",
          "how",
          "all",
          "any",
          "both",
          "each",
          "few",
          "more",
          "most",
          "other",
          "some",
          "such",
          "no",
          "nor",
          "not",
          "only",
          "own",
          "same",
          "so",
          "than",
          "too",
          "very",
          "say",
          "says",
          "single",
          "using",
          "high",
          "host",
          "-",
          "versus",
          "1",
          "type",
          "data",
          "multiple",
          "era",
          "based",
          "two",
          "said",
          "among",
          "role",
          "la",
          "first",
          "controlled",
          "shall",
          "not",
          "use",
          "en",
          "2",
          "new",
          "by",
          "its",
          "or",
          "its",
          "do",
          "as",
          "or",
          "de"
        ]
      }
    }
  },
  {
    "$match": {
      "words": {
        "$nin": [
          null,
          "",
          "A",
          "AN",
          "AND",
          "FOR",
          "FROM",
          "IN",
          "OF",
          "ON",
          "THE",
          "TO",
          "WITH"
        ]
      }
    }
  },
  {
    "$group": {
      "_id": {
        "__alias_0": "$words",
        "__alias_1": "$words"
      },
      "__alias_2": {
        "$sum": {
          "$cond": [
            {
              "$ne": [
                {
                  "$type": "$words"
                },
                "missing"
              ]
            },
            1,
            0
          ]
        }
      }
    }
  },
  {
    "$project": {
      "_id": 0,
      "__alias_0": "$_id.__alias_0",
      "__alias_1": "$_id.__alias_1",
      "__alias_2": 1
    }
  },
  {
    "$project": {
      "text": "$__alias_0",
      "size": "$__alias_2",
      "color": "$__alias_1",
      "_id": 0
    }
  },
  {"$sort":{"size":-1}},
  {
    "$limit": 100
  }
]`

	Geo string = `[
  {
    "$match": {
      "metadata.authors.affiliation.location.country": {
        "$nin": []
      }
    }
  },
  {
    "$addFields": {
      "__alias_0": {
        "$arrayElemAt": [
          {
            "$switch": {
              "branches": [
                {
                  "case": {
                    "$isArray": "$metadata.authors.affiliation.location.country"
                  },
                  "then": "$metadata.authors.affiliation.location.country"
                },
                {
                  "case": {
                    "$in": [
                      {
                        "$type": "$metadata.authors.affiliation.location.country"
                      },
                      [
                        "null",
                        "missing",
                        "object"
                      ]
                    ]
                  },
                  "then": []
                }
              ],
              "default": [
                "$metadata.authors.affiliation.location.country"
              ]
            }
          },
          0
        ]
      }
    }
  },
  {
    "$addFields": {
      "__alias_1": {
        "$size": {
          "$filter": {
            "input": {
              "$switch": {
                "branches": [
                  {
                    "case": {
                      "$isArray": "$metadata.authors.affiliation.location.country"
                    },
                    "then": "$metadata.authors.affiliation.location.country"
                  },
                  {
                    "case": {
                      "$in": [
                        {
                          "$type": "$metadata.authors.affiliation.location.country"
                        },
                        [
                          "null",
                          "missing",
                          "object"
                        ]
                      ]
                    },
                    "then": []
                  }
                ],
                "default": [
                  "$metadata.authors.affiliation.location.country"
                ]
              }
            },
            "as": "str",
            "cond": {
              "$eq": [
                "$$str",
                "case-insenstive"
              ]
            }
          }
        }
      }
    }
  },
  {
    "$group": {
      "_id": {
        "__alias_0": "$__alias_0"
      },
      "__alias_1": {
        "$sum": {
          "$cond": [
            {
              "$ne": [
                {
                  "$type": "$__alias_1"
                },
                "missing"
              ]
            },
            1,
            0
          ]
        }
      }
    }
  },
  {
    "$project": {
      "_id": 0,
      "__alias_0": "$_id.__alias_0",
      "__alias_1": 1
    }
  },
  {
    "$project": {
      "location": "$__alias_0",
      "color": "$__alias_1",
      "_id": 0
    }
  },
  {
    "$limit": 5000
  }
]`
)
