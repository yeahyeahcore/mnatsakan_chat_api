package repository

// IRead - interface describes read responsibilities
type IRead[T interface{}] interface {
	Find(arguments interface{}) (*[]T, error)
	FindOne(arguments interface{}) (*T, error)
	FindByID(id string) (*T, error)
}

// IWrite - interface describes write responsibilities
type IWrite[T interface{}] interface {
	Insert(arguments interface{}) (*T, error)
	InsertMany(arguments []interface{}) (*[]T, error)
}

// IRemove - interface describes remove responsibilities
type IRemove[T interface{}] interface {
	DeleteByID(id string) (*T, error)
	DeleteOne(arguments interface{}) (*T, error)
	DeleteMany(arguments interface{}) (*[]T, error)
}

// IUpdate - interface describes update responsibilities
type IUpdate[T interface{}] interface {
	UpdateByID(id string) (*T, error)
	UpdateOne(arguments interface{})
	UpdateMany(arguments interface{})
}
