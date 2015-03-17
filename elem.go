package xsd

import (
	"types"
)

type element interface {
	base() *ElemBase
	init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake)
	Parent() element
}

type ElemBase struct {
	atts         []beforeAfterMake
	parent, self element // self is the struct that embeds ElemBase, rather than the ElemBase pseudo-field
	xsdName      xsdt.NCName
	HasNameAttr  bool
}

func (me *ElemBase) afterMakePkg(bag *PkgBag) {
	if !me.HasNameAttr {
		bag.Stacks.Name.Pop()
	}
	for _, a := range me.atts {
		a.afterMakePkg(bag)
	}
}

func (me *ElemBase) beforeMakePkg(bag *PkgBag) {
	if !me.HasNameAttr {
		bag.Stacks.Name.Push(me.xsdName)
	}
	for _, a := range me.atts {
		a.beforeMakePkg(bag)
	}
}

func (me *ElemBase) base() *ElemBase { return me }

func (me *ElemBase) init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake) {
	me.parent, me.self, me.xsdName, me.atts = parent, self, xsdName, atts
	for _, a := range atts {
		if _, me.HasNameAttr = a.(*HasAttrName); me.HasNameAttr {
			break
		}
	}
}

func (me *ElemBase) longSafeName(bag *PkgBag) (ln string) {
	var els = []element{}
	for el := me.self; (el != nil) && (el != bag.Schema); el = el.Parent() {
		els = append(els, el)
	}
	for i := len(els) - 1; i >= 0; i-- {
		ln += bag.safeName(els[i].base().selfName().String())
	}
	return
}

func (me *ElemBase) selfName() xsdt.NCName {
	if me.HasNameAttr {
		for _, at := range me.atts {
			if an, ok := at.(*HasAttrName); ok {
				return an.Name
			}
		}
	}
	return me.xsdName
}

func (me *ElemBase) Parent() element { return me.parent }

type All struct {
	ElemBase
	//	XMLName xml.Name `xml:"all"`
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasElemAnnotation
	HasElemsElement `xml:"element"`
}

type Annotation struct {
	ElemBase
	//	XMLName xml.Name `xml:"annotation"`
	HasElemsAppInfo
	HasElemsDocumentation
}

type Any struct {
	ElemBase
	//	XMLName xml.Name `xml:"any"`
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasAttrNamespace
	HasAttrProcessContents
	HasElemAnnotation
}

type AnyAttribute struct {
	ElemBase
	//	XMLName xml.Name `xml:"anyAttribute"`
	HasAttrId
	HasAttrNamespace
	HasAttrProcessContents
	HasElemAnnotation
}

type AppInfo struct {
	ElemBase
	//	XMLName xml.Name `xml:"appinfo"`
	HasAttrSource
	HasCdata
}

type Attribute struct {
	ElemBase
	//	XMLName xml.Name `xml:"attribute"`
	HasAttrDefault
	HasAttrFixed
	HasAttrForm
	HasAttrId
	HasAttrName
	HasAttrRef
	HasAttrType
	HasAttrUse
	HasElemAnnotation
	HasElemsSimpleType
}

type AttributeGroup struct {
	ElemBase
	//	XMLName xml.Name `xml:"attributeGroup"`
	HasAttrId
	HasAttrName
	HasAttrRef
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
}

type Choice struct {
	ElemBase
	//	XMLName xml.Name `xml:"choice"`
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasElemAnnotation
	HasElemsAny
	HasElemsChoice
	HasElemsElement `xml:"element"`
	HasElemsGroup
	HasElemsSequence
}

type ComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"complexContent"`
	HasAttrId
	HasAttrMixed
	HasElemAnnotation
	HasElemExtensionComplexContent
	HasElemRestrictionComplexContent
}

type ComplexType struct {
	ElemBase
	//	XMLName xml.Name `xml:"complexType"`
	HasAttrAbstract
	HasAttrBlock
	HasAttrFinal
	HasAttrId
	HasAttrMixed
	HasAttrName
	HasElemAll
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
	HasElemChoice
	HasElemComplexContent
	HasElemGroup
	HasElemSequence
	HasElemSimpleContent
}

type Documentation struct {
	ElemBase
	//	XMLName xml.Name `xml:"documentation"`
	HasAttrLang
	HasAttrSource
	HasCdata
}

type Element struct {
	ElemBase
	//	XMLName xml.Name `xml:"element"`
	HasAttrAbstract
	HasAttrBlock
	HasAttrDefault
	HasAttrFinal
	HasAttrFixed
	HasAttrForm
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasAttrName
	HasAttrNillable
	HasAttrRef
	HasAttrSubstitutionGroup
	HasAttrType
	HasElemAnnotation
	HasElemComplexType
	HasElemsKey
	HasElemKeyRef
	HasElemsSimpleType
	HasElemUnique
}

type ExtensionComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"extension"`
	HasAttrBase
	HasAttrId
	HasElemAll
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
	HasElemsChoice
	HasElemsGroup
	HasElemsSequence
}

type ExtensionSimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"extension"`
	HasAttrBase
	HasAttrId
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
}

type Field struct {
	ElemBase
	//	XMLName xml.Name `xml:"field"`
	HasAttrId
	HasAttrXpath
	HasElemAnnotation
}

type Group struct {
	ElemBase
	//	XMLName xml.Name `xml:"group"`
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasAttrName
	HasAttrRef
	HasElemAll
	HasElemAnnotation
	HasElemChoice
	HasElemSequence
}

type Include struct {
	ElemBase
	//	XMLName xml.Name `xml:"include"`
	HasAttrId
	HasAttrSchemaLocation
	HasElemAnnotation
}

type Import struct {
	ElemBase
	//	XMLName xml.Name `xml:"import"`
	HasAttrId
	HasAttrNamespace
	HasAttrSchemaLocation
	HasElemAnnotation
}

type Key struct {
	ElemBase
	//	XMLName xml.Name `xml:"key"`
	HasAttrId
	HasAttrName
	HasElemAnnotation
	HasElemField
	HasElemSelector
}

type KeyRef struct {
	ElemBase
	//	XMLName xml.Name `xml:"keyref"`
	HasAttrId
	HasAttrName
	HasAttrRefer
	HasElemAnnotation
	HasElemField
	HasElemSelector
}

type List struct {
	ElemBase
	//	XMLName xml.Name `xml:"list"`
	HasAttrId
	HasAttrItemType
	HasElemAnnotation
	HasElemsSimpleType
}

type Notation struct {
	ElemBase
	//	XMLName xml.Name `xml:"notation"`
	HasAttrId
	HasAttrName
	HasAttrPublic
	HasAttrSystem
	HasElemAnnotation
}

type Redefine struct {
	ElemBase
	//	XMLName xml.Name `xml:"redefine"`
	HasAttrId
	HasAttrSchemaLocation
	HasElemAnnotation
	HasElemsAttributeGroup
	HasElemsComplexType
	HasElemsGroup
	HasElemsSimpleType
}

type RestrictionComplexContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction"`
	HasAttrBase
	HasAttrId
	HasElemAll
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
	HasElemsChoice
	HasElemsSequence
}

type RestrictionSimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction"`
	HasAttrBase
	HasAttrId
	HasElemAnnotation
	HasElemsAnyAttribute
	HasElemsAttribute
	HasElemsAttributeGroup
	HasElemsEnumeration
	HasElemFractionDigits
	HasElemLength
	HasElemMaxExclusive
	HasElemMaxInclusive
	HasElemMaxLength
	HasElemMinExclusive
	HasElemMinInclusive
	HasElemMinLength
	HasElemPattern
	HasElemsSimpleType
	HasElemTotalDigits
	HasElemWhiteSpace
}

type RestrictionSimpleEnumeration struct {
	ElemBase
	//	XMLName xml.Name `xml:"enumeration"`
	HasAttrValue
}

