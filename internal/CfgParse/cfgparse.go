package CfgParse

type Header struct {
    Name string `json:"name"`
    Val string `json:"val"`
}

type URLEncoded struct {
    Var string `json:"var"`
    Val string `json:"val"`
}

type BookmarkBody struct {
    Mode string `json:"mode"`
    Urlencoded []URLEncoded `json:"urlencoded"`
    Jsonencoded string `json:"jsonencoded"`
}

type Bookmark struct {
    Id string `json:"id"` 
    Name string `json:"name"`
    Request string `json:"request"`
    Method string `json:"method"`
    Headers []Header `json:"headers"` 
    Body BookmarkBody `json:"body"`
}

type KeyVal struct {
    Key string
    Val string
}

type Environment struct {
    Name string `json:"Name"`
    Vars map[string]string  `json:"vars"`
}
