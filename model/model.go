package model

type model struct {
	collections map[string]string
}

var context model

func Init(collections map[string]string) (err error) {
	context = model{
		collections: collections,
	}

	WishModel = wishModel{
		Name: collections["wish"],
	}
	return
}