type RestrictionSimpleFractionDigits struct {
	ElemBase
	//	XMLName xml.Name `xml:"fractionDigits"`
	HasAttrValue
}

type RestrictionSimpleLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"length"`
	HasAttrValue
}

type RestrictionSimpleMaxExclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxExclusive"`
	HasAttrValue
}

type RestrictionSimpleMaxInclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxInclusive"`
	HasAttrValue
}

type RestrictionSimpleMaxLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"maxLength"`
	HasAttrValue
}

type RestrictionSimpleMinExclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"minExclusive"`
	HasAttrValue
}

type RestrictionSimpleMinInclusive struct {
	ElemBase
	//	XMLName xml.Name `xml:"minInclusive"`
	HasAttrValue
}

type RestrictionSimpleMinLength struct {
	ElemBase
	//	XMLName xml.Name `xml:"minLength"`
	HasAttrValue
}

type RestrictionSimplePattern struct {
	ElemBase
	//	XMLName xml.Name `xml:"pattern"`
	HasAttrValue
}

type RestrictionSimpleTotalDigits struct {
	ElemBase
	//	XMLName xml.Name `xml:"totalDigits"`
	HasAttrValue
}

type RestrictionSimpleType struct {
	ElemBase
	//	XMLName xml.Name `xml:"restriction"`
	HasAttrBase
	HasAttrId
	HasElemAnnotation
	HasElemsEnumeration
	HasElemFractionDigits
	HasElemLength
	HasElemMaxExclusive
	HasElemMaxInclusive
	HasElemMaxLength
	HasElemMinExclusive
	HasElemMinInclusive
	HasElemMinLength
	HasElemPattern
	HasElemsSimpleType
	HasElemTotalDigits
	HasElemWhiteSpace
}

type RestrictionSimpleWhiteSpace struct {
	ElemBase
	//	XMLName xml.Name `xml:"whiteSpace"`
	HasAttrValue
}

type Selector struct {
	ElemBase
	//	XMLName xml.Name `xml:"selector"`
	HasAttrId
	HasAttrXpath
	HasElemAnnotation
}

type Sequence struct {
	ElemBase
	//	XMLName xml.Name `xml:"sequence"`
	HasAttrId
	HasAttrMaxOccurs
	HasAttrMinOccurs
	HasElemAnnotation
	HasElemsAny
	HasElemsChoice
	HasElemsElement `xml:"element"`
	HasElemsGroup
	HasElemsSequence
}

type SimpleContent struct {
	ElemBase
	//	XMLName xml.Name `xml:"simpleContent"`
	HasAttrId
	HasElemAnnotation
	HasElemExtensionSimpleContent
	HasElemRestrictionSimpleContent
}

type SimpleType struct {
	ElemBase
	//	XMLName xml.Name `xml:"simpleType"`
	HasAttrFinal
	HasAttrId
	HasAttrName
	HasElemAnnotation
	HasElemList
	HasElemRestrictionSimpleType
	HasElemUnion
}

type Union struct {
	ElemBase
	//	XMLName xml.Name `xml:"union"`
	HasAttrId
	HasAttrMemberTypes
	HasElemAnnotation
	HasElemsSimpleType
}

type Unique struct {
	ElemBase
	//	XMLName xml.Name `xml:"unique"`
	HasAttrId
	HasAttrName
	HasElemAnnotation
	HasElemField
	HasElemSelector
}

func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence) {
	var tmpChoices []*Choice
	var tmpSeqs []*Sequence
	for _, ch := range choices {
		if ch != nil {
			allChoices = append(allChoices, ch)
			tmpChoices, tmpSeqs = Flattened(ch.Choices, ch.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	for _, seq := range seqs {
		if seq != nil {
			allSeqs = append(allSeqs, seq)
			tmpChoices, tmpSeqs = Flattened(seq.Choices, seq.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
	}
	return
}
