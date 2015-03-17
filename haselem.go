package xsd

type HasCdata struct {
	CDATA string `xml:",chardata"`
}

type HasElemAll struct {
	All *All `xml:"all"`
}

type HasElemAnnotation struct {
	Annotation *Annotation `xml:"annotation"`
}

type HasElemsAny struct {
	Anys []*Any `xml:"any"`
}

type HasElemsAnyAttribute struct {
	AnyAttributes []*AnyAttribute `xml:"anyAttribute"`
}

type HasElemsAppInfo struct {
	AppInfos []*AppInfo `xml:"appinfo"`
}

type HasElemsAttribute struct {
	Attributes []*Attribute `xml:"attribute"`
}

type HasElemsAttributeGroup struct {
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}

type HasElemChoice struct {
	Choice *Choice `xml:"choice"`
}

type HasElemsChoice struct {
	Choices []*Choice `xml:"choice"`
}

type HasElemComplexContent struct {
	ComplexContent *ComplexContent `xml:"complexContent"`
}

type HasElemComplexType struct {
	ComplexType *ComplexType `xml:"complexType"`
}

type HasElemsComplexType struct {
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type HasElemsDocumentation struct {
	Documentations []*Documentation `xml:"documentation"`
}

type HasElemsElement struct {
	Elements []*Element `xml:"element"`
}

type HasElemsEnumeration struct {
	Enumerations []*RestrictionSimpleEnumeration `xml:"enumeration"`
}

type HasElemExtensionComplexContent struct {
	ExtensionComplexContent *ExtensionComplexContent `xml:"extension"`
}

type HasElemExtensionSimpleContent struct {
	ExtensionSimpleContent *ExtensionSimpleContent `xml:"extension"`
}

type HasElemField struct {
	Field *Field `xml:"field"`
}

type HasElemFractionDigits struct {
	FractionDigits *RestrictionSimpleFractionDigits `xml:"fractionDigits"`
}

type HasElemGroup struct {
	Group *Group `xml:"group"`
}

type HasElemsGroup struct {
	Groups []*Group `xml:"group"`
}

type HasElemsImport struct {
	Imports []*Import `xml:"import"`
}

type HasElemsInclude struct {
	Includes []*Include `xml:"include"`
}

type HasElemsKey struct {
	Keys []*Key `xml:"key"`
}

type HasElemKeyRef struct {
	KeyRef *KeyRef `xml:"keyref"`
}

type HasElemLength struct {
	Length *RestrictionSimpleLength `xml:"length"`
}

type HasElemList struct {
	List *List `xml:"list"`
}

type HasElemMaxExclusive struct {
	MaxExclusive *RestrictionSimpleMaxExclusive `xml:"maxExclusive"`
}

type HasElemMaxInclusive struct {
	MaxInclusive *RestrictionSimpleMaxInclusive `xml:"maxInclusive"`
}

type HasElemMaxLength struct {
	MaxLength *RestrictionSimpleMaxLength `xml:"maxLength"`
}

type HasElemMinExclusive struct {
	MinExclusive *RestrictionSimpleMinExclusive `xml:"minExclusive"`
}

type HasElemMinInclusive struct {
	MinInclusive *RestrictionSimpleMinInclusive `xml:"minInclusive"`
}

type HasElemMinLength struct {
	MinLength *RestrictionSimpleMinLength `xml:"minLength"`
}

type HasElemsNotation struct {
	Notations []*Notation `xml:"notation"`
}

type HasElemPattern struct {
	Pattern *RestrictionSimplePattern `xml:"pattern"`
}

type HasElemsRedefine struct {
	Redefines []*Redefine `xml:"redefine"`
}

type HasElemRestrictionComplexContent struct {
	RestrictionComplexContent *RestrictionComplexContent `xml:"restriction"`
}

type HasElemRestrictionSimpleContent struct {
	RestrictionSimpleContent *RestrictionSimpleContent `xml:"restriction"`
}

type HasElemRestrictionSimpleType struct {
	RestrictionSimpleType *RestrictionSimpleType `xml:"restriction"`
}

type HasElemSelector struct {
	Selector *Selector `xml:"selector"`
}

type HasElemSequence struct {
	Sequence *Sequence `xml:"sequence"`
}

type HasElemsSequence struct {
	Sequences []*Sequence `xml:"sequence"`
}

type HasElemSimpleContent struct {
	SimpleContent *SimpleContent `xml:"simpleContent"`
}

type HasElemsSimpleType struct {
	SimpleTypes []*SimpleType `xml:"simpleType"`
}

type HasElemTotalDigits struct {
	TotalDigits *RestrictionSimpleTotalDigits `xml:"totalDigits"`
}

type HasElemUnion struct {
	Union *Union `xml:"union"`
}

type HasElemUnique struct {
	Unique *Unique `xml:"unique"`
}

type HasElemWhiteSpace struct {
	WhiteSpace *RestrictionSimpleWhiteSpace `xml:"whiteSpace"`
}
