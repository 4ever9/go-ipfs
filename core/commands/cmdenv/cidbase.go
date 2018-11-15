package cmdenv

import (
	cidenc "gx/ipfs/QmWf8NwKFLbTBvAvZst3bYF7WEEetzxWyMhvQ885cj9MM8/go-cidutil/cidenc"
	cmds "gx/ipfs/Qma6uuSyjkecGhMFFLfzyJDPyoDtNJSHJNweDccZhaWkgU/go-ipfs-cmds"
	cmdkit "gx/ipfs/Qmde5VP1qUkyQXKCfmEUA7bP64V2HAptbJ7phuPp7jXWwg/go-ipfs-cmdkit"
	mbase "gx/ipfs/QmekxXDhCxCJRNuzmHreuaT3BsuJcsjcXWNrtV9C8DRHtd/go-multibase"
)

var OptionCidBase = cmdkit.StringOption("cid-base", "Multi-base encoding used for version 1 CIDs in output.")
var OptionOutputCidV1 = cmdkit.BoolOption("output-cidv1", "Upgrade CID version 0 to version 1 in output.")

// ProcCidBase processes the `cid-base` and `output-cidv1` options and
// returns a encoder to use based on those parameters.
func ProcCidBase(req *cmds.Request) (cidenc.Encoder, error) {
	base, _ := req.Options["cid-base"].(string)
	upgrade, upgradeDefined := req.Options["output-cidv1"].(bool)

	var e cidenc.Encoder = cidenc.Default

	if base != "" {
		var err error
		e.Base, err = mbase.EncoderByName(base)
		if err != nil {
			return e, err
		}
		if !upgradeDefined {
			e.Upgrade = true
		}
	}

	if upgradeDefined {
		e.Upgrade = upgrade
	}

	return e, nil
}

// ProcCidBaseClientSide processes the `cid-base` and `output-cidv1`
// options and sets the default encoder based on those options
func ProcCidBaseClientSide(req *cmds.Request) error {
	enc, err := ProcCidBase(req)
	if err != nil {
		return err
	}
	cidenc.Default = enc
	return nil
}
