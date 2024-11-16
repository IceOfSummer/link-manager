package localizer

// command.add
const (
	CommandAddShort       = "command.add.short"
	CommandAddLong        = "command.add.long"
	CommandAddLinkUse     = "command.add.link_use"
	CommandAddLinkShort   = "command.add.link_short"
	CommandAddLinkLong    = "command.add.link_long"
	CommandAddLinkSuccess = "command.add.link_add_success"
	CommandAddTagUse      = "command.add.tag_use"
	CommandAddTagShort    = "command.add.tag_short"
	CommandAddTagLong     = "command.add.tag_long"
	CommandAddBindUse     = "command.add.bind_use"
	CommandAddBindShort   = "command.add.bind_short"
	CommandAddBindLong    = "command.add.bind_long"
	CommandAddBindSuccess = "command.add.bind_success"
)

// command.delete
const (
	CommandDeleteShort     = "command.delete.short"
	CommandDeleteLong      = "command.delete.long"
	CommandDeleteLinkUse   = "command.delete.link_use"
	CommandDeleteLinkShort = "command.delete.link_short"
	CommandDeleteTagUse    = "command.delete.tag_use"
	CommandDeleteTagShort  = "command.delete.tag_short"
	CommandDeleteTagLong   = "command.delete.tag_long"
	CommandDeleteBindUse   = "command.delete.bind_use"
	CommandDeleteBindShort = "command.delete.bind_short"
)

// command.get
const (
	CommandGetShort      = "command.get.short"
	CommandGetLinksShort = "command.get.links_short"
	CommandGetTagUse     = "command.get.tag_use"
	CommandGetTagShort   = "command.get.tag_short"
	CommandGetBindUse    = "command.get.bind_use"
	CommandGetBindShort  = "command.get.bind_short"
	CommandGetUsing      = "command.get.using_short"
)

// command.root
const (
	CommandRootShort = "command.root.short"
	CommandRootLong  = "command.root.long"
)

// command.update
const (
	CommandUpdateShort     = "command.update.short"
	CommandUpdateLong      = "command.update.long"
	CommandUpdateLinkUse   = "command.update.link_use"
	CommandUpdateLinkShort = "command.update.link_short"
	CommandUpdateTagUse    = "command.update.tag_use"
	CommandUpdateTagShort  = "command.update.tag_short"
	CommandUpdateBindUse   = "command.update.bind_use"
	CommandUpdateBindShort = "command.update.bind_short"
)

// command.update.flag
const (
	UpdateFlagName = "command.update.flag.name"
	UpdateFlagPath = "command.update.flag.path"
	UpdateFlagTag  = "command.update.flag.tag"
)

// command.use

const (
	CommandUse        = "command.use.use"
	CommandUseShort   = "command.use.short"
	CommandUseSuccess = "command.use.success"
)

// errors
const (
	ErrorInvalidNTPair    = "error.invalid_name_mark_pair"
	ErrorNoSuchBind       = "error.no_such_bind"
	ErrorBindNotExist     = "error.bind_not_exist"
	ErrorBindAlreadyExist = "error.bind_already_exit"
)

// message
const (
	MessageSuccess               = "message.update_success"
	MessageDeleteSuccessPrefix   = "message.delete_success_prefix"
	NothingChanged               = "message.nothing_changed"
	NothingFound                 = "message.nothing_found"
	LinkDeclarationDeleteSuccess = "message.delete_link_delaration_success"
	MessageHelp                  = "message.help_message"
)
