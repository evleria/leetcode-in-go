package _551_student_attendance_record_I

func checkRecord(s string) bool {
	wasAbsent, late := false, 0
	for _, r := range s {
		switch r {
		case 'A':
			if wasAbsent {
				return false
			}
			wasAbsent, late = true, 0
		case 'L':
			if late == 2 {
				return false
			}
			late++
		default:
			late = 0
		}
	}

	return true
}
