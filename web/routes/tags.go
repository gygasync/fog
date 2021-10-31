package routes

// type TagRoute struct {
// 	logger     common.Logger
// 	tplEngine  TplEngine
// 	tagService services.ITagService

// 	internalTplEngine *template.Template
// }

// func NewTagRoute(logger common.Logger, tplEngine TplEngine, tagService services.ITagService) *TagRoute {
// 	return &TagRoute{
// 		logger:            logger,
// 		tplEngine:         tplEngine,
// 		tagService:        tagService,
// 		internalTplEngine: template.Must(template.ParseFiles("./web/static/templates/tagList.template.html"))}
// }
// func (i *TagRoute) generateComponent(tags []models.Tag) string {
// 	var buf bytes.Buffer
// 	i.internalTplEngine.ExecuteTemplate(&buf, "tagList", tags)
// 	return buf.String()
// }

// func (i *TagRoute) HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	data, err := i.tagService.List()
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		i.logger.Warn("could not find tags ", err)
// 		return
// 	}

// 	var bodyContent []template.HTML
// 	component := i.generateComponent(data)
// 	bodyContent = append(bodyContent, template.HTML(component))

// 	page := Page{Header: Header{Title: "Tags"}, Body: Body{Content: bodyContent}}
// 	i.tplEngine.Render(w, "main", &page)
// }

// func (i *TagRoute) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	tagName := r.FormValue("tagName")
// 	if tagName != "" {
// 		i.tagService.Add(&models.Tag{Name: tagName})
// 		http.Redirect(w, r, "tags", http.StatusFound)
// 	}
// }
