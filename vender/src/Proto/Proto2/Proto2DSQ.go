package Proto2

//  G_GameDSQ_Proto == 10    斗兽棋
const (
	INITDSQ               = iota //  INITDSQ == 0
	GW2DSQ_InitGameProto2        //  GW2DSQ_InitGameProto2   == 1  初始化协议
	DSQ2GW_InitGameProto2        //  GW2DSQ_InitGameProto2   == 1
)

// 斗兽棋的棋子类型
const (
	DSQINIT_QZ = iota // DSQ_QZ == 0
	Elephant          // elephant == 1  大象
	Lion              // lion == 2 		狮子
	Tiger             // tiger == 3 	老虎
	Leopard           // leopard == 4 	豹子
	Wolf              // wolf == 5 		狼
	Dog               // dog == 6 		狗
	Cat               // cat == 7 		猫
	Mouse             // mouse == 8 	老鼠
)

// 棋子的行动的方向
const (
	FANGXIANGINIT = iota // FANGXIANGINIT == 0
	UP                   // UP 		== 1
	DOWN                 // DOWN 	== 2
	LEFT                 // LEFT 	== 3
	RIGHT                // RIGHT 	== 4
)

//------------------------------------------------------------------------------
// GW2DSQ_InitGameProto2
type GW2DSQ_InitGame struct {
	Protocol  int
	Protocol2 int
	OpenID    string
	RoomID    string
}

// DSQ2GW_InitGameProto2
type DSQ2GW_InitGame struct {
	Protocol  int
	Protocol2 int
	Isucc     bool       // 是否初始化成功
	InitData  [4][4]*int // 斗兽棋的棋盘的数据
}

//------------------------------------------------------------------------------
