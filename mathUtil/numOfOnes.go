package mathUtil

var (
	numOfOnesArray = [256]int8{}
)

func init() {
	numOfOnesArray[0] = 0
	numOfOnesArray[1] = 1
	numOfOnesArray[2] = 1
	numOfOnesArray[3] = 2
	numOfOnesArray[4] = 1
	numOfOnesArray[5] = 2
	numOfOnesArray[6] = 2
	numOfOnesArray[7] = 3
	numOfOnesArray[8] = 1
	numOfOnesArray[9] = 2
	numOfOnesArray[10] = 2
	numOfOnesArray[11] = 3
	numOfOnesArray[12] = 2
	numOfOnesArray[13] = 3
	numOfOnesArray[14] = 3
	numOfOnesArray[15] = 4
	numOfOnesArray[16] = 1
	numOfOnesArray[17] = 2
	numOfOnesArray[18] = 2
	numOfOnesArray[19] = 3
	numOfOnesArray[20] = 2
	numOfOnesArray[21] = 3
	numOfOnesArray[22] = 3
	numOfOnesArray[23] = 4
	numOfOnesArray[24] = 2
	numOfOnesArray[25] = 3
	numOfOnesArray[26] = 3
	numOfOnesArray[27] = 4
	numOfOnesArray[28] = 3
	numOfOnesArray[29] = 4
	numOfOnesArray[30] = 4
	numOfOnesArray[31] = 5
	numOfOnesArray[32] = 1
	numOfOnesArray[33] = 2
	numOfOnesArray[34] = 2
	numOfOnesArray[35] = 3
	numOfOnesArray[36] = 2
	numOfOnesArray[37] = 3
	numOfOnesArray[38] = 3
	numOfOnesArray[39] = 4
	numOfOnesArray[40] = 2
	numOfOnesArray[41] = 3
	numOfOnesArray[42] = 3
	numOfOnesArray[43] = 4
	numOfOnesArray[44] = 3
	numOfOnesArray[45] = 4
	numOfOnesArray[46] = 4
	numOfOnesArray[47] = 5
	numOfOnesArray[48] = 2
	numOfOnesArray[49] = 3
	numOfOnesArray[50] = 3
	numOfOnesArray[51] = 4
	numOfOnesArray[52] = 3
	numOfOnesArray[53] = 4
	numOfOnesArray[54] = 4
	numOfOnesArray[55] = 5
	numOfOnesArray[56] = 3
	numOfOnesArray[57] = 4
	numOfOnesArray[58] = 4
	numOfOnesArray[59] = 5
	numOfOnesArray[60] = 4
	numOfOnesArray[61] = 5
	numOfOnesArray[62] = 5
	numOfOnesArray[63] = 6
	numOfOnesArray[64] = 1
	numOfOnesArray[65] = 2
	numOfOnesArray[66] = 2
	numOfOnesArray[67] = 3
	numOfOnesArray[68] = 2
	numOfOnesArray[69] = 3
	numOfOnesArray[70] = 3
	numOfOnesArray[71] = 4
	numOfOnesArray[72] = 2
	numOfOnesArray[73] = 3
	numOfOnesArray[74] = 3
	numOfOnesArray[75] = 4
	numOfOnesArray[76] = 3
	numOfOnesArray[77] = 4
	numOfOnesArray[78] = 4
	numOfOnesArray[79] = 5
	numOfOnesArray[80] = 2
	numOfOnesArray[81] = 3
	numOfOnesArray[82] = 3
	numOfOnesArray[83] = 4
	numOfOnesArray[84] = 3
	numOfOnesArray[85] = 4
	numOfOnesArray[86] = 4
	numOfOnesArray[87] = 5
	numOfOnesArray[88] = 3
	numOfOnesArray[89] = 4
	numOfOnesArray[90] = 4
	numOfOnesArray[91] = 5
	numOfOnesArray[92] = 4
	numOfOnesArray[93] = 5
	numOfOnesArray[94] = 5
	numOfOnesArray[95] = 6
	numOfOnesArray[96] = 2
	numOfOnesArray[97] = 3
	numOfOnesArray[98] = 3
	numOfOnesArray[99] = 4
	numOfOnesArray[100] = 3
	numOfOnesArray[101] = 4
	numOfOnesArray[102] = 4
	numOfOnesArray[103] = 5
	numOfOnesArray[104] = 3
	numOfOnesArray[105] = 4
	numOfOnesArray[106] = 4
	numOfOnesArray[107] = 5
	numOfOnesArray[108] = 4
	numOfOnesArray[109] = 5
	numOfOnesArray[110] = 5
	numOfOnesArray[111] = 6
	numOfOnesArray[112] = 3
	numOfOnesArray[113] = 4
	numOfOnesArray[114] = 4
	numOfOnesArray[115] = 5
	numOfOnesArray[116] = 4
	numOfOnesArray[117] = 5
	numOfOnesArray[118] = 5
	numOfOnesArray[119] = 6
	numOfOnesArray[120] = 4
	numOfOnesArray[121] = 5
	numOfOnesArray[122] = 5
	numOfOnesArray[123] = 6
	numOfOnesArray[124] = 5
	numOfOnesArray[125] = 6
	numOfOnesArray[126] = 6
	numOfOnesArray[127] = 7
	numOfOnesArray[128] = 1
	numOfOnesArray[129] = 2
	numOfOnesArray[130] = 2
	numOfOnesArray[131] = 3
	numOfOnesArray[132] = 2
	numOfOnesArray[133] = 3
	numOfOnesArray[134] = 3
	numOfOnesArray[135] = 4
	numOfOnesArray[136] = 2
	numOfOnesArray[137] = 3
	numOfOnesArray[138] = 3
	numOfOnesArray[139] = 4
	numOfOnesArray[140] = 3
	numOfOnesArray[141] = 4
	numOfOnesArray[142] = 4
	numOfOnesArray[143] = 5
	numOfOnesArray[144] = 2
	numOfOnesArray[145] = 3
	numOfOnesArray[146] = 3
	numOfOnesArray[147] = 4
	numOfOnesArray[148] = 3
	numOfOnesArray[149] = 4
	numOfOnesArray[150] = 4
	numOfOnesArray[151] = 5
	numOfOnesArray[152] = 3
	numOfOnesArray[153] = 4
	numOfOnesArray[154] = 4
	numOfOnesArray[155] = 5
	numOfOnesArray[156] = 4
	numOfOnesArray[157] = 5
	numOfOnesArray[158] = 5
	numOfOnesArray[159] = 6
	numOfOnesArray[160] = 2
	numOfOnesArray[161] = 3
	numOfOnesArray[162] = 3
	numOfOnesArray[163] = 4
	numOfOnesArray[164] = 3
	numOfOnesArray[165] = 4
	numOfOnesArray[166] = 4
	numOfOnesArray[167] = 5
	numOfOnesArray[168] = 3
	numOfOnesArray[169] = 4
	numOfOnesArray[170] = 4
	numOfOnesArray[171] = 5
	numOfOnesArray[172] = 4
	numOfOnesArray[173] = 5
	numOfOnesArray[174] = 5
	numOfOnesArray[175] = 6
	numOfOnesArray[176] = 3
	numOfOnesArray[177] = 4
	numOfOnesArray[178] = 4
	numOfOnesArray[179] = 5
	numOfOnesArray[180] = 4
	numOfOnesArray[181] = 5
	numOfOnesArray[182] = 5
	numOfOnesArray[183] = 6
	numOfOnesArray[184] = 4
	numOfOnesArray[185] = 5
	numOfOnesArray[186] = 5
	numOfOnesArray[187] = 6
	numOfOnesArray[188] = 5
	numOfOnesArray[189] = 6
	numOfOnesArray[190] = 6
	numOfOnesArray[191] = 7
	numOfOnesArray[192] = 2
	numOfOnesArray[193] = 3
	numOfOnesArray[194] = 3
	numOfOnesArray[195] = 4
	numOfOnesArray[196] = 3
	numOfOnesArray[197] = 4
	numOfOnesArray[198] = 4
	numOfOnesArray[199] = 5
	numOfOnesArray[200] = 3
	numOfOnesArray[201] = 4
	numOfOnesArray[202] = 4
	numOfOnesArray[203] = 5
	numOfOnesArray[204] = 4
	numOfOnesArray[205] = 5
	numOfOnesArray[206] = 5
	numOfOnesArray[207] = 6
	numOfOnesArray[208] = 3
	numOfOnesArray[209] = 4
	numOfOnesArray[210] = 4
	numOfOnesArray[211] = 5
	numOfOnesArray[212] = 4
	numOfOnesArray[213] = 5
	numOfOnesArray[214] = 5
	numOfOnesArray[215] = 6
	numOfOnesArray[216] = 4
	numOfOnesArray[217] = 5
	numOfOnesArray[218] = 5
	numOfOnesArray[219] = 6
	numOfOnesArray[220] = 5
	numOfOnesArray[221] = 6
	numOfOnesArray[222] = 6
	numOfOnesArray[223] = 7
	numOfOnesArray[224] = 3
	numOfOnesArray[225] = 4
	numOfOnesArray[226] = 4
	numOfOnesArray[227] = 5
	numOfOnesArray[228] = 4
	numOfOnesArray[229] = 5
	numOfOnesArray[230] = 5
	numOfOnesArray[231] = 6
	numOfOnesArray[232] = 4
	numOfOnesArray[233] = 5
	numOfOnesArray[234] = 5
	numOfOnesArray[235] = 6
	numOfOnesArray[236] = 5
	numOfOnesArray[237] = 6
	numOfOnesArray[238] = 6
	numOfOnesArray[239] = 7
	numOfOnesArray[240] = 4
	numOfOnesArray[241] = 5
	numOfOnesArray[242] = 5
	numOfOnesArray[243] = 6
	numOfOnesArray[244] = 5
	numOfOnesArray[245] = 6
	numOfOnesArray[246] = 6
	numOfOnesArray[247] = 7
	numOfOnesArray[248] = 5
	numOfOnesArray[249] = 6
	numOfOnesArray[250] = 6
	numOfOnesArray[251] = 7
	numOfOnesArray[252] = 6
	numOfOnesArray[253] = 7
	numOfOnesArray[254] = 7
	numOfOnesArray[255] = 8
}

/*
goos: linux
goarch: amd64
pkg: github.com/Jordanzuo/goutil/mathUtil
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkGetNumOfOnes1_uint32-12        171454423                6.981 ns/op
BenchmarkGetNumOfOnes2_uint32-12        1000000000               0.6923 ns/op
*/

func GetNumOfOnes1_uint32(value uint32) int {
	if value == 0 {
		return 0
	}

	num := 0
	for {
		num++
		value = value & (value - 1)
		if value == 0 {
			break
		}
	}

	return num
}

// According to benchmark test result, BenchmarkGetNumOfOnes2_uint32 is ten times faster than BenchmarkGetNumOfOnes1_uint32.
func GetNumOfOnes2_uint32(value uint32) int {
	value1 := (value & 0xFF000000) >> 24
	value2 := (value & 0x00FF0000) >> 16
	value3 := (value & 0x0000FF00) >> 8
	value4 := (value & 0x000000FF)

	num1 := int(numOfOnesArray[value1])
	num2 := int(numOfOnesArray[value2])
	num3 := int(numOfOnesArray[value3])
	num4 := int(numOfOnesArray[value4])

	return num1 + num2 + num3 + num4
}
