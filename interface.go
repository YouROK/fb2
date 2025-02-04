package fb2

import (
	"encoding/xml"
)

// List of interfaces for integration

// FB2 represents FB2 structure
//proteus:generate
type FB2 struct {
	ID          string      `bson:"_id"`
	FictionBook xml.Name    `xml:"FictionBook" bson:"FictionBook"`
	Stylesheet  []string    `xml:"stylesheet" bson:"stylesheet"`
	Description Description `xml:"description" bson:"description"`
	Body        Body        `xml:"body" bson:"body"`
	Binary      []Binary    `xml:"binary" bson:"binary"`
}

type Description struct {
	TitleInfo    TitleInfo    `xml:"title-info" bson:"title-info"`
	DocumentInfo DocumentInfo `xml:"document-info" bson:"document-info"`
	PublishInfo  PublishInfo  `xml:"PublishInfo" bson:"PublishInfo"`
	CustomInfo   []CustomInfo `xml:"custom-info" bson:"custom-info"`
}

type TitleInfo struct {
	Genre      []string     `xml:"genre" bson:"genre"`
	GenreType  []string     `xml:"genreType" bson:"genreType"`
	Author     []AuthorType `xml:"author" bson:"author"`
	BookTitle  string       `xml:"book-title" bson:"book-title"`
	Annotation Title       `xml:"annotation" bson:"annotation"`
	Keywords   string       `xml:"keywords" bson:"keywords"`
	Date       string       `xml:"date" bson:"date"`
	Coverpage  Coverpage    `xml:"coverpage" bson:"coverpage"`
	Lang       string       `xml:"lang" bson:"lang"`
	SrcLang    string       `xml:"src-lang" bson:"src-lang"`
	Translator AuthorType   `xml:"translator" bson:"translator"`
	Sequence   string       `xml:"sequence" bson:"sequence"`
}

type Coverpage struct {
	Image Image `xml:"image,allowempty" bson:"image"`
}

type Image struct {
	Href string `xml:"xlink:href,attr" bson:"href"`
}

type DocumentInfo struct {
	Author      []AuthorType `xml:"author" bson:"author"`
	ProgramUsed string       `xml:"program-used" bson:"program-used"`
	Date        string       `xml:"date" bson:"date"`
	SrcURL      []string     `xml:"src-url" bson:"src-url"`
	SrcOcr      string       `xml:"src-ocr" bson:"src-ocr"`
	ID          string       `xml:"id" bson:"id"`
	Version     float64      `xml:"version" bson:"version"`
	History     Title        `xml:"history" bson:"history"`
}

type PublishInfo struct {
	BookName  string `xml:"book-name" bson:"book-name"`
	Publisher string `xml:"publisher" bson:"publisher"`
	City      string `xml:"city" bson:"city"`
	Year      int    `xml:"year" bson:"year"`
	ISBN      string `xml:"isbn" bson:"isbn"`
	Sequence  string `xml:"sequence" bson:"sequence"`
}

type CustomInfo struct {
	InfoType string `xml:"info-type" bson:"info-type"`
}

type Binary struct {
	Value       string `xml:",chardata" bson:"value"`
	ContentType string `xml:"content-type,attr" bson:"content-type"`
	ID          string `xml:"id,attr" bson:"id"`
}

type Sections struct {
	P     []string `xml:"p" bson:"p"`
	Title Title    `xml:"title"`
	Image struct {
		Href string `xml:"href,attr" bson:"href"`
	} `xml:"image"`
	Subtitle string `xml:"subtitle"`
}

type Title struct {
	P []string `xml:"p"`
}

type Body struct {
	Sections []Sections `xml:"section" bson:"section"`
	Title    Title     `xml:"title"`
	Image    struct {
		Href string `xml:"href,attr" bson:"href"`
	} `xml:"image"`
	Subtitle string `xml:"subtitle"`
}

// AuthorType embedded fb2 type, represents author info
type AuthorType struct {
	FirstName  string `xml:"first-name"`
	MiddleName string `xml:"middle-name"`
	LastName   string `xml:"last-name"`
	Nickname   string `xml:"nickname"`
	HomePage   string `xml:"home-page"`
	Email      string `xml:"email"`
}

// UnmarshalCoverpage func
func (f *FB2) UnmarshalCoverpage(data []byte) {
	tagOpened := false
	coverpageStartIndex := 0
	coverpageEndIndex := 0
	// imageHref := ""
	tagName := ""
_loop:
	for i, v := range data {
		if tagOpened {
			switch v {
			case '>':
				if tagName != "p" && tagName != "/p" {
				}
				tagOpened = false
				if tagName == "coverpage" {
					coverpageStartIndex = i + 1
				} else if tagName == "/coverpage" {
					coverpageEndIndex = i - 11
					break _loop
				}
				tagName = ""
				break
			default:
				tagName += string(v)
			}
		} else {
			if v == '<' {
				tagOpened = true
			}
		}
	}

	if coverpageEndIndex > coverpageStartIndex {
		href := parseImage(data[coverpageStartIndex:coverpageEndIndex])
		f.Description.TitleInfo.Coverpage.Image.Href = href
	}
}

// TextFieldType embedded fb2 type, represents text field
type TextFieldType struct {
}

// TitleType embedded fb2 type, represents title type fields
type TitleType struct {
	P         []string `xml:"p"`
	EmptyLine []string `xml:"empty-line"`
}

// PType embedded fb2 type, represents paragraph
type PType struct {
}
