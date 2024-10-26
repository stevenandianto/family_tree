package main

import (
	"family-tree/model"
	"strings"
)

type BrotherInLaw struct {
}

func (this *BrotherInLaw) getRelationship(familyTree *FamilyTree, name string) string {
	result := "NONE"
	resultArr := []string{}
	me := familyTree.FamilyCache[name]

	if me == nil {
		return result
	} else {
		//wive of siblings
		resultArr = append(resultArr, findHusbandSiblings(familyTree, me)...)

		//spouse of sisters
		resultArr = append(resultArr, findSpouseBrothers(familyTree, me)...)
	}

	if len(resultArr) > 0 {
		return strings.Join(resultArr, ",")
	}

	return result
}

func findHusbandSiblings(familyTree *FamilyTree, me *model.FamilyNode) []string {
	result := []string{}

	mother := me.Mother
	if mother == nil {
		return result
	} else {
		//husbands of siblings
		for _, v := range mother.Children {
			if v.Spouse != nil && v.Spouse.Gender == "Male" && v.Name != me.Name {
				result = append(result, v.Spouse.Name)
			}
		}
	}

	return result
}

func findSpouseBrothers(familyTree *FamilyTree, me *model.FamilyNode) []string {
	result := []string{}
	//spouse validation
	if me.Spouse == nil || me.Spouse.Mother == nil {
		return result
	} else {
		for _, v := range me.Spouse.Mother.Children {
			if v.Gender == "Male" && v.Name != me.Spouse.Name {
				result = append(result, v.Name)
			}
		}
	}

	return result
}
