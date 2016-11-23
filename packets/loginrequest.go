package packets

type LoginRequest struct {
	Version    uint32
	Username   string
	Password   string
	ClientType byte
}

func (r *LoginRequest) Parse(db *PacketDatabase, d *Definition, p *RawPacket) error {
	p.Read(&r.Version)
	p.ReadString(24, &r.Username)
	p.ReadString(24, &r.Password)
	p.Read(&r.ClientType)

	return nil
}
