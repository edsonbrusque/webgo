----------
Type Template

template.Template

----------
Parsing templates

template.ParseFiles

func ParseFiles(filenames ...string) (*Template, error)

template.ParseGlob

func ParseGlob(pattern string) (*Template, error)

template.Parse

func (t *Template) Parse(text string) (*Template, error)

template.ParseFiles

func (t *Template) ParseFiles(filenames ...string) (*Template, error)

template.ParseGlob

func (t *Template) ParseGlob(pattern string) (*Template, error)

----------
Executing templates

template.Execute

func (t *Template) Execute(wr io.Writer, data interface{}) error

template.ExecuteTemplate

func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

----------
Helpful template functions

template.Must

func Must(t *Template, err error) *Template

template.New

func New(name string) *Template

----------
The init function

func init()
