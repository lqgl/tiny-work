package player

// HandlerRegister 注册事件
func (p *Player) HandlerRegister() {
	p.handlers["addFriend"] = p.AddFriend
	p.handlers["delFriend"] = p.DelFriend
	p.handlers["rsChatMsg"] = p.ResolveChatMsg
}
