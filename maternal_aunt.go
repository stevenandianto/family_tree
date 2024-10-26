package main

import (
	"strings"
)

type MaternalAunt struct {
}

func (this *MaternalAunt) getRelationship(familyTree *FamilyTree, name string) string {
	result := "NONE"
	resultArr := []string{}

	if familyTree.FamilyCache[name] == nil || familyTree.FamilyCache[name].Mother == nil {
		return result
	} else {
		mother := familyTree.FamilyCache[name].Mother
		grandMother := mother.Mother
		if grandMother == nil {
			return result
		} else {
			for _, v := range grandMother.Children {
				if v.Gender == "Female" && v.Name != mother.Name {
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
