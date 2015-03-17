package xsd

import (
	//	xsdt "github.com/metaleap/go-xsd/types"
"types"
)

type HasAttrAbstract struct {
	Abstract bool `xml:"abstract,attr"`
}

type HasAttrAttributeFormDefault struct {
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`
}

type HasAttrBase struct {
	Base xsdt.Qname `xml:"base,attr"`
}

type HasAttrBlock struct {
	Block string `xml:"block,attr"`
}

type HasAttrBlockDefault struct {
	BlockDefault string `xml:"blockDefault,attr"`
}

type HasAttrDefault struct {
	Default string `xml:"default,attr"`
}

type HasAttrFinal struct {
	Final string `xml:"final,attr"`
}

type HasAttrFinalDefault struct {
	FinalDefault string `xml:"finalDefault,attr"`
}

type HasAttrFixed struct {
	Fixed string `xml:"fixed,attr"`
}

type HasAttrForm struct {
	Form string `xml:"form,attr"`
}

type HasAttrElementFormDefault struct {
	ElementFormDefault string `xml:"elementFormDefault,attr"`
}

type HasAttrId struct {
	Id xsdt.Id `xml:"id,attr"`
}

type HasAttrItemType struct {
	ItemType xsdt.Qname `xml:"itemType,attr"`
}

type HasAttrLang struct {
	Lang xsdt.Language `xml:"lang,attr"`
}

type HasAttrMaxOccurs struct {
	MaxOccurs string `xml:"maxOccurs,attr"`
}

func (me *HasAttrMaxOccurs) Value() (l xsdt.Long) {
	if len(me.MaxOccurs) == 0 {
		l = 1
	} else if me.MaxOccurs == "unbounded" {
		l = -1
	} else {
		l.Set(me.MaxOccurs)
	}
	return
}

type HasAttrMemberTypes struct {
	MemberTypes string `xml:"memberTypes,attr"`
}

type HasAttrMinOccurs struct {
	MinOccurs uint64 `xml:"minOccurs,attr"`
}

type HasAttrMixed struct {
	Mixed bool `xml:"mixed,attr"`
}

type HasAttrName struct {
	Name xsdt.NCName `xml:"name,attr"`
}

type HasAttrNamespace struct {
	Namespace string `xml:"namespace,attr"`
}

type HasAttrNillable struct {
	Nillable bool `xml:"nillable,attr"`
}

type HasAttrProcessContents struct {
	ProcessContents string `xml:"processContents,attr"`
}

type HasAttrPublic struct {
	Public string `xml:"public,attr"`
}

type HasAttrRef struct {
	Ref xsdt.Qname `xml:"ref,attr"`
}

type HasAttrRefer struct {
	Refer xsdt.Qname `xml:"refer,attr"`
}

type HasAttrSchemaLocation struct {
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr"`
}

type HasAttrSource struct {
	Source xsdt.AnyURI `xml:"source,attr"`
}

type HasAttrSubstitutionGroup struct {
	SubstitutionGroup xsdt.Qname `xml:"substitutionGroup,attr"`
}

type HasAttrSystem struct {
	System xsdt.AnyURI `xml:"system,attr"`
}

type HasAttrTargetNamespace struct {
	TargetNamespace xsdt.AnyURI `xml:"targetNamespace,attr"`
}

type HasAttrType struct {
	Type xsdt.Qname `xml:"type,attr"`
}

type HasAttrUse struct {
	Use string `xml:"use,attr"`
}

type HasAttrValue struct {
	Value string `xml:"value,attr"`
}

type HasAttrVersion struct {
	Version xsdt.Token `xml:"version,attr"`
}

type HasAttrXpath struct {
	Xpath string `xml:"xpath,attr"`
}
