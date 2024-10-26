package main

import (
	"family-tree/model"
	"strings"
)

type SisterInLaw struct {
}

func (this *SisterInLaw) getRelationship(familyTree *FamilyTree, name string) string {
	result := "NONE"
	resultArr := []string{}
	me := familyTree.FamilyCache[name]

	if me == nil {
		return result
	} else {
		//wive of siblings
		resultArr = append(resultArr, findWiveSiblings(familyTree, me)...)

		//spouse of sisters
		resultArr = append(resultArr, findSpouseSisters(familyTree, me)...)
	}

	if len(resultArr) > 0 {
		return strings.Join(resultArr, ",")
	}

	return result
}

func findWiveSiblings(familyTree *FamilyTree, me *model.FamilyNode) []string {
	result := []string{}

	mother := me.Mother
	if mother == nil {
		return result
	} else {
		//wives of siblings
		for _, v := range mother.Children {
			if v.Spouse != nil && v.Spouse.Gender == "Female" && v.Name != me.Name {
				result = append(result, v.Spouse.Name)
			}
		}
	}

	return result
}

func findSpouseSisters(familyTree *FamilyTree, me *model.FamilyNode) []string {
	result := []string{}
	//spouse validation
	if me.Spouse == nil || me.Spouse.Mother == nil {
		return result
	} else {
		for _, v := range me.Spouse.Mother.Children {
			if v.Gender == "Female" && v.Name != me.Spouse.Name {
				result = append(result, v.Name)
			}
		}
	}

	return result
}
