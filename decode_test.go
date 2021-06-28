package script

import (
	"encoding/hex"
	"testing"
)

var scripts []string = []string{
	"5101400176018801a901ac580114011458011451515a01145458011452795279935a7951799300795b799300795b799300795b799300795b799300795b799300795b7993007959799301240120028000520120540114580000795579930079557993007955799300795579930079011f7993007901217993012a7900876361012f796100792097dfd76851bf465e8f715593b217714858bbe9570ff3bd5e33840a34e20ff0262102ba79df5f8ae7604a9830f03c7933028186aede0675a16f025dc4f8be8eec0382210ac407f0e4bd44bfc207355a778b046225a7068fc59ee7eda43ad905aadbffc800206c266b30e6a1319c66dc401e5bd6b432ba49688eecd118297041da8074ce0810201008ce7480da41702918d1ec8e6849ba32b4d65b1e40dc669c31a1e6306b266c013079013079855679aa616100790079517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e01007e81517a756157795679567956795679537956795479577995939521414136d08c5ed2bf3ba048afe6dcaebafeffffffffffffffffffffffffffffff0061517951795179517997527a75517a5179009f635179517993527a75517a685179517a75517a7561527a75517a517951795296a0630079527994527a75517a68537982775279827754527993517993013051797e527e53797e57797e527e52797e5579517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7e56797e0079517a75517a75517a75517a75517a75517a75517a75517a75517a75517a75517a75517a756100795779ac517a75517a75517a75517a75517a75517a75517a75517a75517a7561517a756169012d79aa61013079007901247f75547f77517a75618769012d79012f796079955f79937f75012f796079957f77012d79aa517987696101307961007901687f7700005279517f75007f77007901fd8763615379537f75517f77007901007e81517a7561537a75527a527a5379535479937f75537f77527a75517a67007901fe8763615379557f75517f77007901007e81517a7561537a75527a527a5379555479937f75557f77527a75517a67007901ff8763615379597f75517f77007901007e81517a7561537a75527a527a5379595479937f75597f77527a75517a67615379517f75007f77007901007e81517a7561537a75527a527a5379515479937f75517f77527a75517a6868685179517a75517a75517a75517a7561517a756100798277615179517951795179011b7994012679937f755179011b79947f77517a75517a75616101307901307954615279517951937f7551797f77007901007e81517a756151795193527a75517a00000000547953a16961537955799f635579012493567a75557a557a557a557a557a615779567951937f7556797f77007901007e81517a7561527a75517a517902fd009f635179537a75527a527a557951935379935493567a75557a557a557a557a557a67517902fd009c63615779567953937f75567951937f77007901007e81517a7561537a75527a527a557953935379935493567a75557a557a557a557a557a67517902fe009c63615779567955937f75567951937f77007901007e81517a7561537a75527a527a557955935379935493567a75557a557a557a557a557a67615779567959937f75567951937f77007901007e81517a7561537a75527a527a557959935379935493567a75557a557a557a557a557a68686853795193547a75537a537a537a6861537955799f635579012493567a75557a557a557a557a557a615779567951937f7556797f77007901007e81517a7561527a75517a517902fd009f635179537a75527a527a557951935379935493567a75557a557a557a557a557a67517902fd009c63615779567953937f75567951937f77007901007e81517a7561537a75527a527a557953935379935493567a75557a557a557a557a557a67517902fe009c63615779567955937f75567951937f77007901007e81517a7561537a75527a527a557955935379935493567a75557a557a557a557a557a67615779567959937f75567951937f77007901007e81517a7561537a75527a527a557959935379935493567a75557a557a557a557a557a68686853795193547a75537a537a537a6861537955799f635579012493567a75557a557a557a557a557a615779567951937f7556797f77007901007e81517a7561527a75517a517902fd009f635179537a75527a527a557951935379935493567a75557a557a557a557a557a67517902fd009c63615779567953937f75567951937f77007901007e81517a7561537a75527a527a557953935379935493567a75557a557a557a557a557a67517902fe009c63615779567955937f75567951937f77007901007e81517a7561537a75527a527a557955935379935493567a75557a557a557a557a557a67615779567959937f75567951937f77007901007e81517a7561537a75527a527a557959935379935493567a75557a557a557a557a557a68686853795193547a75537a537a537a68615779567951937f7556797f77007901007e81517a756156795193577a75567a567a567a567a567a567a00557a75547a547a547a547a007953a16961547951799f6356795893577a75567a567a567a567a567a567a615879577951937f7557797f7701007e007901007e81517a7561537a75527a527a527902fd009f635279547a75537a537a537a56795193547993577a75567a567a567a567a567a567a67527902fd009c63615879577953937f75577951937f77007901007e81517a7561547a75537a537a537a56795393547993577a75567a567a567a567a567a567a67527902fe009c63615879577955937f75577951937f77007901007e81517a7561547a75537a537a537a56795593547993577a75567a567a567a567a567a567a67615879577959937f75577951937f77007901007e81517a7561547a75537a537a537a56795993547993577a75567a567a567a567a567a567a686868547958799c63587957797f7557795579947f77527a75517a6854795193557a75547a547a547a547a6861547951799f6356795893577a75567a567a567a567a567a567a615879577951937f7557797f7701007e007901007e81517a7561537a75527a527a527902fd009f635279547a75537a537a537a56795193547993577a75567a567a567a567a567a567a67527902fd009c63615879577953937f75577951937f77007901007e81517a7561547a75537a537a537a56795393547993577a75567a567a567a567a567a567a67527902fe009c63615879577955937f75577951937f77007901007e81517a7561547a75537a537a537a56795593547993577a75567a567a567a567a567a567a67615879577959937f75577951937f77007901007e81517a7561547a75537a537a537a56795993547993577a75567a567a567a567a567a567a686868547958799c63587957797f7557795579947f77527a75517a6854795193557a75547a547a547a547a6861547951799f6356795893577a75567a567a567a567a567a567a615879577951937f7557797f7701007e007901007e81517a7561537a75527a527a527902fd009f635279547a75537a537a537a56795193547993577a75567a567a567a567a567a567a67527902fd009c63615879577953937f75577951937f77007901007e81517a7561547a75537a537a537a56795393547993577a75567a567a567a567a567a567a67527902fe009c63615879577955937f75577951937f77007901007e81517a7561547a75537a537a537a56795593547993577a75567a567a567a567a567a567a67615879577959937f75577951937f77007901007e81517a7561547a75537a537a537a56795993547993577a75567a567a567a567a567a567a686868547958799c63587957797f7557795579947f77527a75517a6854795193557a75547a547a547a547a685179517a75517a75517a75517a75517a75517a75517a75517a75517a75610079a9527987777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777767012a79518763610139796100792097dfd76851bf465e8f715593b217714858bbe9570ff3bd5e33840a34e20ff0262102ba79df5f8ae7604a9830f03c7933028186aede0675a16f025dc4f8be8eec0382210ac407f0e4bd44bfc207355a778b046225a7068fc59ee7eda43ad905aadbffc800206c266b30e6a1319c66dc401e5bd6b432ba49688eecd118297041da8074ce0810201008ce7480da41702918d1ec8e6849ba32b4d65b1e40dc669c31a1e6306b266c013079013079855679aa616100790079517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e01007e81517a756157795679567956795679537956795479577995939521414136d08c5ed2bf3ba048afe6dcaebafeffffffffffffffffffffffffffffff0061517951795179517997527a75517a5179009f635179517993527a75517a685179517a75517a7561527a75517a517951795296a0630079527994527a75517a68537982775279827754527993517993013051797e527e53797e57797e527e52797e5579517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f517f7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7c7e7e56797e0079517a75517a75517a75517a75517a75517a75517a75517a75517a75517a75517a75517a756100795779ac517a75517a75517a75517a75517a75517a75517a75517a75517a7561517a756169013579aa61013a79007901247f75547f77517a7561876901367900a06901307900a0696101397961007901687f7700005279517f75007f77007901fd8763615379537f75517f77007901007e81517a7561537a75527a527a5379535479937f75537f77527a75517a67007901fe8763615379557f75517f77007901007e81517a7561537a75527a527a5379555479937f75557f77527a75517a67007901ff8763615379597f75517f77007901007e81517a7561537a75527a527a5379595479937f75597f77527a75517a67615379517f75007f77007901007e81517a7561537a75527a527a5379515479937f75517f77527a75517a6868685179517a75517a75517a75517a7561517a756100798277615179517951795179011a7994012579937f755179011a79947f77517a75517a7561013b79a951798769013a79013c79ac69000000013c7953a169013c7951a063615279013d799f63013979537951935879957f7553795879957f77610139795279011479937f7552797f77007901007e81517a756100517900a063013b795479011679935379937f755479011679937f77517a7568537901157993527993547a75537a537a537a61527951797e0079a8007982775296517951797f75a8527952797f77a87e01007e81517a75517a75517a756161013b7957795193011879957f755779011879957f77007901007e81517a75610079517995014079975279014179979c690140795879011b7995011a79937f755879011b79957f776101417959795193011c79957f755979011c7995011b79937f77007901007e81517a756151796157790079011679011b79937f750116797f77517a7561876900796157790079011579011a79937f750115797f7781517a75619c696156790079011179013079937f750111797f7701007e81517a756159795179935a7a75597a597a597a597a597a597a597a597a597a75757575757575756852795193537a75527a527a615279013d799f63013979537951935879957f7553795879957f77610139795279011479937f7552797f77007901007e81517a756100517900a063013b795479011679935379937f755479011679937f77517a7568537901157993527993547a75537a537a537a61527951797e0079a8007982775296517951797f75a8527952797f77a87e01007e81517a75517a75517a756161013b7957795193011879957f755779011879957f77007901007e81517a75610079517995014079975279014179979c690140795879011b7995011a79937f755879011b79957f776101417959795193011c79957f755979011c7995011b79937f77007901007e81517a756151796157790079011679011b79937f750116797f77517a7561876900796157790079011579011a79937f750115797f7781517a75619c696156790079011179013079937f750111797f7701007e81517a756159795179935a7a75597a597a597a597a597a597a597a597a597a75757575757575756852795193537a75527a527a615279013d799f63013979537951935879957f7553795879957f77610139795279011479937f7552797f77007901007e81517a756100517900a063013b795479011679935379937f755479011679937f77517a7568537901157993527993547a75537a537a537a61527951797e0079a8007982775296517951797f75a8527952797f77a87e01007e81517a75517a75517a756161013b7957795193011879957f755779011879957f77007901007e81517a75610079517995014079975279014179979c690140795879011b7995011a79937f755879011b79957f776101417959795193011c79957f755979011c7995011b79937f77007901007e81517a756151796157790079011679011b79937f750116797f77517a7561876900796157790079011579011a79937f750115797f7781517a75619c696156790079011179013079937f750111797f7701007e81517a756159795179935a7a75597a597a597a597a597a597a597a597a597a75757575757575756852795193537a75527a527a6761557955796151795179011f7994012a79937f755179011f79947f77007901007e81517a7561517a75517a7561527a75517a680000547a75537a537a537a0001387953a1696154790139799f6301377955795193012979957f755579012979957f776101377956795193012b79957f755679012b79957f77007901007e81517a75615279517993537a75527a527a61597959795379537953795379012479947f75007f7752797e517958807e547954797f755479012779947f777e517a75517a75517a75517a7561610138795879519358957f75587958957f77007901007e81517a756161517951790079013279806152790079827700517902fd009f63615179515179517951938000795179827751947f75007f77517a75517a75517a7561517a75675179030000019f6301fd615279525179517951938000795179827751947f75007f77517a75517a75517a75617e517a756751790500000000019f6301fe615279545179517951938000795179827751947f75007f77517a75517a75517a75617e517a75675179090000000000000000019f6301ff615279585179517951938000795179827751947f75007f77517a75517a75517a75617e517a7568686868007953797e517a75517a75517a75617e517a75517a7561567951797e577a75567a567a567a567a567a567a75757575756854795193557a75547a547a547a547a6154790139799f6301377955795193012979957f755579012979957f776101377956795193012b79957f755679012b79957f77007901007e81517a75615279517993537a75527a527a61597959795379537953795379012479947f75007f7752797e517958807e547954797f755479012779947f777e517a75517a75517a75517a7561610138795879519358957f75587958957f77007901007e81517a756161517951790079013279806152790079827700517902fd009f63615179515179517951938000795179827751947f75007f77517a75517a75517a7561517a75675179030000019f6301fd615279525179517951938000795179827751947f75007f77517a75517a75517a75617e517a756751790500000000019f6301fe615279545179517951938000795179827751947f75007f77517a75517a75517a75617e517a75675179090000000000000000019f6301ff615279585179517951938000795179827751947f75007f77517a75517a75517a75617e517a7568686868007953797e517a75517a75517a75617e517a75517a7561567951797e577a75567a567a567a567a567a567a75757575756854795193557a75547a547a547a547a6154790139799f6301377955795193012979957f755579012979957f776101377956795193012b79957f755679012b79957f77007901007e81517a75615279517993537a75527a527a61597959795379537953795379012479947f75007f7752797e517958807e547954797f755479012779947f777e517a75517a75517a75517a7561610138795879519358957f75587958957f77007901007e81517a756161517951790079013279806152790079827700517902fd009f63615179515179517951938000795179827751947f75007f77517a75517a75517a7561517a75675179030000019f6301fd615279525179517951938000795179827751947f75007f77517a75517a75517a75617e517a756751790500000000019f6301fe615279545179517951938000795179827751947f75007f77517a75517a75517a75617e517a75675179090000000000000000019f6301ff615279585179517951938000795179827751947f75007f77517a75517a75517a75617e517a7568686868007953797e517a75517a75517a75617e517a75517a7561567951797e577a75567a567a567a567a567a567a75757575756854795193557a75547a547a547a547a537951799c6901347900a06361013379013079012f797e012c797e51797e0130797e012e797e517a75616100790136790079012f79806152790079827700517902fd009f63615179515179517951938000795179827751947f75007f77517a75517a75517a7561517a75675179030000019f6301fd615279525179517951938000795179827751947f75007f77517a75517a75517a75617e517a756751790500000000019f6301fe615279545179517951938000795179827751947f75007f77517a75517a75517a75617e517a75675179090000000000000000019f6301ff615279585179517951938000795179827751947f75007f77517a75517a75517a75617e517a7568686868007953797e517a75517a75517a75617e517a75517a7561537951797e547a75537a537a537a7575685179aa007961014379007982775179517958947f7551790128947f77517a75517a75618777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777670068686a4c70f70849e48ed22aae08262a4a3ca34efdfd644949546573742042544300000000000000000000000054425443000000000000000809c86063546bdb7f6e55714176a271aac687750c00ca9a3b00000000598d220eaecb68cf783cbc6cc6295d042897874f010000006f7261636c657376",
}

func TestDecode(t *testing.T) {
	for _, scriptHex := range scripts {
		script, _ := hex.DecodeString(scriptHex)

		txo := &TxoData{
			IsNFT:      false,
			CodeHash:   empty,
			GenesisId:  empty,
			AddressPkh: empty,

			MetaTxId: empty,
			TokenIdx: 0,

			Name:    "",
			Symbol:  "",
			Amount:  0,
			Decimal: 0,
		}

		DecodeSensibleTxo(script, txo)
		t.Logf("txo: %#v", txo)
	}
}
