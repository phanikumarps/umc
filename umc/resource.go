package umc

func getResource(name *string, id *string) *string {
	switch n := *name; n {
	case "accounts":
		u := "Accounts" + "('" + *id + "')"
		return &u
	default:
		u := "$metadata"
		return &u

	}
}
