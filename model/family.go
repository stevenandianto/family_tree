package model

type FamilyTree struct {
	FamilyName  string
	RootFather  *FamilyNode
	RootMother  *FamilyNode
	FamilyCache map[string]*FamilyNode
}

type FamilyNode struct {
	Name     string
	Gender   string
	Father   *FamilyNode
	Mother   *FamilyNode
	Spouse   *FamilyNode
	Children []*FamilyNode
}
