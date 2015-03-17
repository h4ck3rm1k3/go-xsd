package xsd
//import "fmt"
func (me *HasElemAll) makePkg(bag *PkgBag) {
	if me.All != nil {
		me.All.makePkg(bag)
	}
}

func (me *HasElemAnnotation) makePkg(bag *PkgBag) {
	if me.Annotation != nil {
		me.Annotation.makePkg(bag)
	}
}

func (me *HasElemsAny) makePkg(bag *PkgBag) {
	for _, any := range me.Anys {
		any.makePkg(bag)
	}
}

func (me *HasElemsAnyAttribute) makePkg(bag *PkgBag) {
	for _, aa := range me.AnyAttributes {
		aa.makePkg(bag)
	}
}

func (me *HasElemsAppInfo) makePkg(bag *PkgBag) {
	for _, ai := range me.AppInfos {
		ai.makePkg(bag)
	}
}

func (me *HasElemsAttribute) makePkg(bag *PkgBag) {
	for _, ea := range me.Attributes {
		ea.makePkg(bag)
	}
}

func (me *HasElemsAttributeGroup) makePkg(bag *PkgBag) {
	//fmt.Printf("makePkg HasElemsAttributeGroup %v\n", me)
	for _, ag := range me.AttributeGroups {
		ag.makePkg(bag)
	}
}

func (me *HasElemChoice) makePkg(bag *PkgBag) {
	if me.Choice != nil {
		me.Choice.makePkg(bag)
	}
}

func (me *HasElemsChoice) makePkg(bag *PkgBag) {
	for _, ch := range me.Choices {
		ch.makePkg(bag)
	}
}

func (me *HasElemComplexContent) makePkg(bag *PkgBag) {
	if me.ComplexContent != nil {
		me.ComplexContent.makePkg(bag)
	}
}

func (me *HasElemComplexType) makePkg(bag *PkgBag) {
	if me.ComplexType != nil {
		me.ComplexType.makePkg(bag)
	}
}

func (me *HasElemsComplexType) makePkg(bag *PkgBag) {
	for _, ct := range me.ComplexTypes {
		ct.makePkg(bag)
	}
}

func (me *HasElemsDocumentation) makePkg(bag *PkgBag) {
	for _, doc := range me.Documentations {
		doc.makePkg(bag)
	}
}

func (me *HasElemsElement) makePkg(bag *PkgBag) {
	for _, el := range me.Elements {
		el.makePkg(bag)
	}
}

func (me *HasElemsEnumeration) makePkg(bag *PkgBag) {
	for _, enum := range me.Enumerations {
		enum.makePkg(bag)
	}
}

func (me *HasElemExtensionComplexContent) makePkg(bag *PkgBag) {
	if me.ExtensionComplexContent != nil {
		me.ExtensionComplexContent.makePkg(bag)
	}
}

func (me *HasElemExtensionSimpleContent) makePkg(bag *PkgBag) {
	if me.ExtensionSimpleContent != nil {
		me.ExtensionSimpleContent.makePkg(bag)
	}
}

func (me *HasElemField) makePkg(bag *PkgBag) {
	if me.Field != nil {
		me.Field.makePkg(bag)
	}
}

func (me *HasElemFractionDigits) makePkg(bag *PkgBag) {
	if me.FractionDigits != nil {
		me.FractionDigits.makePkg(bag)
	}
}

func (me *HasElemGroup) makePkg(bag *PkgBag) {
	if me.Group != nil {
		me.Group.makePkg(bag)
	}
}

func (me *HasElemsGroup) makePkg(bag *PkgBag) {
	for _, gr := range me.Groups {
		gr.makePkg(bag)
	}
}

func (me *HasElemsImport) makePkg(bag *PkgBag) {
	for _, imp := range me.Imports {
		imp.makePkg(bag)
	}
}

func (me *HasElemsKey) makePkg(bag *PkgBag) {
	for _, k := range me.Keys {
		k.makePkg(bag)
	}
}

func (me *HasElemKeyRef) makePkg(bag *PkgBag) {
	if me.KeyRef != nil {
		me.KeyRef.makePkg(bag)
	}
}

func (me *HasElemLength) makePkg(bag *PkgBag) {
	if me.Length != nil {
		me.Length.makePkg(bag)
	}
}

func (me *HasElemList) makePkg(bag *PkgBag) {
	if me.List != nil {
		me.List.makePkg(bag)
	}
}

func (me *HasElemMaxExclusive) makePkg(bag *PkgBag) {
	if me.MaxExclusive != nil {
		me.MaxExclusive.makePkg(bag)
	}
}

func (me *HasElemMaxInclusive) makePkg(bag *PkgBag) {
	if me.MaxInclusive != nil {
		me.MaxInclusive.makePkg(bag)
	}
}

