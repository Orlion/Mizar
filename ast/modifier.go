package ast

type MemberModifierType int8

const (
	ModifierPublic MemberModifierType = iota + 1
	ModifierProtected
	ModifierPrivate
	ModifierAbstract
)

type MemberModifier struct {
	Type MemberModifierType
}

func (memberMod *MemberModifier) accept(vistor ASTVistor) {

}
