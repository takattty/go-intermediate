package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Post Article..\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List..\n")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId := 1
	resString := fmt.Sprintf("Article No.%d\n", articleId)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice..\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment..\n")
}
