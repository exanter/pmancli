package CfgParse

type BookmarkHeader struct {
    Key string `yaml:"key"`
    Value string `yaml:"value"`
}

type BookmarkInfo struct {
    PostmanId string `yaml:"_postman_id"`
    Name string `yaml:"name"`
    Schema string `yaml:"schema"`
    ExporterId string `yaml:"_exporter_id"`
}

type BookmarkURL struct {
    Raw string `yaml:"raw"`
    Protocol string `yaml:"protocol"`
    Host []string `yaml:"host"`
    Path []string `yaml:"path"`
}

type BookmarkRequest struct {
    Method string `yaml:"method"`
    Header []BookmarkHeader `yaml:"header"`
    Body map[string]interface{} `yaml:"body"`
    Url BookmarkURL `yaml:"url"`
}

type BookmarkItem struct {
    Name string `yaml:"name"`
    Request BookmarkRequest `yaml:"request"`
    Response []map[string]interface{} `yaml:"response,omitempty"`
}

type Bookmark struct {
    Info BookmarkInfo `yaml:"info"`
    Item []BookmarkItem `yaml:"item"`
}

type EnvValues struct {
    Key string `yaml:"key"`
    Value string `yaml:"value"`
    Type string `yaml:"type"`
    Enabled bool `yaml:"enabled"`
}

type Environment struct {
    Id string `yaml:"id"`
    Name string `yaml:"Name"`
    Values []EnvValues  `yaml:"values"`
    PostmanVariableScope string `yaml:"_postman_veriable_scope"`
    PostmanExportedAt string `yaml:"_postman_exported_at"`
    PostmanExportedUsing string `yaml:"_postman_exported_using"`
}

