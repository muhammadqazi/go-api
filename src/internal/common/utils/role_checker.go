package utils

func RoleChecker(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if role == userRole {
			return true
		}
	}
	return false
}
