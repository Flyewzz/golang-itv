package main

func CompareSets(slice1, slice2 []Task) bool {
	findElement := func(elem Task, slice []Task) bool {
		for _, curElem := range slice {
			if elem == curElem {
				return true
			}
		}
		return false
	}
	for _, element := range slice1 {
		if !findElement(element, slice2) {
			return false
		}
	}
	return true
}
