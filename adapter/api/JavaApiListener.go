package api

import (
	"github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
	"strings"
)

var clz []models.JClassNode

type RestApi struct {
	Uri            string
	HttpMethod     string
	MethodName     string
	ResponseStatus string
	Body           []string
	MethodParams   map[string]string
}

var hasEnterClass = false
var isSpringRestController = false
var hasEnterRestController = false
var baseApiUrlName = ""

var currentRestApi RestApi
var RestApis []RestApi

func NewJavaApiListener() *JavaApiListener {
	isSpringRestController = false
	return &JavaApiListener{}
}

type JavaApiListener struct {
	BaseJavaParserListener
}

func (s *JavaApiListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	hasEnterClass = true
}

func (s *JavaApiListener) EnterAnnotation(ctx *AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "RestController" {
		isSpringRestController = true
	}

	if !isSpringRestController {
		return
	}

	if !hasEnterClass {
		if annotationName == "RequestMapping" {
			if ctx.ElementValuePairs() != nil {
				firstPair := ctx.ElementValuePairs().GetChild(0).(*ElementValuePairContext)
				if firstPair.IDENTIFIER().GetText() == "value" {
					baseApiUrlName = firstPair.ElementValue().GetText()
				}
			} else {
				baseApiUrlName = "/"
			}
		}
	}

	if !(annotationName == "GetMapping" || annotationName == "PutMapping" || annotationName == "PostMapping" || annotationName == "DeleteMapping") {
		return
	}

	hasEnterRestController = true
	uri := ""
	if ctx.ElementValue() != nil {
		uri = baseApiUrlName + ctx.ElementValue().GetText()
	} else {
		uri = baseApiUrlName
	}

	uriRemoveQuote := strings.ReplaceAll(uri, "\"", "")

	currentRestApi = RestApi{uriRemoveQuote, "", "", "", nil, nil}
	if hasEnterClass {
		switch annotationName {
		case "GetMapping":
			currentRestApi.HttpMethod = "GET"
		case "PutMapping":
			currentRestApi.HttpMethod = "PUT"
		case "PostMapping":
			currentRestApi.HttpMethod = "POST"
		case "DeleteMapping":
			currentRestApi.HttpMethod = "DELETE"
		}
	}
}

func (s *JavaApiListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	if hasEnterRestController {
		RestApis = append(RestApis, currentRestApi)
		hasEnterRestController = false
	}
}

func (s *JavaApiListener) appendClasses(classes []models.JClassNode) {
	clz = classes
}


func (s *JavaApiListener) getApis() []RestApi {
	return RestApis
}
