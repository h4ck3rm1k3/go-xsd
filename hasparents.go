package xsd
//import "fmt"
func (me *HasElemAll) initChildren(p element) {
	if me.All != nil {
		me.All.initElement(p)
	}
}

func (me *HasElemAnnotation) initChildren(p element) {
	if me.Annotation != nil {
		me.Annotation.initElement(p)
	}
}

func (me *HasElemsAny) initChildren(p element) {
	for _, any := range me.Anys {
		any.initElement(p)
	}
}

func (me *HasElemsAnyAttribute) initChildren(p element) {
	for _, aa := range me.AnyAttributes {
		aa.initElement(p)
	}
}

func (me *HasElemsAppInfo) initChildren(p element) {
	for _, ai := range me.AppInfos {
		ai.initElement(p)
	}
}

func (me *HasElemsAttribute) initChildren(p element) {
	for _, ea := range me.Attributes {
		ea.initElement(p)
	}
}

func (me *HasElemsAttributeGroup) initChildren(p element) {
	//fmt.Printf("initChildren %v\n", me)
	for _, ag := range me.AttributeGroups {
		ag.initElement(p)
	}
}

func (me *HasElemChoice) initChildren(p element) {
	if me.Choice != nil {
		me.Choice.initElement(p)
	}
}

func (me *HasElemsChoice) initChildren(p element) {
	for _, ch := range me.Choices {
		ch.initElement(p)
	}
}

func (me *HasElemComplexContent) initChildren(p element) {
	if me.ComplexContent != nil {
		me.ComplexContent.initElement(p)
	}
}

func (me *HasElemComplexType) initChildren(p element) {
	if me.ComplexType != nil {
		me.ComplexType.initElement(p)
	}
}

func (me *HasElemsComplexType) initChildren(p element) {
	for _, ct := range me.ComplexTypes {
		ct.initElement(p)
	}
}

func (me *HasElemsDocumentation) initChildren(p element) {
	for _, doc := range me.Documentations {
		doc.initElement(p)
	}
}

func (me *HasElemsElement) initChildren(p element) {
	for _, el := range me.Elements {
		el.initElement(p)
	}
}

func (me *HasElemsEnumeration) initChildren(p element) {
	for _, enum := range me.Enumerations {
		enum.initElement(p)
	}
}

func (me *HasElemExtensionComplexContent) initChildren(p element) {
	if me.ExtensionComplexContent != nil {
		me.ExtensionComplexContent.initElement(p)
	}
}

func (me *HasElemExtensionSimpleContent) initChildren(p element) {
	if me.ExtensionSimpleContent != nil {
		me.ExtensionSimpleContent.initElement(p)
	}
}

func (me *HasElemField) initChildren(p element) {
	if me.Field != nil {
		me.Field.initElement(p)
	}
}

func (me *HasElemFractionDigits) initChildren(p element) {
	if me.FractionDigits != nil {
		me.FractionDigits.initElement(p)
	}
}

func (me *HasElemGroup) initChildren(p element) {
	if me.Group != nil {
		me.Group.initElement(p)
	}
}

func (me *HasElemsGroup) initChildren(p element) {
	for _, gr := range me.Groups {
		gr.initElement(p)
	}
}

func (me *HasElemsImport) initChildren(p element) {
	for _, imp := range me.Imports {
		imp.initElement(p)
	}
}

func (me *HasElemsKey) initChildren(p element) {
	for _, k := range me.Keys {
		k.initElement(p)
	}
}

func (me *HasElemKeyRef) initChildren(p element) {
	if me.KeyRef != nil {
		me.KeyRef.initElement(p)
	}
}

func (me *HasElemLength) initChildren(p element) {
	if me.Length != nil {
		me.Length.initElement(p)
	}
}

func (me *HasElemList) initChildren(p element) {
	if me.List != nil {
		me.List.initElement(p)
	}
}

func (me *HasElemMaxExclusive) initChildren(p element) {
	if me.MaxExclusive != nil {
		me.MaxExclusive.initElement(p)
	}
}

func (me *HasElemMaxInclusive) initChildren(p element) {
	if me.MaxInclusive != nil {
		me.MaxInclusive.initElement(p)
	}
}

func (me *HasElemMaxLength) initChildren(p element) {
	if me.MaxLength != nil {
		me.MaxLength.initElement(p)
	}
}

func (me *HasElemMinExclusive) initChildren(p element) {
	if me.MinExclusive != nil {
		me.MinExclusive.initElement(p)
	}
}

func (me *HasElemMinInclusive) initChildren(p element) {
	if me.MinInclusive != nil {
		me.MinInclusive.initElement(p)
	}
}

func (me *HasElemMinLength) initChildren(p element) {
	if me.MinLength != nil {
		me.MinLength.initElement(p)
	}
}

func (me *HasElemsNotation) initChildren(p element) {
	for _, not := range me.Notations {
		not.initElement(p)
	}
}

func (me *HasElemPattern) initChildren(p element) {
	if me.Pattern != nil {
		me.Pattern.initElement(p)
	}
}

func (me *HasElemsRedefine) initChildren(p element) {
	for _, rd := range me.Redefines {
		rd.initElement(p)
	}
}

func (me *HasElemRestrictionComplexContent) initChildren(p element) {
	if me.RestrictionComplexContent != nil {
		me.RestrictionComplexContent.initElement(p)
	}
}

func (me *HasElemRestrictionSimpleContent) initChildren(p element) {
	if me.RestrictionSimpleContent != nil {
		me.RestrictionSimpleContent.initElement(p)
	}
}

func (me *HasElemRestrictionSimpleType) initChildren(p element) {
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.initElement(p)
	}
}

func (me *HasElemSelector) initChildren(p element) {
	if me.Selector != nil {
		me.Selector.initElement(p)
	}
}

func (me *HasElemSequence) initChildren(p element) {
	if me.Sequence != nil {
		me.Sequence.initElement(p)
	}
}

func (me *HasElemsSequence) initChildren(p element) {
	for _, seq := range me.Sequences {
		seq.initElement(p)
	}
}

func (me *HasElemSimpleContent) initChildren(p element) {
	if me.SimpleContent != nil {
		me.SimpleContent.initElement(p)
	}
}

func (me *HasElemsSimpleType) initChildren(p element) {
	for _, st := range me.SimpleTypes {
		st.initElement(p)
	}
}

func (me *HasElemTotalDigits) initChildren(p element) {
	if me.TotalDigits != nil {
		me.TotalDigits.initElement(p)
	}
}

func (me *HasElemUnion) initChildren(p element) {
	if me.Union != nil {
		me.Union.initElement(p)
	}
}

func (me *HasElemUnique) initChildren(p element) {
	if me.Unique != nil {
		me.Unique.initElement(p)
	}
}

func (me *HasElemWhiteSpace) initChildren(p element) {
	if me.WhiteSpace != nil {
		me.WhiteSpace.initElement(p)
	}
}
