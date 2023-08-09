package listfactory

import "practice-go/design-pattern/abstract-factory/factory"

type ListFactory struct {
	factory.Factory
}

func (lf ListFactory) CreateLink(caption, url string) factory.LinkInterface {
	return NewListLink(caption, url)
}

func (lf ListFactory) CreateTray(caption string) factory.TrayInterface {
	return NewListTray(caption)
}

func (lf ListFactory) CreatePage(title, author string) factory.PageInterface {
	return NewListPage(title, author)
}

type ListLink struct {
	factory.Link
}

func NewListLink(caption, url string) ListLink {
	return ListLink{factory.NewLink(caption, url)}
}

func (ll ListLink) MakeHTML() string {
	return "  <li><a href=\"" + ll.URL + "\">" + ll.Caption + "</a></li>\n"
}

type ListTray struct {
	*factory.Tray
}

func NewListTray(caption string) ListTray {
	t := factory.NewTray(caption)
	return ListTray{&t}
}

func (lt ListTray) MakeHTML() string {
	buf := make([]byte, 0)
	buf = append(buf, []byte("<li>\n")...)
	buf = append(buf, []byte(lt.Caption+"\n")...)
	buf = append(buf, []byte("<ul>\n")...)

	it := lt.Iterator()
	for it.HasNext() {
		//item := it.Next().(factory.ItemInterface)
		item := it.Next().(factory.Item)
		buf = append(buf, []byte(item.MakeHTML())...)
	}

	buf = append(buf, []byte("</ul>\n")...)
	buf = append(buf, []byte("</li>\n")...)
	return string(buf)
}

type ListPage struct {
	*factory.Page
}

func NewListPage(title, author string) ListPage {
	p := factory.NewPage(title, author)
	return ListPage{&p}
}

func (lp ListPage) MakeHTML() string {
	buf := make([]byte, 0)
	buf = append(buf, []byte("<html><head><title>"+lp.Title+"</title></head>\n")...)
	buf = append(buf, []byte("<body>\n")...)
	buf = append(buf, []byte("<h1>"+lp.Title+"</h1>\n")...)
	buf = append(buf, []byte("<ul>\n")...)

	it := lp.Iterator()
	for it.HasNext() {
		//item := it.Next().(factory.ItemInterface)
		item := it.Next().(factory.Item)
		buf = append(buf, []byte(item.MakeHTML())...)
	}

	buf = append(buf, []byte("</ul>\n")...)
	buf = append(buf, []byte("<hr><address>"+lp.Author+"</address\n")...)
	buf = append(buf, []byte("</body></html>\n")...)
	return string(buf)
}
