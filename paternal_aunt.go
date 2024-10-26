package main

import (
	"strings"
)

type PaternalAunt struct {
}

func (this *PaternalAunt) getRelationship(familyTree *FamilyTree, name string) string {
	result := "NONE"
	resultArr := []string{}

	if familyTree.FamilyCache[name] == nil || familyTree.FamilyCache[name].Father == nil {
		return result
	} else {
		father := familyTree.FamilyCache[name].Father
		grandMother := father.Mother
		if grandMother == nil {
			return result
		} else {
			for _, v := range grandMother.Children {
				if v.Gender == "Female" && v.Name != father.Name {
					resultArr = append(resultArr, v.Name)
				}
			}
		}
	}

	if len(resultArr) > 0 {
		return strings.Join(resultArr, ",")
	}

	return result
}