func (me *HasElemMaxLength) makePkg(bag *PkgBag) {
	if me.MaxLength != nil {
		me.MaxLength.makePkg(bag)
	}
}

func (me *HasElemMinExclusive) makePkg(bag *PkgBag) {
	if me.MinExclusive != nil {
		me.MinExclusive.makePkg(bag)
	}
}

func (me *HasElemMinInclusive) makePkg(bag *PkgBag) {
	if me.MinInclusive != nil {
		me.MinInclusive.makePkg(bag)
	}
}

func (me *HasElemMinLength) makePkg(bag *PkgBag) {
	if me.MinLength != nil {
		me.MinLength.makePkg(bag)
	}
}

func (me *HasElemsNotation) makePkg(bag *PkgBag) {
	for _, not := range me.Notations {
		not.makePkg(bag)
	}
}

func (me *HasElemPattern) makePkg(bag *PkgBag) {
	if me.Pattern != nil {
		me.Pattern.makePkg(bag)
	}
}

func (me *HasElemsRedefine) makePkg(bag *PkgBag) {
	for _, rd := range me.Redefines {
		rd.makePkg(bag)
	}
}

func (me *HasElemRestrictionComplexContent) makePkg(bag *PkgBag) {
	if me.RestrictionComplexContent != nil {
		me.RestrictionComplexContent.makePkg(bag)
	}
}

func (me *HasElemRestrictionSimpleContent) makePkg(bag *PkgBag) {
	if me.RestrictionSimpleContent != nil {
		me.RestrictionSimpleContent.makePkg(bag)
	}
}

func (me *HasElemRestrictionSimpleType) makePkg(bag *PkgBag) {
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.makePkg(bag)
	}
}

func (me *HasElemSelector) makePkg(bag *PkgBag) {
	if me.Selector != nil {
		me.Selector.makePkg(bag)
	}
}

func (me *HasElemSequence) makePkg(bag *PkgBag) {
	if me.Sequence != nil {
		me.Sequence.makePkg(bag)
	}
}

func (me *HasElemsSequence) makePkg(bag *PkgBag) {
	for _, seq := range me.Sequences {
		seq.makePkg(bag)
	}
}

func (me *HasElemSimpleContent) makePkg(bag *PkgBag) {
	if me.SimpleContent != nil {
		me.SimpleContent.makePkg(bag)
	}
}

func (me *HasElemsSimpleType) makePkg(bag *PkgBag) {
	for _, st := range me.SimpleTypes {
		st.makePkg(bag)
	}
}

func (me *HasElemTotalDigits) makePkg(bag *PkgBag) {
	if me.TotalDigits != nil {
		me.TotalDigits.makePkg(bag)
	}
}

func (me *HasElemUnion) makePkg(bag *PkgBag) {
	if me.Union != nil {
		me.Union.makePkg(bag)
	}
}

func (me *HasElemUnique) makePkg(bag *PkgBag) {
	if me.Unique != nil {
		me.Unique.makePkg(bag)
	}
}

func (me *HasElemWhiteSpace) makePkg(bag *PkgBag) {
	if me.WhiteSpace != nil {
		me.WhiteSpace.makePkg(bag)
	}
}

func (me *HasAttrAbstract) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrBase) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrBlock) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrDefault) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrFinal) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrFixed) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrForm) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrId) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrLang) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrMixed) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrName) beforeMakePkg(bag *PkgBag) {
	bag.Stacks.Name.Push(me.Name)
}

func (me *HasAttrNamespace) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrNillable) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrPublic) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrRef) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrRefer) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrSource) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrSystem) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrType) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrUse) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrValue) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrVersion) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrXpath) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrBlockDefault) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrFinalDefault) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrItemType) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrMaxOccurs) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrMemberTypes) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrMinOccurs) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrProcessContents) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrSchemaLocation) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrSubstitutionGroup) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrTargetNamespace) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrAttributeFormDefault) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrElementFormDefault) beforeMakePkg(bag *PkgBag) {
}

func (me *HasAttrAbstract) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrBase) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrBlock) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrDefault) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrFinal) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrFixed) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrForm) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrId) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrLang) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrMixed) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrName) afterMakePkg(bag *PkgBag) {
	bag.Stacks.Name.Pop()
}

func (me *HasAttrNamespace) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrNillable) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrPublic) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrRef) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrRefer) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrSource) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrSystem) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrType) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrUse) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrValue) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrVersion) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrXpath) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrBlockDefault) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrFinalDefault) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrItemType) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrMaxOccurs) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrMemberTypes) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrMinOccurs) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrProcessContents) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrSchemaLocation) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrSubstitutionGroup) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrTargetNamespace) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrAttributeFormDefault) afterMakePkg(bag *PkgBag) {
}

func (me *HasAttrElementFormDefault) afterMakePkg(bag *PkgBag) {
}
