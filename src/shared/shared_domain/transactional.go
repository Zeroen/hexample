package shared_domain

type Transactional interface {
	Begin() error
	Commit() error
	Rollback() error
}

func BeginTX(tr Transactional) error {
	return tr.Begin()
}

func EndTX(tr Transactional, ptrErr *error) {
	err := recover()
	if err != nil {
		_ = tr.Rollback()
		panic(err)
	}

	if *ptrErr == nil {
		_ = tr.Commit()
	}

	if *ptrErr != nil {
		_ = tr.Rollback()
	}
}
