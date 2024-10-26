package main

import (
	"family-tree/model"
	"fmt"
)

type FamilyTree model.FamilyTree
type FamilyNode model.FamilyNode

func Constructor(familyName, fatherName, motherName string) *FamilyTree {
	rootFather := &model.FamilyNode{
		Name:   fatherName,
		Gender: "Male",
	}
	rootMother := &model.FamilyNode{
		Name:   motherName,
		Gender: "Female",
	}

	rootFather.Spouse = rootMother
	rootMother.Spouse = rootFather

	familyCache := make(map[string]*model.FamilyNode)
	familyCache[fatherName] = rootFather
	familyCache[motherName] = rootMother

	return &FamilyTree{
		FamilyName:  familyName,
		RootFather:  rootFather,
		RootMother:  rootMother,
		FamilyCache: familyCache,
	}
}

func (this *FamilyTree) CheckNameExist(name string) bool {
	if _, ok := this.FamilyCache[name]; ok {
		return true
	}

	return false
}

func (this *FamilyTree) AddSpouse(groomName, brideName string) {
	//Check if groom is from other family
	if !this.CheckNameExist(groomName) {
		groom := &model.FamilyNode{
			Name:   groomName,
			Gender: "Male",
		}

		this.FamilyCache[brideName].Spouse = groom
		this.FamilyCache[groomName] = groom
		this.FamilyCache[groomName].Spouse = this.FamilyCache[brideName]
	}

	//Check if bride is from other family
	if !this.CheckNameExist(brideName) {
		bride := &model.FamilyNode{
			Name:   brideName,
			Gender: "Female",
		}

		this.FamilyCache[groomName].Spouse = bride
		this.FamilyCache[brideName] = bride
		this.FamilyCache[brideName].Spouse = this.FamilyCache[groomName]
	}
}

func (this *FamilyTree) AddChild(motherName, name, gender string) {
	//Check if the mother exist and make sure the mother is female (add child through the mother)
	if this.CheckNameExist(motherName) {
		if this.FamilyCache[motherName].Gender == "Female" {
			newChild := &model.FamilyNode{
				Name:   name,
				Gender: gender,
				Father: this.FamilyCache[motherName].Spouse,
				Mother: this.FamilyCache[motherName],
			}
			this.FamilyCache[motherName].Children = append(this.FamilyCache[motherName].Children, newChild)
			this.FamilyCache[name] = newChild
			fmt.Println("CHILD_ADDED")
		} else {
			//Error handling the mother is Male
			fmt.Sprintf("[CHILD_ADDITION_FAILED] %s is Male \n", motherName)
		}
	} else {
		//Error handling the mother is not exist
		fmt.Println("PERSON_NOT_FOUND")
	}
}

func (this *FamilyTree) GetRelationship(name, relationship string) {
	relationshipFactory, _ := GetRelationshipFactory(relationship)
	fmt.Println(relationshipFactory.getRelationship(this, name))
}

func main() {
	//Construct the Lengaburu family tree with the root fater "King Arthur" and root mother "Queen Margaret"
	lengaburuFamilyTree := Constructor("Lengaburu", "King Arthur", "Queen Margaret")

	//First Line Family
	lengaburuFamilyTree.AddChild("Queen Margaret", "Bill", "Male")
	lengaburuFamilyTree.AddChild("Queen Margaret", "Charlie", "Male")
	lengaburuFamilyTree.AddChild("Queen Margaret", "Percy", "Male")
	lengaburuFamilyTree.AddChild("Queen Margaret", "Ronald", "Male")
	lengaburuFamilyTree.AddChild("Queen Margaret", "Ginerva", "Female")

	//Add Spouse to First Line
	lengaburuFamilyTree.AddSpouse("Bill", "Flora")
	lengaburuFamilyTree.AddSpouse("Percy", "Audrey")
	lengaburuFamilyTree.AddSpouse("Ronald", "Helen")
	lengaburuFamilyTree.AddSpouse("Harry", "Ginerva")

	//Second Line Family
	//Add Children of Bill and Flora
	lengaburuFamilyTree.AddChild("Flora", "Victoire", "Female")
	lengaburuFamilyTree.AddChild("Flora", "Dominique", "Female")
	lengaburuFamilyTree.AddChild("Flora", "Louis", "Male")

	//Add Children of Percy and Audrey
	lengaburuFamilyTree.AddChild("Audrey", "Molly", "Female")
	lengaburuFamilyTree.AddChild("Audrey", "Lucy", "Female")

	//Add Children of Ronald and Helen
	lengaburuFamilyTree.AddChild("Helen", "Rose", "Female")
	lengaburuFamilyTree.AddChild("Helen", "Hugo", "Male")

	//Add Children of Harry and Ginerva
	lengaburuFamilyTree.AddChild("Ginerva", "James", "Male")
	lengaburuFamilyTree.AddChild("Ginerva", "Albus", "Male")
	lengaburuFamilyTree.AddChild("Ginerva", "Lily", "Female")

	//Add Spouse to Second Line
	lengaburuFamilyTree.AddSpouse("Ted", "Victoire")
	lengaburuFamilyTree.AddSpouse("Malloy", "Rose")
	lengaburuFamilyTree.AddSpouse("James", "Darcy")
	lengaburuFamilyTree.AddSpouse("Albus", "Alice")

	//Third Line Family
	//Add Children of Ted and Victoire
	lengaburuFamilyTree.AddChild("Victoire", "Remus", "Male")

	//Add Children of Malloy and Rose
	lengaburuFamilyTree.AddChild("Rose", "Draco", "Male")
	lengaburuFamilyTree.AddChild("Rose", "Aster", "Female")

	//Add Children of James and Darcy
	lengaburuFamilyTree.AddChild("Darcy", "William", "Male")

	//Add Children of Albus and Alice
	lengaburuFamilyTree.AddChild("Alice", "Ron", "Male")
	lengaburuFamilyTree.AddChild("Alice", "Ginny", "Female")

	//Check Relationship
	lengaburuFamilyTree.GetRelationship("Dominique", "Paternal-Uncle")
	lengaburuFamilyTree.GetRelationship("James", "Maternal-Uncle")
	lengaburuFamilyTree.GetRelationship("Dominique", "Paternal-Aunt")
	lengaburuFamilyTree.GetRelationship("Remus", "Maternal-Aunt")
	lengaburuFamilyTree.GetRelationship("Lily", "Sister-In-Law")
	lengaburuFamilyTree.GetRelationship("Hugo", "Brother-In-Law")
}
