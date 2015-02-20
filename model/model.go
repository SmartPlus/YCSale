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

	UserModel = userModel{
		Name: collections["user"],
	}

	CustomerModel = customerModel{
		Name: collections["customer"],
	}

	CourseModel = courseModel{
		Name: collections["course"],
	}
	return
}
