package visitor

func NewBuilder() Builder { return Builder{visitor: New()} }

type Builder struct {
	visitor Visitor
}

func (b *Builder) Build() Visitor { return b.visitor }

func (b *Builder) SetVisitCommentNode(f Func) *Builder  { b.visitor.visitCommentNode = f; return b }
func (b *Builder) SetVisitDoctypeNode(f Func) *Builder  { b.visitor.visitDoctypeNode = f; return b }
func (b *Builder) SetVisitDocumentNode(f Func) *Builder { b.visitor.visitDocumentNode = f; return b }
func (b *Builder) SetVisitElementNode(f Func) *Builder  { b.visitor.visitElementNode = f; return b }
func (b *Builder) SetVisitErrorNode(f Func) *Builder    { b.visitor.visitErrorNode = f; return b }
func (b *Builder) SetVisitTextNode(f Func) *Builder     { b.visitor.visitTextNode = f; return b }
