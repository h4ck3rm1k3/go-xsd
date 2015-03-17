package xsd
//import "fmt"

func (me *All) initElement(parent element) {
	me.ElemBase.init(parent, me, "all", &me.HasAttrId, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsElement.initChildren(me)
}

func (me *Annotation) initElement(parent element) {
	me.ElemBase.init(parent, me, "annotation")
	me.HasElemsAppInfo.initChildren(me)
	me.HasElemsDocumentation.initChildren(me)
}

func (me *Any) initElement(parent element) {
	me.ElemBase.init(parent, me, "any", &me.HasAttrId, &me.HasAttrNamespace, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs, &me.HasAttrProcessContents)
	me.HasElemAnnotation.initChildren(me)
}

func (me *AnyAttribute) initElement(parent element) {
	me.ElemBase.init(parent, me, "anyAttribute", &me.HasAttrId, &me.HasAttrNamespace, &me.HasAttrProcessContents)
	me.HasElemAnnotation.initChildren(me)
}

func (me *AppInfo) initElement(parent element) {
	me.ElemBase.init(parent, me, "appInfo", &me.HasAttrSource)
}

func (me *Attribute) initElement(parent element) {
	me.ElemBase.init(parent, me, "attribute", &me.HasAttrDefault, &me.HasAttrFixed, &me.HasAttrForm, &me.HasAttrId, &me.HasAttrName, &me.HasAttrRef, &me.HasAttrType, &me.HasAttrUse)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *AttributeGroup) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "attributeGroup", &me.HasAttrId, &me.HasAttrName, &me.HasAttrRef)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
}

func (me *Choice) initElement(parent element) {
	me.ElemBase.init(parent, me, "choice", &me.HasAttrId, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsAny.initChildren(me)
	me.HasElemsChoice.initChildren(me)
	me.HasElemsElement.initChildren(me)
	me.HasElemsGroup.initChildren(me)
	me.HasElemsSequence.initChildren(me)
}

func (me *ComplexContent) initElement(parent element) {
	me.ElemBase.init(parent, me, "complexContent", &me.HasAttrId, &me.HasAttrMixed)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemExtensionComplexContent.initChildren(me)
	me.HasElemRestrictionComplexContent.initChildren(me)
}

func (me *ComplexType) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "complexType", &me.HasAttrAbstract, &me.HasAttrBlock, &me.HasAttrFinal, &me.HasAttrId, &me.HasAttrMixed, &me.HasAttrName)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemAll.initChildren(me)
	me.HasElemChoice.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemGroup.initChildren(me)
	me.HasElemSequence.initChildren(me)
	me.HasElemComplexContent.initChildren(me)
	me.HasElemSimpleContent.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
}

func (me *Documentation) initElement(parent element) {
	me.ElemBase.init(parent, me, "documentation", &me.HasAttrLang, &me.HasAttrSource)
}

func (me *Element) initElement(parent element) {
	me.ElemBase.init(parent, me, "element", &me.HasAttrAbstract, &me.HasAttrBlock, &me.HasAttrDefault, &me.HasAttrFinal, &me.HasAttrFixed, &me.HasAttrForm, &me.HasAttrId, &me.HasAttrName, &me.HasAttrNillable, &me.HasAttrRef, &me.HasAttrType, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs, &me.HasAttrSubstitutionGroup)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemUnique.initChildren(me)
	me.HasElemsKey.initChildren(me)
	me.HasElemComplexType.initChildren(me)
	me.HasElemKeyRef.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *ExtensionComplexContent) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "extension", &me.HasAttrBase, &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemAll.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsChoice.initChildren(me)
	me.HasElemsGroup.initChildren(me)
	me.HasElemsSequence.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
}

func (me *ExtensionSimpleContent) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "extension", &me.HasAttrBase, &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
}

func (me *Field) initElement(parent element) {
	me.ElemBase.init(parent, me, "field", &me.HasAttrId, &me.HasAttrXpath)
	me.HasElemAnnotation.initChildren(me)
}

