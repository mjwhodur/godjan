package godjan

type FieldType int

const (
	SerialPrimaryKey FieldType = iota
	CharField
	TextField
	BooleanField
	NullBooleanField
	ForeignKey
	ReverseRelation
	OneToOneField
	InetField
	JSONField
	JSONBField
	FloatField
)

type PostgresType int

const (
	Bigint PostgresType = iota
	Bigserial
	Bit
	BitVarying
	Boolean
	Box
	Bytea
	Character
	CharacterVarying
	CIDR
	Circle
	Date
	DoublePrecision
	Inet
	Interval
	JSON
	JSONB
	Line
	Lseg
	MACAddr
	MACAddr8
	Money
	Numeric
	Path
	PG_lsn
	PG_Snapshot
	Point
	Polygon
	Real
	SmallInt
	SmallSerial
	Serial
	Text
	Time
	TimeWithTimezone
	TimeStamp
	TimeStampWithTimezone
	TSQuery
	TSVector
	TXID_Snapshot
	UUID
	XML
)
