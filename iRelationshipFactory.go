package main

import (
	"fmt"
)

type IRelationshipFactory interface {
	getRelationship(familyTree *FamilyTree, name string) string
}

func GetRelationshipFactory(relationship string) (IRelationshipFactory, error) {
	switch relationship {
	case "Paternal-Uncle":
		return &PaternalUncle{}, nil
	case "Maternal-Uncle":
		return &MaternalUncle{}, nil
	case "Paternal-Aunt":
		return &PaternalAunt{}, nil
	case "Maternal-Aunt":
		return &MaternalAunt{}, nil
	case "Sister-In-Law":
		return &SisterInLaw{}, nil
	case "Brother-In-Law":
		return &BrotherInLaw{}, nil
	// case "Son":
	// 	return &Son{}, nil
	// case "Daughter":
	// 	return &Daughter{}, nil
	// case "Siblings":
	// 	return &Siblings{}, nil
	default:
		return nil, fmt.Errorf("Wrong relationship type passed")
	}
}