func (me *Group) initElement(parent element) {
	me.ElemBase.init(parent, me, "group", &me.HasAttrId, &me.HasAttrName, &me.HasAttrRef, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemAll.initChildren(me)
	me.HasElemChoice.initChildren(me)
	me.HasElemSequence.initChildren(me)
}

func (me *Import) initElement(parent element) {
	me.ElemBase.init(parent, me, "import", &me.HasAttrId, &me.HasAttrNamespace, &me.HasAttrSchemaLocation)
	me.HasElemAnnotation.initChildren(me)
}

func (me *Key) initElement(parent element) {
	me.ElemBase.init(parent, me, "key", &me.HasAttrId, &me.HasAttrName)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemField.initChildren(me)
	me.HasElemSelector.initChildren(me)
}

func (me *KeyRef) initElement(parent element) {
	me.ElemBase.init(parent, me, "keyref", &me.HasAttrId, &me.HasAttrName, &me.HasAttrRefer)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemField.initChildren(me)
	me.HasElemSelector.initChildren(me)
}

func (me *List) initElement(parent element) {
	me.ElemBase.init(parent, me, "list", &me.HasAttrId, &me.HasAttrItemType)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *Notation) initElement(parent element) {
	me.ElemBase.init(parent, me, "notation", &me.HasAttrId, &me.HasAttrName, &me.HasAttrPublic, &me.HasAttrSystem)
	me.HasElemAnnotation.initChildren(me)
}

func (me *Redefine) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)

	me.ElemBase.init(parent, me, "redefine", &me.HasAttrId, &me.HasAttrSchemaLocation)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsGroup.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
	me.HasElemsComplexType.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *RestrictionComplexContent) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "restriction", &me.HasAttrBase, &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemAll.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsChoice.initChildren(me)
	me.HasElemsSequence.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
}

func (me *RestrictionSimpleContent) initElement(parent element) {
	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)
	me.ElemBase.init(parent, me, "restriction", &me.HasAttrBase, &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemLength.initChildren(me)
	me.HasElemPattern.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsEnumeration.initChildren(me)
	me.HasElemFractionDigits.initChildren(me)
	me.HasElemMaxExclusive.initChildren(me)
	me.HasElemMaxInclusive.initChildren(me)
	me.HasElemMaxLength.initChildren(me)
	me.HasElemMinExclusive.initChildren(me)
	me.HasElemMinInclusive.initChildren(me)
	me.HasElemMinLength.initChildren(me)
	me.HasElemTotalDigits.initChildren(me)
	me.HasElemWhiteSpace.initChildren(me)
	me.HasElemsAnyAttribute.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *RestrictionSimpleEnumeration) initElement(parent element) {
	me.ElemBase.init(parent, me, "enumeration", &me.HasAttrValue)
}

func (me *RestrictionSimpleFractionDigits) initElement(parent element) {
	me.ElemBase.init(parent, me, "fractionDigits", &me.HasAttrValue)
}

func (me *RestrictionSimpleLength) initElement(parent element) {
	me.ElemBase.init(parent, me, "length", &me.HasAttrValue)
}

func (me *RestrictionSimpleMaxExclusive) initElement(parent element) {
	me.ElemBase.init(parent, me, "maxExclusive", &me.HasAttrValue)
}

func (me *RestrictionSimpleMaxInclusive) initElement(parent element) {
	me.ElemBase.init(parent, me, "maxInclusive", &me.HasAttrValue)
}

func (me *RestrictionSimpleMaxLength) initElement(parent element) {
	me.ElemBase.init(parent, me, "maxLength", &me.HasAttrValue)
}

func (me *RestrictionSimpleMinExclusive) initElement(parent element) {
	me.ElemBase.init(parent, me, "minExclusive", &me.HasAttrValue)
}

func (me *RestrictionSimpleMinInclusive) initElement(parent element) {
	me.ElemBase.init(parent, me, "minInclusive", &me.HasAttrValue)
}

