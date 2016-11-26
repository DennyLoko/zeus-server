package packets

var incomingTypeMap = map[string]IncomingPacket{
	"SS_NULL": &NullPacket{},
	"SS_PING": &Ping{},

	"CA_LOGIN": &LoginRequest{},

	"CH_ENTER":       &CharEnter{},
	"CH_SELECT_CHAR": &SelectChar{},

	"CZ_ENTER": &ZoneEnter{},
}

var outgoingTypeMap = map[string]OutgoingPacket{
	"SS_NULL": &NullPacket{},

	"AC_ACCEPT_LOGIN": &AcceptLoginResponse{},

	"HC_ACCEPT_ENTER":        &AcceptCharEnter{},
	"HC_REFUSE_ENTER":        &RefuseCharEnter{},
	"HC_ACCEPT2":             &AcceptCharEnter2{},
	"HC_BLOCK_CHARACTER":     &BlockCharacter{},
	"HC_SECOND_PASSWD_LOGIN": &SecondPasswordLogin{},
	"HC_NOTIFY_ZONESVR":      &NotifyZoneServer{},

	"ZC_AID":          &ZoneAid{},
	"ZC_ACCEPT_ENTER": &ZoneAcceptEnter{},
}