func (me *RestrictionSimpleMinLength) initElement(parent element) {
	me.ElemBase.init(parent, me, "minLength", &me.HasAttrValue)
}

func (me *RestrictionSimplePattern) initElement(parent element) {
	me.ElemBase.init(parent, me, "pattern", &me.HasAttrValue)
}

func (me *RestrictionSimpleTotalDigits) initElement(parent element) {
	me.ElemBase.init(parent, me, "totalDigits", &me.HasAttrValue)
}

func (me *RestrictionSimpleType) initElement(parent element) {
	me.ElemBase.init(parent, me, "restriction", &me.HasAttrBase, &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemLength.initChildren(me)
	me.HasElemPattern.initChildren(me)
	me.HasElemsEnumeration.initChildren(me)
	me.HasElemFractionDigits.initChildren(me)
	me.HasElemMaxExclusive.initChildren(me)
	me.HasElemMaxInclusive.initChildren(me)
	me.HasElemMaxLength.initChildren(me)
	me.HasElemMinExclusive.initChildren(me)
	me.HasElemMinInclusive.initChildren(me)
	me.HasElemMinLength.initChildren(me)
	me.HasElemTotalDigits.initChildren(me)
	me.HasElemWhiteSpace.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *RestrictionSimpleWhiteSpace) initElement(parent element) {
	me.ElemBase.init(parent, me, "whiteSpace", &me.HasAttrValue)
}

func (me *Schema) initElement(parent element) {

	//fmt.Printf("initElement %v\n", me)
	//fmt.Printf("initElement %v\n", parent)

	me.ElemBase.init(parent, me, "schema", &me.HasAttrId, &me.HasAttrLang, &me.HasAttrVersion, &me.HasAttrBlockDefault, &me.HasAttrFinalDefault, &me.HasAttrSchemaLocation, &me.HasAttrTargetNamespace, &me.HasAttrAttributeFormDefault, &me.HasAttrElementFormDefault)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsAttribute.initChildren(me)
	me.HasElemsElement.initChildren(me)
	me.HasElemsGroup.initChildren(me)
	me.HasElemsImport.initChildren(me)
	me.HasElemsNotation.initChildren(me)
	me.HasElemsRedefine.initChildren(me)
	me.HasElemsAttributeGroup.initChildren(me)
	me.HasElemsComplexType.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *Selector) initElement(parent element) {
	me.ElemBase.init(parent, me, "selector", &me.HasAttrId, &me.HasAttrXpath)
	me.HasElemAnnotation.initChildren(me)
}

func (me *Sequence) initElement(parent element) {
	me.ElemBase.init(parent, me, "sequence", &me.HasAttrId, &me.HasAttrMaxOccurs, &me.HasAttrMinOccurs)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsAny.initChildren(me)
	me.HasElemsChoice.initChildren(me)
	me.HasElemsElement.initChildren(me)
	me.HasElemsGroup.initChildren(me)
	me.HasElemsSequence.initChildren(me)
}

func (me *SimpleContent) initElement(parent element) {
	me.ElemBase.init(parent, me, "simpleContent", &me.HasAttrId)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemExtensionSimpleContent.initChildren(me)
	me.HasElemRestrictionSimpleContent.initChildren(me)
}

func (me *SimpleType) initElement(parent element) {
	me.ElemBase.init(parent, me, "simpleType", &me.HasAttrFinal, &me.HasAttrId, &me.HasAttrName)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemRestrictionSimpleType.initChildren(me)
	me.HasElemList.initChildren(me)
	me.HasElemUnion.initChildren(me)
}

func (me *Union) initElement(parent element) {
	me.ElemBase.init(parent, me, "union", &me.HasAttrId, &me.HasAttrMemberTypes)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemsSimpleType.initChildren(me)
}

func (me *Unique) initElement(parent element) {
	me.ElemBase.init(parent, me, "unique", &me.HasAttrId, &me.HasAttrName)
	me.HasElemAnnotation.initChildren(me)
	me.HasElemField.initChildren(me)
	me.HasElemSelector.initChildren(me)
}
